package main //Importando pacote principal

import ( //Importando bibliotecas

	"fmt"
	"os"
)

func CriarArquivos(dirBase string, arquivos ...string) { //Criando função que cria arquivos usando como parâmetros a base e o nome do arquivo
	for _, nome := range arquivos { //Percorrendo arquivos
		caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt") //Concatena os dados, criando um caminho para o arquivo

		arq, err := os.Create(caminhoArquivo) //Cria arquivo

		defer arq.Close() //O fecha

		if err != nil { //Verifica se há erro
			fmt.Print("Erro ao criar arquivo %s: %v\n",
				nome, err)
			os.Exit(1)
		}

		fmt.Printf("Arquivo %s criado.\n", arq.Name()) //Demonstra o arquivo que fora criado

	}
}

func main() { //Chamando função principal
	tmp := os.TempDir() //Passando diretorio

	CriarArquivos(tmp) //Chamando funções
	CriarArquivos(tmp, "teste1")
	CriarArquivos(tmp, "teste2", "teste3", "teste4")
}
