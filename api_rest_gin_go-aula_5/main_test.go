package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

//Gera um sistema de setup de rotas chamando gin
func SetupDasRotasDeTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

/*func TestFalhador(t *testing.T) {
	t.Fatalf("Teste falhou de proposito pobre!")
}
*/

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
