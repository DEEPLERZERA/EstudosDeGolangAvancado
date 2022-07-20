package main //Importando pacote principal

//IMPORTANDO BIBLIOTECAS PRINCIPAIS
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() { //Chamando função main
	inicio := time.Now()         //Pega o tempo de agora
	rand.Seed(inicio.UnixNano()) //Randomiza

	var controle sync.WaitGroup //Garante que todas as goroutines sejam sincronizadas

	for i := 0; i < 5; i++ {
		controle.Add(1)
		go executar(&controle) //Chamando goroutine
	}

	controle.Wait() //Espera todas as goroutines terminarem sua tarefas

	fmt.Printf("Finalizado em %s.\n", time.Since(inicio))
}

//CRIANDO FUNÇÃO DE EXECUTAR QUE TEM COMO PARÂMETRO UM WAITGROUP
func executar(controle *sync.WaitGroup) {
	defer controle.Done()

	duracao := time.Duration(1+rand.Intn(5)) * time.Second //Randomiza duração
	fmt.Printf("Dormindo por %s...\n", duracao)
	time.Sleep(duracao) //Da delay entre uma ação e outra
}
