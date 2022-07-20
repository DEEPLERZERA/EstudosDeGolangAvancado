package main //pacote principal

import ( //Importando bilbiotecas

	"fmt"
	"regexp"
	"strings"
)

func main() { //Chamando função principal
	expr := regexp.MustCompile("\\b\\w") //Atribuindo a expr função MustCompile

	transformadora := func(s string) string { //Atribuindo função anonima que deixa tudo maiúsculo a variável transformadora
		return strings.ToUpper(s)
	}

	texto := "Antonio carlos jobim"    //Atribuindo valor a variável
	fmt.Println(transformadora(texto)) //Imprimindo na tela de acordo com as funções
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}
