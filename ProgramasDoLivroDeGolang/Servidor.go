package main

import (
	"fmt"
	"log"
	"net/http"

	"programasdolivrodegolang/url"
)

var (
	porta   int
	urlBase string
)

type Headers map[string]string

func init() {
	porta = 8888
	urlBase = fmt.Sprintf("http://localhost:%d", porta)
}

func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, headers{
			"Allow": "POST",
		})
		return
	}

	url, nova, err := url.BuscarOuCriarNovaUrl(extrairUrl(r))

	if err != nil {
		responderCom(w, http.StatusBadRequest, nil)

		return
	}

	var status int

	if nova {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}
	urlCurta := fmt.Sprintf("%s/r%s", urlBase, url.Id)
	responderCom(w, status, Headers{"Location": urlCurta})
}

func responderCom(
	w http.ResponseWriter,
	status int,
	headers Headers,
) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func extrairUrl(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func main() {
	http.HandleFunc("/api/encurtar", Encurtador)
	http.HandleFunc("/r/", Redirecionador)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", porta), nil))
}
