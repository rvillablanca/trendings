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

type SearchEngine interface {
	Search(location string) ([]string, error)
}

type DataStore interface {
	AddSearchResult(location string, hashtags []string)
	Clean()
}

type Commander struct {
	store  DataStore
	engine SearchEngine
	reader io.Reader
}

func NewCommander(store DataStore, engine SearchEngine, reader io.Reader) *Commander {
	return &Commander{store: store, engine: engine, reader: reader}
}

func (c *Commander) Listen() error {
	scanner := bufio.NewScanner(c.reader)
	for scanner.Scan() {
		cmd := scanner.Text()
		switch cmd {
		case exitCommand:
			return nil

		case cleanCommand:
			c.store.Clean()

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

				log.Printf("Location: %s => %v", cmd, hashtags)
				c.store.AddSearchResult(cmd, hashtags)
			}()
		}
	}

	return scanner.Err()
}
