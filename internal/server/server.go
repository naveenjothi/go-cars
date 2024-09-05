package server

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"go-cars/internal/database"
)

type FiberServer struct {
	*fiber.App

	client *mongo.Client
}

func New() *FiberServer {
	client,_:=database.InitializeMongoClient()
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-cars",
			AppName:      "go-cars",
		}),
		client:client,
	}

	return server
}
