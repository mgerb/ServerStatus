package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Variables used for command line parameters
var Config configStruct

type configStruct struct {
	Token           string        `json:"Token"`
	RoomIDList      []string      `json:"RoomIDList"`
	RolesToNotify   []string      `json:"RolesToNotify"`
	Servers         []Server      `json:"Servers"`
	GameStatus      string        `json:"GameStatus"`
	PollingInterval time.Duration `json:"PollingInterval"`
}

type Server struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
	Online  bool   `json:"Online,omitempty"`
	// OnlineTimestamp - time of when the server last came online
	OnlineTimestamp  time.Time
	OfflineTimestamp time.Time
}

func Configure() {

	fmt.Println("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		log.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	err := json.Unmarshal(file, &Config)

	if err != nil {
		log.Println(err)
	}

	if Config.PollingInterval == 0 {
		log.Fatal("Please set your PollingInterval > 0 in your config file.")
	}

}
