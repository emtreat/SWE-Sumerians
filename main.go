package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/emtreat/SWE-Sumerians/models"
	"github.com/emtreat/SWE-Sumerians/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

type Env struct {
	users models.UserModel
}

func main() {

    db := utils.ConectToDb(); // gets the database connection
    defer db.Disconnect(context.Background()) // defers disconnecting from the server until after the function is closed


	collection = db.Database("project_db").Collection("users")

	env := &Env{ //not to be confused with the poorly named ".env" file that is totally unrelated
		users: models.UserModel{DB: collection},
	}

	app := fiber.New()

	app.Get("/api/users", env.users.GetUsers)
	app.Post("/api/users", env.users.AddUser)
	app.Delete("/api/users/:id", env.users.DeleteUser)

	port := os.Getenv("PORT")

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
