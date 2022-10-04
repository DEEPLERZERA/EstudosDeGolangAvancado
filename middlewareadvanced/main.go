package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc  //Declarando Middleware


func Logging() Middleware{  //Middleware de log
	return func(f http.HandlerFunc) http.HandlerFunc { //Retorna uma função que recebe um http.HandlerFunc
		return func (w http.ResponseWriter, r *http.Request) { //Retorna uma função que recebe um http.ResponseWriter e um http.Request

			start := time.Now() //Pega o tempo atual
			defer func() { log.Println(r.URL.Path, time.Since(start)) }() //Imprime o tempo que levou para executar a função

			f(w, r) 
	}

 }
}

func Method(m string) Middleware { //Middleware de método
	return func(f http.HandlerFunc) http.HandlerFunc { //Retorna uma função que recebe um http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) { //Retorna uma função que recebe um http.ResponseWriter e um http.Request
			if r.Method != m { //Se o método da requisição for diferente do método passado como parâmetro
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) //Retorna um erro
				return
			}

			f(w, r)
		}
	}
}


func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc { //Função que recebe um http.HandlerFunc e uma lista de middlewares
	for _, m := range middlewares { //Percorre a lista de middlewares
		f = m(f) //Executa o middleware passando a função como parâmetro
	}
	return f //Retorna a função
}

func Hello(w http.ResponseWriter, r *http.Request) { //Função que será executada de hello world
	fmt.Println(w, "Hello, world!") //Imprime hello world
}

func main() { //Função principal
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging())) //Chama a função Chain passando a função Hello, o middleware Method e o middleware Logging
	http.ListenAndServe(":8080", nil) //Inicia o servidor na porta 8080
}