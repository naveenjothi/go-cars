package services

import (
	"encoding/json"
	"go-cars/internal/models"
	"go-cars/internal/repos"
	"go-cars/internal/utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	repository *repos.UserRepository
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		repository: repos.NewUserRepository(collection),
	}
}

func (s *UserService) CreateUser(ctx *fiber.Ctx) error {
	dto := new(models.UserModel)
	body := ctx.Body()
	if err := json.Unmarshal(body, dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Unable to parse body"})
	}
	dto.InitiliseDefaultValue()
	insertResult, err := s.repository.InsertOne(dto)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Profile name already exists"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unable to create user"})
	}
	dto.Id = insertResult.InsertedID.(primitive.ObjectID)
	return ctx.Status(fiber.StatusCreated).JSON(dto)
}

func (s *UserService) FindOneUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user := new(models.UserModel)
	resp, err := s.repository.FindById(id)
	if err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}

	if err := resp.Decode(user); err != nil {
		return utils.HandleMongoError(ctx, err, id)
	}
	return ctx.Status(fiber.StatusOK).JSON(user)
}
