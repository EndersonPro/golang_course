package helpers

import "golang.org/x/crypto/bcrypt"

// EncryptPassword metodo para encriptar la contraseña del usuario
func EncryptPassword(p string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), costo);
	return string(bytes), err
} 