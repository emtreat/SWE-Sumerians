package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Emails struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Emil string             `json:"email"`
}


type EmailModel struct {
	DB *mongo.Collection
}

func (m EmailModel) GetEmail(cx *fiber.Ctx) error {
	var emails []Emails

	pointer, err := m.DB.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer pointer.Close(context.Background())

	for pointer.Next(context.Background()) {
		var email Emails
		if err := pointer.Decode(&email); err != nil {
			return err
		}
		emails = append(emails, email)
	}
	return cx.JSON(emails)
}
