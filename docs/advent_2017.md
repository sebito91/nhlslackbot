# Create a Slack bot with golang 

### Introduction

In this post we'll look at how to set up a quick Slack bot that receives messages (either direct or
from channel) and replies to the user. I've been an IRC user for many years and always loved setting up 
bots, whether for sports scores, weather, or something else entirely. Recently I've actually had an 
opportunity to implement my first Slack bot and figured I would document the process for others! You 
can find all of the code for this post listed [here](https://github.com/sebito91/nhlslackbot), 
and PRs are certainly welcome :D 

For this assignment we'll need a few things, not all of which are covered in this post. I invite the
reader to take a look at the installation practices for the other software dependencies based on their
specific environment needs. Here I'll be using Fedora 26 (4.14.6-200.fc26.x86_64) along with these tools:

1. ngrok for Slack API replies -- https://ngrok.com/docs#expose 
2. NHL statsapi to collect hockey scores -- https://statsapi.web.nhl.com/api/v1/schedule
3. the excellent golang slack library from nlopes -- https://github.com/nlopes/slack

You'll either need to set up an `ngrok` listener for your chosen localhost port, or develop on a server 
that it externally routable (e.g. DigitalOcean droplet). In my case here I'm developing on my laptop but 
would deploy permanently on a droplet.

### The Slack API

#### Initial Configuration

The Slack [API](https://api.slack.com/slack-apps) is well flushed out and spells out what specific
payloads to anticipate for any particular object. There are a number of calls you can develop
your bot to address, but in our case here we'll look at using the Real Time Messaging API 
([RTM](https://api.slack.com/rtm)) and specifically the `chat.postMessage` and `chat.postEphemeral`
methods.

Before any of our code is working we'll need to set up an app within slack itself. Navigate to the
app registration [tool](https://api.slack.com/apps?new_app=1) to create a new application within your
workspace. Here I've created the `NHL Scores` app within my workspace.

create_app.png

Once done you'll be presented with a number of options for your new application. Here we'll need to create
a `Bot User` that will act as our listener within the workspace. My example is called `nhlslackbot` and
will be visible to all users within the workspace once mounted.

bot_user.png

We'll need to generate an OAuth token for our user in order to actually connect with the Slack API. To do so
click on the `OAuth & Permissions` section to `Install App to Workspace` which will prompt you to authorize
access and generate the tokens you'll use. You'll need to copy the `Bot User OAuth Access Token` somewhere local,
but always make sure this is not shared anywhere! This token is secret and should be treated like your
password!

authorize.png

Lastly we'll need to set up the `Interative Components` of our application and specify the ngrok (or other) 
endpoint that the API will send responses to. In my case, I've added a custom ngrok value here called `https://sebtest.ngrok.io/`. This endpoint is where we'll receive all correspondence from Slack itself, and this is how
we'll be able to process any incoming messages from the channels.

interactive.png

With that all sorted, we can finally dig into the code!

#### Code components

The crux of the code is how we handle receiving messages from the slack connection. Using the `Bot User OAuth
Access Token` to establish the initial connection, we must continuously poll the system for incoming messages.
The API gives us the ability to trigger off of a number of event types, such as:

1. Hello Events
2. Connected Events
3. Presence Change Events
4. Message Events
5. and many more

The beauty of this verbosity is that we can trigger messages on a number of different use-cases, really
giving us the ability to tailor the bot to our specific needs. For this example, we'll look at using the
`*slack.MessageEvent` type to support both indirect (within channel using `@`) or direct messages. From the 
library, The primary poll for message events leverages the `websocket` handler and just loops over events
until we've received one that we want:

```go
func (s *Slack) run(ctx context.Context) {
    slack.SetLogger(s.Logger)

    rtm := s.Client.NewRTM()
    go rtm.ManageConnection()

    s.Logger.Printf("[INFO]  now listening for incoming messages...")
    for msg := range rtm.IncomingEvents {
        switch ev := msg.Data.(type) {
        case *slack.MessageEvent:
            if len(ev.User) == 0 {
                continue
            }

            // check if we have a DM, or standard channel post
            direct := strings.HasPrefix(ev.Msg.Channel, "D")

            if !direct && !strings.Contains(ev.Msg.Text, "@"+s.UserID) {
                // msg not for us!
                continue
            }

            user, err := s.Client.GetUserInfo(ev.User)
            if err != nil {
                s.Logger.Printf("[WARN]  could not grab user information: %s", ev.User)
                continue
            }

            s.Logger.Printf("[DEBUG] received message from %s (%s)\n", user.Profile.RealName, ev.User)

            err = s.askIntent(ev)
            if err != nil {
                s.Logger.Printf("[ERROR] posting ephemeral reply to user (%s): %+v\n", ev.User, err)
            }
        case *slack.RTMError:
            s.Logger.Printf("[ERROR] %s\n", ev.Error())
        }
    }
}
```

Once we confirm that the message is indeed directed to us, we pass the event handler along to our `askIntent`
function. Remember that this is a contrived example that's just going to send back NHL game scores to the
user, iff they acknowledge that specific intent. We could build up an entire workflow around this user
interaction that would send different paths depending on user choices to our prompts, or have no prompts
at all! Those different cases are outside the scope of this introductory post, so for now we just want to 
send back a quick `Yes` v `No` prompt and handle accordingly.

To do precisely that, our handler `askIntent` will process the message and genreate an `chat.postEphemeral`
message to send back to the event user (aka the person asking for details). The "ephemeral" post is one that's
directed _only_ to the requester. Though other users will see the initial request to the bot if within the 
same channel, the subsequent interaction with the bot will only be done between the user and the bot. From
the docs:

> This method posts an ephemeral message, which is visible only to the assigned user in a specific public channel, private channel, or private conversation.


### The NHL API
