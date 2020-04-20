package routers

import (
	"encoding/json"
	"net/http"

	"github.com/EndersonPro/golang_course/db"
	"github.com/EndersonPro/golang_course/helpers"
	"github.com/EndersonPro/golang_course/models"
)

// SignUp es para crear un nuevo usuario
func SignUp(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Error: " + err.Error())
		return
	}

	if len(t.Email) == 0 {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Debe ingresar una contraseña")
		return
	}

	if len(t.Password) < 6 {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "La contraseña debe tener al menos 6 caracteres")
		return
	}


	_, encontrado, _ := db.CheckIsExistUser(t.Email)

	if encontrado {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Ya el usuario existe")
		return
	}

	_, status, err := db.CreateUser(t)

	if err != nil {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Ocurrio un error: " + err.Error())
		return
	}

	if !status {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "No se pudo realizar el registro.")
		return
	} 
	helpers.HandleResponse(w, r, http.StatusCreated, "Usuario creado con exito")
}