package utils

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

// HandleNotFoundError handles the not found error and sends a 404 response
func HandleNotFoundError(ctx *fiber.Ctx, id string) error {
	log.Printf("Resource not found: %s", id)
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": "Resource not found",
	})
}

// HandleInternalServerError handles unexpected errors and sends a 500 response
func HandleInternalServerError(ctx *fiber.Ctx, err error) error {
	log.Printf("Internal server error: %v", err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": "Internal server error",
	})
}

// HandleMongoError checks for MongoDB specific errors and sends appropriate responses
func HandleMongoError(ctx *fiber.Ctx, err error, id string) error {
	if err == mongo.ErrNoDocuments {
		return HandleNotFoundError(ctx, id)
	}
	return HandleInternalServerError(ctx, err)
}
