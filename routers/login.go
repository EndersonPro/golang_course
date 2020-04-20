package routers

import (
	"encoding/json"
	"net/http"

	"github.com/EndersonPro/golang_course/db"
	"github.com/EndersonPro/golang_course/helpers"
	"github.com/EndersonPro/golang_course/jwt"
	"github.com/EndersonPro/golang_course/models"
	"golang.org/x/crypto/bcrypt"
)

// Login metodo para iniciar sesion
func Login(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Error: " + err.Error())
		return
	}

	if len(t.Email) == 0 {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "El correo es requerido")
		return
	}

	user, exist, _ := db.CheckIsExistUser(t.Email)
	if !exist {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Correo y/o contraseña incorrecto/s")
		return
	}

	passwordBytes := []byte(t.Password)
	passwordDbBytes := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(passwordDbBytes, passwordBytes)

	if err != nil {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Correo y/o contraseña incorrecto/s")
		return
	}

	token, errorToken := jwt.GenerateJWT(t)

	if errorToken != nil {
		helpers.HandleResponse(w, r, http.StatusBadRequest, "Correo y/o contraseña incorrecto/s")
		return
	}

	response := models.ResponseToken{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}