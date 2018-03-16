# Server Status
Scans a list of TCP servers checking checking which are currently online.
This bot will send a chat notification when the status of a server changes (goes on or offline).

I originally made this bot to check if private World of Warcraft servers were up or not.
It's actually much more useful than that and can be used for most servers.

NOTE: This bot currently does not work for UDP servers.

## Configuration
- Download the latest release [here](https://github.com/mgerb/ServerStatus/releases)
- Add your bot token as well as other configurations to config.json
- Execute the OS specific binary!

### Mentioning Roles/Users
- you must first get your role/user id
- for user `<@userid>`
- for role `<@&roleid>`

## Usage
To get the current status of your servers simply type `!ServerStatus` in chat.

![Server Status](https://i.imgur.com/ZzQSBJp.png)

## Compiling from source
- Make sure Go and Make are installed
- make all

### How to get the bot token
https://github.com/reactiflux/discord-irc/wiki/Creating-a-discord-bot-&-getting-a-token

### How to get your room ID
To get IDs, turn on Developer Mode in the Discord client (User Settings -> Appearance) and then right-click your name/icon anywhere in the client and select Copy ID.

<img src="https://camo.githubusercontent.com/9f759ec8b45a6e9dd2242bc64c82897c74f84a25/687474703a2f2f692e696d6775722e636f6d2f47684b70424d512e676966"/>

