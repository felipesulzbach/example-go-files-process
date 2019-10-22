package main

import (
	"time"

	"github.com/_dev/exemplo-files-process/sistema"
)

// Determina o tempo em segundos, que o sistema ira esperar para executar novamente.
const delaySecound = 5

func main() {
	for {
		sistema.ProcessarArquivos()
		time.Sleep(delaySecound * time.Second)
	}
}
