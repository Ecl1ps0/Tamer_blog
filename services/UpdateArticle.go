package services

import (
	"Tamer/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func UpdateArticle(id string, updatedArticle *model.Article, db *model.DBCollection) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ObjectID format: %v\n", err)
		return nil, errors.New("Invalid ObjectID format")
	}

	filter := bson.D{{"_id", objectID}}
	update := bson.D{{"$set", bson.D{
		{"title", updatedArticle.Title},
		{"textContent", updatedArticle.TextContent},
		{"imageContent", updatedArticle.ImageContent},
		{"updatedAt", updatedArticle.UpdatedAt},
	}}}

	result, err := db.Collection.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}
