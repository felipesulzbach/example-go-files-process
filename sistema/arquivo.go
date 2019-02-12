package sistema

import (
	"exemplo-files-process/sistema/layout"
	"exemplo-files-process/util"
	"log"
	"strconv"
	"strings"
)

// ProcessarArquivos Faz a leitura dos arquivos, processa os dados e move cada arquivo para outra pasta.
func ProcessarArquivos() {
	for _, pessoa := range processarArquivosPessoa0001(pathOrigem, extencaotxt) {
		log.Println(layout.ShowLayout(pessoa))
		if util.VerificarArquivoExiste(pathOrigem + "/" + pessoa.Nomearquivo) {
			pessoa.MoverArquivo(pathOrigem, pathDestino, pathInvalido)
		}
	}

	for _, pessoa := range processarArquivosPessoa0001(pathOrigem, extencaocsv) {
		log.Println(layout.ShowLayout(pessoa))
		if util.VerificarArquivoExiste(pathOrigem + "/" + pessoa.Nomearquivo) {
			pessoa.MoverArquivo(pathOrigem, pathDestino, pathInvalido)
		}
	}

	for _, pessoa := range processarArquivosPessoa0002(pathOrigem, extencaotxt) {
		log.Println(layout.ShowLayout(pessoa))
		if util.VerificarArquivoExiste(pathOrigem + "/" + pessoa.Nomearquivo) {
			pessoa.MoverArquivo(pathOrigem, pathDestino, pathInvalido)
		}
	}

	for _, pessoa := range processarArquivosPessoa0002(pathOrigem, extencaocsv) {
		log.Println(layout.ShowLayout(pessoa))
		if util.VerificarArquivoExiste(pathOrigem + "/" + pessoa.Nomearquivo) {
			pessoa.MoverArquivo(pathOrigem, pathDestino, pathInvalido)
		}
	}
}

func processarArquivosPessoa0001(pathArquivos string, extencao string) []layout.Pessoa0001 {
	var pessoaList []layout.Pessoa0001

	arquivoList := RetornarListaNomeArquivo(pathArquivos, extencao)
	for _, arquivo := range arquivoList {
		pessoaArquivoList := obterLinhasArquivoPessoa0001(pathArquivos, arquivo)
		pessoaList = append(pessoaList, pessoaArquivoList...)
	}

	return pessoaList
}

func obterLinhasArquivoPessoa0001(pathArquivos string, arquivo string) []layout.Pessoa0001 {
	linhasArquivo := RetornarLinhasArquivo(pathArquivos + "/" + arquivo)
	var pessoa layout.Pessoa0001
	var colunas []string
	var pessoaList []layout.Pessoa0001

	for countlinha, linha := range linhasArquivo {
		if countlinha == 0 {
			continue // ignora os titulos
		}
		colunas = nil
		colunas = strings.Split(linha, ";")
		if len(colunas) == 3 {
			err := pessoa.New(arquivo, strconv.Itoa(countlinha), true, "", colunas[0], colunas[1], colunas[2])
			if err != nil {
				log.Panic(err)
			}
			pessoaList = append(pessoaList, pessoa)
		}
	}

	return pessoaList
}

func processarArquivosPessoa0002(pathArquivos string, extencao string) []layout.Pessoa0002 {
	var pessoaList []layout.Pessoa0002

	arquivoList := RetornarListaNomeArquivo(pathArquivos, extencao)
	for _, arquivo := range arquivoList {
		pessoaArquivoList := obterLinhasArquivoPessoa0002(pathArquivos, arquivo)
		pessoaList = append(pessoaList, pessoaArquivoList...)
	}

	return pessoaList
}

func obterLinhasArquivoPessoa0002(pathArquivos string, arquivo string) []layout.Pessoa0002 {
	linhasArquivo := RetornarLinhasArquivo(pathArquivos + "/" + arquivo)
	var pessoa layout.Pessoa0002
	var colunas []string
	var pessoaList []layout.Pessoa0002

	for countlinha, linha := range linhasArquivo {
		if countlinha == 0 {
			continue // ignora os titulos
		}
		colunas = nil
		colunas = strings.Split(linha, ";")
		if len(colunas) == 4 {
			err := pessoa.New(arquivo, strconv.Itoa(countlinha), true, "", colunas[0], colunas[1], colunas[2], colunas[3])
			if err != nil {
				log.Panic(err)
			}
			pessoaList = append(pessoaList, pessoa)
		}
	}

	return pessoaList
}
