package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/otel"
	"log"
	"redfox/model"
)

var tracer = otel.Tracer("db/users")

func CreateUser(creatingUser model.User, tCtx context.Context) model.User {
	tCtx, span := tracer.Start(tCtx, "CreateUser")
	defer span.End()

	client = GetClient()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	col := client.Database("redfox").Collection("users")

	result, err := col.InsertOne(ctx, creatingUser)
	if err != nil {
		cancel()
		// TODO: [E11000 duplicate key error collection
		panic(err)
	}

	if id, success := result.InsertedID.(primitive.ObjectID); success {
		creatingUser.Id = id
	} else {
		// TODO: What if id is missing ??
	}

	return creatingUser
}

func GetUser(tCtx context.Context) model.User {
	tCtx, span := tracer.Start(tCtx, "GetUser")
	defer span.End()

	client = GetClient()
	col := client.Database("redfox").Collection("users")

	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	err := col.FindOne(ctx, bson.D{}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// TODO: How to handle nil properly?
		return model.User{}
	}
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	//jsonData, err := json.MarshalIndent(user, "", " ")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%s\n", jsonData)

	return user
}
