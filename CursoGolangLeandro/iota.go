package main //Chamando pacote principal

import "fmt" //Chamando biblioteca fmt

const (  //Declarando variáveis constantes com iota

	x = iota
	y = iota
	z = iota



)

func main() {  //Criando função principal
	fmt.Println(x, y, z)  //Imprimindo na tela
}