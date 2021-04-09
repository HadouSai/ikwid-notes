package db

import (
	"context"
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertTweet
func InsertTweet(tweet models.InsertTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConnection.Database("ikwid-notes")
	collection := db.Collection("tweets")

	register := bson.M{
		"userid":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := collection.InsertOne(ctx, register)

	if err != nil {
		return "", false, err
	}

	objectId, _ := result.InsertedID.(primitive.ObjectID)

	return objectId.Hex(), true, nil

}
