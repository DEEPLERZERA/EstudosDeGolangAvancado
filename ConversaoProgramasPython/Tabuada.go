package main //Importando pacote main

import "fmt" //Importando biblioteca fmt

func main() { //Criando função principal
	var num int                       //Criando variável
	fmt.Println("Digite um número: ") //Imprimindo na tela
	fmt.Scan(&num)                    //Lendo input do usuário

	for x := 0; x < 11; x++ { //Faça enquanto
		var multi int = num * x                      //Criando variável e fazendo a multiplicação
		fmt.Printf("%v x %v = %v \n", num, x, multi) //Imprimindo na tela
	}

}
