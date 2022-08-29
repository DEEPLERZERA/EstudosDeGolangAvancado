package main //Importando pacote main

import (   ///Importando biblioteca fmt
	"fmt"
)

type Pessoa struct { //Criando tipo Pessoa
	nome string 
	idade int 
	endereco string
}

func main() {  //Criando função principal
	pessoa := Pessoa{"João", 20, "Rua A"} //Criando variável pessoa e atribuindo valores
	fmt.Println(pessoa) //Imprimindo na tela a variável pessoa
	falar() //Chamando função falar
	comer("Pizza", pessoa.nome) //Chamando função comer

}

func falar() { //Criando função falar
	fmt.Println("Oi") //Imprimindo na tela
}

func comer(alimento, nome string)  { //Criando função comer
	fmt.Printf("%v esta comendo %v",nome, alimento) //Imprimindo na tela
}

