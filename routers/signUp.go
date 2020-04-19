package routers

import (
	"encoding/json"
	"net/http"

	"github.com/EndersonPro/golang_course/db"
	"github.com/EndersonPro/golang_course/models"
)

// SignUp es para crear un nuevo usuario
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error: " + err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Debe ingresar una contraseña", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener al menos 6 caracteres", 400)
		return
	}


	_, encontrado, _ := db.CheckIsExistUser(t.Email)

	if encontrado {
		http.Error(w, "Ya el usuario existe", 400)
		return
	}

	_, status, err := db.CreateUser(t)

	if err != nil {
		http.Error(w, "Ocurrio un error: " + err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se pudo realizar el registro.", 400)
		return
	} 

	w.WriteHeader(http.StatusCreated)


}