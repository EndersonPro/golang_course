package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/EndersonPro/golang_course/middlewares"
	"github.com/EndersonPro/golang_course/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers Añadiendo puerto y pongo a escuchar mi  
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlewares.CheckDatabase(routers.SignUp)).Methods("POST")


	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}