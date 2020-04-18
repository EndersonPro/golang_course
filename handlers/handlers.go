package handlers

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

// Handlers Añadiendo puerto y pongo a escuchar mi  
func Handlers() {
	router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}