package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client  *mongo.Client
	Ctx     context.Context
	Options MongoOptions
}

type MongoOptions struct {
	database string
	url      string
}

func newMongoConnection() *Mongo {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	mongoOptions := getMongoOptionsFromEnv()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoOptions.url))
	if err != nil {
		panic(err)
	}

	log.Printf("App is connected to MongoDB !")
	return &Mongo{client, ctx, mongoOptions}
}

func getMongoOptionsFromEnv() MongoOptions {
	url, declared := os.LookupEnv("MONGO_URL")
	if !declared {
		url = "mongodb://localhost:27017"
	}

	database, declared := os.LookupEnv("MONGO_DATABASE")
	if !declared {
		database = "csv-to-db"
	}

	return MongoOptions{database, url}
}

func (m *Mongo) getCollection(name string) *mongo.Collection {
	return m.Client.Database(m.Options.database).Collection(name)
}
