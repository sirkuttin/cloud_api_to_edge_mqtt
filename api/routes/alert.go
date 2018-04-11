package routes

import (
	"github.com/sirkuttin/mqtt"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/sirkuttin/edge_vehicle_data"
	"encoding/binary"
	"bytes"
)

func SendAlert(mqttClient mqtt.Mqtt) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		body_bytes := GetPayloadBytes(request.Body)
		var alert data.Alert
		json.Unmarshal(body_bytes,&alert)
		err := alert.Validate()
		if err != nil {
			fmt.Fprintf(responseWriter, err.Error())
			return
		}

		buf := &bytes.Buffer{}
		err = binary.Write(buf, binary.LittleEndian, alert)
		if err != nil {
			panic(err)
		}

		mqttClient.PublishToTopic("vehicle-alert", buf.Bytes())
	}
}
