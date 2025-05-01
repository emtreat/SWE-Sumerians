package models_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emtreat/SWE-Sumerians/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)




/*
func TestGetFile() (t *testing.T) {

    env := &TestEnv{
        files: models.FileModel{DB: collection_test},
    }

    app := fiber.New()
    app.Post("/", env.files.GetFile)

    req, err := http.NewRequest(http.MethodGet, "/", nil)
    if err != nil {
        t.Fatalf("Couldn't create get request: %v\n", err)
    }

    w := httptest.NewRecorder()

    app.Test()

}*/

func TestFileFunctions(t *testing.T) {
    // Test server values
    app := fiber.New()
    defer app.Shutdown()

    env := &TestEnv{
        files: models.FileModel{DB: collection_test},
    }

    app.Post("/test", env.files.AddFile)

    // Test adding a File 
    testFile := map[string] interface{} {
        "ID": "TestID",
        "FileName": "TestFileName",
        "FileSize": 34,
        "FileBlob": 0xa20c,
    }
    fileData, _ := json.Marshal(testFile) // convert to json
    // Check Post
    req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(fileData))
//    req.Header.Set("Content-Type", "application/json")

    resp, err := app.Test(req)
    print(resp.StatusCode)

    assert.Nil(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)

    // Check Get

    // Check Delete
}
