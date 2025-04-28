package models

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//JSON Format for files
// {
// filename: "name",
// filesize: 123,
// fileblob: binary
//}

type File struct {
	FileName string `json:"file_name" bson:"filename"`
	FileSize int32  `json:"file_size" bson:"filesize"`
	FileBlob []byte `json:"file_blob,omitempty" bson:"fileblob,omitempty"` // New field for blob data
}

type FileUpload struct {
	FileName string `json:"file_name"`
	FileSize int32  `json:"file_size"`
	FileBlob []byte `json:"file_blob"` // For receiving binary data
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

	file, err := c.FormFile("file")
	if err != nil {
		var fileUpload FileUpload
		if err := c.BodyParser(&fileUpload); err != nil {
			return c.Status(BadRequest).JSON(fiber.Map{
				"error": "Invalid file data",
			})
		}

		newFile := File{
			FileName: fileUpload.FileName,
			FileSize: fileUpload.FileSize,
			FileBlob: fileUpload.FileBlob,
		}

		if newFile.FileSize <= 0 && len(fileUpload.FileBlob) > 0 {
			newFile.FileSize = int32(len(fileUpload.FileBlob))
		}

		//Files to match the JSON tag in Users struct
		filter := bson.M{"email": email}
		update := bson.M{
			"$push": bson.M{"files": newFile}, //F to match Users struct
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

	fileData, err := file.Open()
	if err != nil {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Failed to open uploaded file",
		})
	}
	defer fileData.Close()

	buffer := make([]byte, file.Size)
	_, err = fileData.Read(buffer)
	if err != nil {
		return c.Status(BadRequest).JSON(fiber.Map{
			"error": "Failed to read file data",
		})
	}

	newFile := File{
		FileName: file.Filename,
		FileSize: int32(file.Size),
		FileBlob: buffer,
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
		"file": fiber.Map{
			"file_name": newFile.FileName,
			"file_size": newFile.FileSize,
		},
	})
}

// we have actually never assigned this function to an endpoint
// in the front end we extract files through the user struct

func (m FileModel) GetFiles(cx *fiber.Ctx) error {
	email := cx.Params("email")
	if email == "" {
		return cx.Status(400).JSON(fiber.Map{
			"error": "Email parameter is required",
		})
	}

	filter := bson.M{"email": email}
	var user Users

	err := m.DB.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return cx.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return cx.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}
	// returns same JSON as before of file and size
	fileInfos := make([]map[string]interface{}, len(user.Files))
	for i, file := range user.Files {
		fileInfos[i] = map[string]interface{}{
			"file_name": file.FileName,
			"file_size": file.FileSize,
		}
	}

	return cx.JSON(fileInfos)
}

// GetFile is to get information on one singular file
// given: email and filename
// if you want a blob then use get file

func (m FileModel) GetFile(cx *fiber.Ctx) error {
	email := cx.Params("email")
	fileName := cx.Params("fileName")

	if email == "" || fileName == "" {
		return cx.Status(400).JSON(fiber.Map{
			"error": "Email and file name parameters are required",
		})
	}

	filter := bson.M{"email": email}
	var user Users

	err := m.DB.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return cx.Status(404).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return cx.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	for _, file := range user.Files {
		if file.FileName == fileName {
			cx.Set("Content-Disposition", "attachment; filename="+file.FileName)
			// This sets the Content-Disposition header to prompt a download
			// with the original file name. The browser will show a download dialog
			// with the file name specified in the header.
			cx.Set("Content-Type", "application/octet-stream")
			// Set the content type to application/octet-stream for binary files
			// You can also set it to a specific type if you know the file type
			// for a PDF file, you can use "application/pdf"
			// text file, you can use "text/plain"
			// image, you can use "image/jpeg" or "image/png"
			// zip file, you can use "application/zip"
			// JSON file, you can use "application/json"
			// CSV file, you can use "text/csv"
			// Word document, you can use "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
			// Excel file, you can use "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
			// PowerPoint file, you can use "application/vnd.openxmlformats-officedocument.presentationml.presentation"
			// PDF file, you can use "application/pdf"

			return cx.Send(file.FileBlob)
			// Send the file blob as the response
		}
	}

	return cx.Status(404).JSON(fiber.Map{
		"error": "File not found",
	})
}
