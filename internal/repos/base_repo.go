package repos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct{
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection}
}

func (r *Repository) FindById(id string) (*mongo.SingleResult,error) {
	log.Printf("Finding %s in the collection", id)

    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        log.Printf("Invalid ID format: %v", err)
		return nil,err
    }
    
    filter := bson.M{"_id": objectID}

	result := r.collection.FindOne(context.Background(), filter)

	return result,nil
}

func (r *Repository) FindOne(filter interface{}) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

func (r *Repository) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(),document)
}

func (r *Repository) UpdateOne(filter, document interface{}) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(context.Background(),filter,document)
}

func (r *Repository) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(context.Background(),filter)
}