package main //Chamando pacote principal

import "fmt"  //Importando biblioteca fmt



func main() {  //Chamando função main
	x := (5 != 5)  //Atribuindo valores booleanos as variáveis
	y := (10 <= 5)
	z := (15 > 10)
	b := (19 == 15)
	c := (9 > 12)
	d := (15 < 20)

	fmt.Println(x, y, z, b, c, d)  //Imprimindo na tela
}