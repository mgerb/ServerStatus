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

// Start - add command, start port scanner and bot listeners
func Start() {
    //add command
    _, err := bot.Session.ApplicationCommandCreate(bot.Session.State.User.ID, "", &discordgo.ApplicationCommand {
        Name: "server-status",
        Description: "Get the status of the servers.",
    })

    if err != nil {
        log.Panicf("Cannot create status command '%v'", err)
    }

    //set each server status as online to start
	for i := range config.Config.Servers {
		config.Config.Servers[i].Online = true
		config.Config.Servers[i].OnlineTimestamp = time.Now()
		config.Config.Servers[i].OfflineTimestamp = time.Now()
	}

	err = bot.Session.UpdateStatusComplex(discordgo.UpdateStatusData { 
        Status: "online", 
        Activities: []*discordgo.Activity {
            &discordgo.Activity {
                Type: discordgo.ActivityTypeGame,
                Name: config.Config.GameStatus,
            },
        },
    })

	sendMessageToRooms(blue, "Server Status", "Bot started! Type /server-status to see the status of your servers :smiley:", false)

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

// InteractionHandler will be called every time an interaction from a user occurs
// Command interaction handling requires bot command scope
func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    // A user is calling us with our status command
    if i.ApplicationCommandData().Name == "server-status" {
        online := ""
        offline := ""

        for _, server := range config.Config.Servers {
            if server.Online {
                online = online + server.Name + "  :  " + fmtDuration(time.Since(server.OnlineTimestamp)) + "\n"
            } else {
                offline = offline + server.Name + "  :  " + fmtDuration(time.Since(server.OfflineTimestamp)) + "\n"
            }
        }

        // Only one message can be an interaction response. Messages can only contain up to 10 embeds.
        // Our message will therefore instead be two embeds (online and offline), each with a list of servers in text.
        // Embed descriptions can be ~4096 characters, so no limits should get hit with this.
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse {
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData {
                Embeds: []*discordgo.MessageEmbed {
                    {
                        Title: ":white_check_mark: Online",
                        Color: green,
                        Description: online,
                    },
                    {
                        Title: ":x: Offline",
                        Color: red,
                        Description: offline,
                    },
                },
            },
        })
    }
}

func fmtDuration(d time.Duration) string {

	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}
