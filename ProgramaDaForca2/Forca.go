package main  //Importando pacote principal

import (  //Importando bibliotecas
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strings"
	"io"

)
func main() {  //Função principal
	tempoAgora := time.Now()  //Atribue tempo de agora
	rand.Seed(tempoAgora.UnixNano())  //Gera um seed aleatória a cada nano segundo
	palavraEscolhida := escolhePalavra()  //Atribue palavra escolhida para a variável
	chances := len(palavraEscolhida)  //Pegando tamanho da palavra e atribuindo a nGuesses
	encontrado := []string{} //Criando slice de encontrado

	for i := 0; i < len(palavraEscolhida); i++ {  //Criando laço de repetição
		encontrado = append(encontrado, "_") //Vai adicionando _ _ _ _ a tela do usuário
	}
	for chances > 0 {  //Para tamanho da palavra maior que zero
		fmt.Println("Você tem", chances, "Chances sobrando!")  //Mostra número de chances sobrando
		letra, err := pegaLetra(encontrado)  //Chamando função de pegar letra e atribuindo a letra
		if err != nil {  //Se for nulo faça
			fmt.Println("Erro de leitura no console!")  //Imprime erro na tela
			return
		}
		if !contem(palavraEscolhida, []string{letra}) {  //Se palavraEscolhida não estiver contida em string letras faça
			chances--  //Diminue nGuesses
		}
		if atualizandoEncontrado(encontrado, palavraEscolhida, letra) {  //Se encontrar palavra escolhida
			fmt.Println("Você ganhou ganhou! a palavra era:", palavraEscolhida)  //Imprime na tela
			return  //Da retorno
		}
	}
	fmt.Println("Você perdeu! a palavra era:", palavraEscolhida)  //Imprime na tela
}

func pegaLetra(encontrou []string) (string, error) {  //Criando função de pegar letra
	ALFABETO := "aáàãbcçdeéêfghiíîjklmnoôõópqrstuúvwxyz"  //Atribuindo alfabeto

	for true {  //Laço de retição infinito
		letra, err := leInput("Escolha uma letra:", entrada(encontrou, " "))  //Recebe input do usuário e atribue a letra
		if err != nil {  //Se erro for diferente de nulo faça
			return "", err  //Retorne erro
		}
		if len(letra) == 1 && contem(ALFABETO, []string{letra}) {  //Condicional que verifica se está tudo certo com a letra entregue pelo usuário
			return letra, nil  //Retorna letra
		}
		fmt.Println("Input invalido: Tente com uma única letra minúsucula!")  //Se não imprime input invalido

	}
	return "", nil  //Retorna nulo
}


func atualizandoEncontrado(encontrado []string, palavra string, letra string) bool {  //Criando função de atualizar encontrado
	completo := true   //Completo passa a ser true
	for i, r := range palavra {  //Percorrendo palavra com r faça
		if letra == string(r) {  //Se letra for igual a r faça
			encontrado[i] = letra  //Atribue letra a encontrado
		}
		if encontrado[i] == "_"  {  //Se encontrado for igual a _ faça
			completo = false  //Completo ser false
		}

	}
	return completo //Retorne completo
}


func abreArquivo() []string {   //Criando função de abrir arquivo
	var palavras []string //Chamando var palavras

	arquivo, err := os.Open("palavras.txt") //Abrindo arquivo de palavras e atribuindo a arquivo

	if err != nil { //Se houver erro faça
		fmt.Println("Ocorreu um erro:", err) //Imprime na tela mensagem de erro
	}

	leitor := bufio.NewReader(arquivo) //Lê linha a linha do arquivo de texto e atribue a variável leitor

	for { //Criando laço for de repetição
		linha, err := leitor.ReadString('\n') //Fazendo leitura individual de cada linha
		linha = strings.TrimSpace(linha)      //Tirando os \n
		palavras = append(palavras, linha)          //Atribuindo strings a slice

		if err == io.EOF { //Se for fim de arquivo
			break //Saia do loop for
		}
	}

	arquivo.Close() //Feche o arquivo aberto a cima(Questão de boa prática)

	return palavras //Retornando palavras
}


func escolhePalavra() string{  //Criando função de escolher palavras
	conteudo := abreArquivo()  //Atribue retorno da função abre_arquivo() a conteudo
	indexAleatorio := rand.Intn(len(conteudo))  //Randomizando o conteudo
	escolha := conteudo[indexAleatorio]  //Fazendo um escolha2 aleatorio 
	return escolha   //Retornando escolha2
}

func leInput(vals ...interface{}) (string, error) {  //Criando função de leInput
	if len(vals) != 0 {  //Se tamanho de vals for diferente de zero faça
		fmt.Println(vals...)  //imprima vals
	}
	leitura := bufio.NewScanner(os.Stdin)  //leitura recebe bufio.NewScanner
	leitura.Scan()  //Vai lendo input
	err := leitura.Err()   //Se tiver erro atribue a err
	if err != nil {  //Se tiver erro
		return "", err  //Imprime erro
	}
	
	return leitura.Text(), nil  //Retorna o texto scanneado e nil se tiver
}


func contem(valor1 string, caracteres []string) bool {  //Criando função pra verificar se contem
	for _, ca := range valor1 {  //Percorrendo valor1 com ca
		for _, ca2 := range caracteres {  //Percorrendo caracteres que foi recebido como parametro por ca2
			if string(ca) == ca2 {  //Fazendo uma comparação
				return true  //Retornando verdadeiro
			}
		}
	}
	return false  //Retornando falso
}


func entrada(strings []string, separador string) string {  //Função que recebe entrada
	if len(strings) == 0 {  //Verifica se tamanho das strings é igual a 0
		return ""  //Retorna string vazia
	}
	s := ""   //Atribue vazio a s
	ultimoIdx := len(strings) - 1 //Verifica qual é o último index
	for _, v := range strings[0:ultimoIdx] {  //Percorre o slice strings com v
		s += v + separador  //Vai concatenando
	}
	return s + strings[ultimoIdx]  //Retorne a string
}