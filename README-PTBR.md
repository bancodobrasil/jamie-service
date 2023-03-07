[![Go Reference](https://pkg.go.dev/badge/github.com/abu-lang/goabu.svg)](https://pkg.go.dev/github.com/bancodobrasil/jamie-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/bancodobrasil/jamie-service/blob/develop/LICENSE)

# Jamie Service [![About_en](https://github.com/yammadev/flag-icons/blob/master/png/US.png?raw=true)](https://github.com/bancodobrasil/jamie-service/blob/develop/README.md)



## Software Necessário

-   Será necessário ter instalado em sua máquina a **Linguagem de Programação Go** para rodar o projeto. Você pode fazer o download na pádina oficial [aqui](https://go.dev/doc/install).


## Inicializando o Projeto

-   Clone o projeto para sua máquina local.
- Com a pasta do projeto aberto (_../jamie-service/main.go_), abra o arquivo _main.go_ e o terminal integrado, digite o comando `go run main.go`. Se voce utiliza o sistema OS ou windows, voce tambem pode dar build e executar o projeto com os comandos: `go build && ./jamie-service.exe`, caso use windows, ou `go build -o service && ./service $@` se utiliza Mac/Linux.

# Usando principais endpoints

_Por padrão a porta utilizada será a :8005_

-   GET **http://localhost:SUAPORTAESCOLHIDA/**

    -   Retornará uma mensagem simples ao cliente, como: "Jamie Service Works!!!"

<!-- -   POST **http://localhost:SUAPORTAESCOLHIDA/api/v1/eval**
    -   Neste ponto final você deve ter que passar um corpo, que são os parâmetros definidos pela pasta rulesheet no featws-transpiler. Usando o case 0001, por exemplo, o corpo deve ser:
        ```json
        {
        	"mynumber": "45"
        }
        ```
    -   Após a solicitação ter sido enviada, a resposta deve ser algo assim: (essa resposta é definida pelo arquivo .featws na pasta ruleshet, nesse caso é false porque a condição é meunúmero> 12)
        ```json
        {
        	"myboolfeat": false
        } 
        ```-->
-   GET **http://localhost:SUAPORTAESCOLHIDA/swagger/index.html**

    -   No seu navegador, você pode ver a documentação do swagger da api.

-   GET **http://localhost:SUAPORTAESCOLHIDA/health/live?full=1**

    -   Este endpoint verificará o status ativo do aplicação:

-   GET **http://localhost:SUAPORTAESCOLHIDA/health/ready?full=1**
    -   Este endpoint verificará o status se o serviços está pronto para ser usado ​​pelo projeto da ruller.
    ```json
    {
    	"goroutine-threshold": "OK",
    	"resource-loader": "OK"
    }
    ```
