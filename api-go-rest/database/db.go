package database  //Importando pacote database

import (  //Importando bibliotecas
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (   //Criando variáveis
    DB  *gorm.DB
    err error
)

func ConectaComBancoDeDados() {  //Criando função que conecta com banco de dados
    stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"  //Atribuindo dados
    DB, err = gorm.Open(postgres.Open(stringDeConexao)) //Iniciando conexão com postgres
    if err != nil {     //Verifica se há erro
        log.Panic("Erro ao conectar com banco de dados")
    }
}