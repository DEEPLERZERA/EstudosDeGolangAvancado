package models

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

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	DeletaOProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}

	DeletaOProduto.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	ProdutoDoBanco, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for ProdutoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = ProdutoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar

}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
