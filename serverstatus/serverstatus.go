package serverstatus

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	portscanner "github.com/anvie/port-scanner"
	"github.com/bwmarrin/discordgo"
	steam "github.com/kidoman/go-steam"
	"github.com/mgerb/ServerStatus/bot"
	"github.com/mgerb/ServerStatus/config"
)

const (
	red   = 0xf4425c
	green = 0x42f477
	blue  = 0x42adf4
)

// Start - start port scanner and bot listeners
func Start() {
	//set each server status as online to start
	for i := range config.Config.Servers {
		config.Config.Servers[i].Online = true
		config.Config.Servers[i].OnlineTimestamp = time.Now()
		config.Config.Servers[i].OfflineTimestamp = time.Now()
	}

	err := bot.Session.UpdateStatus(0, config.Config.GameStatus)

	sendMessageToRooms(blue, "Server Status", fmt.Sprintf("Bot started! Type %sServerStatus to see the status of your servers :smiley:", config.Config.BotPrefix), false)

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

		// use waitgroup to scan all servers concurrently
		var wg sync.WaitGroup

		for index := range config.Config.Servers {
			wg.Add(1)
			go worker(&config.Config.Servers[index], &wg)
		}

		wg.Wait()

		time.Sleep(time.Second * config.Config.PollingInterval)
	}
}

func worker(server *config.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	prevServerUp := server.Online //set value to previous server status

	var serverUp bool
	retryCounter := 0

	// try reconnecting 5 times if failure persists (every 2 seconds)
	for {
		serverScanner := portscanner.NewPortScanner(server.Address, time.Second*2, 1)
		serverUp = serverScanner.IsOpen(server.Port) //check if the port is open

		// if server isn't up check RCON protocol (UDP)
		if !serverUp {
			host := server.Address + ":" + strconv.Itoa(server.Port)
			steamConnection, err := steam.Connect(host)
			if err == nil {
				defer steamConnection.Close()
				_, err := steamConnection.Ping()
				if err == nil {
					serverUp = true
				}
			}
		}

		if serverUp || retryCounter >= 5 {
			break
		}

		retryCounter++
		time.Sleep(time.Second * 2)
	}

	if serverUp && serverUp != prevServerUp {
		server.OnlineTimestamp = time.Now()
		sendMessageToRooms(green, server.Name, "Is now online :smiley:", true)
	} else if !serverUp && serverUp != prevServerUp {
		server.OfflineTimestamp = time.Now()
		sendMessageToRooms(red, server.Name, "Has gone offline :frowning2:", true)
	}

	server.Online = serverUp
}

func sendMessageToRooms(color int, title, description string, mentionRoles bool) {
	for _, roomID := range config.Config.RoomIDList {
		if mentionRoles {
			content := strings.Join(config.Config.RolesToNotify, " ")
			bot.Session.ChannelMessageSend(roomID, content)
		}
		sendEmbeddedMessage(roomID, color, title, description)
	}
}

func sendEmbeddedMessage(roomID string, color int, title, description string) {

	embed := &discordgo.MessageEmbed{
		Color:       color,
		Title:       title,
		Description: description,
	}

	bot.Session.ChannelMessageSendEmbed(roomID, embed)
}

// MessageHandler will be called every time a new
// message is created on any channel that the autenticated bot has access to.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == bot.BotID {
		return
	}

	if m.Content == config.Config.BotPrefix+"ServerStatus" {
		for _, server := range config.Config.Servers {
			if server.Online {
				sendEmbeddedMessage(m.ChannelID, green, server.Name, "Online!\nUptime: "+fmtDuration(time.Since(server.OnlineTimestamp)))
			} else {
				sendEmbeddedMessage(m.ChannelID, red, server.Name, "Offline!\nDowntime: "+fmtDuration(time.Since(server.OfflineTimestamp)))
			}
		}
	}
}

func fmtDuration(d time.Duration) string {

	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}
