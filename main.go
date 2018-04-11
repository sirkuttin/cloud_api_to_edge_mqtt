package main

import (
	"github.com/sirkuttin/cloud_api_to_edge_mqtt/api"
	"github.com/sirkuttin/mqtt"
	"fmt"
	"errors"
	"os"
)

func main() {

	mqttClient, err := mqtt.New("tcp://127.0.0.1:1883", "cloud-api")
	defer mqttClient.Disconnect(2000)
	if err != nil {
		panic(err)
	}

	errChan := make(chan error)

	go func() {
		api.Start(mqttClient)
		errChan <- errors.New("api exited")
	}()

	fmt.Println("error:", <- errChan)
	os.Exit(1)
}
