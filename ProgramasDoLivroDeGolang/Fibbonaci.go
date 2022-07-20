package main //Pacote principal

import "fmt" //Importando biblioteca

func main() { //Chamando função principal
	a, b := 0, 1 //Atribuindo valores as variáveis

	fib := func() int { //Criando função de fibonnaci que é um closure
		a, b = b, a+b

		return a
	}

	for i := 0; i < 8; i++ { //Criando laço de repetição
		fmt.Printf("%v ", fib()) //Imprimindo na tela de acordo com valor retornado pela função
	}
}
