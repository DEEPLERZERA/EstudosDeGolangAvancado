package main //Chamando pacote principal

import "fmt"  //Importando biblioteca fmt

type ListaDeCompras []string  //Criando meu tipo de slice

func imprimirSlice(slice []string) {  //Criando função de imprimir slice que recebe slice como parâmetro
	fmt.Println("Slice:", slice)  //Imprime slice na tela
}

func imprimirLista(lista ListaDeCompras) {  //Criando função de imprimir lista recebendo ListaDeCompras como parâmetro
	fmt.Println("Lista de compras:", lista)  //Imprimindo lista
}

func main() {  //Criando função main
	lista := ListaDeCompras{"Alface", "Atum", "Azeite"}  //Atribuindo um ListaDeCompras a variável lista
	slice := []string{"Alface", "Atum", "Azeite"}  //Repetindo o processo normalmente sem o nosso tipo

	imprimirSlice([]string(lista))  //Imprimindo tudo na tela
	imprimirLista(ListaDeCompras(slice))
}