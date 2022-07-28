package database //DEFININDO PACOTE DATABASE

//IMPORTANDO BIBLIOTECAS
import (
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ( //Criando variáveis
	DB  *gorm.DB
	err error
)

//Definindo função que conecta com banco de dados
func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432  sslmode=disable" //Passa os dados do banco
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil { //Verifica se há erro
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{}) //Imigra todos os dados para o banco de dados
}
