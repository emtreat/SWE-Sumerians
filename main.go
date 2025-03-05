package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/emtreat/SWE-Sumerians/models" 
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var collection *mongo.Collection

type Env struct {
    users models.UserModel
}

func main() {
    err := godotenv.Load(".env") // Get the environment set up (currently just using localhost and the DB I set up)

    if err != nil { // Return an error if the environment isn't set up 
        log.Fatal("Error loading environment: check \".env\" file", err)
    }

    URI := os.Getenv("URI") // Gets the database's URI from the ".env" 

    dbOpts := options.Client().ApplyURI(URI)
    db, err := mongo.Connect(context.Background(), dbOpts)

    if err != nil {
        log.Fatal("Error connecting to database", err)
    }

    defer db.Disconnect(context.Background())

    err = db.Ping(context.Background(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to database")


    collection = db.Database("project_db").Collection("users")

    env := &Env{ //not to be confused with the poorly named ".env" file that is totally unrelated
        users: models.UserModel{DB: collection},
    }

    app := fiber.New()


    app.Get("/api/users", env.users.GetUsers)
    app.Post("/api/users", env.users.AddUser)
    app.Delete("/api/users/:id", env.users.DeleteUser)

    port := os.Getenv("PORT")

    log.Fatal(app.Listen("0.0.0.0:"+port))
}
