package controllers //Esse controller é necessário para testar

//http://localhost:5000/api/v1 antes da variavel pra testar

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
)

func ShowBook(c *gin.Context) { //o gin ja injeta direto, portanto com esse objeto "gin.Context", podemos fazer oq quisermos, tendo mais liberdade
	id := c.Param("id") //é daq que pegamos o ID, usando o "c"

	newid, err := strconv.Atoi(id) //é importante, para converter o ID, que vem em string, para ID

	if err != nil { //Caso tenha erro, retornamos a JSON, com o status 400, e o objeto de.h Passando o erro
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return //se não tiver esse return, ele vai retornar uma JSON com vários campos
	}

	db := database.GetDatabase()

	var book models.Book               //Essa é a variavel do nosso livro
	err = db.First(&book, newid).Error //Esa lista vai retornar um onjeto do tipo gorm.db, ele tem esse .error que consegue definir se tem um erro ou não.
	if err != nil {
		c.JSON(400, gin.H{ //Caso tenha erro, retornamos a JSON, com o status 400, e o objeto de.h Passando o erro
			"error": "cannot find product by id:" + err.Error(),
		})
		return //se não tiver esse return, ele vai retornar uma JSON com vários campos
	}
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) { //criação de livro
	db := database.GetDatabase()

	var book models.Book //esse é o nosso book. Ele é do tipo models.Book

	err := c.ShouldBindJSON(&book) //Nosso body deve ser passado para dentro do obejto book(ser passado para JSON)

	if err != nil {
		c.JSON(400, gin.H{ //Caso tenha erro, retornamos a JSON, com o status 400, e o objeto de.h Passando o erro
			"error": "cannot bind JSON:" + err.Error(),
		})
		return //se não tiver esse return, ele vai retornar uma JSON com vários campos
	}

	err = db.Create(&book).Error //Ele cria nossa entidade pela primeira vez, passando nosso book.Error

	if err != nil { //Caso tenha erro, retornamos a JSON, com o status 400, e o objeto de.h Passando o erro
		c.JSON(400, gin.H{
			"error": "cannot creat book: " + err.Error(),
		})
		return //se não tiver esse return, ele vai retornar uma JSON com vários campos
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {

	db := database.GetDatabase()
	var books []models.Book
	err := db.Find(&books).Error
	if err != nil { //Caso tenha erro, retornamos a JSON, com o status 400, e o objeto de.h Passando o erro
		c.JSON(400, gin.H{
			"error": "cannot list book: " + err.Error(),
		})
		return //se não tiver esse return, ele vai retornar uma JSON com vários campos
	}

}
