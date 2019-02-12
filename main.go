package main

import (
	"exemplo-files-process/sistema"
	"time"
)

// Determina o tempo em segundos, que o sistema ira esperar para executar novamente.
const delaySecound = 5

func main() {
	for {
		sistema.ProcessarArquivos()
		time.Sleep(delaySecound * time.Second)
	}
}
