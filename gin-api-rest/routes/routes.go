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
	r.POST("/aluno", controllers.CriaNovoAluno)             //Cria novos dados na base de dados
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)       //Busca por ID
	r.DELETE("/alunos/:id", controllers.DeletaAluno)        //Deleta aluno
	r.PATCH("/alunos/:id", controllers.EditaAluno)          //Edita aluno
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF) //Busca por CPF
	r.Run()                                                 //Rodando gin
}
