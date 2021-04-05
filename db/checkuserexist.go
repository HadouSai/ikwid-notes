package db

import (
	"context"
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* CheckUserExist verify if an user exists*/
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("ikwid-notes")

	collection := db.Collection("users")

	conditionemail := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, conditionemail).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
