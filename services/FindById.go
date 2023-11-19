package services

import (
	"Tamer/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func FindBYId(id string, db *model.DBCollection) (*model.Article, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID format: %v\n", err)
		return nil, errors.New("Invalid ObjectID format")
	}

	filter := bson.D{{"_id", objectID}}

	var result model.Article
	if err := db.Collection.FindOne(db.Ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
