package db

import (
	"context"
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	// necesario cerrar esas addicciones al context

	db := MongoConnection.Database("ikwid-notes")
	collection := db.Collection("users")

	user.Password, _ = Encrypt(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjectID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjectID.String(), true, nil

}
