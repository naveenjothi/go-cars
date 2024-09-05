package server

import (
	"go-cars/internal/handlers/user"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Post("/user",s.withClient(user.CreateUserHandler))
	s.App.Get("/user/:id",s.withClient(user.GetUserHandler))
}

func (s *FiberServer) withClient(handler func(*fiber.Ctx, *mongo.Client) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return handler(c, s.client)
	}
}
