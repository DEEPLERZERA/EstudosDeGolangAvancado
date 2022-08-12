package controllers

import (
	"ProjetosDeGolang/minhaapi/db"
	"ProjetosDeGolang/minhaapi/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

//Esta função exibe todos os alunos com base no contexto do gin
func ExibeTodosAlunos(c *gin.Context) {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados
	AlunoDoBanco, err := db.Query("select * from alunos") //Faz uma busca na base de dados de acordo com id

	if err != nil { //Verifica se há erro
		panic(err.Error())
	}

	c.JSON(200, AlunoDoBanco)       //Retorna 200 e pega os dados da slice Alunos
}

//Criando função que cria aluno
func CriaNovoAluno(c *gin.Context) {
	var Nome string 
	var Cpf, Rg, ID int
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados
	var aluno models.Aluno                           //Define aluno
	if err := c.ShouldBindJSON(&aluno); err != nil { //Atribue json a aluno e verifica se há erro
		c.JSON(http.StatusBadRequest, gin.H{ //Se houver erro retorna http status bad request
			"error": err.Error()}) //Imprime erro
		return
	}
	insereDadosNoBanco, err := db.Prepare("insert into alunos(id, nome, rg, cpf) values($1, $2, $3, $4)") //Prepara para inserir dados no BD
	if err != nil {                                                                                                          //Verifica se há erro
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(ID, Nome, Rg, Cpf) //Se não executa ordem e insere os dados
	defer db.Close()                                            //Fecha o banco de dados
	
	c.JSON(http.StatusCreated, insereDadosNoBanco) //De volve status code de criado
}

//Criando função que busca aluno por id
func BuscaAlunoPorID(c *gin.Context) {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados
	var aluno models.Aluno        //Variável alunos do tipo models.Aluno
	id := c.Params.ByName("id")   //Busca todos os parâmetros por id
	ProdutoDoBanco, err := db.Query("select * from alunos where id = $1", id) //Faz uma busca na base de dados de acordo com id

	if err != nil { //Verifica se há erro
		panic(err.Error())
	}

	if aluno.ID == 0 { //Verifica se ID está igual a zero
		c.JSON(http.StatusNotFound, gin.H{ //Retorna que não foi found
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, ProdutoDoBanco) //retorna status code de sucesso e o aluno
}


//Criando função de deletar
func DeletaAluno(c *gin.Context) {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados
	id := c.Params.ByName("id")                                               //Busca parâmetros por ID
	DeletaOProduto, err := db.Prepare("delete from alunos where id = $1") //Prepara para deletar dados do banco de dados
	if err != nil {                                                         //Verifica se há erro
		panic(err.Error())
	}

	DeletaOProduto.Exec(id) //Executa ordem de deletar produto
	defer db.Close()        //Fecha banco de dados
	c.JSON(http.StatusNoContent, gin.H{"data": "Aluno deletado com sucesso"}) //Retorna 204 e imprime na tela mensagem de deletado
}



//Criando função de editar alunos
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno        //Variável alunos do tipo models.Aluno
	id := c.Params.ByName("id")   //Busca parâmetros por ID
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados

	ProdutoDoBanco, err := db.Query("select * from alunos where id = $1", id) //Faz uma busca na base de dados de acordo com id

	if err := c.ShouldBindJSON(&aluno); err != nil { //Verifica se não há erro
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	Alunos:= models.Aluno{} //Atribue slice para produtoParaAtualizar

	for ProdutoDoBanco.Next() { //percorre os dados do produto
		var ID, Rg, Cpf int
		var nome string

		err = ProdutoDoBanco.Scan(&ID, &nome, &Rg, &Cpf) //Scanneia todos um por um

		if err != nil { //Verifica se há erro
			panic(err.Error())
		}
		Alunos.ID = ID //Atribue os dados para produtoParaAtualizar que contém um slice em sua memória
		Alunos.Nome = nome
		Alunos.RG = Rg
		Alunos.CPF = Cpf
	}
	defer db.Close() //Fecha o banco de dados
	c.JSON(http.StatusOK, Alunos)                   //Retorna status ok
}

