package main //Chamando pacote

import (
	"fmt"
	"os"
) //Importando bibliotecas

func main() { //Função main
	var x, y, resultado, op int //Criando variáveis 
	
	fmt.Println("Seja bem vindo a calculadora básica: ")  //Imprimindo na tela menu etc

	fmt.Println("Escolha uma opção: ")

	fmt.Println("1- Soma")
	fmt.Println("2- Subtração")
	fmt.Println("3- Sair")

	fmt.Scan(&op)  //Lendo input do usuário

	switch op { //Chamando switch-case
	case 1:  //Caso seja 1 faça
		fmt.Println("Digite o primeiro número: ")  //Imprima na tela
		fmt.Scan(&x)  //Leia input
		fmt.Println("Digite o segundo número: ") //Imprima na tela
		fmt.Scan(&y) //Leia input
		resultado = x + y  //Faz soma e atribui a resultado
		fmt.Printf("O resultado desta soma é: %v", resultado) //Imprima na tela
	case 2: //Caso seja 2 faça
		fmt.Println("Digite o primeiro número: ") //Imprima na tela
		fmt.Scan(&x)  //Leia input
		fmt.Println("Digite o segundo número: ") //Imprima na tela
		fmt.Scan(&y) //Leia input
		resultado = x - y //Faz subtração e atribui a resultado
		fmt.Printf("O resultado desta subtração é: %v", resultado)
	case 3: //Caso seja 3 faça
		fmt.Println("Ok encerrando programa....") //Imprima na tela
		os.Exit(0)  //Fecha programa com exito
	default:
		fmt.Println("Não reconheço essa opção! Tente novamente")  //Imprima na tela	
		os.Exit(-1) //Encerra programa com erro	
	}
	
	

}