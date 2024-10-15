package models

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server-mqtt/db"
	"time"
)

type TaskModel struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name  string             `bson:"name" json:"name"`
	Value string             `bson:"value" json:"value"`
	Date  string             `bson:"date" json:"date"`
}

func (r TaskModel) TableName() string {
	return "task"
}

func (r TaskModel) Insert(data TaskModel) (*mongo.InsertOneResult, error) {

	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.Date = time.Now().Format("2006-01-02 15:04:05")
	}
	println("isi date ", r.Date)
	return r.Collection().InsertOne(r.ctx(), data)
}

func (r TaskModel) FindAll(filter any, findOptions *options.FindOptions) []TaskModel {
	if filter == nil {
		filter = gin.H{}
	}
	c, err := r.Collection().Find(r.ctx(), filter, findOptions)
	if err != nil {
		println("error ", err.Error())
		return nil
	}
	var result []TaskModel
	err = c.All(r.ctx(), &result)
	if err != nil {
		println("error all ", err.Error())
	}
	return result
}

func (r TaskModel) ctx() context.Context {
	_ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	return _ctx
}

func (r TaskModel) Collection() *mongo.Collection {
	return db.DB().Collection(r.TableName())
}
