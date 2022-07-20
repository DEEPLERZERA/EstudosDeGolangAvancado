package main     //Chamando pacote principal

//Importando bibliotecas
import "fmt"  

//Definindo tipo da função
type Agregadora func(n, m int) int

func Agregar(   //Criando função que tem como um de seus parâmetros uma outra função
	valores []int,
	inicial int,
	fn Agregadora,) int {
		agregado := inicial 
		
		for _, v := range valores {
			agregado = fn(v, agregado)
		}

		return agregado
	}

//Criando função de calcularsoma com base numa outra função que faz esse tipo de calculo	
func CalcularSoma(valores []int) int {
	soma := func(n, m int) int {
		return n + m
	}

	return Agregar(valores, 0, soma)
}	

//Criando função de calcularproduto com base numa outra função que faz esse tipo de calculo	
func CalcularProduto(valores []int) int {
	multiplicacao := func(n, m int) int {
		return n * m
	}

	return Agregar(valores, 1, multiplicacao)
}	


//Chamando função principal
func main() {
	valores := []int{3, -2, 5, 7, 8, 22, 32, -1} 
	//Imprimindo resultados na tela com base nas funções especificadas
	fmt.Println(CalcularSoma(valores))   
	fmt.Println(CalcularProduto(valores))
}