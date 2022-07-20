package main  //Importando pacote principal

import "fmt"  //Importando biblioteca fmt

func main() { //Chamando função main
	var irmao string  //Criando variável
	fmt.Println("Você tem irmãos? : ")  //Fazendo pergunta ao usuário
	fmt.Scan(&irmao) //Coletando input

	if irmao == "sim" { //Fazendo condicional
		var quantos int //Criando outra variável
		fmt.Println("Quantos irmãos você tem? : ") //Fazendo outra pergunta ao usuário
		fmt.Scan(&quantos)  //Recebendo input 
		fmt.Printf("Então você tem: %d irmãos!!!!", quantos) //Imprimindo na tela de acordo com inputs do usuário Nota %d serve para imprimir valores inteiros!!!!
	} else if irmao == "não" {  //Se não faça
		var gostaria string  //Cria variável
		fmt.Println("Gostaria de ter? : ") //Imprime pergunta na tela 
		fmt.Scan(&gostaria) //Recebe input do usuário
		fmt.Printf("Então você %s gostaria de ter um irmão!!!!", gostaria)  //Imprime na tela de acordo com inputs do usuário Nota %s serve para imprimir valores do tipo string!!!!

	} else {  //Se não faça
		fmt.Println("Opção invalida!!!!")  //Imprime opção invalida!!!
	}
}
