package main  //Importando pacote principal

import (  //Chamando funções
	"fmt"
	"sort"

)


func main() {  //Chamando função principal
	quadrados := make(map[int]int, 15) //Criando um mapa e atribuindo a quadrados

	for n := 1; n <= 15; n++ {  //Criando um laço de repetição for
		quadrados[n] = n * n  //Vamos aumentando o tamanho do slice quadrados
	}
 
	numeros := make([]int, 0, len(quadrados))  //Criamos um slice e atribuimos a numeros e recebe como parametro o tamanho do quadrados

	for n:= range quadrados {  //Criando uma iteração de n em cima de quadrados
		numeros = append(numeros, n)  //Vamos adicionando n ao slice numeros
	}
	sort.Ints(numeros)  //Depois deixamos tudo em ordem

	for _, numero := range numeros {  //Por fim fazemos outra iteração agora percorrendo numeros de novo
		quadrado := quadrados[numero]  //Vamos adicionando numero ao slice quadrado
		fmt.Printf("%d^2 = %d\n", numero, quadrado)  //Imprimimos tudo na tela
	}


}