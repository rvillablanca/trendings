package trends

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

const (
	trends24URL = "https://trends24.in/"
)

type Trends24 struct {
}

func (t *Trends24) Search(location string) ([]string, error) {
	resp, err := http.Get(trends24URL + location)
	if err != nil {
		return nil, fmt.Errorf("unable to search for %s: %w", location, err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var hashtags []string
	doc.Find("div#trend-list div.trend-card").First().Find("ol li a").Each(func(i int, selection *goquery.Selection) {
		hashtags = append(hashtags, selection.Text())
	})

	return hashtags, nil
}
