package models  //Definindo pacote models

//IMPORTANDO BIBLIOTECAS
import (
	"ProjetosDeGolang/aplicacaoweb/db"
)

//DEFININDO STRUCT
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//CRIANDO FUNÇÃO DE BUSCAR TODOS OS PRODUTOS
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados() //Gera conexão com db

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc") //Faz pesquisa no banco de dados

	if err != nil { //Verifica se há erro
		panic(err.Error())
	}

	p := Produto{} //Cria slice
	produtos := []Produto{}

	//AGORA IREMOS PERCORRER OS DADOS
	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade) //Scanneia os campos

		if err != nil { //Verifica se há erro
			panic(err.Error())
		}

		//ATRIBUINDO DADOS AOS CAMPOS
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p) //ADICIONA TODOS OS DADOS AO SLICE PRODUTOS
	}
	defer db.Close() //Fecha BD
	return produtos
}

//FUNÇÃO DE CRIAR PRODUTOS
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)") //Prepara para inserir dados no BD
	if err != nil { //Verifica se há erro
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade) //Se não executa ordem e insere os dados
	defer db.Close() //Fecha o banco de dados
}

//FUNÇÃO DE DELETAR PRODUTO
func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()  //Define conexão com banco de dados

	DeletaOProduto, err := db.Prepare("delete from produtos where id = $1") //Prepara para deletar dados do banco de dados
	if err != nil {  //Verifica se há erro
		panic(err.Error())
	}

	DeletaOProduto.Exec(id)  //Executa ordem de deletar produto
	defer db.Close()  //Fecha banco de dados
}

//FUNÇÃO DE EDITAR PRODUTO QUE RETORNA PRODUTO
func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados

	ProdutoDoBanco, err := db.Query("select * from produtos where id = $1", id)  //Faz uma busca na base de dados de acordo com id

	if err != nil { //Verifica se há erro
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}  //Atribue slice para produtoParaAtualizar

	for ProdutoDoBanco.Next() {  //percorre os dados do produto
		var id, quantidade int  
		var nome, descricao string
		var preco float64

		err = ProdutoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)  //Scanneia todos um por um

		if err != nil { //Verifica se há erro
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id  //Atribue os dados para produtoParaAtualizar que contém um slice em sua memória
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()  //Fecha o banco de dados
	return produtoParaAtualizar

}


//FUNÇÃO DE ATUALIZAR PRODUTO
func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados() //Define conexão com banco de dados

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5") //Prepara para dar update no BD
	if err != nil {  //Verifica se há erro
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)  //Se não executa ordem e atualiza dados do produto
	defer db.Close()  //Fecha o banco de dados
}
