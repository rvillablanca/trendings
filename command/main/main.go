package main

import (
	"log"
	"os"
	"trendings/command"
	"trendings/trends"
)

func main() {

	commander := command.NewCommander(&command.StdOutPublisher{}, &trends.Trends24{}, os.Stdin)
	err := commander.Listen()
	if err != nil {
		log.Printf("error while listening commands: %v", err)
	}
}
