package api

import (
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirkuttin/mqtt"
	"github.com/Sirupsen/logrus"
	"io"
	"bytes"
)

var (
	log        logrus.Logger
	mqttClient mqtt.Client
)

func Start(incomingMqttClient mqtt.Client, logger *logrus.Logger) {

	log = *logger
	mqttClient = incomingMqttClient

	log.Info("Starting API")

	router := mux.NewRouter()

	router.HandleFunc("/alert", sendAlert()).Methods("POST")
	router.HandleFunc("/weather", sendWeather()).Methods("POST")

	err := http.ListenAndServe(":8000", handlers.CORS(createCorsOptions()...)(router))

	if err != nil {
		panic(err.Error())
	}
}

func createCorsOptions() (corsOptions []handlers.CORSOption) {
	corsOptions = append(corsOptions, handlers.AllowedHeaders([]string{"X-Requested-With"}))
	corsOptions = append(corsOptions, handlers.AllowedOrigins([]string{"*"}))
	corsOptions = append(corsOptions, handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}))
	return
}

func GetPayloadBytes(closer io.ReadCloser) ([]byte) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(closer)
	return buf.Bytes()
}
