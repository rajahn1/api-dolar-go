## API COTAÇÃO DÓLAR
Projeto realizado para o curso de GO EXPERT da FULLCYCLE.

### DESCRIÇÃO
- Server conecta com uma API de cotação do dólar para real e retorna o valor em JSON, salva em um banco de dados sqlite3. Roda localmente na porta 3232.
- Client recebe o valor vindo do server, e cria um arquivo cotacao.txt para salvar.

### RODAR A APLICAÇÃO:
1. Clone o reposítorio na sua máquina local:
```
$ git clone https://github.com/rajahn1/api-dolar-go.git
```
2. Execute o arquivo **server.go**
```
$ cd server
$ go run server.go
```
3. Execute o arquivo **client.go**
```
$ cd client
$ go run client.go 
```

### REQUISITOS
- Ter o GOPATH instalado na sua máquina.
- Git


