package main

import (
	"ProjetosDeGolang/ProjetoChi/platform/newsfeed"
	"ProjetosDeGolang/ProjetoChi/httpd/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":3000"
	feed := newsfeed.New()
	feed.Add(newsfeed.Item{
		Title: "Andrey cara", 
		Post: "Um jogador caro",
	})

	r := chi.NewRouter()

	r.Get("/newsfeed", handler.NewsFeedGet(feed))

	r.Post("/newsfeed", handler.NewsFeedPost(feed))


	fmt.Println("Servidor iniciado na porta",port)
	http.ListenAndServe(port, r)
}
