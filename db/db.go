package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var _db *mongo.Database
var _ctx *context.Context

func DB() *mongo.Database {
	if _db == nil {
		_db = connect()
	}
	return _db
}

func Context() context.Context {
	return *_ctx
}

func connect() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	clientopt := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientopt)
	if err != nil {
		println("mongodb.connect error ", err)
		return nil
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		println("mongodb.connect error ", err)
		return nil
	}
	db := client.Database("db_sensor")
	_ctx = &ctx
	return db
}
