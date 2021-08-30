package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	tpe "github.com/mukherjeearnab/gotpe"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func SendReading(positive bool, server string) bool {
	// Load Config JSON
	configJSON, _ := ioutil.ReadFile("ClientConfig.json")
	ClientConfig := clientConfig{}
	err := json.Unmarshal(configJSON, &ClientConfig)
	if err != nil {
		fmt.Println("Error parsing Client Config")
	}

	fmt.Println("\nLoading Client Config.")
	// fmt.Println(ClientConfig)

	// Init TPE instance
	var TPE tpe.TPE

	// Setup TPE instance
	fmt.Println("Creating TPE instance.")
	TPE.Setup(ClientConfig.TPE.N, ClientConfig.TPE.Theta)

	//Load Key JSON
	keyBytes, _ := ioutil.ReadFile("ClientKey.json")
	keyJSON := string(keyBytes)

	// Load Key into TPE instance
	TPE.ImportKey(keyJSON)
	fmt.Println("Imported Key into TPE instance.")

	if positive {
		// Positive Case
		// Create Vector X
		x := []float64{1, 1, 9, 1, 1}

		// Encrypt Vector X using Secret Key
		cipher := TPE.Encrypt(x)
		fmt.Println("Generated Positive Case Cipher.")

		// Send Cipher to server
		flag := sendReadingHTTP(ClientConfig.NetID, cipher, server)
		return flag
	} else {
		// Negative Case
		// Create Vector X
		x := []float64{1, 1, 11, 1, 1}

		// Encrypt Vector X using Secret Key
		cipher := TPE.Encrypt(x)
		fmt.Println("Generated Negative Case Cipher.")

		// Send Cipher to server
		flag := sendReadingHTTP(ClientConfig.NetID, cipher, server)
		return flag
	}
}

func sendReadingHTTP(NetID string, cipher string, server string) bool {
	// Get JWT Authentication Token
	jwt := loginJWT("p1", "1234", server)
	// fmt.Println(jwt)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", server+"/api/detection/check/"+NetID+"/?cipher="+url.QueryEscape(cipher), nil)
	req.Header.Set("x-access-token", jwt)
	res, err := client.Do(req)

	//Handle Error
	if err != nil {
		fmt.Printf("An Error Occured %v\n", err)
	}
	defer res.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	out := string(body)

	fmt.Println("Sent HTTP POST request containing Cipher. RES=" + out)
	if strings.Contains(out, "1") {
		return true
	} else {
		return false
	}
}

func loginJWT(username string, password string, server string) string {
	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"username":     username,
		"password":     password,
		"organization": "patient",
	})

	responseBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(server+"/api/auth/loginpt", "application/json", responseBody)

	//Handle Error
	if err != nil {
		fmt.Printf("An Error Occured %v\n", err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	jwt := string(body)
	fmt.Println("Obtained JWT from Server.")

	return jwt
}
