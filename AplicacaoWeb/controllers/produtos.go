package controllers //DEFININDO PACOTE CONTROLLERS
//Importando bibliotecas
import (
	"ProjetosDeGolang/AplicacaoWeb/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) //Criando conexão com index.html

//CRIANDO FUNÇÃO INDEX
func Index(w http.ResponseWriter, r *http.Request) {
	TodosOsProdutos := models.BuscaTodosOsProdutos()  //Chamando função que vai buscar todos os produtos
	temp.ExecuteTemplate(w, "Index", TodosOsProdutos) //Chama pág html e passa os dados do BD

}

//FUNÇÃO EXIBE PÁGINA COM FORMULÁRIO 
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

//FUNÇÃO QUE RECEBE FORMULÁRIO E INSERE NO BANCO DE DADOS
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //Verifica se é método post
		nome := r.FormValue("nome") //Busca valores do forms e passa os valores
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64) //Faz conversão para float
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade) //Faz conversao para int
		if err != nil {
			log.Println("Erro na conversão do quantidade:", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt) //Chama função de criar novo produto e passa valores
	}
	http.Redirect(w, r, "/", 301) //Redireciona para página main e retorna status code 301 de retorno com sucesso
}

//FUNÇÃO DE DELETAR
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id") //Pega id
	models.DeletaProduto(idDoProduto)      //Chama função de deletar produto
	http.Redirect(w, r, "/", 301)          //Redireciona para página main e retorna status code 301 de retorno com sucesso
}

//CRIANDO FUNÇÃO DE EDITAR
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")      //Pega id
	produto := models.EditaProduto(idDoProduto) //Chama função de editar produto
	temp.ExecuteTemplate(w, "Edit", produto)    //Atribue a página o que irá editar
}

//FUNÇÃO DE ATUALIZAR DADOS
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { //Verifica se é método post
		id := r.FormValue("id") //Pega valores do forms
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id) //Converte para inteiro
		if err != nil {
			log.Println("Erro na conversão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64) //Converte para float
		if err != nil {
			log.Println("Erro na conversão do preco para float:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade) //Converte para int
		if err != nil {
			log.Println("Erro na conversão da quantidade para int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt) //Chama função de atualizar produto

	}
	http.Redirect(w, r, "/", 301) //Redireciona para página main e retorna status code 301 de retorno com sucesso

}
