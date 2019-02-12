package util

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// ObterExtencao Retorna a extensao do arquivo.
func ObterExtencao(descArquivo string) (string, error) {
	index := strings.Index(descArquivo, ".")
	if index == -1 {
		return "", errors.New("Arquivo sem extensao ou nao informado")
	}
	extencao := descArquivo[index+1:]
	return extencao, nil
}

// ValidarString Valida se valor da string foi informado.
func ValidarString(valor string, nomeColuna string) error {
	if valor == "" {
		return errors.New(obterMsgErroNaoInformado(nomeColuna))
	}
	return nil
}

// ValidarStringInt64 Valida se valor da string foi informado e se o valor e do tipo int64.
func ValidarStringInt64(valor string, nomeColuna string) error {
	if valor == "" {
		return errors.New(obterMsgErroNaoInformado(nomeColuna))
	} else if _, err := strconv.ParseInt(valor, 10, 64); err != nil {
		return errors.New(obterMsgErroInvalido(nomeColuna))
	}
	return nil
}

// ValidarStringInt Valida se valor da string foi informado e se o valor e do tipo int.
func ValidarStringInt(valor string, nomeColuna string) error {
	if valor == "" {
		return errors.New(obterMsgErroNaoInformado(nomeColuna))
	} else if _, err := strconv.Atoi(valor); err != nil {
		return errors.New(obterMsgErroInvalido(nomeColuna))
	}
	return nil
}

func obterMsgErroNaoInformado(nomeColuna string) string {
	var buffer bytes.Buffer
	buffer.WriteString("Valor coluna ")
	buffer.WriteString(nomeColuna)
	buffer.WriteString(" nao informado")
	return buffer.String()
}

func obterMsgErroInvalido(nomeColuna string) string {
	var buffer bytes.Buffer
	buffer.WriteString("Valor coluna ")
	buffer.WriteString(nomeColuna)
	buffer.WriteString(" invalido")
	return buffer.String()
}
