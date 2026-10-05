package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			break
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}
func main() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)
	fmt.Println("Type something and press Enter (enter Q to quit)")
	for {
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)
		if userInput == "Q" || userInput == "q" {
			break
		}
		ping <- userInput
		//wait responde
		response := <-pong
		fmt.Println("Response: ", response)
	}
	fmt.Println("Bye")
	close(ping)
	close(pong)
}
