package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Emails struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Emil string             `json:"email"`
}
