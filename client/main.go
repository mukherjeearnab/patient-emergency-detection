package main

import (
	"fmt"
)

func main() {
	NetID := GenerateKey("ClientConfig.json")
	fmt.Printf("Network ID: %s\n", NetID)

}
