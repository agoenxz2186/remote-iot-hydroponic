package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server-mqtt/app/models"
	"strconv"
	"strings"
)

type SensorController struct {
}

func NewSensorController() *SensorController {
	return &SensorController{}
}

func (receiver *SensorController) Index(c *gin.Context) {
	page := c.Query("page")
	sensor := c.Query("sensor")

	ipage, e := strconv.Atoi(page)
	if e != nil {
		ipage = 1
	}

	skip := (ipage - 1) * 10
	findOpt := options.Find()
	findOpt.SetSort(bson.D{{"date", -1}})
	findOpt.SetLimit(10)
	findOpt.SetSkip(int64(skip))

	filter := gin.H{}
	if len(strings.Trim(sensor, " ")) > 1 {
		filter = gin.H{
			"sensor": gin.H{
				"$regex":   sensor,
				"$options": "i",
			},
		}
	}

	c.JSON(200, gin.H{
		"data": models.SensorModel{}.FindAll(filter, findOpt),
	})
}

func (receiver *SensorController) Show(c *gin.Context) {
	idparam := c.Param("id")

	id, _ := primitive.ObjectIDFromHex(idparam)

	c.JSON(200, gin.H{
		"data": models.SensorModel{}.FindAll(gin.H{"_id": id}, options.Find()),
	})
}
