package main //Chamando pacote main

import "fmt" //Importando biblioteca fmt

func main() { //Chamando função main
	var i int //Criando variável

loop: //Declarando variável de loop
	for i = 0; i < 10; i++ { //Chamando loop que se repete 10 vezes
		fmt.Printf("for i = %d\n", i) //imprime na tela a cada iteração
		switch i {                    //Criando Switch-case
		case 2, 3: //Caso seja 2 ou 3 faça
			fmt.Printf("Quebrando switch, i == %d .\n", i) //Imprima na tela
			break                                          //Quebre switch
		case 5: //Caso seja 5 faça
			fmt.Println("Quebrando loop, i == 5.") //Imprima na tela
			break loop                             //Quebre loop geral
		}

	}
	fmt.Println("Fim.") //Imprima na tela
}
