package controllers //DEFINE PACOTE CONTROLLERS

//IMPORTANDO BIBLIOTECAS
import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Esta função exibe todos os alunos com base no contexto do gin
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos) //Retorna 200 e pega os dados da slice Alunos
}

//Criando função de saudação segundo nome passado como parâmetro na url
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome") //Pega o nome que foi passado no campo url
	c.JSON(200, gin.H{              //Retorna 200 e imprime na tela o que tem que ser impresso
		"API diz:": "E ai " + nome + ",Tudo em beleza?",
	})
}

//Criando função que cria aluno
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno                           //Define aluno
	if err := c.ShouldBindJSON(&aluno); err != nil { //Atribue json a aluno e verifica se há erro
		c.JSON(http.StatusBadRequest, gin.H{ //Se houver erro retorna http status bad request
			"error": err.Error()}) //Imprime erro
		return
	}
	database.DB.Create(&aluno)        //Cria database com base no aluno
	c.JSON(http.StatusCreated, aluno) //De volve status code de criado
}
