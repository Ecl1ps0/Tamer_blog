package services

import (
	"Tamer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateArticle(id string, updatedArticle *model.Article, db *model.DBCollection) (*mongo.UpdateResult, error) {
	filter := bson.D{{"_id", id}}
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
