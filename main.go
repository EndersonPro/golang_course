package main

import (
	"log"

	"github.com/EndersonPro/golang_course/db"
	"github.com/EndersonPro/golang_course/handlers"
)

func main(){
	if db.CheckConnection() == 0 {
		log.Fatal("Ocurrio un error al conectar a la base de datos")
		return
	}
	handlers.Handlers()
}