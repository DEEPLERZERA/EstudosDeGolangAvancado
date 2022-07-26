package handler

import (
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed"
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed/mock_http"
	"net/http"
	"testing"
)

func TestNewsfeedPost(t *testing.T) {
	feed := newsfeed.New()

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")

	w := &mock_http.ResponseWriter{}
	r := &http.Request{
		Header: headers,
	}

	r.Body = mock_http.RequestBody(map[string]string{
		"title" : "Legal",
		"post"  : "anime",
	})

	handler := NewsFeedPost(feed)

	handler(w, r)

	result:= w.GetBodyString()

	if result != "Good job!" {
		t.Errorf("Handler did not complete")
	}

	if len(feed.GetAll()) != 1 {
		t.Errorf("Item did not add")
	}

	if feed.GetAll()[0].Title != "hello" {
		t.Errorf("Item bad")
	}

}
