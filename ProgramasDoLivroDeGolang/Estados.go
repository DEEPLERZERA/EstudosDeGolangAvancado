package main //Chamando pacote principal

import "fmt" //Importando biblioteca fmt

type Estado struct { //Criando uma estrutura de dados Estado
	nome      string //Atribuindo valores
	populacao int
	capital   string
}

func main() { //Chamando função principal
	estados := make(map[string]Estado, 6) //Criando um map que recebe valores de string da struct Estado e atribuindo a estados

	estados["GO"] = Estado{"Goias", 6434052, "Goiânia"} //Passando os dados para o slice
	estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
	estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
	estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
	estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
	estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}

	fmt.Println(estados) //Imprimindo o slice
}
