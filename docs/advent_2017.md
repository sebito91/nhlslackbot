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
specific environment needs. For the purposes of this test I'll be using Fedora 26 (4.14.6-200.fc26.x86_64),
using these tools:

1. ngrok for Slack API replies -- https://api.slack.com/tutorials/tunneling-with-ngrok
2. NHL statsapi to collect hockey scores -- https://statsapi.web.nhl.com/api/v1/schedule
3. the excellent golang slack library from nlopes -- https://github.com/nlopes/slack

### The Slack API

The Slack [API](https://api.slack.com/slack-apps) is well flushed out and spells out what specific
payloads to anticipate for any particular object. There are a number of calls you can develop
your bot to address, but in our case here we'll look at using the Real Time Messaging API 
([RTM](https://api.slack.com/rtm)) and specifically the `chat.postMessage` and `chat.postEphemeral`
methods. For this to work successfully you'll either need to set up an `ngrok` listener for your chosen
localhost port, or develop on a server that it externally routable (e.g. DigitalOcean droplet). In my
case here I'm developing on my local laptop but would deploy permanently on an droplet.
