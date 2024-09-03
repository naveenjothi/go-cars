package server

import (
	"github.com/gofiber/fiber/v2"

	"go-cars/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-cars",
			AppName:      "go-cars",
		}),

		db: database.New(),
	}

	return server
}
