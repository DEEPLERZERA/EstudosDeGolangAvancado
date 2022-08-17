package handler

import (
	"encoding/json"
	"net/http"
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed"
)

func NewsFeedPost(feed newsfeed.Adder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"], 
			Post:  request["post"],
		})	
		w.Write([]byte("Good job little boy!"))

	}
}
