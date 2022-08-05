package db //Definindo pacote db

//Importando bibliotecas
import (
	"database/sql"

	_ "github.com/lib/pq"
)

//CRIANDO FUNÇÃO DE CONECTAR COM BANCO DE DADOS
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=root dbname=root password = root host = localhost port = 5432 sslmode = disable" //Passa parâmetros de coneo
	db, err := sql.Open("postgres", conexao)                                                       //Abre db
	if err != nil {                                                                                //Verifica se há erro
		panic(err.Error())
	}
	return db
}
