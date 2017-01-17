package serverstatus

import (
	"../bot"
	"../config"
	"fmt"
	"github.com/anvie/port-scanner"
	"github.com/bwmarrin/discordgo"
	"time"
)

func Start() {
	//set each server status as online to start
	for i, _ := range config.Config.Servers {
		config.Config.Servers[i].Online = true
	}

	//start a new go routine
	go loop()
}

func loop() {

	//check if server are in config file
	if len(config.Config.Servers) < 1 {
		fmt.Println("No servers in config file.")
		return
	}

	for {

		for index, server := range config.Config.Servers {
			prevServerUp := server.Online //set value to previous server status

			elysiumPvP := portscanner.NewPortScanner(server.Address, time.Second*2)
			serverUp := elysiumPvP.IsOpen(server.Port) //check if the port is open

			if serverUp && serverUp != prevServerUp {
				sendMessage("@here " + server.Name + " is now online!")
			} else if !serverUp && serverUp != prevServerUp {
				sendMessage("@here " + server.Name + " went offline!")
			}

			config.Config.Servers[index].Online = serverUp
		}

		time.Sleep(time.Second * 5)
	}
}

func sendMessage(message string) {
	for _, roomID := range config.Config.RoomIDList {
		bot.Session.ChannelMessageSend(roomID, message)
	}
}

// This function will be called every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == bot.BotID {
		return
	}

	if m.Content == "!ServerStatus" {
		for _, server := range config.Config.Servers {
			if server.Online {
				s.ChannelMessageSend(m.ChannelID, server.Name+" is online!")
			} else {
				s.ChannelMessageSend(m.ChannelID, server.Name+" is down!")
			}
		}
	}
}
