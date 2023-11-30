package services

import (
	"Tamer/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateArticle(article *model.Article, db *model.DBCollection) (*mongo.InsertOneResult, error) {
	result, err := db.Collection.InsertOne(db.Ctx, article)
	if err != nil {
		return result, err
	}

	return result, nil
}
