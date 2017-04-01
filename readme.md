## Server Status

Scans a list of servers checking whether the ports are open or not.
This bot will send a chat notification when the status of a server changes.

I originally made this bot the check if private World of Warcraft servers were up or not.
This bot is actually much more useful than that and can be used for any type of server.

## Configuration

- Download the latest release [here](https://github.com/mgerb/ServerStatus/releases)
- Add your bot token as well as other configurations to config.json
- Execute the OS specific binary!

## Compiling from source

- Make sure Go and Make are installed
- make all

### How to get the bot token
https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token

### How to get your room ID

To get IDs, turn on Developer Mode in the Discord client (User Settings -> Appearance) and then right-click your name/icon anywhere in the client and select Copy ID.

<img src="https://camo.githubusercontent.com/9f759ec8b45a6e9dd2242bc64c82897c74f84a25/687474703a2f2f692e696d6775722e636f6d2f47684b70424d512e676966"/>

## List server status in discord channel

`!ServerStatus`

```
Elysium PvP is online!
Zethkur PvP is online!
Anathema PvP is online!
Darrowshire PvE is online!
Elysium Authentication Server is online!
```
