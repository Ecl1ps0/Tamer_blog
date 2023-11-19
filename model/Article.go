package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Title        string             `bson:"title"`
	TextContent  string             `bson:"textContent"`
	ImageContent string             `bson:"imageContent"`
	CreateAt     uint64             `bson:"createAt"`
	UpdatedAt    uint64             `bson:"updatedAt"`
}
