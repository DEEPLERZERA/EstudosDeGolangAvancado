package main //Importando pacote principal

import "fmt"  //Importando biblioteca fmt

var fat int = 1 //Criando variáveis do tipo global
var n int


func main() { //Chamando função principal
	fmt.Println("Digite um número para eu calcular o fatorial: ")  //Imprimindo na tela
	fmt.Scan(&n)  //Capturando input do usuário
	var y int = 0  //Declarando variável y
	for x := 0; x < n; x ++ { //Criando laço de repetição for
		y += 1 //Fazendo incremento na variável y
		fat *= y //E aumentando meu fat
	}

	fmt.Printf("O fatorial de: %v é %v", y, fat)  //Imprimindo fatorial para o usuário

}