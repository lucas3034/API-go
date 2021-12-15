package main

import (
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/server"
)

func main() {
	database.StartDB() //Qnd rodar o main, vamos ter o nosso banco de dados

	server := server.NewServer() //Aqui é a criação do nosso server

	server.Run() //Aqui é para rodar o server
}
