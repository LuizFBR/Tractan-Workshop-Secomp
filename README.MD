# Introdução
Respositório do workshop 'Ingestão de Dados com Golang para Construir Sensores IOT" da Tractian realizado 05/08/2023 na Secomp

Utilizando Kafka, os membros do workshop aprenderam a linguagem GO, paralelismo e a acessar um servidor Kafka, onde foi realizado um sorteio com os clientes do servidor.
Foi adquirido conhecimento sobre paralelismo em go (veja `main.go`).

# Passos para acessar servidor

1 - `go mod init https://github.com/jpvolt/secomp-golang` para configurar o servidor.
2 - `go get github.com/segmentio/kafka-go para instalação` do pacote Kafka para GO.
3 - `go run server.go` para estabelecer uma conexão com o servidor.

