package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendLocation(location string, server string) {
	// Load Config JSON
	configJSON, _ := ioutil.ReadFile("ClientConfig.json")
	ClientConfig := clientConfig{}
	err := json.Unmarshal(configJSON, &ClientConfig)
	if err != nil {
		fmt.Println("Error parsing Client Config")
	}

	fmt.Println("Loading Client Config.")
	// fmt.Println(ClientConfig)

	fmt.Println("Sending Location.")
	sendLocationHTTP(ClientConfig.NetID, location, server)
}

func sendLocationHTTP(NetID string, location string, server string) {
	// Get JWT Authentication Token
	jwt := loginJWT("p1", "1234", server)
	// fmt.Println(jwt)

	//Encode the data
	postBody, _ := json.Marshal(map[string]string{
		"location": location,
	})

	reqBody := bytes.NewBuffer(postBody)
	//Leverage Go's HTTP Post function to make request
	client := &http.Client{}
	req, _ := http.NewRequest("POST", server+"/api/location/set/"+NetID, reqBody)
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

	fmt.Println("Sent HTTP POST request containing Location. RES=" + out)
}
