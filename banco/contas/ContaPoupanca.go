package contas //DEFININDO PACOTE CONTAS

//IMPORTANDO BIBLIOTECAS
import "ProjetosDeGolang/banco/clientes"


//DEFININDO ESTRUTURA DA POUPANCA
type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}


//CRIANDO FUNÇÃO DE SACAR
func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo //VERIFICA SE PODE SACAR
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//CRIANDO FUNÇÃO DE DEPOSITAR
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {  //VERIFICA SE PODE DEPOSITAR
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

//CRIANDO FUNÇÃO DE OBTER SALDO
func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
