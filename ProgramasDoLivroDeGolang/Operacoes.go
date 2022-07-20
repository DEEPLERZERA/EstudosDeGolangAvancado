package main //Chamando pacote principal

import (
	"fmt"  //Importando biblioteca fmt
	"time" //Importando biblioteca time
)

type Idade struct { //Criando estrutura de idade
	anoNascimento int //Declarando variável
}

func (i Idade) Calcular() int { //Criando função de calcular idade
	return time.Now().Year() - i.anoNascimento //Retorna o ano de agora
}

func (i Idade) String() string { //Função de imprimir idade
	return fmt.Sprintf("Idade desde %d", i.anoNascimento) //Imprimindo na tela
}

type Operacao interface { //Criando tipo operacao interface
	Calcular() int //Chamando função Calcular int
}

type Soma struct { //Criando tipo Soma que é uma struct
	operando1, operando2 int //Declarando valores int
}

func (s Soma) Calcular() int { //Criando função de somar
	return s.operando1 + s.operando2 //Retornando resultado
}

func (s Soma) String() string { //Criando função que imprime
	return fmt.Sprintf("%d + %d", s.operando1, s.operando2) //Imprimindo e concatenando valores
}

type Subtracao struct { // Criando tipo subtracao que é uma struct
	operando1, operando2 int //Declarando variáveis
}

func (s Subtracao) Calcular() int { //Criando função de calcular subtracao
	return s.operando1 - s.operando2 ///Retornando resultado
}

func (s Subtracao) String() string { //Imprimindo na tela
	return fmt.Sprintf("%d - %d", s.operando1, s.operando2) //Imprimindo e concatenando valores
}

func main() { //Chamando função principal
	operacoes := make([]Operacao, 4) //Criando slice e atribuindo a operações
	operacoes[0] = Soma{10, 20}      //Fazendo operações e atribuindo a slice operacoes
	operacoes[1] = Subtracao{30, 15}
	operacoes[2] = Subtracao{10, 50}
	operacoes[3] = Soma{5, 2}

	fmt.Println("Valor acumulado =", acumular(operacoes)) //Imprimindo na tela o valor acumulado com base na função acumular

	idades := make([]Operacao, 3) //Criando um slice de idades
	idades[0] = Idade{1969}       //Atribuindo anos de idade
	idades[1] = Idade{1977}
	idades[2] = Idade{2001}

	fmt.Println("Idades acumuladas =", acumular(idades)) //Imprimindo na tela o valor acumulado com base na função acumular

}

func acumular(operacoes []Operacao) int { //Criando função acumular que recebe como parâmetros um slice Operacao de valor int
	acumulador := 0 //Declara variável

	for _, op := range operacoes { //Op percorrendo operacoes faça
		valor := op.Calcular()             //atribuindo a valor a função Calcular
		fmt.Printf("%v = %d\n", op, valor) //Imprimindo na tela os valores
		acumulador += valor                //Aumentando acumulador
	}

	return acumulador //Retornando acumulador
}
