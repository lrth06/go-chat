package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty, unique"`
	Email    string             `bson:"email,omitempty, unique, index"`
	Password string             `bson:"password,omitempty"`
	Avatar   string             `bson:"avatar,omitempty"`
	Token    string             `bson:"token,omitempty"`
	Roles    []string           `bson:"roles,omitempty, default=['dXNlcjpiYXNpYw==']"`
}
