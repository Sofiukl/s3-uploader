package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// SaveToDB - This is the helper function for saving data in DB
func SaveToDB(database *mongo.Database, collectionName string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := database.Collection(collectionName)
	saveResult, insertErr := collection.InsertOne(context.TODO(), doc)
	if insertErr != nil {
		return saveResult, insertErr
	}
	return saveResult, nil
}
