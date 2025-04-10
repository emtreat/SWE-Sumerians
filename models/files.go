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
