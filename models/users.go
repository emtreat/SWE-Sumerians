package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string             `json: "email"`
	Files []File             `json: "files"`
}
