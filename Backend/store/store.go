package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// this is basically a structure that contains all the collections
type MongoStore struct {
	UsersCollection      *mongo.Collection
	AdminCollection      *mongo.Collection
	AttendanceCollection *mongo.Collection
}

// this creates a new mongo store
func NewMongoStore() *MongoStore {
	return &MongoStore{}
}

// this function opens the connection with mongodb
func (m *MongoStore) OpenConnectionWithMongoDB(connectionString, databaseName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	db := client.Database(databaseName)
	m.UsersCollection = db.Collection("Users")
	m.AdminCollection = db.Collection("Admin")
	m.AttendanceCollection = db.Collection("Attendance")

	return nil
}

// this function disconnects connection with mongodb
func (m *MongoStore) Close() {
	if err := m.UsersCollection.Database().Client().Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting from MongoDB: %v", err)
	}
}

// this function opens particular collection
func OpenCollection(client *mongo.Client, databaseName string, collectionName string) *mongo.Collection {

	var collection *mongo.Collection = client.Database(databaseName).Collection(collectionName)

	return collection
}
