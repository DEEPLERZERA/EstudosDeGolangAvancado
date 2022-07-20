package main //Cuamando pacote principal

//IMPORTANDO BIBLIOTECAS
import "fmt"

//CRIANDO FUNÇÃO PRODUZIR  QUE RECEBE UM CANAL
func produzir(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c) //Fechando canal após termino do programa
}


//CHAMANDO FUNÇÃO PRINCIPAL
func main() {
	c := make(chan int, 3) 
	
	//CHAMANDO FUNÇÃO PRODUZIR
	go produzir(c)

	//PERCORRENDO C E ATRIBUINDO A VALOR
	for valor := range c {
		fmt.Println(valor)
	}
}