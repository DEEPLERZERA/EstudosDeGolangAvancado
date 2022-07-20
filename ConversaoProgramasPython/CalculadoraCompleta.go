package main //Chamando pacote

import (
	"fmt"
	"os"
) //Importando bibliotecas
var x, y, resultado, op int //Criando variáveis 

func main() { //Função main
	
	
	fmt.Println("Seja bem vindo a calculadora básica: ")  //Imprimindo na tela menu etc

	fmt.Println("Escolha uma opção: ")

	fmt.Println("1- Soma")
	fmt.Println("2- Subtração")
	fmt.Println("3- Multiplicação")
	fmt.Println("4- Divisão")
	fmt.Println("5- Resto da divisão")
	fmt.Println("0- Sair")

	fmt.Scan(&op)  //Lendo input do usuário

	switch op { //Chamando switch-case
	case 1:  //Caso seja 1 faça
		soma()  //Chamando função de soma
	case 2: //Caso seja 2 faça
		sub() //Chamando função de subtração
	case 3: 
		multi() //Chamando função de multiplicação
	case 4:
		div() //Chamando função de divisão
	case  5:
		resto() //Chamando função de resto da divisão
	case 0: //Caso seja 3 faça
		fmt.Println("Ok encerrando programa....") //Imprima na tela
		os.Exit(0)  //Fecha programa com exito
	default:
		fmt.Println("Não reconheço essa opção! Tente novamente")  //Imprima na tela	
		os.Exit(-1) //Encerra programa com erro	
	}
	
	

}

func soma() {  //Criando função de soma
	fmt.Println("Digite o primeiro número: ")  //Imprima na tela
	fmt.Scan(&x)  //Leia input
	fmt.Println("Digite o segundo número: ") //Imprima na tela
	fmt.Scan(&y) //Leia input
	resultado = x + y  //Faz soma e atribui a resultado
	fmt.Printf("O resultado desta soma é: %v", resultado) //Imprima na tela
}

func sub() { //Criando função de subtração
	fmt.Println("Digite o primeiro número: ") //Imprima na tela
	fmt.Scan(&x)  //Leia input
	fmt.Println("Digite o segundo número: ") //Imprima na tela
	fmt.Scan(&y) //Leia input
	resultado = x - y //Faz subtração e atribui a resultado
	fmt.Printf("O resultado desta subtração é: %v", resultado)
}

func multi() {  //Criando função de multiplicação
	fmt.Println("Digite o primeiro número: ")  //Imprima na tela
	fmt.Scan(&x)  //Leia input
	fmt.Println("Digite o segundo número: ") //Imprima na tela
	fmt.Scan(&y) //Leia input
	resultado = x * y  //Faz soma e atribui a resultado
	fmt.Printf("O resultado desta multiplicação é: %v", resultado) //Imprima na tela	
}

func div() { //Criando função de divisão
	fmt.Println("Digite o primeiro número: ")  //Imprima na tela
	fmt.Scan(&x)  //Leia input
	fmt.Println("Digite o segundo número: ") //Imprima na tela
	fmt.Scan(&y) //Leia input
	resultado = x / y  //Faz soma e atribui a resultado
	fmt.Printf("O resultado desta divisão é: %v", resultado) //Imprima na tela	
}

func resto() {  //Criando função de resto da divisão
	fmt.Println("Digite o primeiro número: ")  //Imprima na tela
	fmt.Scan(&x)  //Leia input
	fmt.Println("Digite o segundo número: ") //Imprima na tela
	fmt.Scan(&y) //Leia input
	resultado = x % y  //Faz soma e atribui a resultado
	fmt.Printf("O resultado deste resto da divisão é: %v", resultado) //Imprima na tela	
}