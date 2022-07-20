package main //Chamando pacote principal

import "fmt" //Importando biblioteca fmt
 
var notamedia float32  //Declarando variáveis
var notasomatotal float32 = 0
var nota float32 

func main () {  //Criando função principal
	fmt.Println("Seja bem vindo ao programa de calcular sua média escolar!")  //Imprimindo na tela
	fmt.Println("Digite a baixo suas notas para obter uma média:\n")

	CalculaMedia() //Chamando função de calcular Média
}



func CalculaMedia() { //Criando função de calcular media
	for x := 0; x < 5; x++ { //Criando laço de repetição de 5
		fmt.Println("Digite sua nota: ") //Imprimindo na tela
		fmt.Scan(&nota) //Pegando input do usuário
		notasomatotal = notasomatotal + nota //Fazendo soma total das notas
	}
	notamedia = notasomatotal / 5 //Calculando a nota media

	if notamedia > 6 { //Fazendo condicional
		fmt.Println("Passou meu querido aluno!") //Imprimindo na tela resultado
	} else { //Se não faça
		fmt.Println("Reprovou meu querido aluno!")  //Imprime na tela resultado
	}
	
	fmt.Printf("Sua média é: %.2f", notamedia)  //Imprime na tela media final
}