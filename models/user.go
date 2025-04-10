package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	Ok                int = 200
	Created           int = 201
	NotFound          int = 404
	ExpectationFailed int = 417
	LengthRequired    int = 411
)

type User struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"`
}

type UserModel struct {
	DB *mongo.Collection
}

func (m UserModel) DeleteUser(cx *fiber.Ctx) error {
	id := cx.Params("id")

	usrId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return cx.Status(NotFound).JSON(fiber.Map{"error": "Id doesn't exist"})
	}

	filter := bson.M{"_id": usrId}
	_, err = m.DB.DeleteOne(context.Background(), filter)

	if err != nil {
		return cx.Status(ExpectationFailed).JSON(fiber.Map{"error": "failed to delete user"})
	}

	return cx.Status(Ok).JSON(fiber.Map{"user successfully deleted": true})
}

func (m UserModel) AddUser(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if user.Name == "" {
		return c.Status(LengthRequired).JSON(fiber.Map{"error:": "User must have a name"})
	}

	result, err := m.DB.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)

	return c.Status(Created).JSON(user)

}

func (m UserModel) GetUsers(c *fiber.Ctx) error {
	var users []User

	pointer, err := m.DB.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer pointer.Close(context.Background())

	for pointer.Next(context.Background()) {
		var user User
		if err := pointer.Decode(&user); err != nil {
			return err
		}
		users = append(users, user)
	}

	return c.JSON(users)
}
