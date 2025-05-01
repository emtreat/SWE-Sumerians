package models_test

import (
	"testing"

	"github.com/emtreat/SWE-Sumerians/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const (
	Ok                int = 200
	Created           int = 201
	NotFound          int = 404
	ExpectationFailed int = 417
	LengthRequired    int = 411
)


type mockUsers struct {}

type mockUsersModel struct {}

func (m *mockUsersModel) GetUsers(cx *fiber.Ctx) error {
    ...
}

func (m *mockUsersModel) DeletUser(cx *fiber.Ctx) error {
    ...
}

func (m mockUsersModel) AddUser(cx *fiber.Ctx) error {
    ...
}

