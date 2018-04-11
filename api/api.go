package api

import (
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/sirkuttin/mqtt"
	"github.com/sirkuttin/cloud_api_to_edge_mqtt/api/routes"
)

func Start(mqttClient mqtt.Mqtt){
	fmt.Println("Starting API")

	router := mux.NewRouter()

	router.HandleFunc("/alert", routes.SendAlert(mqttClient)).Methods("POST")
	router.HandleFunc("/weather", routes.SendWeather(mqttClient)).Methods("POST")

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

