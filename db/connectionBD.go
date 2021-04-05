package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoConnection has the connection to Database */
var MongoConnection = ConnectionBD()

const username = "adminx1"
const password = "adminx1"
const uri = "mongodb+srv://" + username + ":" + password + "@ikwid.jwa6i.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

var clientOptions = options.Client().ApplyURI(uri)

/* ConnectionBD connect to DB  */
func ConnectionBD() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	// comprabation DB
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Connection to DB Succesfull")

	return client
}

/* CheckConnection cehck the connection to the database  */
func CheckConnection() int {
	if err := MongoConnection.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
