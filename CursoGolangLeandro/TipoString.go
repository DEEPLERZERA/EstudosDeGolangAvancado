package main //Chamando pacote principal

import "fmt" //Chamando biblioteca fmt

func main() {  //Criando função principal
	s := "Hello World"  //Atribuindo string a s
	sb := []byte(s) //Criando um slice de byte e atribuindo a sb

	for _, v := range sb { //Percorrendo sb faça
		fmt.Printf("%v- %T - %#U - %#x\n", v, v, v, v)  //Imprima na tela
	}
}