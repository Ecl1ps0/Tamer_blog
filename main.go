package main

import (
	"Tamer/controller"
	"Tamer/database"
	"Tamer/model"
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	if err := godotenv.Load("config/.env"); err != nil {
		errorLogger.Fatalf("Load dotenv error: %v\n", err)
	}

	client := database.DBInit(errorLogger, infoLogger)
	collection := client.Database("Tamer").Collection("posts")
	ctx := context.Background()
	defer func(client *mongo.Client, ctx context.Context) {
		if err := client.Disconnect(ctx); err != nil {
			errorLogger.Fatalf("Collection disconnection error: %v\n", err)
		}
	}(client, ctx)

	db := &model.DBCollection{
		Collection: collection,
		Ctx:        ctx,
	}

	controller.StartServer(db, errorLogger)
}
