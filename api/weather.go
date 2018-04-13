package api

import (
	"net/http"
	"github.com/sirkuttin/edge_vehicle_data"
	"encoding/json"
	"fmt"
	"bytes"
	"encoding/binary"
)

func sendWeather() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		bodyBytes := GetPayloadBytes(request.Body)
		log.Debug("weather post body = ", string(bodyBytes))

		var weather data.Weather
		json.Unmarshal(bodyBytes,&weather)
		err := weather.Validate()
		if err != nil {
			fmt.Fprintf(responseWriter, err.Error())
			return
		}

		buf := &bytes.Buffer{}
		err = binary.Write(buf, binary.LittleEndian, weather)
		if err != nil {
			panic(err)
		}

		mqttClient.PublishToTopic("weather", buf.Bytes())
	}
}
