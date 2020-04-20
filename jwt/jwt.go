package jwt

import (
	"time"

	"github.com/EndersonPro/golang_course/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT metodo para genera token
func GenerateJWT(u models.User) (string, error) {
	jwtKey := []byte("miclavesecreta")

	payload := jwt.MapClaims{
		"_id": 		u.ID.Hex(),
		"email":	u.Email,
		"name": 	u.Name,
		"lastname": u.LastName,
		"exp": 		time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(jwtKey)

	if err != nil {
		return tokenStr, err		
	}
	return tokenStr, nil		
}