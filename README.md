# Exemplo de processamento de arquivos

Exemplo simples de um sistema em GoLang para processamento de arquivos.

___

## Pré-requisitos
* Visual Studio Code ou outro;
* GoLang (https://golang.org/);

## Como o sistema funciona?

### Início da execução
O sistema inicia pelo arquivo *main.go*, método *main()*.

Arquivo *main.go*
```go
// Determina o tempo em segundos, que o sistema ira esperar para executar novamente.
const delaySecound = 5

func main() {
	for {
		sistema.ProcessarArquivos()
		time.Sleep(delaySecound * time.Second)
	}
}
```

Ele ficará ativo em um looping e será executado em um intervalo de 5 segundos (pré determinado na constante *delaySecound*).

Será possível determinar o caminho do diretório raiz onde os arquivos serão processados, modificando as contantes do arquivo *sistema.go*.

Arquivo *sistema.go*
```go
//const pathroot = "/home/felipe/go/diretorioarquivos"
//const pathroot = "C:/temp/diretorioarquivos"
const pathroot = "diretorioarquivos"
const pathOrigem = pathroot + "/origem"
const pathDestino = pathroot + "/destino"
const pathInvalido = pathroot + "/invalido"
const extencaocsv = "csv"
const extencaotxt = "txt"
```


### Processamento de arquivos
O sistema irá processar os arquivos que se encontram na pasta *origem*. Serão considerados todos arquivos com extenção *txt* e *cvs*, que possuem as extruturas conforme exemplos abaixo:

Arquivo com dados de pessoa Física
```sql
ID;Nome;CPF;Idade
1;Nome 1;12345678901;15
2;Nome 2;23456789012;25
3;Nome 3;34567890123;18
4;Nome 4;45678901234;41
5;Nome 5;56789012345;32
```

Arquivo com dados de pessoa Jurídica
```sql
ID;Nome;CNPJ
6;Razao Social 6;12345678901234
7;Razao Social 7;23456789012345
8;Razao Social 8;34567890123456
9;Razao Social 9;45678901234567
10;Razao Social 10;56789012345678
```

Todos os arquivos que não respeitarem essas estruturas, não serão processados.

### Arquivos processamento
Os arquivos que não possuem nenhuma inconsistência em seus dados, serão processados com sucesso, seus dados serão exibidos no log e serão movidos para a pasta *destino*. Os arquivos que possuem alguma inconsistência em seus dados, serão movidos para pasta *invalido*.

Arquivos para serem processados
![Arquivos para serem processados](https://github.com/felipesulzbach/exemplo-files-process/blob/master/img/aprocessar.png)

Log com os dados extraidos
![Log com os dados extraidos](https://github.com/felipesulzbach/exemplo-files-process/blob/master/img/logdados.png)

Arquivos processados
![Arquivos processados](https://github.com/felipesulzbach/exemplo-files-process/blob/master/img/processados.png)

**OBS:** Como o sistema fica executando em um looping, no momento em que um arquivo é incluido na pasta *origem*, ele será processado (considerando o intervalo determinado).

## Considerações finais
O GoLang é uma linguagem que foi criada para ser mais performática possível, superando até mesmo linguagens como o Pyton, que atualmente são utilizados como solução para sistemas que necessitam de maior performance.
GoLang é utilizado em diversos produtos da empresa que à criou, a Google e vem ganhando posição mais forte nas empresas. Seguem algumas empresas que já utilizam o GoLang:
* Uber;
* Docker;
* Dropbox;
* OpenShift;
* Twitter.
