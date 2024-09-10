package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetBodyPayload(ctx *fiber.Ctx, dto interface{}) error {
	body := ctx.Body()
	if err := json.Unmarshal(body, dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Unable to parse body"})
	}
	return nil
}
