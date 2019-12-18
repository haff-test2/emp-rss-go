package main

import (
	"encoding/json"
	"fmt"
	"github.com/haff-test2/emp-rss-go/reader"
	"io/ioutil"
	"os"
)

func main() {
	feedUrls := os.Args[1:]

	var items []reader.RssItem
	items = reader.Parse(feedUrls)
	serialized, err := writeJson("./last_feeds.json", items)
	if err == nil {
		fmt.Println(string(serialized))
	} else {
		panic(err)
	}
}

func writeJson(filePath string, items []reader.RssItem) ([]byte, error) {
	serializedItems, err := json.Marshal(items)
	err = ioutil.WriteFile(filePath, serializedItems, 0644)

	return serializedItems, err
}
