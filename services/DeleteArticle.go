package services

import (
	"Tamer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteArticle(id uint64, db *model.DBCollection) (*mongo.DeleteResult, error) {
	filter := bson.D{{"_id", id}}

	result, err := db.Collection.DeleteOne(db.Ctx, filter)
	if err != nil {
		return result, err
	}

	return result, nil
}
