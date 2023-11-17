package model

type Article struct {
	ID           uint64 `bson:"_id, omitempty"`
	Title        string `bson:"title"`
	TextContent  string `bson:"textContent"`
	ImageContent string `bson:"imageContent"`
	CreateAt     uint64 `bson:"createAt"`
	UpdatedAt    uint64 `bson:"updatedAt"`
}
