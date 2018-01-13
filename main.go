package main

import (
	"github.com/mgerb/serverstatus/bot"
	"github.com/mgerb/serverstatus/config"
	"github.com/mgerb/serverstatus/serverstatus"
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

	// add handlers
	bot.AddHandler(serverstatus.MessageHandler)

	//start websocket to listen for messages
	bot.Start()

	//start server status task
	serverstatus.Start()

	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
}
