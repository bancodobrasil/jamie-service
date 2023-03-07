[![Go Reference](https://pkg.go.dev/badge/github.com/abu-lang/goabu.svg)](https://pkg.go.dev/github.com/bancodobrasil/jamie-service)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/bancodobrasil/jamie-service/blob/develop/LICENSE)

# Jamie Service [![About_de](https://github.com/yammadev/flag-icons/blob/master/png/BR.png?raw=true)](https://github.com/bancodobrasil/jamie-service/blob/develop/README-PTBR.md)

## Required Software

-   You must have the **Go Programming Language** installed in your machine to run this project. You can get the official download [here](https://go.dev/doc/install).

## Initializate the project

-   Clone this project to your local machine.
- On _main.go_ folder (_../jamie-service/main.go_), open your local terminal and type the command `go run main.go`, if your OS is windows, you can by build and run the executable `go build && ./jamie-service.exe`, or if your OS is mac or linux, type the command `go build -o service && ./service $@`.

# Using main endpoints

_By default the port will be :8005_

-   GET **http://localhost:YOURSETTEDPORT/**

    -   Will return simple message to client such as: "Jamie Service Works!!!"

<!-- -   POST **http://localhost:YOURSEETEDPORT/api/v1/eval**
    -   On this end point you must have to pass a body, witch is the parameters setted by rulesheet folder on the featws-transpiler. Using case 0001 for example the body should be:
        ```json
        {
        	"mynumber": "45"
        }
        ```
        -   Sending the request, response should be something like that: (this response difined by .featws file on ruleshet folder, in that case is false because the condition is mynumber > 12)
        ```json
        {
        	"myboolfeat": false
        }
        ``` -->
-   GET **http://localhost:YOURSEETEDPORT/swagger/index.html**

    -   On your browser, you can see the swagger documentation of the api.

-   GET **http://localhost:YOURSEETEDPORT/health/live?full=1**

    -   This endpoint will check the live status of the application, just like that:
        ```json
        {
        	"goroutine-threshold": "OK"
        }
        ```

-   GET **http://localhost:YOURSEETEDPORT/health/ready?full=1**
    -   This endpoint will check the ready status of the services used by ruller project
    ```json
    {
    	"goroutine-threshold": "OK",
    	"resource-loader": "OK"
    }
    ```
