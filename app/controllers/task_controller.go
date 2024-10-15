package controllers

import (
	"github.com/gin-gonic/gin"
	"server-mqtt/app/helper"
	"server-mqtt/app/models"
)

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (r *TaskController) Index(c *gin.Context) {

	c.JSON(200, gin.H{
		"data": models.TaskModel{}.FindAll(gin.H{}, nil),
	})
}

func (receiver *TaskController) Doanything(c *gin.Context) {
	var data models.TaskModel
	c.BindJSON(&data)

	token := (*helper.Client()).Publish("device", 2, false, helper.JsonEncoded(data))
	token.Wait()
	if token.Error() != nil {
		println("Publish error : ", token.Error())
		c.JSON(501, gin.H{
			"message": token.Error(),
		})
	} else {
		r, e := models.TaskModel{}.Insert(data)
		if e != nil {
			c.JSON(502, gin.H{
				"message": e.Error(),
			})
			return
		}
		c.JSON(200, r)
	}

}
