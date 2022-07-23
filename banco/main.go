package main

import (
	//"ProjetosDeGolang/banco/clientes"
	"ProjetosDeGolang/banco/contas"
	"fmt"
)

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	contaDoDenis.Sacar(55)

	fmt.Println(contaDoDenis.ObterSaldo())

}
