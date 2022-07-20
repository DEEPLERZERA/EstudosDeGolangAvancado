package main  //Importando pacote principal

import "fmt"  //Importando biblioteca fmt

type otaku int  //Declarando novo tipo de variável

var y int //Criando variável y do tipo int

var x otaku //Criando variável x do tipo otaku

 
func main() {  //Criando função principal
x = 52  //Atribuindo valor a variável

fmt.Printf("otaku: %v, %T\n", x, x)  //Imprimindo na tela valor e tipo da variável

x = 42  //Mudando valor da variável x
fmt.Println(x)  //Imprimindo na tela

y = int(x) //Atribuindo valor de int a variável y

fmt.Println(y) //Imprimindo valor y
fmt.Printf("y: %T", y)  //Imprimindo na tela tipo da variável



}