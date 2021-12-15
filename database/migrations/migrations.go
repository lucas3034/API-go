package migrations

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) { //Recebe o banco do tipo gorm.DB
	db.AutoMigrate(models.Book{}) //Recebemos o models.Book
}
