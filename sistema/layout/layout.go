package layout

import (
	"bytes"
	"exemplo-files-process/util"
)

const separatorFront = ": '"
const separatorBehind = "', "
const separatorBehind2 = "'"

// Layout Informacoes padrao layout.
type Layout struct {
	Nomearquivo   string
	Linhaarquivo  string
	Flgprocessado bool
	Msgerro       string
}

type layout interface {
	show() string
}

// ShowLayout Retorna a estrutura/valores do layout.
func ShowLayout(l layout) string {
	return l.show()
}

func toString(entidade string, campos map[string]string) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	buffer.WriteString(entidade)
	buffer.WriteString(" = {")

	qtd := len(campos)
	count := 0
	for campo, valorCampo := range campos {
		buffer.WriteString(campo)
		buffer.WriteString(separatorFront)
		buffer.WriteString(valorCampo)
		count++
		if qtd != count {
			buffer.WriteString(separatorBehind)
		} else {
			buffer.WriteString(separatorBehind2)
		}
	}
	buffer.WriteString("}]")
	return buffer.String()
}

// MoverArquivo Move o arquivo processado.
func (l *Layout) MoverArquivo(pathOrigem string, pathDestino string, pathInvalido string) {
	if l.Flgprocessado {
		err := util.MoverArquivo(pathOrigem+"/"+l.Nomearquivo, pathDestino+"/"+l.Nomearquivo)
		if err != nil {
			panic(err)
		}
	} else {
		err := util.MoverArquivo(pathOrigem+"/"+l.Nomearquivo, pathInvalido+"/"+l.Nomearquivo)
		if err != nil {
			panic(err)
		}
	}
}
