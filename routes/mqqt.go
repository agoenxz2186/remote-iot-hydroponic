package routes

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"server-mqtt/app/controllers"
)

func MqqtRoutes(client mqtt.Client) {
	ctrl := controllers.NewMQQTController()
	if token := client.Subscribe("sensor", 2, ctrl.Sensor); token.Wait() && token.Error() != nil {
		fmt.Println("Subscribe error sensor, ", token.Error())
	}

}
