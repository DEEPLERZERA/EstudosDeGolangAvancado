package main//Importando pacote principal

//IMPORTANDO BIBLIOTECAS
import (
	"fmt"
	"time"

)

//CRIANDO FUNÇÃO DE EXECUTAR QUE TEM UM TEMPO DE ESPERA DE 5 SEGUNDOS
func executar(c chan <- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}

//CHAMANDO FUNÇÃO PRINCIPAL
func main() {
	c := make(chan bool, 1) //Criando um canal
	go executar(c)  //Chamando uma goroutine

	fmt.Println("Esperando...")

	fim := false
	for !fim {  //Enquanto não for fim faça
		select {  //Select-case
		case fim = <- c:
			fmt.Println("Fim!")

		case <-  time.After(2 * time.Second):
			fmt.Println("Timeout!")
			fim = true	
		}
	}
}