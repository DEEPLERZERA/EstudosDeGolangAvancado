package main //Importando pacote main

import "fmt" //Importando biblioteca fmt

var y = "Olá bom dia" //Declarando variável fora do escopo main para ser global

var d = 15

var k int
var j float32
var h string
var g bool

type andrey string //Criando meu próprio tipo

func main() { //Criando função principal

	x := 10 //Declarando variável x local

	fmt.Printf("x: %v, %T\n", x, y) //Imprimindo na tela
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30 //Declarando novos valores para x e z

	fmt.Printf("x: %v, %T\n", x, x) //Imprimindo na tela

	fmt.Printf("z: %v, %T\n", z, z)

	b := 19          //Declarando variável b local
	qualquercoisa(b) //Chamando função qualquercoisa e passando b como parâmetro

	fmt.Printf("k: %v, %T\n", k, k) //Imprimindo na tela os valores zeros destas variáveis
	fmt.Printf("j: %v, %T\n\n", j, j)
	fmt.Printf("h: %v, %T\n\n", h, h)
	fmt.Printf("g: %v, %T\n\n", g, g)

	n := "oi" //Criando variáveis
	m := "bom dia"

	c := fmt.Sprint(n, " ", m) //Usando Sprint para concatenar elas e atribuir a c

	fmt.Println(c) //Imprimindo na tela

	var o andrey = "Andrey é um cara legal" //Atribue valor string a variável o que é do tipo andrey

	fmt.Println(o) //Imprime na tela
	fmt.Printf("o: %v, %T\n", o, o)

	m = string(o) //Passando valor de o que é do tipo andrey para m que é do tipo string

	fmt.Println(m) //Imprime na tela
	fmt.Printf("m: %v, %T\n", m, m)
}

func qualquercoisa(x int) { //Criando função que recebe um x como parâmetro
	fmt.Println(d) //Imprimindo na tela
	fmt.Println(x)
}
