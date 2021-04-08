package db

import (
	"context"
	"time"

	"github.com/HadouSai/ikwid-notes/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EditProfile : edit a profile for a user
func EditProfile(user models.User, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConnection.Database("ikwid-notes")
	collection := db.Collection("users")

	register := make(map[string]interface{})

	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		register["lastName"] = user.Lastname
	}
	if len(user.Email) > 0 {
		register["email"] = user.Email
	}
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}

	update := bson.M{
		"$set": register,
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": bson.M{"$eq": objectId}}

	if _, err := collection.UpdateOne(ctx, query, update); err != nil {
		return false, err
	}

	return true, nil
}
