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
    files models.FileModel
}

func ServerTest() (*fiber.App, *mongo.Collection, error) {
    err := godotenv.Load("../.env")

    if err != nil {
        fmt.Println("Test environment setup error")
        return nil, nil, err
    }

    URI := os.Getenv("URI")

    dbOpts := options.Client().ApplyURI(URI)
    db, err := mongo.Connect(context.Background(), dbOpts)

    if err != nil {
        fmt.Println("Error test connecting")
        return nil, nil, err
    }

    collection_test := db.Database("project_db").Collection("test_users")

    _, err = collection_test.DeleteMany(context.Background(), bson.M{})

    if err != nil {
        fmt.Println("error deleting")
        return nil, nil, err
    }

    env := &TestEnv{
        files: models.FileModel{DB: collection_test},
    }

    app := fiber.New()

    app.Post("/api/users/:email/files", env.files.AddFile)
    return app, collection_test, nil
}

func TestFileFunctions(t *testing.T) {
    // Test server values
    app, collection, err := ServerTest()

    if err != nil {
        fmt.Println("Test server didn't start")
        t.FailNow()
    }


    defer collection.Database().Client().Disconnect(context.Background())

    defer collection.DeleteMany(context.Background(), bson.M{})


    // Test adding a File TODO
    testFile := map[string] interface{} {
        "ID": "TestID",
        "FileName": "TestFileName",
        "FileSize": 34,
        "FileBlob": 01001,
    }

    fileData, _ := json.Marshal(testFile) // convert to json

    // Check Post
    req := httptest.NewRequest(http.MethodPost, "/api/users/:email/files", bytes.NewReader(fileData))

    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)

    assert.Nil(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    // Check Get


    // Check Delete
}
