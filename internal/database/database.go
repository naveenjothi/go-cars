package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func InitializeMongoClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	return client, nil
}




// func (c *mongo.Client) Health() map[string]string {
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	err := c.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatalf(fmt.Sprintf("db down: %v", err))
// 	}

// 	return map[string]string{
// 		"message": "It's healthy",
// 	}
// }


func getCollection(client *mongo.Client, dbName,collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

func GetUserCollection(client *mongo.Client) *mongo.Collection{
	dbName := os.Getenv("DB_NAME")
	return getCollection(client,dbName, "users")
}