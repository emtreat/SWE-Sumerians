package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware

	"github.com/emtreat/SWE-Sumerians/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/emtreat/SWE-Sumerians/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var users_collection *mongo.Collection

type Env struct {
	users models.UserModel
    emails models.EmailModel
    files models.FileModel
}

func main() {

	db := utils.ConectToDb()                  // gets the database connection
	defer db.Disconnect(context.Background()) // defers disconnecting from the server until after the function is closed

	collection = db.Database("project_db").Collection("users") // subject to deletion
	users_collection = db.Database("project_db").Collection("users")

	env := &Env{ //not to be confused with the poorly named ".env" file that is totally unrelated
		users: models.UserModel{DB: collection},
	}

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",      // Allow requests from React frontend (vite-app server)
		AllowMethods: "GET,POST,DELETE",            // Allowed HTTP methods
		AllowHeaders: "Origin,Content-Type,Accept", // Allowed headers
	}))

	// app.Get("/api/users", env.users.GetUsers) // may not need in the future

	app.Post("/api/users", AddUser)
	app.Delete("/api/users/:id", env.users.DeleteUser)
	app.Get("/api/users", getUsers)
	app.Get("/api/users/:email", GetEmail)
	app.Post("/api/users/:email/files", AddFile)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getUsers(cx *fiber.Ctx) error {
	var files []models.Users

	pointer, err := users_collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer pointer.Close(context.Background())

	for pointer.Next(context.Background()) {
		var file models.Users
		if err := pointer.Decode(&file); err != nil {
			return err
		}
		files = append(files, file)
	}

	return cx.JSON(files)

}

func AddUser(c *fiber.Ctx) error {
	const (
		Ok                int = 200
		Created           int = 201
		NotFound          int = 404
		ExpectationFailed int = 417
		LengthRequired    int = 411
	)

	var user = new(models.Users)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if user.Email == "" {
		return c.Status(LengthRequired).JSON(fiber.Map{"error:": "User must have a valid email"})
	}

	result, err := users_collection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)

	return c.Status(Created).JSON(user)

}

func GetEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	if email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email parameter is required",
		})
	}

	filter := bson.M{
		"email": bson.M{"$regex": "^" + email + "$", "$options": "i"},
	}

	var user models.Users
	err := users_collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	return c.JSON(user)
}

func AddFile(c *fiber.Ctx) error {
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

	var newFile models.File
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

	result, err := users_collection.UpdateOne(
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
