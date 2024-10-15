package main

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"server-mqtt/app/helper"
	"server-mqtt/db"
	"server-mqtt/routes"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if db.DB() != nil {
		println("mongodb connected")
	}

	app := gin.Default()
	go func() {
		routes.Routes(app)
		if err := app.Run(":8000"); err != nil {
			panic(err)
		}
	}()

	clientChannel := make(chan mqtt.Client)
	go func() {
		client := helper.MQTT_connect()
		if client != nil {
			routes.MqqtRoutes(*client)
			clientChannel <- *client
		} else {
			close(clientChannel)
		}
	}()

	// Menangani graceful shutdown dengan sinyal sistem
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	client, ok := <-clientChannel
	if ok {
		client.Disconnect(2500)
	}

	// Mematikan server dengan timeout
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Server exiting")

}
