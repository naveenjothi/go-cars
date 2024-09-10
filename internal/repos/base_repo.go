package repos

import (
	"context"
	"go-cars/internal/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection}
}

func (r *Repository) FindById(id string) (*mongo.SingleResult, error) {
	log.Printf("Finding %s in the collection", id)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	result := r.collection.FindOne(context.Background(), filter)

	return result, nil
}

func (r *Repository) FindOne(filter interface{}) *mongo.SingleResult {
	return r.collection.FindOne(context.Background(), filter)
}

func (r *Repository) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(context.Background(), document)
}

func (r *Repository) UpdateOne(id string, document interface{}) (*models.UserModel, error) {
	// Convert the ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return nil, err
	}

	// Construct the filter to find the document
	filter := bson.M{"_id": objectID}

	// Construct the update document using $set to update specific fields
	update := bson.M{"$set": document}

	// Define the result variable to store the updated document
	result := new(models.UserModel)

	// Find the document and apply the update with options to return the updated document
	err = r.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After), // Return the updated document
	).Decode(result)

	if err != nil {
		log.Printf("Error updating document: %v", err)
		return nil, err
	}

	return result, nil
}

func (r *Repository) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(context.Background(), filter)
}
