package main //DEFININDO 

import (
	//"ProjetosDeGolang/banco/clientes"
	"ProjetosDeGolang/banco/contas"
	"fmt"
)

//DEFININDO INTERFACE DE VERIFICAR CONTA
type verificarConta interface {
	Sacar(valor float64) string //ATRIBUINDO FUNÇÃO DE SACAR
}


//DEFININDO FUNÇÃO DE PAGAR BOLETO DE ACORDO COM CONTA E VALOR DO BOLETO
func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto) //CHAMANDO FUNÇÃO DE SACAR
}



//DEFININDO FUNÇÃO MAIN
func main() {
	contaDoDenis := contas.ContaPoupanca{} //CRIANDO POUPANÇA
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDoDaniel := contas.ContaCorrente{} //CRIANDO FUNÇÃO
	contaDoDaniel.Depositar(1000000)
	PagarBoleto(&contaDoDaniel, 15000)

	fmt.Println(contaDoDaniel.ObterSaldo())

}
