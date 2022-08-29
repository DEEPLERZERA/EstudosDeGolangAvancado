package main //Criando pacote main

import "fmt" //Importando biblioteca fmt

type contaCorrente struct { //Criando tipo contaCorrente
	nome  string
	banco string
	conta string
	saldo float64
}

func main() { //Criando função principal
	p1 := contaCorrente{"João", "Banco do Brasil", "123-4", 5000.00} //Criando variável p1 e atribuindo valores
	saque(500, p1) //Chamando função saque
	deposito(1000, p1) //Chamando função deposito
	consultar_saldo(p1) //Chamando função consultar_saldo

}

func saque(valor float64, conta contaCorrente) { //Criando função saque
	if conta.saldo > 0 { //Se o saldo da conta for maior que 0
		conta.saldo -= valor //Subtraia o valor do saldo
		fmt.Printf("Saque efetuado com sucesso, agora você tem %v reais na sua conta\n", conta.saldo) //Imprima na tela
	} else { 
		fmt.Println("Saldo insuficiente") //Se não, imprima na tela
	}

}

func deposito(valor float64, conta contaCorrente) { //Criando função deposito
	conta.saldo += valor //Adicione o valor ao saldo
	fmt.Printf("Deposito efetuado com sucesso, agora você tem %v reais na sua conta\n", conta.saldo) //Imprima na tela
}

func consultar_saldo(conta contaCorrente) { //Criando função consultar_saldo
	fmt.Printf("O saldo da conta é %v\n", conta.saldo) //Imprima na tela
}


