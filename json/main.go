package main

import (   //Importando dependencias
	"encoding/json"
	"fmt"
	"net/http"
)
 
type User struct {  //Definindo a estrutura do objeto
	Firstname string `json:"firstname"` //Definindo o nome do campo no json
	Lastname  string `json:"lastname"` 
	Age       int    `json:"age"` 
}


func main() { //Função principal
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) { //Definindo a rota decode
		var user User //Instanciando o objeto
		json.NewDecoder(r.Body).Decode(&user) //Decodificando o json e atribuindo ao objeto

		fmt.Fprintf(w, "%s %s is %d years old", user.Firstname, user.Lastname, user.Age) //Imprimindo o objeto
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) { //Definindo a rota encode
		peter := User { //Instanciando o objeto
			Firstname: "Jonh",
			Lastname:  "Doe",
			Age:       25,
		}

		json.NewEncoder(w).Encode(peter) //Codificando o objeto e imprimindo no json
	})
 
	http.ListenAndServe(":8080", nil) //Iniciando o servidor
}