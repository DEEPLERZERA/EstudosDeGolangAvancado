package main //Chamando pacote principal

import "fmt" //Chamando biblioteca fmt



func main() {  //Criando função principal
	y := 24   //Deslocando bits e declarando variáveis
	
	x := y << 2

	z := y >> 2


	fmt.Printf("%b\n", y)  //Imprimindo na tela
	fmt.Printf("%b\n", x)  //Imprimindo na tela
	fmt.Printf("%b\n", z)  //Imprimindo na tela
}