package command

import (
	"bufio"
	"io"
	"log"
)

const (
	exitCommand  = "exit"
	cleanCommand = "clean"
)

type TwitterEngine interface {
	Search(location string) ([]string, error)
}

type Publisher interface {
	PublishResult(location string, hashtags []string)
	CleanAll()
}

type Commander struct {
	store  Publisher
	engine TwitterEngine
	reader io.Reader
}

func NewCommander(publisher Publisher, engine TwitterEngine, reader io.Reader) *Commander {
	return &Commander{store: publisher, engine: engine, reader: reader}
}

func (c *Commander) Listen() error {
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		cmd := scanner.Text()
		switch cmd {
		case exitCommand:
			return nil

		case cleanCommand:
			c.store.CleanAll()

		default:
			go func() {
				hashtags, err := c.engine.Search(cmd)
				if err != nil {
					log.Printf("unable to search for %s: %v", cmd, err)
					return
				}

				if len(hashtags) == 0 {
					log.Printf("No hashtags found for %s", cmd)
					return
				}

				if cmd == "" {
					cmd = "Worldwide"
				}

				c.store.PublishResult(cmd, hashtags)
			}()
		}
	}

	return scanner.Err()
}
