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
	var alunos []models.Aluno
	database.DB.Find(&alunos) ///Acha o dado alunos na base de dados com um ponteiro
	c.JSON(200, alunos)       //Retorna 200 e pega os dados da slice Alunos
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

//Criando função que busca aluno por id
func BuscaAlunoPorID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")   //Busca todos os parâmetros por id
	database.DB.First(&aluno, id) //Retorna o primeiro que achar que bateu

	if aluno.ID == 0 { //Verifica se ID está igual a zero
		c.JSON(http.StatusNotFound, gin.H{ //Retorna que não foi found
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno) //retorna status code de sucesso e o aluno
}

//Criando função de deletar
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")    //Busca parâmetros por ID
	database.DB.Delete(&aluno, id) //Deleta da base de dados
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

//Criando função de editar alunos
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")   //Busca parâmetros por ID
	database.DB.First(&aluno, id) //Procura no banco

	if err := c.ShouldBindJSON(&aluno); err != nil { //Verifica se não há erro
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno) //Atualiza o modelo na base de dados
	c.JSON(http.StatusOK, aluno)                   //Retorna status ok
}

//Criando função que retorna por busca de cpf
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno

	cpf := c.Param("cpf")                                    //Busca por parâmetro cpf
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno) //O primeiro cpf que acha equivalente retorna

	if aluno.ID == 0 { //Verificando se existe
		c.JSON(http.StatusNotFound, gin.H{ //Retorna que não foi found
			"Not found": "Aluno não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno) //Retorna status code de sucesso e o aluno
}
