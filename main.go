package main

import (
	"Tamer/controller"
	"Tamer/database"
	"Tamer/model"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Printf("Load dotenv error: %v\n", err)
	}

	client := database.DBInit()
	collection := client.Database("Tamer").Collection("posts")
	ctx := context.Background()
	defer func(client *mongo.Client, ctx context.Context) {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Collection disconnection error: %v\n", err)
		}
	}(client, ctx)

	db := &model.DBCollection{
		Collection: collection,
		Ctx:        ctx,
	}

	database.DBInit()
	controller.StartServer(db)
}
