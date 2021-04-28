package services

import (
	"log"

	"github.com/mmcdole/gofeed"
)

// RetrieveItemsFromRss : retrieves items from any rss feed
func RetrieveItemsFromRss(rssURI string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL("https://" + rssURI)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return feed.Items, nil
}
