package main //Chamando pacote principal

import ( //Importando bibliotecas
	"Net/http"
	"fmt"
	"os"
	"time"
)

const monitoramentos = 3 //Declarando variáveis constantes
const delay = 5

func main() { //Chamando função main
	for { //Para while true
		ExibeIntroducao() //Chamando funções
		ExibeMenu()

		comando := LeComando() //Atribuindo o comando escolhido a variável comando

		switch comando { //Criando switch-case
		case 1: //Caso seja 1 faça
			IniciarMonitoramento() //Chamando função
		case 2: //Caso seja 2 faça
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
	fmt.Println("2 - Sair do programa")

}

func LeComando() int { //Cria funcao de ler comando
	var comandoLido int    //Criando variavel de comando lido
	fmt.Scan(&comandoLido) //Recebendo input do usuário

	fmt.Println("O comando escolhido foi", comandoLido) //Imprimindo na tela

	return comandoLido //Dando retorno pois é uma função que exige retornar valores
}

var sites string //Criando uma variável sites

func IniciarMonitoramento() { //Cria função de monitoramento
	fmt.Println("Digite o seu site de escolha: ") //Dando opção de escolha ao usuário
	fmt.Scan(&sites)                              //Lendo escolha do usuário
	fmt.Println("Monitorando.......")             //Imprimindo na tela

	for i := 0; i < monitoramentos; i++ { //Quantas vezes se deve monitorar
		for i, site := range sites { //Percorrendo meu slice
			fmt.Println("Testando site", i, ":", site) //Imprimindo na tela
			testaSite(string(sites))                   //Chamando função de testar site
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

	} else { //Se não faz
		fmt.Println("Site:", site, "está com problemas. Status code:", resp.StatusCode) //Imprime na tela que deu erro e imprime o statuscode do site

	}
}
