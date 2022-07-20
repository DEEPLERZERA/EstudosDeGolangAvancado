package main  //Importando pacote principal

import (  //Importando bibliotecas

	"fmt"
	"io"
)

type LeitorDeStrings struct {} //Criando tipo struct

func (l LeitorDeStrings) Read(p []byte) (int, error) {  //Criando uma função que lê as strings com base na func Read que recebe um parâmetro de byte p
	p[0] = 'A'  //Declarando os itens do slice p
	p[1] = 'B'
	p[2] = 'C'
	p[3] = 'D'

	return len(p), nil  //Retorna tamanho do slice p e seu respectivo erro se houver
}

func lerString(r io.Reader) string { //Cria função de ler string com base no r io.Reader
	p := make([]byte, 4)  //Cria slice p
	r.Read(p)  //Lê os dados do slice
	return string(p) //Retorna a string p
}

func main() {  //Chamando função principal
	leitor := LeitorDeStrings{}  //Atribuindo struct a leitor
	fmt.Println(lerString(leitor)) //Imprimindo na tela os dados do slice
}
