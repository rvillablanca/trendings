package main

import (
	"log"
	"os"
	"trendings/command"
	"trendings/server"
	"trendings/trends"
)

func main() {
	log.Println("creating server")
	srv := server.NewServer()

	srvCh := make(chan error)
	go func() {
		log.Println("running server")
		err := srv.Start()
		srvCh <- err
	}()

	log.Println("creating commander")
	commander := command.NewCommander(srv, &trends.Trends24{}, os.Stdin)
	cmdCh := make(chan error)
	go func() {
		log.Println("listening commands")
		err := commander.Listen()
		cmdCh <- err
	}()

	select {
	case serverErr := <-srvCh:
		log.Printf("unable to run server: %v", serverErr)
		return
	case cmdErr := <-cmdCh:
		if cmdErr != nil {
			log.Printf("an error ocurred while listening commands: %v", cmdErr)
		}
		return
	}
}
