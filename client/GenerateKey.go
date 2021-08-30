package main

import (
	"encoding/json"
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"io/ioutil"
	"time"
)

type tpeConfig struct {
	N     int     `json:"N"`
	Theta float64 `json:"Theta"`
}

type clientConfig struct {
	TPE   tpeConfig `json:"TPE"`
	Y     []float64 `json:"Y"`
	NetID string    `json:"NetID"`
}

func GenerateKey(config string) string {
	// Load Config JSON
	configJSON, _ := ioutil.ReadFile(config)
	ClientConfig := clientConfig{}
	err := json.Unmarshal(configJSON, &ClientConfig)
	if err != nil {
		fmt.Println("Error parsing Client Config")
	}

	fmt.Println("Loading Client Config.")
	fmt.Println(ClientConfig)

	// Create a seed
	seed := time.Now().UnixNano()

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	fmt.Println("Creating TPE instance.")
	TPE.Setup(ClientConfig.TPE.N, ClientConfig.TPE.Theta)

	// Generate a new Secret Key
	fmt.Println("Generating Key.")
	TPE.KeyGen(seed)
	exportText(TPE.ExportKey(), "ClientKey.json")
	fmt.Println("Exported Key to ClientKey.json")

	// Generate a new Token using Y and Secret Key
	fmt.Println("Generating Token.")
	token := TPE.TokenGen(ClientConfig.Y)
	exportText(token, "ClientToken.txt")
	fmt.Println("Exported Key to Token")

	return ClientConfig.NetID
}

func exportText(content string, filename string) {
	data := []byte(content)

	err := ioutil.WriteFile(filename, data, 0)

	if err != nil {
		fmt.Println(err)
	}
}
