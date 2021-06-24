package main

import (
	"fmt"
	"log"
	"os"
	"trendings/trends"
)

func main() {
	var location string
	if len(os.Args) > 1 {
		location = os.Args[1]
	}

	t := trends.Trends24{}
	hashtags, err := t.Search(location)
	if err != nil {
		log.Printf("failed to get trends: %v", err)
		return
	}

	fmt.Println(hashtags)
}
