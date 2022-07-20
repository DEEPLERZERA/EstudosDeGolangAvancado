package main //Chamando pacote principal

//Importando bibliotecas
import (
	"fmt"
	"time"
)

//Criando função de goroutine 5 segundos de espera
func dormir() {
	fmt.Println("Goroutine dormindo por 5 segundos...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine finalizada.")
}


//Chamando função principal
func main() {
	go dormir() //Chamando função dormir com goroutine
	fmt.Println("Main dormindo por 3 segundos....")
	time.Sleep(3 * time.Second)
	fmt.Println("Main finalizada")
}