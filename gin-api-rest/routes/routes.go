package routes //DEFININDO PACOTE DE ROTAS

//IMPORTANDO BIBLIOTECAS
import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

//Criando função que define os caminhos e rotas a serem seguidos
func HandleRequests() {
	r := gin.Default()                             //Starta gin
	r.GET("/alunos", controllers.ExibeTodosAlunos) //Criando rotas que tem como parâmetros suas funções que estão em controllers
	r.GET("/:nome", controllers.Saudacao)
	r.Run() //Rodando gin
}
