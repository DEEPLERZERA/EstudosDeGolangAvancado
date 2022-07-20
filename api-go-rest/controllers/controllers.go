package controllers //Chamando pacote de controllers

import ( //Importando bibliotecas
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) { //Criando função home que consegue escrever e fazer request
	fmt.Fprint(w, "Home Page") //Imprime na página
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) { //Criando função que escreve e faz request no http
	w.Header().Set("Content-type", "application/json") //Deixa o arquivo mais organizado
	var p []models.Personalidade                       //Criando slice p que recebe personalidades
	database.DB.Find(&p)                               //O banco de dados faz uma varredura
	json.NewEncoder(w).Encode(p)                       //Chamando models.Personalidade e aplicando encode
}

//CRIANDO FUNÇÃO QUE RETORNA A PERSONALIDADE DE ACORDO COM ID PASSADO PELO USUÁRIO NA URL
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade   //Variável personalidade recebe os modelos
	database.DB.Find(&personalidade, id)     //Database acha a personalidade de acordo com id
	json.NewEncoder(w).Encode(personalidade) //Imprime na tela os dados

}

//CRIANDO FUNÇÃO QUE CRIA PERSONALIDADES
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated) //Troca status code para 201 significando que foi criado com sucesso
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade) //Recebe dados
	database.DB.Create(&novaPersonalidade)             //Criando BD
	json.NewEncoder(w).Encode(novaPersonalidade)       //Demonstrando
}

//CRIANDO FUNÇÃO QUE DELETA PERSONALIDADES
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id) //Deletando
	json.NewEncoder(w).Encode(personalidade)

}

//CRIANDO FUNÇÃO QUE EDITA PERSONALIDADES
func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id) //Busca personalidade no banco pelo ID
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade) //Salva o banco de dados(UPDATE)
	json.NewEncoder(w).Encode(personalidade)
}

//CRIANDO FUNÇÃO QUE EDITA APENAS UM DADO DA PERSONALIDADE
func EditaDadoUnicoPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id) //Busca personalidade no banco pelo ID
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade) //Salva o banco de dados(UPDATE)
	json.NewEncoder(w).Encode(personalidade)
}
