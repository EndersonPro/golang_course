package middlewares

import (
	"net/http"
	"github.com/EndersonPro/golang_course/db"
)

// CheckDatabase Middleware para checkar la conexion
func CheckDatabase(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Lost database conection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}