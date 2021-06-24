package command

import "fmt"

type StdOutPublisher struct {
}

func (s *StdOutPublisher) PublishResult(location string, hashtags []string) {
	fmt.Printf("Location: %s: Hashtags: %v\n", location, hashtags)
}

func (s *StdOutPublisher) CleanAll() {

}
