package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("1. Generate Keys.\n2. Send Reading.\nEnter Choice: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Entered Choice: " + text)
	if strings.Contains(text, "1") {
		t_keyGen()
	} else if strings.Contains(text, "2") {
		t_reading()
	} else {
		fmt.Println("Exiting.")
	}
}

func t_reading() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Reading Type (1/0): ")
		text, _ := reader.ReadString('\n')
		fmt.Println("Entered Choice: " + text)
		if strings.Contains(text, "1") {
			locFlag := SendReading(true, "http://192.168.1.100:3000")
			if locFlag {
				SendLocation("1.123,2.456", "http://192.168.1.100:3000")
			}
		} else if strings.Contains(text, "0") {
			SendReading(false, "http://192.168.1.100:3000")
		} else {
			fmt.Println("Exiting.")
			break
		}
	}
}

func t_keyGen() {
	NetID := GenerateKey("ClientConfig.json")
	fmt.Printf("Network ID: %s\n", NetID)
}
