package handler

import (
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed"
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed/mock_http"
	"net/http"
	"testing"
)

func TestNewsFeedGet(t *testing.T) {
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{"hello", "World"}) // add an item to the feed

	handler := NewsFeedGet(feed)

	w := &mock_http.ResponseWriter{}

	r := &http.Request{}

	handler(w, r)

	result := w.GetBodyJSONArray()

	if len(result) !=  1 {
		t.Errorf("Item was not added to the datastore")
	}

	if result[0]["title"] != "hello" {
		t.Errorf("Item was not properly set")
	}

}
