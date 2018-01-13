package serverstatus

import (
	"log"
	"time"

	"github.com/anvie/port-scanner"
	"github.com/bwmarrin/discordgo"
	"github.com/mgerb/serverstatus/bot"
	"github.com/mgerb/serverstatus/config"
)

func Start() {
	//set each server status as online to start
	for i, _ := range config.Config.Servers {
		config.Config.Servers[i].Online = true
	}

	err := bot.Session.UpdateStatus(0, config.Config.GameStatus)

	if err != nil {
		log.Println(err)
	}

	//start a new go routine
	go scanServers()
}

func scanServers() {

	//check if server are in config file
	if len(config.Config.Servers) < 1 {
		log.Println("No servers in config file.")
		return
	}

	for {

		for index, server := range config.Config.Servers {
			prevServerUp := server.Online //set value to previous server status

			serverScanner := portscanner.NewPortScanner(server.Address, time.Second*2, 1)
			serverUp := serverScanner.IsOpen(server.Port) //check if the port is open

			if serverUp && serverUp != prevServerUp {
				sendMessage(config.Config.RoleToNotify + " " + server.Name + " is now online!")
			} else if !serverUp && serverUp != prevServerUp {
				sendMessage(config.Config.RoleToNotify + " " + server.Name + " went offline!")
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
