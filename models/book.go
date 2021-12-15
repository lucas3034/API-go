package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct { //struct do tipo book
	ID          uint           `json:"id" gorm:"primaryKey"` //aq recebemos um ID(ele tem q ser todo maisculo por causa do gorm), ele recebe a instrução: "gorm:"primaryKey"". Ele é refereten ao campo id, importante para quando exportar para JSON
	Name        string         `json:"name"`                 //o Nome é uma string que recebe o json:name
	Description string         `json:"description"`          //Descrição do livro
	MediumPrice float32        `json:"medium_price"`         //Preço médio do livro
	Author      string         `json:"author"`               //Autor do livro
	ImageURL    string         `json:"img_url"`              //URL da imagem do livro
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"created"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted"` //Serve para acontecer o soft delete, que não deleta totalmente, só da um "fake" delete, que da pra ver qnd ta deletado ou n. Por padrão ele não lista q estão com os campos deletados diferentes de null
}
