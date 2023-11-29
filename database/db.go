package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func DBInit(errorLogger *log.Logger, infoLogger *log.Logger) *mongo.Client {
	password := os.Getenv("DB_PASS")
	uri := fmt.Sprintf("mongodb+srv://tamer:%s@tamer.ci0zyf2.mongodb.net/?retryWrites=true&w=majority", password)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		errorLogger.Fatalf("Database connection error: %v\n", err)
	}

	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{"ping", 1}}).Err(); err != nil {
		errorLogger.Fatalf("Ping to database error: %v\n", err)
	}

	infoLogger.Printf("Connection is succeed!")

	return client
}
