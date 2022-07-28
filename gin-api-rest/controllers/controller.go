package controllers //DEFINE PACOTE CONTROLLERS

//IMPORTANDO BIBLIOTECAS
import (
	"gin-api-rest/models"

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
