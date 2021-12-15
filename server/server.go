package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/server/routes"
)

type Server struct { //essa estrutura recebe a porta e o server.
	port   string
	server *gin.Engine // o gin precisa desse gin.Engine pra iniciar o server
}

func NewServer() Server { //esse é o construtor do server q retorna o server, e escolhemos a porta que queremos.
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() { //esse é o método instanciado pelo server. o método Run é o responsavel por rodar o server
	router := routes.ConfigRoutes(s.server)

	log.Print("server está rodando na porta: ", s.port)
	log.Fatal(router.Run(":" + s.port)) //Aqui nos incicamos qual a porta que queremos rodar
}
