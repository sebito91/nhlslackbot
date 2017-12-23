# initial cut of the blog post

A few parts to this...

1. set up ngrok
2. set up slackbot in api
3. get/confirm NHL api works
4. set up nlopes/slack RTM
5. test out e2e

### Introduction

In this post we'll look at how to set up a quick Slack bot that receives messages (either direct or
from channel) and replies to the user. I've been an IRC user for many years and always loved setting up 
bots, whether for sports scores, weather, or something else entirely. Recently I've actually had an 
opportunity to implement my first Slack bot and figured I would document the process for others! 

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

------



### The NHL API
