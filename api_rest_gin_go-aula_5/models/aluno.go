package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"` //Implementando formas de validação
	CPF  string `json:"cpf"  validate:"len=9",   regexp=^[0-9]*$"`
	RG   string `json:"rg"   validate:"len=11,   regexp=^[0-9]*$"`
}

//Criando função que valida alunos recebendo como parâmetro os dados da struct e retornando um error caso necessário
func ValidaDadosDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}

	return nil
}
