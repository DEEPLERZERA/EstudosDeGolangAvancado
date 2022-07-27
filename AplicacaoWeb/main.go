package main //Definindo pacote principal

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/aplicacaoweb/routes"
	"net/http"
)

//CHAMANDO FUNÇÃO MAIN
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil) //Fica ouvindo requisições na porta 8000
}
