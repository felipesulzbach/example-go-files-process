package util

import (
	"io"
	"log"
	"os"
)

// MoverArquivo Move arquivo de uma pasta para outra.
func MoverArquivo(origem string, destino string) error {
	origemPosicaoMemoria, informacaoArquivo, err := obterPosicaoMemoriaOrigem(origem)
	if err != nil {
		return err
	}

	destinoPosicaoMemoria, err := obterPosicaoMemoriaDestino(destino, informacaoArquivo)
	if err != nil {
		return err
	}

	err = moverArquivoOrigemDestino(origem, destino, origemPosicaoMemoria, destinoPosicaoMemoria)
	if err != nil {
		return err
	}

	return nil
}

func obterPosicaoMemoriaOrigem(origem string) (*os.File, os.FileInfo, error) {
	origemPosicaoMemoria, err := os.Open(origem)
	if err != nil {
		log.Panic("obterPosicaoMemoriaOrigem", err)
		return nil, nil, err
	}

	informacaoArquivo, err := origemPosicaoMemoria.Stat()
	if err != nil {
		log.Panic("obterPosicaoMemoriaOrigem", err)
		origemPosicaoMemoria.Close()
		return nil, nil, err
	}
	return origemPosicaoMemoria, informacaoArquivo, nil
}

func obterPosicaoMemoriaDestino(destino string, informacaoArquivo os.FileInfo) (*os.File, error) {
	flag := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	permissoes := informacaoArquivo.Mode() & os.ModePerm
	destinoPosicaoMemoria, err := os.OpenFile(destino, flag, permissoes)
	if err != nil {
		log.Panic("obterPosicaoMemoriaDestino", err)
		return nil, err
	}
	return destinoPosicaoMemoria, nil
}

func moverArquivoOrigemDestino(origem string, destino string, origemPosicaoMemoria *os.File, destinoPosicaoMemoria *os.File) error {
	_, err := io.Copy(destinoPosicaoMemoria, origemPosicaoMemoria)
	if err != nil {
		log.Panic("moverArquivoOrigemDestino", err)
		destinoPosicaoMemoria.Close()
		os.Remove(destino)
		return err
	}
	err = destinoPosicaoMemoria.Close()
	if err != nil {
		log.Panic("moverArquivoOrigemDestino", err)
		return err
	}
	err = origemPosicaoMemoria.Close()
	if err != nil {
		log.Panic("moverArquivoOrigemDestino", err)
		return err
	}
	err = os.Remove(origem)
	if err != nil {
		log.Panic("moverArquivoOrigemDestino", err)
		return err
	}
	return nil
}

// VerificarArquivoExiste Verifica se arquivo exite no local informado.
func VerificarArquivoExiste(patharquivo string) bool {
	if _, err := os.Stat(patharquivo); !os.IsNotExist(err) {
		return true
	}
	return false
}
