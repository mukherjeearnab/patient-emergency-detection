package main

import (
	"encoding/json"
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"io/ioutil"
	"os"
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
  // Remove Old Crypto Files
  os.Remove("ClientKey.json")
	os.Remove("ClientToken.txt")

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

	x := []float64{1, 1, 8, 1, 1}
	fmt.Println(x)

	// Encrypt Vector X using Secret Key
	cipher := TPE.Encrypt(x)
	fmt.Println("Generated Positive Case Cipher.\n" + cipher)

	// Decrypt Cipher and obtain result
	dec := TPE.Decrypt(cipher, token)
	fmt.Println("\n\nDecrypted Result: " + fmt.Sprint(dec))

	////////////////////////////////////////////////////////////////////////////////////////////////

	// // Load Config JSON
	// configJSON, _ := ioutil.ReadFile("ClientConfig.json")
	// ClientConfig := clientConfig{}
	// err := json.Unmarshal(configJSON, &ClientConfig)
	// if err != nil {
	// 	fmt.Println("Error parsing Client Config")
	// }

	for a := 0; a < 10; a++ {
		// Load Token
		tokenBytes, _ := ioutil.ReadFile("ClientToken.txt")
		token0 := string(tokenBytes)

		// fmt.Println("Loading Client Config.")
		// fmt.Println(ClientConfig)

		// Init TPE instance
		var TPE0 tpe.TPE

		// Setup TPE instance
		fmt.Println("Creating TPE instance.")
		TPE0.Setup(ClientConfig.TPE.N, ClientConfig.TPE.Theta)

		//Load Key JSON
		keyBytes, _ := ioutil.ReadFile("ClientKey.json")
		keyJSON := string(keyBytes)

		// Load Key into TPE instance
		TPE0.ImportKey(keyJSON)
		fmt.Println("Imported Key into TPE instance.")

		// Encrypt Vector X using Secret Key
		cipher0 := TPE0.Encrypt(x)
		// fmt.Println("Generated Positive Case Cipher.\n" + cipher)

		// Decrypt Cipher and obtain result
		dec0 := TPE0.Decrypt(cipher0, token0)
		fmt.Println("\n\nDecrypted Result: " + fmt.Sprint(dec0))
	}
	//////////////////////////////////////////////////////////

	return ClientConfig.NetID
}

func exportText(content string, filename string) {
	data := []byte(content)

	err := ioutil.WriteFile(filename, data, 0)

	if err != nil {
		fmt.Println(err)
	}
}
