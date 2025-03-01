package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	// "github.com/emtreat/SWE-Sumerians/models/user" TODO add external module for models and connect it here
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type User struct{
    Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Name string `json:"name"`
}


var collection *mongo.Collection

func main() {
    err := godotenv.Load(".env") // Get the environment set up (currently just using localhost and the DB I set up)

    if err != nil { // Return an error if the environment isn't set up 
        log.Fatal("Error loading environment: check \".env\" file", err)
    }

    URI := os.Getenv("URI") // Gets the database's URI from the ".env" 

    clientOpts := options.Client().ApplyURI(URI)
    client, err := mongo.Connect(context.Background(), clientOpts)

    if err != nil {
        log.Fatal("Error connecting to database", err)
    }

    defer client.Disconnect(context.Background())

    err = client.Ping(context.Background(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to database")

    collection = client.Database("project_db").Collection("users")

    app := fiber.New()

    app.Get("/api/users", getUser)
    app.Post("/api/users", addUser)
    app.Delete("/api/users/:id", deleteUser)

    port := os.Getenv("PORT")

    log.Fatal(app.Listen("0.0.0.0:"+port))
}

func addUser(cx *fiber.Ctx) error {
    user := new(User)

    if err := cx.BodyParser(user); err != nil {
        return err
    }

    if user.Name == "" {
        return cx.Status(411).JSON(fiber.Map{"error": "User must have a name"})
    }

    result, err := collection.InsertOne(context.Background(), user)

    if err != nil {
        return err
    }

    user.Id = result.InsertedID.(primitive.ObjectID)

    return cx.Status(201).JSON(user)
}

func deleteUser(cx *fiber.Ctx) error {
    id := cx.Params("id")

    usrId, err := primitive.ObjectIDFromHex(id)
    
    if err != nil {
        return cx.Status(404).JSON(fiber.Map{"error": "Id doesn't exist"})
    }

    filter := bson.M{"_id": usrId}
    _, err = collection.DeleteOne(context.Background(), filter)

    if err != nil {
        return cx.Status(417).JSON(fiber.Map{"error": "failed to delete user"})
    }

    return cx.Status(200).JSON(fiber.Map{"user successfully deleted":true})
}

func getUser(cx *fiber.Ctx) error {
    var users []User

    pointer, err := collection.Find(context.Background(), bson.M{})

    if err != nil {
        return err
    }

    defer pointer.Close(context.Background())

    for pointer.Next(context.Background()){
        var user User 
        if err := pointer.Decode(&user); err != nil {
            return err
        }
        users = append(users, user)
    }

    return cx.JSON(users)

}

