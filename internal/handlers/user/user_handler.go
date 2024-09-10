package user

import (
	"go-cars/internal/database"
	"go-cars/internal/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserHandler(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := database.GetUserCollection(client)

	userService := services.NewUserService(collection)

	return userService.CreateUser(ctx)
}

func GetUserHandler(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := database.GetUserCollection(client)

	userService := services.NewUserService(collection)

	return userService.FindOneUserByID(ctx)
}

func UpdateUserHandler(ctx *fiber.Ctx, client *mongo.Client) error {
	collection := database.GetUserCollection(client)

	userService := services.NewUserService(collection)

	return userService.UpdateUser(ctx)
}
