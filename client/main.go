package main

import (
// "fmt"
)

func main() {
	// NetID := GenerateKey("ClientConfig.json")
	// fmt.Printf("Network ID: %s\n", NetID)
	SendReading(true, "http://192.168.1.100:3000")
}
