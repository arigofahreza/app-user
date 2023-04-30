package configs

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Collection

func MongoConfig() (*mongo.Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	uri := "mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	MongoClient = client.Database(os.Getenv("MONGO_DB")).Collection(os.Getenv("MONGO_USER_COLLECTION"))
	return client, nil
}
