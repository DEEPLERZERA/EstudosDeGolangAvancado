package main //Importando pacote principal

//IMPORTANDO BIBLIOTECAS
import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

//CRIANDO FUNÇÃO DE CALCULAR QUE RECEBE COMO PARÂMETRO UMA BASE EM UM CONTROLE
func calcular(base float64, controle *sync.WaitGroup) {
	defer controle.Done()
	n := 0.0

	//Faz o calculo com base na base
	for i := 0; i < 100000000; i++ {
		n += base / math.Pi * math.Sin(2)
	}
	fmt.Println(n)
}

//CHAMANDO FUNÇÃO PRINCIPAL
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //Usuflui de todos os cores da CPU na hora da execução
	inicio := time.Now()
	var controle sync.WaitGroup
	controle.Add(3)

	//CHAMANDO ROUTINES
	go calcular(9.37, &controle)
	go calcular(6.94, &controle)
	go calcular(42.57, &controle)

	controle.Wait()
	fmt.Printf("Finalizado em %s.\n", time.Since(inicio))
}
