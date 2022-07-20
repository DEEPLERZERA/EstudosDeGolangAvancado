package main //Chamando pacote principal

import "fmt"  //Importando biblioteca fmt



func main() {  //Chamando função main
	x := 10 //Declarando variáveis
	y := 15
	z := 35

	fmt.Printf("X: %v e é um valor %T\n", x, x)  //Imprimindo na tela valores
	fmt.Printf("Y: %#x e é um valor %T\n", y, y)
	fmt.Printf("Z: %b e é um valor %T", z, z)
}