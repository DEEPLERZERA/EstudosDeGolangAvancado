package contas //DEFININDO PACOTE DE CONTAS

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/banco/clientes"
)

//DEFININDO STRUCT DE CONTA CORRENTE
type ContaCorrente struct {
	Titular     clientes.Titular
	NumeroConta int
	saldo       float64
}

//CRIANDO FUNÇÃO DE SACAR
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo //VERIFICANDO SE PODE SACAR
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//CRIANDO FUNÇÃO DE DEPOSITAR
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 { //VERIFICANDO SE PODE DEPOSITAR
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

//CRIANDO FUNÇÃO DE TRANSFERIR
func (c *ContaCorrente) Tranferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 { //VERIFICANDO SE PODE TRANSFERIR
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia) //APONTA PARA CONTA DESTINO E DEPOSITA
		return true
	} else {
		return false
	}
}

//CRIANDO FUNÇÃO DE OBTER SALDO
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
