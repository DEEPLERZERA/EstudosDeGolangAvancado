package main //DEFININDO PACOTE MAIN

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/gin-api-rest/routes"
	"gin-api-rest/database"
	"gin-api-rest/models"
)

//DEFININDO FUNÇÃO PRINCIPAL
func main() {
	database.ConectaComBancoDeDados() //Conecta com banco de dados
	models.Alunos = []models.Aluno{   //Atribue dados ao models
		{Nome: "Daniel Borges", CPF: "0002301230", RG: "2131231232"},
		{Nome: "Andrey Ferrari", CPF: "111111111", RG: "2124124214214"},
	}
	routes.HandleRequests() //Chama requisição de rotas
}
