package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //Esse é o banco em sim. Obgs: é minúsculo pois não é exportavel

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=books sslmode=disable password=123456" //Host, a porta q ta rodando o banco, o usuario, o nome do banco, se está em ssl mode e a senha

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{}) //esse erro ṕe pego do gomr.Open que passa o postgres.Open.

	if err != nil { //Aq é validade se teve algum erro
		log.Fatal("error: ", err) //Esse indica o erro, caso tenha algum
	}

	db = database // a variavel db possui o conteudo da database

	config, _ := db.DB() //aq chamamos o metodo db do proprio gorm

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100) //definir quantas conexões simultaneas
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB { //a função getdatabase retorna o o noisoso banco
	return db //nos vamos instaciar essa função uma vez no main e vamos conseguir pegar toda vez o msm banco sem precisar abrir varias conexões
}
