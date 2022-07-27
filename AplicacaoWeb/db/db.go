package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

//CRIANDO FUNÇÃO DE CONECTAR COM BANCO DE DADOS
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=postgres password = 12345 host = localhost sslmode = disable" //Passa parâmetros de conexão
	db, err := sql.Open("postgres", conexao)                                                       //Abre db
	if err != nil {                                                                                //Verifica se há erro
		panic(err.Error())
	}
	return db
}
