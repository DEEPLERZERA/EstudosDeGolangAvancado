package main //Exportando pacote main

import "fmt" //Importando fmt

var x, y int = 0, 1 //Declarando variáveis globais

func main() { //Chamando função principal
	var termo int                                                      //Declarando variável local
	fmt.Println("Digite quantas vezes o fibbonaci deve rodar: ")       //Imprimindo menuzinho na tela do user
	fmt.Scan(&termo)                                                   //Pegando input do usuário
	fmt.Println("---------------------------------------------------") //Barrinha :D
	fmt.Println("Sequência Fibbonaci a baixo:")
	fmt.Println("0")
	fmt.Println("1")

	for i := 0; i < termo; i++ { //Criando laço de repetição que usa como parâmetro a condição i < termo dado pelo user
		soma := x + y     //Aumenta o termo
		x = y             //Atribue y a x
		y = soma          //Atribue soma a y
		fmt.Println(soma) //Imprime soma

	}

}
