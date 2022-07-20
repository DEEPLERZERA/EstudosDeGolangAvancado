package main //Chamando pacote principal

import "fmt"  //Importando biblioteca fmt



func main() {  //Chamando função main

	x := 30   //Declando variável x
	fmt.Printf("%v\n%#x\n%b\n", x, x, x)  //Imprimindo na tela
	y := x << 2  //Descolando bits
	fmt.Printf("%v\n%#x\n%b", y, y, y)  //Imprimindo na tela novamente


}