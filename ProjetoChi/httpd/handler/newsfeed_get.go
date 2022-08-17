package handler

import (
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed"
	"encoding/json"
	"net/http"
)


func NewsFeedGet(feed newsfeed.Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	}
}