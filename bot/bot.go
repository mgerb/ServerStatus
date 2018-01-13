package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	BotID   string
	Session *discordgo.Session
)

func Connect(token string) {
	// Create a new Discord session using the provided bot token.
	var err error
	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get the account information.
	u, err := Session.User("@me")
	if err != nil {
		fmt.Println("error obtaining account details,", err)
	}

	// Store the account ID for later use.
	BotID = u.ID

	fmt.Println("Bot connected")
}

func Start() {
	// Open the websocket and begin listening.
	err := Session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	return
}

func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}
