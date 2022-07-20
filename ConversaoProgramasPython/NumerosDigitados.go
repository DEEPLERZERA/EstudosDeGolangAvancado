package main //Exportando pacote main

import "fmt" //Importando fmt

func main() { //Função principal
	var numero, digitados, soma int = 0, 0, 0  //Criando variáveis
	for numero != 999 {  //Criando laço de repetição for
		fmt.Println("Digite um número Onegaishimasu: ")  //Imprime na tela
		fmt.Scan(&numero)  //Recolhe input do usuário

		if numero == 999 {  //Faz verificação
			digitados += 1  //Soma 
			soma = soma + numero  //Soma de novo -__-
			numero = 999 //Fala que número é igual a 999 para sair do laço de repetição infinito
		} else {  //Se não faça
			digitados += 1  //Soma
			soma = soma + numero  //E soma de novo
		}

	}

	fmt.Printf("Você digitou %v e  a soma da %v", digitados, soma)  //Imprime na tela resultado final

}
