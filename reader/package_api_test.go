package reader

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("Expected 'GET' request, got '%s'", r.Method)
		}

		// if r.URL.EscapedPath() != expected_path {
		// t.Errorf("Expected request to '%s', got '%s'", expected_path, r.URL.EscapedPath())
		// }

		dat, err := ioutil.ReadFile("./test/wired_test.xml")
		if err != nil {
			panic(err)
		}
		// io.Write(w, dat)
		io.WriteString(w, string(dat))
	}))
	defer ts.Close()

	// urls := []string{"http://feeds.wired.com/wired/index"}
	urls := []string{ts.URL}
	result := Parse(urls)
	if len(result) != 2 {
		t.Errorf("expect result count to eq 2, but got %d", len(result))
	}

	expected := []RssItem{
		RssItem{
			Title:       "Earth's Largest Scientific Structure, a WhatsApp Flaw, and More News",
			Description: "Catch up on the most important news from today in two minutes or less.",
			Link:        "https://www.wired.com/story/largest-scientific-structure-whatsapp-group-chat-hack",
			PublishDate: time.Unix(1576624776, 0),
		},
		RssItem{
			Title:       "What a 5,700-Year-Old Piece of Gum Reveals About Its Chewer",
			Description: "From a wad of pitch less than an inch long, researchers have painted a detailed portrait of an ancient human—and added another layer to the story of human evolution.",
			Link:        "https://www.wired.com/story/5700-year-old-piece-of-gum",
			PublishDate: time.Unix(1576598400, 0),
		},
	}

	for i, res := range result {
		if res.Title != expected[i].Title {
			t.Errorf("got unexpected Title %s instead of %s", res.Title, expected[i].Title)
		}
	}

	for i, res := range result {
		if res.Description != expected[i].Description {
			t.Errorf("got unexpected Description %s instead of %s", res.Description, expected[i].Description)
		}
	}

	for i, res := range result {
		if res.Link != expected[i].Link {
			t.Errorf("got unexpected Link %s instead of %s", res.Link, expected[i].Link)
		}
	}

	for i, res := range result {
		if res.PublishDate.Unix() != expected[i].PublishDate.Unix() {
			t.Errorf("got unexpected Link %s instead of %s", res.PublishDate, expected[i].PublishDate)
		}
	}
}
