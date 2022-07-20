package main //Chamando pacote principal

import ( //Importando bibliotecas
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() { //Chamando função principal

	models.Personalidades = []models.Personalidade{ //Atribuindo dados ao nosso slice de personalidades
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
		{Id: 3, Nome: "Nome3", Historia: "Historia3"},
	}
	database.ConectaComBancoDeDados()              //Gera a conexão
	fmt.Println("Iniciando nossa api rest com go") //Imprimindo na tela
	routes.HandleRequest()                         //Chamando api de handleRequest
}
