package services

import (
	"Tamer/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func DeleteArticle(id string, db *model.DBCollection) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID format: %v\n", err)
		return nil, errors.New("Invalid ObjectID format")
	}

	filter := bson.D{{"_id", objectID}}

	result, err := db.Collection.DeleteOne(db.Ctx, filter)
	if err != nil {
		return result, err
	}

	return result, nil
}
