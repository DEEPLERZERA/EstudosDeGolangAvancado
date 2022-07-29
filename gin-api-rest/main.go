package main //DEFININDO PACOTE MAIN

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/gin-api-rest/routes"
	"gin-api-rest/database"
)

//DEFININDO FUNÇÃO PRINCIPAL
func main() {
	database.ConectaComBancoDeDados() //Conecta com banco de dados
	routes.HandleRequests()           //Chama requisição de rotas
}
