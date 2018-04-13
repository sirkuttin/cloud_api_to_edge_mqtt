package main

import (
	"github.com/sirkuttin/cloud_api_to_edge_mqtt/api"
	"github.com/sirkuttin/mqtt"
	"errors"
	"os"
	"github.com/Sirupsen/logrus"
)

var log = logrus.New();

func init() {
	log.SetLevel(logrus.DebugLevel)
}

func main() {

	mqttClient, err := mqtt.New("tcp://127.0.0.1:1884", "cloud-api")
	defer mqttClient.Disconnect(2000)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	errChan := make(chan error)

	go func() {
		api.Start(mqttClient, log)
		errChan <- errors.New("api exited")
	}()

	log.Error("error: ", <- errChan)
	os.Exit(1)
}
