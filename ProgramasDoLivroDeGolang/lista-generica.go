package main  //Importando pacote principal

import "fmt"  //Importando biblioteca fmt

type ListaGenerica []interface {}  //Criando listaGenerica do tipo interface

func (lista *ListaGenerica) RemoverIndice(  //Criando função de removerIndice que é uma interface
	indice int) interface {} {

		l := *lista //Criando variáveis
		removido := l[indice]
		*lista = append(l[0:indice], l[indice+1:]...)  //Aumentando tamanho do slice
		return removido  //Retornando dado removido


	}

	func (lista * ListaGenerica) RemoverInicio() interface{} {  //Criando função de remover inicio
		return lista.RemoverIndice(0)
	}

	func (lista * ListaGenerica) RemoverFim() interface {} { //Criando função de remover Fim 
			return lista.RemoverIndice(len(*lista)-1) 
	}

	func main() {  //Chamando função principal
		lista := ListaGenerica{  //Atribuindo slice a variável lista
			1, "Café", 42, true, 23, "Bola", 3.14, false,
		}

		fmt.Printf(                                                                                   //Imprimindo na tela os dados
			"Removendo do início: %v, após remoção:\n%v\n", lista.RemoverInicio(),lista)

		fmt.Printf(
				"Removendo do fim: %v, após remoção:\n%v\n", lista.RemoverFim(),lista)
		fmt.Printf(
				"Removendo do índice 3: %v, após remoção:\n%v\n", lista.RemoverIndice(3),lista)	

		fmt.Printf(
				"Removendo do índice 0: %v, após remoção:\n%v\n", lista.RemoverIndice(0),lista)
		fmt.Printf(
				"Removendo do último índice: %v, após remoção:\n%v\n", lista.RemoverIndice(len(lista)-1), lista)		


	}