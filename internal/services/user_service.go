package services

import (
	"encoding/json"
	"fmt"
	"go-cars/internal/models"
	"go-cars/internal/repos"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserService struct{
	repository *repos.UserRepository
}

func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{
		repository: repos.NewUserRepository(collection),
	}
}

func (s *UserService)CreateUser(ctx *fiber.Ctx) error {
	dto := new(models.UserModel)
	body := ctx.Body()
	if err := json.Unmarshal(body, dto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Unable to parse body"})
	}
	fmt.Println("dto",dto)
	insertResult,err:=s.repository.InsertOne(dto)
	if(err != nil){
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Unable to create user"})
	}
	dto.Id = insertResult.InsertedID.(primitive.ObjectID)
	return ctx.Status(fiber.StatusCreated).JSON(dto)
}