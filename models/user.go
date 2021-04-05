package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User model for users in database*/
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name,omitempty"`
	Lastname  string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	BirthDay  time.Time          `bson:"birthday,omitempty" json:"birthday,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Location  string             `bson:"Location" json:"location,omitempty"`
	WebSite   string             `bson:"WebSite" json:"webSite,omitempty"`
}
