package services

import (
	"Tamer/model"
	"go.mongodb.org/mongo-driver/bson"
)

func FindBYId(id string, db *model.DBCollection) (*model.Article, error) {
	filter := bson.D{{"_id", id}}

	var result model.Article
	err := db.Collection.FindOne(db.Ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
