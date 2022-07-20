package main //Cria biblioteca

import (
	"Net/http" //Importa biblioteca Net/http
	"bufio"    //Importando  biblioteca bufio
	"fmt"      //Importa biblioteca fmt
	"io"
	"io/ioutil"
	"os"      //Importa biblioteca de os
	"strconv" //Importando biblioteca strconv
	"strings"
	"time" //Importando  biblioteca time
)

const monitoramentos = 3
const delay = 5

func main() { //Criando função principal
	for { //Faz com que se repita infinitamente
		ExibeIntroducao() //Chamando função
		ExibeMenu()       //Chamando função

		comando := LeComando() //Atribuindo função a comando

		/*if comando == 1 {    //Criando condicionais
			fmt.Println("Monitorando.......") //Imprimindo na tela
		} else if comando == 2 {  //Mais condicionais
			fmt.Println("Exibindo logs.....")  //Imprime na tela
		}else if comando == 0 {  //Condicional
			fmt.Println("Saindo do programa....")  //Imprime na tela
		}else {  //Se não faça
			fmt.Println("Não reconheço esse comando.....")  //Imprima na tela
		}
		*/

		switch comando { //Criando switch-case
		case 1: //Caso seja 1 faça
			IniciarMonitoramento() //Chamando função
		case 2: //Caso seja 2 faça
			fmt.Println("Exibindo logs.....") //Imprime na tela
			imprimeLogs() //Chamando função de imprimir logs
		case 0: //Caso seja 0 faça
			fmt.Println("Saindo do programa....") //Imprime na tela
			os.Exit(0)                            //Sai do programa sem erro
		default: //Se não for nenhuma das opções a cima faça isto
			fmt.Println("Não reconheço esse comando.....") //Imprima na tela
			os.Exit(-1)                                    //Sai do programa com erro
		}
	}

}

func ExibeIntroducao() { //Cria função de exibir introducao
	nome := "Daniel" //Criando variáveis
	versao := 1.1
	fmt.Println("Hello World!") //Imprimindo na tela
	fmt.Println("Olá Senhor ", nome)
	fmt.Println("Estamos na versão: ", versao)
}

func ExibeMenu() { //Cria funcao de exibir menu
	fmt.Println("1 - Iniciar ") //Imprimindo menu na tela
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

}

func LeComando() int { //Cria funcao de ler comando
	var comandoLido int    //Criando variavel de comando lido
	fmt.Scan(&comandoLido) //Recebendo input do usuário

	fmt.Println("O comando escolhido foi", comandoLido) //Imprimindo na tela

	return comandoLido //Dando retorno pois é uma função que exige retornar valores
}

func IniciarMonitoramento() { //Cria função de monitoramento
	fmt.Println("Monitorando.......") //Imprimindo na tela

	sites := leSiteDoArquivo()            //Chamando função de ler site do arquivo e atribuindo a variável sites
	for i := 0; i < monitoramentos; i++ { //Quantas vezes se deve monitorar
		for i, site := range sites { //Percorrendo meu slice
			fmt.Println("Testando site", i, ":", site) //Imprimindo na tela
			testaSite(site)                            //Chamando função de testar site
		}
		time.Sleep(delay * time.Second) //Definindo tempo entre cada ação do for
		fmt.Println("")
	}

	fmt.Println("") //Imprimindo na tela espaços
}

func testaSite(site string) { //Criando função de testar site
	resp, err := http.Get(site) //Pegando resposta http do site e atribuindo a variável

	if err != nil { //Se houver erro faça
		fmt.Println("Ocorreu um erro:", err) //Imprime na tela mensagem de erro
	}

	if resp.StatusCode == 200 { //Verifica se acesso ao site teve exito
		fmt.Println("Site:", site, "foi carregado com sucesso!") //Imprime na tela que sim
		registraLog(site, true)                                  //Chamando função como true
	} else { //Se não faz
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode) //Imprime na tela que deu erro e imprime o statuscode do site
		registraLog(site, false)                                                        //Chamando função como false
	}
}

func leSiteDoArquivo() []string { //Criando função leSiteDoArquivo
	var sites []string //Chamando var sites

	arquivo, err := os.Open("sites.txt") //Abrindo arquivo de sites e atribuindo a arquivo

	if err != nil { //Se houver erro faça
		fmt.Println("Ocorreu um erro:", err) //Imprime na tela mensagem de erro
	}

	leitor := bufio.NewReader(arquivo) //Lê linha a linha do arquivo de texto e atribue a variável leitor

	for { //Criando laço for de repetição
		linha, err := leitor.ReadString('\n') //Fazendo leitura individual de cada linha
		linha = strings.TrimSpace(linha)      //Tirando os \n
		sites = append(sites, linha)          //Atribuindo strings a slice

		if err == io.EOF { //Se for fim de arquivo
			break //Saia do loop for
		}
	}

	arquivo.Close() //Feche o arquivo aberto a cima(Questão de boa prática)

	return sites //Retornando sites
}

func registraLog(site string, status bool) { //Criando função de registrar log que recebe como parâmetro site e status
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //Arquivo recebe um open file de log.txt que é criado na hora se necessário

	if err != nil { //Se obtiver erro faça
		fmt.Println("Ocorreu um erro:", err) //Imprima o tipo de erro que aconteceu
	}

	fmt.Println(arquivo) //Imprima a log

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n") //Escreve os dados no arquivo log

	arquivo.Close() //Fechando arquivo
}

func imprimeLogs() {  //Criando função de imprimir logs

	arquivo, err := ioutil.ReadFile("log.txt")  //Lê os dados de log.txt e passa para arquivo

	if err != nil {  //Se for diferente de nulo faça
		fmt.Println("Ocorreu um erro:", err)  //Imprime qual erro ocorreu
	}

	fmt.Println(string(arquivo))  //Imprime a log na tela
}
