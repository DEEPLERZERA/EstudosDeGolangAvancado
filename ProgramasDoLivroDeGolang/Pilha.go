package main //Criando pacote principal 

import ( //Importando bibliotecas de erros e fmt
	"errors"
	"fmt"
)

func main() { //Criando função principal
	pilha := Pilha{}  //Criando uma array pilha
	fmt.Println("Pilha criada com tamanho ", pilha.Tamanho())  //Imprimindo na tela o tamanho da pilha
	fmt.Println("Vazia? ", pilha.Vazia())  //Retornando se pilha está vazia ou não

	pilha.Empilhar("Go")  //Empilhando dados
	pilha.Empilhar(2009)
	pilha.Empilhar(3.14)
	pilha.Empilhar("Fim")
	fmt.Println("Tamanho após empilhar 4 valores: ", //Imprimindo na tela
		pilha.Tamanho())
	fmt.Println("Vazia?", pilha.Vazia())  //Imprimindo se pilha é vazia

	for !pilha.Vazia() {  //Se pilha for diferente de vazia faça
		v, _ := pilha.Desempilhar()  //Va desempilhando
		fmt.Println("Desempilhado ", v) //Imprime array desempilhado
		fmt.Println("Tamanho: ", pilha.Tamanho())  //Imprime tamanho da pilha
		fmt.Println("Vazia? ", pilha.Vazia()) //Imprime se está vazia
	}

	_, err := pilha.Desempilhar() //Atribue desempilhar a err
	if err != nil {  //Se err for diferente de nulo faça
		fmt.Println("Erro: ", err)  //Imprime erro
	}


}

type Pilha struct {  //Criando Pilha do tipo struct
	valores []interface{}  //Recebe valores interface
}

func (pilha Pilha) Tamanho() int { // Criando Função pilha tamanho
	return len(pilha.valores) //Retorna o tamanho da pilha
}

func (pilha Pilha) Vazia() bool {  //Criando função pilha vazia
	return pilha.Tamanho() == 0  //Retorna se está vazia com um bool
}

func (pilha *Pilha) Empilhar(valor interface{}) {  //Criando função de empilhar
	pilha.valores = append(pilha.valores, valor)  //Vai adicionando itens ao slice com append
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {  //Função desempilhar
	if pilha.Vazia() {  //Se pilha está vazia
		return nil, errors.New("Pilha vazia!") //Retorne um erro de pilha vazia
	}
	valor := pilha.valores[pilha.Tamanho()-1]  //Atribue os valores da pilha a variável valor
	pilha.valores = pilha.valores[:pilha.Tamanho()-1] //E vai tirando os elementos
	return valor, nil  //No final retorne o valor e nil
}