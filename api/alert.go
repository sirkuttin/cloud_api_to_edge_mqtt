package api

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/sirkuttin/edge_vehicle_data"
	"encoding/binary"
	"bytes"
)

func sendAlert() http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		bodyBytes := GetPayloadBytes(request.Body)
		log.Debug("alert post body = ", string(bodyBytes))

		var alert data.Alert
		json.Unmarshal(bodyBytes, &alert)
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
