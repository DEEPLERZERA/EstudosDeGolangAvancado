package main //Chamando pacote principal

import ( //Importando bibliotecas fmt, os e strconv para conversão de strings
	"fmt"
	"os"
	"strconv"
)

func main() { //Programa principal
	entrada := os.Args[1:]               //Entrada recebe argumentos passado pela linha de comando cmd
	numeros := make([]int, len(entrada)) //numeros recebem um make(array) com as entradas?

	for i, n := range entrada { //Para i, n range entrada faça
		numero, err := strconv.Atoi(n) //Pega os números e converte em inteiros
		if err != nil {                //Caso haja erro faça
			fmt.Printf("%s não é um número valido!\n", n) //Número não é valido!
			os.Exit(1)                                    //Saia com erro
		}
		numeros[i] = numero //Se não crie um array que recebe números
	}

	fmt.Println(quicksort(numeros)) //Depois imprima os números em forma crescente
}

func quicksort(numeros []int) []int { //Criando função quicksort que recebe os valores de numeros como parâmetros
	if len(numeros) <= 1 { //Se números forem menores que 1 faça
		return numeros //Retorne números até mesmo porque não tem como ordernar algo que é único
	}

	n := make([]int, len(numeros)) //N recebe uma array de números
	copy(n, numeros)               //Copie os números para variável n?

	indicePivo := len(n) / 2                    //indicePivo recebe o número de itens n dividido por 2
	pivo := n[indicePivo]                       //Pivo recebe n array de indicePivo
	n = append(n[:indicePivo], n[indicePivo+1]) //Adicionamos a n indicePivo e IndicePivo + 1

	menores, maiores := particionar(n, pivo) //Menores e maiores recebem particionar de n e pivo

	return append( //Retorne mais um incremento no array agora adicionando quicksort menores e quicksort maiores
		append(quicksort(menores), pivo), quicksort(maiores)...)

}

func particionar( //Criando função particionar
	numeros []int, //Criando array numeros e pivo
	pivo int) (menores []int, maiores []int) {
	for _, n := range numeros { //Para n recebendo range números faça
		if n <= pivo { //Se n for menor que pivo faça
			menores = append(menores, n) //Menores recebem um incremento em seu array com n
		} else { //Se não faça
			maiores = append(maiores, n) //Maiores recebem um incremento em seu array com n
		}
	}
	return menores, maiores //Retorne menores e maiores
}
