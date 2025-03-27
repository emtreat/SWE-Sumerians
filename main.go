package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware

	"github.com/emtreat/SWE-Sumerians/models"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/emtreat/SWE-Sumerians/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var collection_emails *mongo.Collection
var collection_files *mongo.Collection

type Env struct {
	users models.UserModel
}

func main() {


    db := utils.ConectToDb(); // gets the database connection
    defer db.Disconnect(context.Background()) // defers disconnecting from the server until after the function is closed

    collection = db.Database("project_db").Collection("users")
    collection_emails = db.Database("project_db").Collection("emails")
    collection_files = db.Database("project_db").Collection("emails_to_users_test")


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

    app.Get("/api/users", env.users.GetUsers)
    app.Post("/api/users", env.users.AddUser)
    app.Delete("/api/users/:id", env.users.DeleteUser)

    app.Get("/api/emails_to_users_test", getFiles)
    app.Get("/api/emails", getEmail)

    port := os.Getenv("PORT")

    log.Fatal(app.Listen("0.0.0.0:" + port))
}

func getEmail(cx *fiber.Ctx) error {
	var emails []models.Emails

	pointer, err := collection_emails.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer pointer.Close(context.Background())

	for pointer.Next(context.Background()) {
		var email models.Emails
		if err := pointer.Decode(&email); err != nil {
			return err
		}
		emails = append(emails, email)
	}

	return cx.JSON(emails)
}

  
func getFiles(cx *fiber.Ctx) error {
	var files []models.Users

	pointer, err := collection_files.Find(context.Background(), bson.M{})

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

