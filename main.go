package main

import (
	"log"
	"github.com/endersonpro/twittor/handlers"
	"github.com/endersonpro/twittor/db"
)

func main(){
	if db.CheckConnection() == 0 {
		log.Fatal("Ocurrio un error al conectar a la base de datos")
		return
	}
	handlers.Handlers()
}