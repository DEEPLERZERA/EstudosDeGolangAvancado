package main //Chamando pacote principal

//Importando bibliotecas
import (  
	"fmt"
	"net/http"
	"time"
)

//Chamando função principal
func main() {
	//Abrindo uma requisição de escrita no endereço http que tem como caminho /tempo
	http.HandleFunc(
		"/tempo",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
		})
	//Especificando a porta do servidor que no nosso caso é 8080
	http.ListenAndServe(":8080", nil)
}