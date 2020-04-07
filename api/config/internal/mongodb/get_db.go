package mongodb

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// GetDB : get mongo database
func GetDB(ctx context.Context) *mongo.Database {

	if client == nil {
		client = getClient(ctx)
	}

	if err := client.Connect(ctx); err != nil {
		panic(err)
	}

	db := client.Database(os.Getenv("MONGO_DB_NAME"))
	if db == nil {
		panic(fmt.Errorf("could not get database '%s'", os.Getenv("MONGO_DB_NAME")))
	}

	return db
}

func getClient(ctx context.Context) *mongo.Client {
	mongoHost := fmt.Sprintf("mongodb://%s:%s", os.Getenv("MONGO_DB_HOST"), os.Getenv("MONGO_DB_PORT"))
	credential := options.Credential{
		Username: os.Getenv("MONGO_DB_USER"),
		Password: os.Getenv("MONGO_DB_PASSWORD"),
	}
	clientOpts := options.Client().ApplyURI(mongoHost).SetAuth(credential)
	mongoClient, err := mongo.NewClient(clientOpts)
	if err != nil {
		panic(err)
	}
	return mongoClient
}
