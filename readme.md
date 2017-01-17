## Server Status

Scans the Elysium servers checking to see if they are up.
This bot will send a chat notification when the status of a server changes.

## Configuration

- Download or Clone the repository
- Add your bot token and room ID to the config.json file in the dist folder
- Execute the Windows or Linux binary!

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

## Compiling

Make sure you have Go installed.

`make clean all`
