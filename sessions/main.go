package main

import (        //Importando dependências
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (     //Criando variáveis
	key = []byte("chave-super-secreta") //Chave para criptografar os cookies
	store = sessions.NewCookieStore(key)    //Criando uma sessão
)

func secret(w http.ResponseWriter, r *http.Request) {   //Função para verificar se o usuário está logado
	session, _ := store.Get(r, "cookie-name")  //Pegando a sessão

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth { //Verificando se o usuário está logado
		http.Error(w, "Forbidden", http.StatusForbidden) //Se não estiver logado, retorna um erro 403
		return
	}

	fmt.Fprintln(w, "Alegrias!") //Se estiver logado, retorna a mensagem de felicidades!
}

func login(w http.ResponseWriter, r *http.Request) {  //Função para fazer o login
	session, _ := store.Get(r, "cookie-name") //Pegando a sessão


	session.Values["authenticated"] = true //Setando o valor da sessão como true
	session.Save(r, w) //Salvando a sessão
}

func logout(w http.ResponseWriter, r *http.Request) { //Função para fazer o logout
	session, _ := store.Get(r, "cookie-name") //Pegando a sessão

	session.Values["authenticated"] = false //Setando o valor da sessão como false
	session.Save(r, w) //Salvando a sessão
}



//Criando rotas e definindo porta como 8080
func main() {  
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080", nil)
}
