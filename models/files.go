package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type File struct {
	FileName string `json:"file_name"`
	FileSize int32  `json:"file_size"`
}

type FileModel struct {
   DB *mongo.Collection; 
}

func (m FileModel) AddFile(c *fiber.Ctx) error {
	const (
		BadRequest    int = 400
		NotFound      int = 404
		InternalError int = 500
		Created       int = 201
	)

	email := c.Params("email")
	if email == "" {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Email is required",
		})
	}

	var newFile File
	if err := c.BodyParser(&newFile); err != nil {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Invalid file data",
		})
	}

	if newFile.FileName == "" || newFile.FileSize <= 0 {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "File name and size (positive integer) are required",
		})
	}

	filter := bson.M{"email": email}
	update := bson.M{
		"$push": bson.M{"files": newFile},
	}

	result, err := m.DB.UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return c.Status(InternalError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	if result.MatchedCount == 0 {
		return c.Status(NotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(Created).JSON(fiber.Map{
		"message": "File added successfully",
		"file":    newFile,
	})
}

func (m FileModel) GetFiles(cx *fiber.Ctx) error {
	var files []File

	pointer, err := m.DB.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer pointer.Close(context.Background())

	for pointer.Next(context.Background()) {
		var file File
		if err := pointer.Decode(&file); err != nil {
			return err
		}
		files = append(files, file)
	}

	return cx.JSON(files)

}
