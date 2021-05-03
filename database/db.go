package database

import (
	"alumnus/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var err error
var db *mongo.Database
var Users *mongo.Collection
var Ctx context.Context
var Cancel context.CancelFunc

func ConnectToDB(callback func(users *mongo.Collection)) {
	client, err = mongo.NewClient(options.Client().ApplyURI(config.Env.DBURL))
	Ctx, Cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer Cancel()
	err = client.Connect(Ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(Ctx)

	err = client.Ping(Ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(config.Env.DBNAME)
	Users = db.Collection("Users")
	callback(Users)
}
