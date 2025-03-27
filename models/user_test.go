package models_test

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
func ServerTest() (*fiber.App, *mongo.Collection, error) {
	err := godotenv.Load("../.env") // Get the environment set up (currently just using localhost and the DB I set up)

	if err != nil { // Return an error if the environment isn't set up
		fmt.Println("environment setup error")
		return nil, nil, err
	}

	URI := os.Getenv("URI") // Gets the database's URI from the ".env"
	dbOpts := options.Client().ApplyURI(URI)
	db, err := mongo.Connect(context.Background(), dbOpts) //connect to the database

	if err != nil {
		fmt.Println("error connecting")
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

	//Delete all documents in collection to start test
	_, err = collection_test.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("error deleting")
		return nil, nil, err
	}

	env := &TestEnv{ //not to be confused with the poorly named ".env" file that is totally unrelated
		users: models.UserModel{DB: collection_test},
	}
	// new fiber app initialized
	app := fiber.New()
	// Routes
	app.Get("/api/users", env.users.GetUsers)
	app.Post("/api/users", env.users.AddUser)
	app.Delete("/api/users/:id", env.users.DeleteUser)
	return app, collection_test, nil
}

// Test for User API calls
func TestUserFunctions(t *testing.T) {
	//extracting values from the Test Server
	fiberapp, collection, err := ServerTest()
	if err != nil {
		fmt.Println("you failed hard, the server didn't start")
		t.FailNow()
	}
	// close connection and clear data base after the function
	defer collection.Database().Client().Disconnect(context.Background())
	defer collection.DeleteMany(context.Background(), bson.M{})

	// Test data for user field
	testUser := map[string]interface{}{
		"name": "Test User",
	}
	// user data converted to JSON
	userData, _ := json.Marshal(testUser)

	//Create a new user (POST) and test
	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(userData))
	req.Header.Set("Content-Type", "application/json")
	resp, err := fiberapp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// check User GET
	req = httptest.NewRequest(http.MethodGet, "/api/users", nil)
	resp, err = fiberapp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// check User DELETE
	// userid := "67c9ef8861422722922d6597"
	// req = httptest.NewRequest(http.MethodDelete, "/api/user/:id"+userid, nil)
	// resp, err = fiberapp.Test(req)
	// assert.Nil(t, err)
	// assert.Equal(t, http.StatusOK, resp.StatusCode)

}
