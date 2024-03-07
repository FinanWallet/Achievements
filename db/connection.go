package db

import (
	"context"
	"fmt"
	"log"
	"os"

    "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database *mongo.Database
var achievementsCollection *mongo.Collection
var userAchievementsCollection *mongo.Collection

func DBConnection() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("No MONGO_URI in .env file")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    }()

    database = client.Database("achievements")
    achievementsCollection = database.Collection("achievements")
    userAchievementsCollection = database.Collection("user_achievements")

    if err := client.Ping(context.TODO(), nil); err != nil {
        panic(err)
    }

    fmt.Println("Connected to MongoDB!")
}
