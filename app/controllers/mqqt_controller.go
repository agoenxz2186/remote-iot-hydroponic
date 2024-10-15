package controllers

import (
	"encoding/json"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"server-mqtt/app/models"
)

type MQQTController struct {
}

func NewMQQTController() *MQQTController {
	return &MQQTController{}
}

func (receiver *MQQTController) Sensor(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	var out models.SensorModel
	json.Unmarshal([]byte(payload), &out)

	_, err := models.SensorModel{}.Insert(out)
	if err != nil {
		println("Insert error : ", err.Error())
	}
}

func (receiver *MQQTController) Handler(client mqtt.Client, msg mqtt.Message) {

}
