package main

import (
	"encoding/json"
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"io/ioutil"
	// "time"
)

func SendReading(positive bool) {
	// Load Config JSON
	configJSON, _ := ioutil.ReadFile("ClientConfig.json")
	ClientConfig := clientConfig{}
	err := json.Unmarshal(configJSON, &ClientConfig)
	if err != nil {
		fmt.Println("Error parsing Client Config")
	}

	fmt.Println("Loading Client Config.")
	fmt.Println(ClientConfig)

	// Create a seed
	// seed := time.Now().UnixNano()

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	fmt.Println("Creating TPE instance.")
	TPE.Setup(ClientConfig.TPE.N, ClientConfig.TPE.Theta)

	// Load Key JSON
	// keyJSON, _ := ioutil.ReadFile("ClientKey.json")

	// if !positive { // Create Vector X
	// 	x := []float64{1, 1, 11, 1, 1}

	// 	// Encrypt Vector X using Secret Key
	// 	cipher := TPE.Encrypt(x)
	// 	fmt.Println("\n\nCipher: " + cipher)
	// } else {
	// 	// Create Vector X
	// 	x = []float64{1, 1, 9, 1, 1}

	// 	// Encrypt Vector X using Secret Key
	// 	cipher = TPE.Encrypt(x)
	// 	fmt.Println("\n\nCipher: " + cipher)
	// }
}
