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

type SensorModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Sensor string             `bson:"sensor" json:"sensor"`
	Value  float64            `bson:"value" json:"value"`
	Date   string             `bson:"date" json:"date"`
}

func (SensorModel) TableName() string {
	return "sensor"
}

func (r SensorModel) Filter() {

}

func (r SensorModel) FindAll(filter any, findOptions *options.FindOptions) []SensorModel {
	if filter == nil {
		filter = gin.H{}
	}
	c, err := r.Collection().Find(r.ctx(), filter, findOptions)
	if err != nil {
		println("error ", err.Error())
		return nil
	}
	var result []SensorModel
	err = c.All(r.ctx(), &result)
	if err != nil {
		println("error all ", err.Error())
	}
	return result
}

func (SensorModel) ctx() context.Context {
	_ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	return _ctx
}

func (r SensorModel) Collection() *mongo.Collection {
	return db.DB().Collection(r.TableName())
}

func (r SensorModel) Insert(data SensorModel) (*mongo.InsertOneResult, error) {

	if data.ID.IsZero() {
		r.ID = primitive.NewObjectID()
		r.Date = time.Now().Format("2006-01-02 15:04:05")
	}
	return r.Collection().InsertOne(r.ctx(), data)
}
