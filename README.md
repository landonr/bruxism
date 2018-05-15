# Bruxism
A chat bot for YouTube, Discord and IRC.

[![GoDoc](https://godoc.org/github.com/landonr/bruxism?status.svg)](https://godoc.org/github.com/landonr/bruxism)
[![Go report](http://goreportcard.com/badge/landonr/bruxism)](http://goreportcard.com/report/landonr/bruxism)
[![Build Status](https://travis-ci.org/landonr/bruxism.svg?branch=master)](https://travis-ci.org/landonr/bruxism)

## Current plugin support:

Commands are prefixed with `@BotName `.

* `emoji <emoji>` - Returns a big version of an emoji. Please read the README in emojiplugin for installation notes.
* `help [<topic>]` - Returns generic help or help for a specific topic. Available topics: `comic,remind`
* `invite <id>` - Provides invite URL for the bot.
* `numbertrivia [<number>]` - Returns trivia for a random number or a specified number if provided.
* `playing` - Set which game the bot is playing on Discord. Only enabled for bot owner.
* `reminder <time> | <reminder>` - Sets a reminder.
* `stats` - Lists bot statistics.
* `streamer <streamername|streamerid>` - Grabs details about a YouTube streamer.
* `topstreamers` - List the current top streamers on YouTube Gaming.

eg: `@BotName help`

Also supports direct invites on Discord and support for announcing streamers going live on YouTube Gaming.

## Usage:

### Installation:

`go get github.com/landonr/bruxism/cmd/bruxism`

`go install github.com/landonr/bruxism/cmd/bruxism`

`cd $GOPATH/bin`

### Setup

You must first generate a YouTube Oauth token in the [Google Developer Console](https://console.developers.google.com/).

Go to Credentials and download the JSON config file and save it in `$GOPATH/bin` as `youtubeoauth2config.json`

`./bruxism -youtubeurl`

You will then be given a URL (and copied to clipboard), visit that URL and copy the code, then run:

`./bruxism -youtubeauth <AUTH CODE>`

Now the bot can be run:

`./bruxism -youtubechannelids <youtube channel id1,youtube channel id2>`

#### Important!
The account you choose during the oauth flow is the account that the bot will use to send messages.

If you are trying to test the bot, make sure the account you are sending messages with is **different** from the bot account.

Bruxism ignores messages from the bot account, so if you are sharing accounts all commands will be ignored.


### Run as a Discord bot

`./bruxism -discordtoken "Bot <discord bot token>"`

It is suggested that you set `-discordapplicationclientid` if you are running a bot account, this will make `inviteplugin` function correctly.

It is suggested that you set `-discordowneruserid` as this prevents anyone from calling `playingplugin`.

To invite your bot to a server, visit: `https://discordapp.com/oauth2/authorize?client_id=<discord client id>&scope=bot`

### Run as an IRC bot

`./bruxism -ircserver <irc server> -ircusername <irc username> -ircchannels <#channel1,#channel2>`

### Run as a Slack bot

`./bruxism -slacktoken <slack token>`

It is suggested that you set `-slackowneruserid`.

#### Image uploads

Some plugins upload images to imgur (such as chartplugin), it's suggested that you run your bot with `-imgurid <imgurid>` to enable this.

## Arguments:

* `youtubeurl` - Outputs a new OAuth URL for YouTube and then exits.
* `youtubeauth` - Exchanges the provided auth code for an oauth2 token.
* `youtubeconfig` - The filename for your YouTube OAuth client JSON. (Download JSON in Google Developers Console -> Credentials).
* `youtubetoken` - The filename to store the oauth2 token.
* `youtubelivechannelids` - Comma separated list of channel ids to poll.
* `discordtoken` - Sets the Discord token.
* `discordemail` - Sets the Discord account email.
* `discordpassword` - Sets the Discord account password.
* `discordclientid` - Sets the Discord client id.
* `ircserver` - Sets the IRC server.
* `ircusername` - Sets the IRC user name.
* `ircpassword` - Sets the IRC password.
* `ircchannels` - Comma separated list of IRC channels.
* `imgurid` - Sets the Imgur client id, used for uploading images to imgur.
* `imguralbum` - Sets an optional the Imgur album id, used for uploading images to imgur.
* `mashablekey` - Sets the mashable oauth key.

## Special Thanks

[Bruce Marriner](https://github.com/bwmarrin/discordgo) - For DiscordGo and the Music Plugin.
