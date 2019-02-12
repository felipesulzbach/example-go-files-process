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
{1:17}(https://github.com/felipesulzbach/exemplo-files-process/blob/master/main.go)

Ele ficará ativo em um looping e será executado em um intervalo de 5 segundos (pré determinado na constante *delaySecound*).

Será possível determinar o caminho do diretório raiz onde os arquivos serão processados, modificando as contantes do arquivo *sistema.go*.

Arquivo *sistema.go*
{12:21}(https://github.com/felipesulzbach/exemplo-files-process/blob/master/sistema.go)

### Processamento de arquivos
O sistema irá processar os arquivos que se encontram na pasta *origem*. Serão considerados todos arquivos com extenção *txt* e *cvs*, que possuem as extruturas conforme exemplos abaixo:

Arquivo com dados de pessoa Física
{1:7}(https://github.com/felipesulzbach/exemplo-files-process/blob/master/diretorioarquivos/origem/testeok01.txt)

Arquivo com dados de pessoa Jurídica
{1:7}(https://github.com/felipesulzbach/exemplo-files-process/blob/master/diretorioarquivos/origem/testeok02.txt)

Todos os arquivos que não respeitarem essas estruturas, não serão processados.

### Arquivos processamento
Os arquivos que não possuem nenhuma inconsistência em seus dados, serão processados com sucesso, seus dados serão exibidos no log e serão movidos para a pasta *destino*. Os arquivos que possuem alguma inconsistência em seus dados, serão movidos para pasta *invalido*.

![Arquivos para serem processados](https://github.com/felipesulzbach/exemplo-files-process/blob/master/img/aprocessar.png)

![Log com os dados extraidos](https://github.com/felipesulzbach/exemplo-files-process/blob/master/img/logdados.png)

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
