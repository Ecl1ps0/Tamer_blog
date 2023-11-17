package services

import (
	"Tamer/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetAllPosts(db *model.DBCollection) ([]model.Article, error) {
	filter := bson.D{}

	cursor, err := db.Collection.Find(db.Ctx, filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Printf("Close cursor error: %v\n", err)
		}
	}(cursor, db.Ctx)

	var articles []model.Article
	if err := cursor.All(db.Ctx, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}
