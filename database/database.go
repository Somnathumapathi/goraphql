package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Somnathumapathi/goraphql/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string

// func init() {

// }

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionString = os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("CONNECTION_STRING is not set")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetEvents() []*model.Events {

}
