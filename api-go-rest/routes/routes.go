package routes //Chamando pacote de rotas

import ( //Importando bibliotecas
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() { //Chamando função de HandleRequest
	r := mux.NewRouter()                                                                               //Criando nova rota com gorilla mux
	r.Use(middleware.ContentTypeMiddleware)                                                            //Usando um middleware em r
	r.HandleFunc("/", controllers.Home)                                                                //Chamando função de imprimir na tela do site
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")                //Imprimindo na tela todas as personalidades
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")       //Imprimindo na tela uma personalidade específica
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("Post")          //Imprimindo na tela a personalidade criada com método post
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")     //Chamando método de deletar personalidades
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")            //Chamando método de editar personalidades
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaDadoUnicoPersonalidade).Methods("Patch") //Chamando método de editar um dado único da personalidadeco método patch
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))  //Passando qual porta devemos ouvir e servir em caso de log.fatal, handlers.CORS permite que meu servidor aceite requisições de qualquer lugar
}
