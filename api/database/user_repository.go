package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string
	Email    string
}

var db *mongo.Database

func ConnectToMongoDB(uri string, dbName string) (*mongo.Database, error) {
	if db != nil {
		return db, nil
	}

	log.Printf("nice %s", uri)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	db = client.Database(dbName)
	fmt.Println("Connected to MongoDB")
	return db, nil
}

func InsertUser(client *mongo.Client, user *User) error {
	collection := client.Database("your_database_name").Collection("users")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}
