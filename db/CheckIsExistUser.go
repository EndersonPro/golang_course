package db

import (
	"context"
	"time"

	"github.com/EndersonPro/golang_course/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckIsExistUser permite revisar si un usuario ya existe en la base de datos a traves de su email
func CheckIsExistUser(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoClient.Database("twittor")
	col := db.Collection("users")

	condition := bson.M{"email":email}

	var result models.User
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}