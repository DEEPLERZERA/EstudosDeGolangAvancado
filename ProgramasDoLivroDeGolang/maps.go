package main //Importando pacote principal

import ( //Importando bibliotecas fmt, os e strings

	"fmt"
	"os"
	"strings"
)

func main() { //Criando função principal
	palavras := os.Args[1:] //Variável palavras recebe input do usuário dado no cmd

	estatisticas := colherEstatisticas(palavras) //Variável estatisticas recebe função colherEstatisticas com parãmetro palavras

	imprimir(estatisticas) //Chamando função imprimir com estatisticas como parâmetro
}

func colherEstatisticas(palavras []string) map[string]int { //Cria função de colher estatisticas que recebe com parâmetro a própria string e faz um map
	estatisticas := make(map[string]int) //Atribuindo a estatistiscas um map string

	for _, palavra := range palavras { //Para palavra dentro do range de palavras faça
		inicial := strings.ToUpper(string(palavra[0])) //Inicial recebe as strings convertidas para maiusculas
		contador, encontrado := estatisticas[inicial]  //Contador e encontrado recebem dados da estatistica inicial

		if encontrado { //Se valor é encontrado
			estatisticas[inicial] = contador + 1 //Aumenta o contador de estatisticas em 1
		} else { //Se não faça
			estatisticas[inicial] = 1 //Estatisticas é igual a 1
		}
	}
	return estatisticas //Retorna valor de estatisticas
}

func imprimir(estatisticas map[string]int) { //Criando função imprimirque recebe como parâmetros estatisticas map string
	fmt.Println("Contagem de palavras iniciadas em cada letra:") //Imprime na tela

	for inicial, contador := range estatisticas { //Para inicial e contador range estatisticas faça
		fmt.Printf("%s = %d\n", inicial, contador) //Imprima na tela a letra e o número de vezes que ela se repetiu
	}
}
