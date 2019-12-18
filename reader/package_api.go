package reader

import (
	"github.com/mmcdole/gofeed"
	"time"
)

type RssItem struct {
	FeedUrl     string
	Title       string
	Source      string
	SourceUrl   string
	Link        string
	PublishDate time.Time // task description typo?)
	Description string
}

func Parse(urls []string) []RssItem {
	var result []RssItem

	ch := make(chan []RssItem)
	for _, url := range urls {
		go parseFeedUrl(url, ch)
	}

	for i := 0; i < len(urls); i++ {
		newItemsBatch := <-ch
		result = append(result, newItemsBatch...)
	}

	return result
}

func parseFeedUrl(url string, ch chan []RssItem) {
	var resultItems []RssItem

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err == nil {
		for _, item := range feed.Items {
			newItem := buildRssItem(*item)
			resultItems = append(resultItems, newItem)
		}
	}

	ch <- resultItems
}

func buildRssItem(item gofeed.Item) RssItem {
	result := RssItem{
		Title: item.Title,
		// Source: ,
		// SourceUrl:,
		Link:        item.Link,
		Description: item.Description,
	}

	if item.PublishedParsed != nil {
		result.PublishDate = *(item.PublishedParsed)
	}

	return result
}
