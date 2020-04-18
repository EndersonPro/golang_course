package db

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoClient Cliente de conexion de mongo
var MongoClient *mongo.Client = mongoInit()
var clientOption = options.Client().ApplyURI("mongodb+srv://endersonpro:jL6pP8d9kt25vEqJ@twittor-hhfe7.mongodb.net/test?retryWrites=true&w=majority")

/* mongoInit: funcion encargada en inicializar la conexion a la base de datos */
func mongoInit() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil);
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa a la base de datos")
	return client
}

// CheckConnection metodo para checkear la conexion a la base de datos
func CheckConnection() int {
	err :=  MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}