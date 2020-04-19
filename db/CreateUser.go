package db

import (
	"context"
	"time"

	"github.com/EndersonPro/golang_course/helpers"
	"github.com/EndersonPro/golang_course/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser metodo para crear un usuario en la base de datos
func CreateUser(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoClient.Database("twittor")
	col := db.Collection("users")

	u.Password, _ =  helpers.EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}