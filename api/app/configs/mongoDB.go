package configs

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx context.Context
var Client *mongo.Client

func MongoConnection() {
	contextCus := context.Background()
	client, err := mongo.Connect(contextCus, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}
	fmt.Println("ติดต่อ MongoDB Server สำเร็จ")
	Ctx = contextCus
	Client = client
}

func GetMongoCollection(collection string) *mongo.Collection {
	coll := Client.Database(os.Getenv("MONGO_DATABASE")).Collection(collection)
	return coll
}
