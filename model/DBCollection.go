package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBCollection struct {
	Collection *mongo.Collection
	Ctx        context.Context
}
