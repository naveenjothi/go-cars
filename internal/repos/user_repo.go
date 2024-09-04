package repos

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct{
	*Repository
}

func NewUserRepository(collection *mongo.Collection) *UserRepository{
	return &UserRepository{
		Repository: NewRepository(collection),
	}
}