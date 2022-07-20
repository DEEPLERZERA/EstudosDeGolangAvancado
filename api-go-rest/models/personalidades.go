package models //Definindo pacote models

type Personalidade struct { //Criando estrutura de dados
	Id       int    `json:"id"`  
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade //Atribuindo variável
