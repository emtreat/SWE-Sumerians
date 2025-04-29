package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware

	"github.com/emtreat/SWE-Sumerians/models"

	"github.com/emtreat/SWE-Sumerians/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var users_collection *mongo.Collection

type Env struct {
	users  models.UsersModel
	emails models.EmailModel
	files  models.FileModel
}

func main() {

	db := utils.ConectToDb()                  // gets the database connection
	defer db.Disconnect(context.Background()) // defers disconnecting from the server until after the function is closed

	users_collection = db.Database("project_db").Collection("users")

	env := &Env{ //not to be confused with the poorly named ".env" file that is totally unrelated
		users:  models.UsersModel{DB: users_collection},
		emails: models.EmailModel{DB: users_collection},
		files:  models.FileModel{DB: users_collection},
	}

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",      // Allow requests from React frontend (vite-app server)
		AllowMethods: "GET,POST,DELETE",            // Allowed HTTP methods
		AllowHeaders: "Origin,Content-Type,Accept", // Allowed headers
	}))

	app.Post("/api/users", env.users.AddUser)
	app.Delete("/api/users/:id", env.users.DeleteUser)
	app.Get("/api/users", env.users.GetUsers)
	app.Get("/api/users/:email", env.emails.GetEmail)
	app.Post("/api/users/:email/files", env.files.AddFile)
	app.Get("/api/users/:email/files/:fileId", env.files.GetFile)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
