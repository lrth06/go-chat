package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnDB(collection string) *mongo.Collection {
	env, err := GetConfig()
	if err != nil {
		panic(err)
	}
	mongoUri := env.MongoURI
	dbName := env.DBName
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	collectionReference := client.Database(dbName).Collection(collection)
	return collectionReference

}
