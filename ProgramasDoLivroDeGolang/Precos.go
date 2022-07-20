package main //Importando pacote principal

import "fmt" //Importando biblioteca

func main() { //Chamando função principal
	precoDolar, precoReal := PrecoFinal(34.99) //Atribuindo valores as duas variáveis com uma função de calcular preço final

	fmt.Printf("Preço final em dólar: %.2f\n"+"Preço final em reais: %.2f\n", //Imprimindo na tela
		precoDolar, precoReal)
}

func PrecoFinal(precoCusto float64) (precoDolar, precoReal float64) { //Criando função de calcular precoFinal com dois retornos
	fatorLucro := 1.33 //Atribuindo valores
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro //Calculando valores
	precoReal = precoDolar * taxaConversao

	return precoDolar, precoReal //Retornando resultados
}
