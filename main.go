package main

import (
	"fmt"

	"github.com/mgerb/ServerStatus/bot"
	"github.com/mgerb/ServerStatus/config"
	"github.com/mgerb/ServerStatus/serverstatus"
)

// Variables used for command line parameters
var (
	BotID string
)

var version = "undefined"

func init() {
	fmt.Println("Starting Server Status " + version)
}

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
