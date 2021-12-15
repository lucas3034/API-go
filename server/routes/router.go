package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine { //O que está dentro do parentes é o que ele recebe, e a direita do parenteses, é o que ele devolve
	main := router.Group("api/v1") //Aqui começamos a criação dos endpoints.  Essa api/v1 fica padrão nas nossas rotas, qnd passamos nossa URL
	{
		books := main.Group("books") //Tudo aq dentro, ja vai ter /api/v1/books
		{
			books.GET("/id", controllers.ShowBook) //Esse método quer dizer que ele vai listar um livro só
			books.GET("/", controllers.ShowBook)
			books.POST("/id", controllers.CreateBook)
		}
	}

	return router
}
