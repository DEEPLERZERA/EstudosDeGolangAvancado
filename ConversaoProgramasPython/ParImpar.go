package main  //Importando pacote

import "fmt" //Importando biblioteca fmt

func main() {  //Criando função principal
	var n int  //Chamando variável
	fmt.Println("Digite um número por favor: ")  //Pedindo entrada
	fmt.Scan(&n)  //Pegando input do usuário 

	if n % 2 == 0 {  //Criando condicional
		fmt.Println("Par!")  //Imprimindo na tela
	} else {   //Se não faça
		fmt.Println("Impar!")  //Imprima na tela
	}
}