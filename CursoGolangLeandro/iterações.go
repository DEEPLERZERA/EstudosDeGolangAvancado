package main //Importando pacote principal

//IMPORTANDO BIBLIOTECAS
import "fmt"

//CHAMANDO FUNÇÃO MAIN
func main() {
	for x := 0; x < 10; x++ { //Criando laço de repetição
		fmt.Println(x)
	}

	//LAÇO DE REPETIÇÃO HIERARQUICO
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for x := 0; x < 60; x++ {
			fmt.Print(" ", x)
		}
		fmt.Println("\n")
	}

	//LAÇO DE REPETIÇÃO COM CONTINUE
	for i := 0; i < 20; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//TABLE ASCII
	for z := 33; z <= 122; z++ {
		fmt.Printf("%d - %v\n", z, string(z))

	}
	x := 9
	if !(x%2 == 0) && x%3 == 0 {
		fmt.Println("é múltiplo de dois e também de três")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("É múltiplo de dois ou de três")
	}

}
