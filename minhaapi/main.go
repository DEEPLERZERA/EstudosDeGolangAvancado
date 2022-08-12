package main //Definindo pacote principal

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/minhaapi/routes" //modulo
	"net/http"
)

//CHAMANDO FUNÇÃO MAIN
func main() {
	routes.CarregaRotas()             //Chama rotas
	http.ListenAndServe(":8000", nil) //Fica ouvindo requisições na porta 8000
}
