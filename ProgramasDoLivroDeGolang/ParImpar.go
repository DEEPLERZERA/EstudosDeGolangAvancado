package main //Importando pacote principal

//IMPORTANDO BIBLIOTECAS
import "fmt"


//Criando função de separar par de impar
func separar(nums []int, i, p chan <- int, pronto chan <- bool) {
	for _, n:= range nums {  //percorrendo nums
		if  n % 2 == 0 {
			p <- n 
		} else {
			i <- n
		}
	}
	pronto <- true
}


//CHMANDO FUNÇÃO PRINCIPAL
func main() {
	i, p := make(chan int), make(chan int) //Criando canais
	pronto := make(chan bool)
	nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100} //Criando slice

	go separar(nums, i, p, pronto)  //Chamando goroutine

	var impares, pares []int 
	fim := false

	for !fim { //Verificando se já é fim se não faça
		//CHAMANDO SELECT - CASE
		select {  
		case n := <- i: 
			impares = append(impares, n)  //Atribuindo a impares e pares
		
		case n := <- p:
			pares = append(pares, n)
			
		case fim = <- pronto:	
		}
	}
	fmt.Printf("Ímpares: %v | Pares: %v \n", impares, pares)
}