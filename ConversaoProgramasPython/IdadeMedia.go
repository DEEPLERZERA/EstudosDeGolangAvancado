package main //Importando pacote main

import (  //Importando biblioteca fmt
	"fmt"
)

var idademtotal, idadehtotal int = 0, 0  //Criando variáveis de nível global

var sexo string

var idadem int

var idadeh int

var idademediam, idademediah, idadegrupototal, idademediagrupo int

func main() {  //Criando função principal
	for x := 0; x < 10; x++ {  //Chamando laço de repetição for
		fmt.Println("Digite o seu sexo 'M' ou 'H': ")  //Imprimindo na tela 
		fmt.Scan(&sexo)  //Capturando input do usuário

		if sexo == "M" {  //Fazendo verificação
			m() //Chamando função de calcular m
		} else if sexo == "H" { //Se não faça
			h() //Chamando função de calcular h
		} else { //Se não faça
			fmt.Println("Gênero invalido!") //Imprima na tela que gênero é invalido!

		}

		idademediam = idademtotal / 10 //Fazendo calculos
		idademediah = idadehtotal / 10
		idadegrupototal = idademtotal + idadehtotal
		idademediagrupo = idadegrupototal / 10

	}
	fmt.Printf("Idade média das mulheres é: %v\n", idademediam) //Imprimindo para o usuário seus respectivos resultados
	fmt.Printf("Idade média dos homens é: %v\n", idademediah)
	fmt.Printf("Idade média do grupo: %v\n", idademediagrupo)

}

func m() { //Criando função m
	fmt.Println("Digite sua idade: ") //Imprimindo na tela
	fmt.Scan(&idadem) //Capturando input do usuário
	idademtotal = idademtotal + idadem //Fazendo soma da idademtotal
}


func h() {  //Criando função h
	fmt.Println("Digite sua idade: ") //Imprimindo na tela
	fmt.Scan(&idadeh) //Lendo input do usuário
	idadehtotal = idadehtotal + idadeh //Fazendo soma de idadehtotal
}