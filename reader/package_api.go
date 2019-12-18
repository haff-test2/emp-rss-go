package reader

import (
	// "github.com/mmcdole/gofeed"
	// "github.com/mmcdole/gofeed/rss"
	"time"
)

type RssItem struct {
	Title       string
	Source      string
	SourceUrl   string
	Link        string
	PublishDate time.Time // task description typo?)
	Description string
}

func Parse(urls []string) []RssItem {
	var result []RssItem

	return result
}
