package models


//DEFININDO ESTRUTURA ALUNO
type Aluno struct {
	Nome string `json:"nome"`
	CPF  int `json:"cpf"`
	RG   int `json:"rg"`
	ID   int `json:"ID"`
}