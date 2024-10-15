package helper

import (
	"crypto/tls"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"os"
)

var _client *mqtt.Client

func Client() *mqtt.Client {
	return _client
}

func MQTT_connect() *mqtt.Client {
	domain := os.Getenv("MQTT_URL")
	port := os.Getenv("MQTT_PORT")
	address := fmt.Sprintf("ssl://%s:%s", domain, port)

	opts := mqtt.NewClientOptions().AddBroker(address)
	opts.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	opts.SetPassword(os.Getenv("MQQT_CLIENT_PASSWORD"))
	opts.SetUsername(os.Getenv("MQQT_CLIENT_USERNAME"))
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         tls.NoClientCert,
	}
	opts.SetTLSConfig(tlsConfig)
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected to MQTT broker")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Println("Disconnected from MQTT broker")
	}

	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Println("Connect mqqt : ", token.Error())
		return nil
	}
	_client = &client
	return _client
}
