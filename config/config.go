package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Variables used for command line parameters
var Config configStruct

type configStruct struct {
	Token         string   `json:"Token"`
	RoomIDList    []string `json:"RoomIDList"`
	RolesToNotify []string `json:"RolesToNotify"`
	Servers       []server `json:"Servers"`
	GameStatus    string   `json:"GameStatus"`
}

type server struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	Port    int    `json:"Port"`
	Online  bool   `json:"Online,omitempty"`
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

}
