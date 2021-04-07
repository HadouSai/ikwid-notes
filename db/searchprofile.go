package db

import (
	"context"
	"log"
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	db := MongoConnection.Database("ikwid-notes")
	collection := db.Collection("users")

	var profile models.User

	objectId, _ := primitive.ObjectIDFromHex(id)
	searchcondition := bson.M{"_id": objectId}

	err := collection.FindOne(ctx, searchcondition).Decode(&profile)
	profile.Password = ""

	if err != nil {
		log.Fatalf("Register not found" + err.Error())
		return profile, err
	}

	return profile, nil
}
