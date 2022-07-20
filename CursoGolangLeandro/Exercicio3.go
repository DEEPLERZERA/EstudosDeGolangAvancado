package main //Chamando pacote principal

import "fmt" //Importando bilbioteca fmt

func main() { //Chamando função principal

	x := 42 //Criando variáveis

	y := "James Bond"

	z := true

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) //Concatenando tudo em uma variável só

	fmt.Println(s) //Imprimindo está variável na tela

}
