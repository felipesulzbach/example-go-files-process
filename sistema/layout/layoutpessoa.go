package layout

import (
	"exemplo-files-process/util"
	"strconv"
)

// Pessoa Layout.
type Pessoa struct {
	Layout
	ID   int64
	Nome string
}

// Pessoa0001 Layout de pessoa juridica.
type Pessoa0001 struct {
	Pessoa
	Cnpj string
}

// Pessoa0002 Layout de pessoa fisica.
type Pessoa0002 struct {
	Pessoa
	Cpf   string
	Idade int
}

func (p Pessoa0001) show() string {
	campos := map[string]string{
		"NomeArquivo":   p.Nomearquivo,
		"Linhaarquivo":  p.Linhaarquivo,
		"Flgprocessado": strconv.FormatBool(p.Flgprocessado),
		"Msgerro":       p.Msgerro,
		"ID":            strconv.FormatInt(p.ID, 10),
		"Nome":          p.Nome,
		"Cnpj":          p.Cnpj,
	}
	retorno := toString("Pessoa0001", campos)
	return retorno
}

func (p Pessoa0002) show() string {
	campos := map[string]string{
		"NomeArquivo":   p.Nomearquivo,
		"Linhaarquivo":  p.Linhaarquivo,
		"Flgprocessado": strconv.FormatBool(p.Flgprocessado),
		"Msgerro":       p.Msgerro,
		"ID":            strconv.FormatInt(p.ID, 10),
		"Nome":          p.Nome,
		"Cpf":           p.Cpf,
		"Idade":         strconv.Itoa(p.Idade),
	}
	retorno := toString("Pessoa0002", campos)
	return retorno
}

// New Cria uma nova estrutura Pessoa0001 com os dados passados por parametro.
func (p *Pessoa0001) New(nomearquivo string, linhaarquivo string, flgprocessado bool, msgerro string, id string, nome string, cnpj string) error {
	if err := util.ValidarStringInt64(id, "ID"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	} else if err := util.ValidarString(nome, "NOME"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	} else if err := util.ValidarString(cnpj, "CNPJ"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	}

	var idP int64
	var err error
	if flgprocessado {
		idP, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}
	}
	*p = Pessoa0001{Pessoa{Layout{nomearquivo, linhaarquivo, flgprocessado, msgerro}, idP, nome}, cnpj}
	return nil
}

// New Cria uma nova estrutura Pessoa0002 com os dados passados por parametro.
func (p *Pessoa0002) New(nomearquivo string, linhaarquivo string, flgprocessado bool, msgerro string, id string, nome string, cpf string, idade string) error {
	if err := util.ValidarStringInt64(id, "ID"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	} else if err := util.ValidarString(nome, "NOME"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	} else if err := util.ValidarString(cpf, "CPF"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	} else if err := util.ValidarStringInt(idade, "IDADE"); err != nil {
		flgprocessado = false
		msgerro = err.Error()
	}

	var idP int64
	var idadeP int
	var err error
	if flgprocessado {
		idP, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}
		idadeP, err = strconv.Atoi(idade)
		if err != nil {
			return err
		}
	}
	*p = Pessoa0002{Pessoa{Layout{nomearquivo, linhaarquivo, flgprocessado, msgerro}, idP, nome}, cpf, idadeP}
	return nil
}
