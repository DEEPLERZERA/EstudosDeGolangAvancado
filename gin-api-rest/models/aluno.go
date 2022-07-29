package models //DEFININDO PACOTE MODELS

//IMPORTANDO BIBLIOTECAS
import "gorm.io/gorm"

//DEFININDO ESTRUTURA ALUNO
type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
