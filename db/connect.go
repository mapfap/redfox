package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var client *mongo.Client

const timeout = 10 * time.Second

func Connect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Missing 'MONGODB_URI' environmental variable.")
	}

	c, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	client = c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	err = client.Connect(ctx)
	if err != nil {
		// TODO: Is this correct?
		cancel()
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	//defer func() {
	//	if err := client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	fmt.Println("[db] connected")
	return client
}

func GetClient() *mongo.Client {
	if client == nil {
		// TODO: Need to handle panic here
		Connect()
	}
	return client
}
