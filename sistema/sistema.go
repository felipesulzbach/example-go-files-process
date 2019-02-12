package sistema

import (
	"bufio"
	"exemplo-files-process/util"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//const pathroot = "/home/felipe/go/diretorioarquivos"
//const pathroot = "C:/temp/diretorioarquivos"
const pathroot = "diretorioarquivos"
const pathOrigem = pathroot + "/origem"
const pathDestino = pathroot + "/destino"
const pathInvalido = pathroot + "/invalido"
const extencaocsv = "csv"
const extencaotxt = "txt"

// RetornarListaNomeArquivo Lista todos os arquivos do diretorio, de extencao especifica.
func RetornarListaNomeArquivo(pathdir string, extencaoarquivo string) []string {
	var files []string
	var arquivos []string

	err := filepath.Walk(pathdir, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		extencao, _ := util.ObterExtencao(file)
		if extencao == extencaoarquivo {
			arquivos = append(arquivos, file)
		}
	}

	return arquivos
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		//*files = append(*files, path)      // obtem nome de todo o caminho ate o arquivo.
		*files = append(*files, info.Name()) // obtem apenas o nome do arquivo.
		return nil
	}
}

// RetornarLinhasArquivo Retorna as linhas do arquivo.
func RetornarLinhasArquivo(pathArquivo string) []string {
	var linhas []string

	arquivo, err := os.Open(pathArquivo)
	if err != nil {
		arquivo.Close()
		log.Panic(err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		linhas = append(linhas, linha)
	}
	arquivo.Close()

	return linhas
}
