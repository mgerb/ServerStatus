package main

import (
	"./bot"
	"./config"
	"./serverstatus"
)

// Variables used for command line parameters
var (
	BotID string
)

func main() {
	//read config file
	config.Configure()

	//connect bot to account with token
	bot.Connect(config.Config.Token)

	//start side tasks
	serverstatus.Start()

	bot.AddHandler(serverstatus.MessageHandler)

	//start websocket to listen for messages
	bot.Start()
}
