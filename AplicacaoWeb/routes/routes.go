package routes //DEFININDO PACOTE ROUTES

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/aplicacaoweb/controllers"
	"net/http"
)

//FUNÇÃO DE CARREGAR ROTAS
func CarregaRotas() {
	http.HandleFunc("/", controllers.Index) //Define caminho
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
