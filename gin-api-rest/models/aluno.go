package models //DEFININDO PACOTE MODELS

//DEFININDO ESTRUTURA ALUNO
type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno //DECLARANDO VARIÁVEL
