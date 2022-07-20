package main //Chamando pacote principal

import ( //Importando bilbiotecas fmt, math/rand e time

	"fmt"
	"math/rand"
	"time"
)

func main() { //Chamando função principal
	rand.Seed(time.Now().UnixNano()) //Randomiza a seed
	n := 0                           //Atribue 0 a n

	for { //Cria um loop infinito
		n++ //Vai aumentando n a cada iteração

		i := rand.Intn(4200) //Atribue um valor aleatório a i a cada passo
		fmt.Println(i)       //Imprime na tela

		if i%42 == 0 { //Verifica se resto da divisão com 42 é zero
			break //Caso positivo saia do loop infinito
		}

	}
	fmt.Printf("Saída após %d iterações. \n", n) //Imprime na tela
}
