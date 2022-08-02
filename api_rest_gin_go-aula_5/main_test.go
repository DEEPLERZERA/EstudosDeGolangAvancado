package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

//Gera um sistema de setup de rotas chamando gin
func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

/*func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de proposito pobre!")
}
*/

//Função que cria aluno para testes
func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "123456789",
		RG: "11232141124"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

//Função que deleta aluno para testes
func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T) { //Retorna um test de saudacao
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/gui", nil) //Gera uma requisição  GET
	resposta := httptest.NewRecorder()            //Grava a resposta
	r.ServeHTTP(resposta, req)                    //Serve os dois status

	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais") //Compara as duas respostas por meio de um assert
	mockDaResposta := `{"API diz:":"E ai gui, tudo beleza?"}`            //Registra sua resposta de teste
	respostaBody, _ := ioutil.ReadAll(resposta.Body)                     //lê resposta
	assert.Equal(t, mockDaResposta, string(respostaBody))                //Compara
}

func TestListandoTodosOsAlunosHandler(t *testing.T) { //Função que testa lista de todos os alunos
	database.ConectaComBancoDeDados()
	CriaAlunoMock()         //Chamando aluno criado para teste
	defer DeletaAlunoMock() //Deletando aluno
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)   //rota
	req, _ := http.NewRequest("GET", "/alunos", nil) //Requisição a  ser feita
	resposta := httptest.NewRecorder()               //Guardando resposta
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code) //Comparando
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) { //Função que testa aluno por CPF
	database.ConectaComBancoDeDados()
	CriaAlunoMock()         //Cria aluno teste
	defer DeletaAlunoMock() //Deleta aluno teste
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)        //Especifica rota
	req, _ := http.NewRequest("GET", "/alunos/cpf/123456789", nil) //Faz requisição
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code) //Compara
}

func TestBuscaAlunoPorIDHandler(t *testing.T) { //Função que busca aluno por ID
	database.ConectaComBancoDeDados()
	CriaAlunoMock() //Cria aluno
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)  //Rota é especificada
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)       //Cria caminho separado
	req, _ := http.NewRequest("GET", pathDaBusca, nil) //Faz requisição
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)                               //Faz parar de ser json
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "O nome deve ser igual") //Começa a comparar
	assert.Equal(t, "123456789", alunoMock.CPF, "CPF deve ser igual")
	assert.Equal(t, "11232141124", alunoMock.RG, "RG deve ser igual")
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestDeletaAlunoHandler(t *testing.T) { //Função que testa deletar aluno
	database.ConectaComBancoDeDados()
	CriaAlunoMock()         //Cria aluno teste
	defer DeletaAlunoMock() //Deleta aluno teste
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno) //Especifica rota
	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil) //Faz requisição
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code) //Compara
}

func testEditaUmAlunoHandler(t *testing.T) { //Função que testa editar aluno
	database.ConectaComBancoDeDados()
	CriaAlunoMock()         //Cria aluno teste
	defer DeletaAlunoMock() //Deleta aluno teste
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)                                          //Cria rota de edição
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47123456789", RG: "123456700"} //Altera dados
	valorJson, _ := json.Marshal(aluno)                                                     //Transforma aluno em json
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)                                         //Faz rota de edição separada
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))          //Cria requisição PATCH e converte json para bytes
	resposta := httptest.NewRecorder()                                                      //Salva respsota
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado) //Retira resposta do formato json
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)     //Começa a comparar
	assert.Equal(t, "123456700", alunoMockAtualizado.RG)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	fmt.Println(alunoMockAtualizado.CPF)
}
