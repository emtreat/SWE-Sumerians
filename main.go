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
	app.Get("/api/users", getFiles)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getFiles(cx *fiber.Ctx) error {
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
