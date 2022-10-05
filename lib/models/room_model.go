package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Room struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty, unique"`
	Owner    primitive.ObjectID `bson:"owner,omitempty"`
	Password string             `bson:"password,omitempty"`
	Settings RoomSettings       `bson:"settings,omitempty"`
}

type RoomSettings struct {
	Private    bool                 `bson:"private,omitempty"`
	Background string               `bson:"background,omitempty"`
	Moderators []primitive.ObjectID `bson:"moderators,omitempty"`
}
