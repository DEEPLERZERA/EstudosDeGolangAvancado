package main //Pacote principal

import ( //Importando bibliotecas
	"fmt"
	"time"
)

func GerarFibonacci(n int) func() { //Criando função de GerarFibonacci que retorna uma função
	return func() { //Retorna função Fibonnaci
		a, b := 0, 1

		fib := func() int {
			a, b = b, a+b

			return a
		}

		for i := 0; i < n; i++ {
			fmt.Printf("%d ", fib())
		}
	}
}

func Cronometrar(funcao func()) { //Criando função de cronometrar
	inicio := time.Now() //Declarando tempo de agora

	funcao() //Uma funcao aleatoria

	fmt.Printf("\nTempo de execução: %s\n\n", time.Since(inicio)) //Declara o tempo de execução
}

func main() { //Chamando função principal
	Cronometrar(GerarFibonacci(8)) //Cronometra de acordo com fibonacci passado
	Cronometrar(GerarFibonacci(48))
	Cronometrar(GerarFibonacci(88))
}
