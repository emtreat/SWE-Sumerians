package models

import (
	"context"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecievedFile struct {
	FileName string `json:"file_name"`
	FileData string `json:"file_data"`
	FileSize int32  `json:"file_size"`
}

type File struct {
	FileName string `json:"file_name" bson:"file_name"`
	FileData []byte `json:"file_data" bson:"file_data"`
	FileSize int32  `json:"file_size" bson:"file_size"`
}

type FileModel struct {
	DB *mongo.Collection
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

	var uploadedFile RecievedFile
	if err := c.BodyParser(&uploadedFile); err != nil {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Invalid file data",
		})
	}

	if uploadedFile.FileName == "" || uploadedFile.FileData == "" {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "File name and size (positive integer) are required",
		})
	}

	blobData, err := base64.StdEncoding.DecodeString(uploadedFile.FileData)
	if err != nil {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Invalid file encoding",
		})
	}

	newFile := File{
		FileName: uploadedFile.FileName,
		FileData: blobData,
		FileSize: uploadedFile.FileSize,
	}

	if newFile.FileSize <= 0 {
		newFile.FileSize = int32(len(blobData))
	}

	filter := bson.M{"email": email}
	update := bson.M{
		"$push": bson.M{"files": uploadedFile},
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
		"file": fiber.Map{
			"file_name": newFile.FileName,
			"file_size": newFile.FileSize,
		},
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
