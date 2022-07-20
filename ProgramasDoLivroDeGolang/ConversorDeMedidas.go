package main //Importando biblioteca main

import ( //Importando fmt, os e strconv para converter strings
	"fmt"
	"os"
	"strconv"
)

func main() { //Criando func main
	if len(os.Args) < 3 { //Verificando se usuário passou na hora da execução os parâmetros necessários para a execução do respectivo programa
		fmt.Println("Uso: conversor <valores> <unidade>") //Imprimindo na tela
		os.Exit(1)                                        //Saindo do programa

	}

	unidadeOrigem := os.Args[len(os.Args)-1]     //Recebendo variável unidade de origem
	valoresOrigem := os.Args[1 : len(os.Args)-1] //Recebendo valores de origem

	var unidadeDestino string //Declarando variável de unidadeDestino

	if unidadeOrigem == "Celsius" { //Fazendo verificador
		unidadeDestino = "Fahrenheit" //Atribuindo valor a variável
	} else if unidadeOrigem == "Quilometros" { //Fazendo verificador
		unidadeDestino = "Milhas" //Atribuindo valor a variável
	} else { //Fazendo verificador
		fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino) //Imprimindo na tela
		os.Exit(1)                                                    //Saindo do programa
	}

	for i, v := range valoresOrigem { //Enquanto v for um range de valores de origem faça
		valorOrigem, err := strconv.ParseFloat(v, 64) //Convertendo valores de origem para float
		if err != nil {                               //Se err for diferente de nil faça
			fmt.Printf( //Imprime na tela a mensagem de erro
				"O valor %s na posição %d não é um número valido!", v, i)
			os.Exit(1) //Saindo do programa
		}

		var valorDestino float64 //Declarando variável de valorDestino que é um float64

		if unidadeOrigem == "Celsius" { //Fazendo verificador de unidade de origem
			valorDestino = valorOrigem*1.8 + 32 //Calculando valorDestino e atribuindo a sua variável
		} else { //Verificador
			valorDestino = valorOrigem / 1.60934 //Calculando valorDestino e atribuindo a sua variável
		}
		fmt.Printf("%.2f %s = %.2f %s \n", valorOrigem, unidadeOrigem, valorDestino, unidadeDestino) //Imprimindo na tela resultado final com aproximador float de 2 casas decimais
	}

}
