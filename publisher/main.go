package main

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var broker string
	var clientID string
	var topicName string

	clientID = "goClient"
	broker = "tcp://localhost:1883"
	topicName = "merchant/android"

	uri, err := url.Parse(broker)
	if err != nil {
		log.Fatal(err)
	}

	connection := connect(clientID, uri)
	publish(connection, topicName)
}

func connect(clientID string, uri *url.URL) mqtt.Client {
	clientOptions := getMqttClientOptions(clientID, uri)
	client := mqtt.NewClient(clientOptions)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected Successfully!!!!")
	fmt.Println(strconv.FormatBool(client.IsConnected()))
	fmt.Println(strconv.FormatBool(client.IsConnectionOpen()))
	return client
}

func publish(client mqtt.Client, topic string) {
	// for i := 0; i < 10; i++ {
	// 	token := client.Publish(topic, 1, false, strconv.Itoa(i))
	// 	if err := token.Error(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("Published..")
	// }
	token := client.Publish(topic, 2, false, "Source : Golang")
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Published..")
}

/**
* method to return ClientOptions object
**/
func getMqttClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	clientOptions.SetClientID(clientID)
	clientOptions.SetCleanSession(false)
	clientOptions.SetUsername("admin")
	clientOptions.SetPassword("public")
	return clientOptions
}
