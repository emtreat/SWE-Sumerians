package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/emtreat/SWE-Sumerians/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection_test *mongo.Collection

type TestEnv struct {
	users models.UserModel
}

// TestServer starts up the mock data collection and returns the fiber, collection, and error
func TestServer() (*fiber.App, *mongo.Collection, error) {
	err := godotenv.Load(".env") // Get the environment set up (currently just using localhost and the DB I set up)

	if err != nil { // Return an error if the environment isn't set up
		return nil, nil, err
	}

	URI := os.Getenv("URI") // Gets the database's URI from the ".env"
	dbOpts := options.Client().ApplyURI(URI)
	db, err := mongo.Connect(context.Background(), dbOpts)

	if err != nil {
		return nil, nil, err
	}

	//defer db.Disconnect(context.Background())

	// err = db.Ping(context.Background(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected to database")
	//collection with test data
	collection_test := db.Database("project_db").Collection("test_users")

	//Delete all documents in collection
	_, err = collection_test.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return nil, nil, err
	}

	env := &TestEnv{ //not to be confused with the poorly named ".env" file that is totally unrelated
		users: models.UserModel{DB: collection_test},
	}

	app := fiber.New()

	app.Get("/api/users", env.users.GetUsers)
	app.Post("/api/users", env.users.AddUser)
	app.Delete("/api/users/:id", env.users.DeleteUser)
	return app, collection_test, nil
}

func TestUserFunctions(t *testing.T) {
	fiberapp, collection, err := TestServer()
	if err != nil {
		fmt.Println("you failed hard, the server didn't start")
		t.FailNow()
	}
	defer collection.Database().Client().Disconnect(context.Background())
	defer collection.DeleteMany(context.Background(), bson.M{})

	// Test data
	testUser := map[string]interface{}{
		"name":     "Test User",
		"email":    "test@example.com",
		"password": "password123",
	}
	userData, _ := json.Marshal(testUser)

	// Step 1: Create a new user (POST)
	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(userData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := fiberapp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

}
