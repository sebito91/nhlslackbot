# initial cut of the blog post

A few parts to this...

1. set up ngrok
2. set up slackbot in api
3. get/confirm NHL api works
4. set up nlopes/slack RTM
5. test out e2e

In this post we'll look at how to set up a quick Slack bot that receives messages (either direct or
from channel) and replies to the user. I've been an IRC user for many years and always loved setting up 
bots, whether for sports scores, weather, or something else entirely. Recently I've actually had an 
opportunity to implement my first Slack bot and figured I would document the process for others! 

The Slack [API](https://api.slack.com/slack-apps) is well flushed out and spells out what specific
payloads to anticipate for any particular object. There are a number of calls you can develop
your bot to address, but in our case here we'll look at using the Real Time Messaging API 
([RTM](https://api.slack.com/rtm)) and specifically the `chat.postMessage` and `chat.postEphemeral`
methods.
