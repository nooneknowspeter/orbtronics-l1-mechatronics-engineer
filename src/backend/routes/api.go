package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nooneknows/orbtronics-l1-mechatronics-engineer/models"
	"sync"
)

var (
	mutex  sync.RWMutex
	device models.DeviceData
)

func deviceData(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	switch request.Method {
	case http.MethodGet:
		mutex.RLock()
		defer mutex.RUnlock()
		json.NewEncoder(writer).Encode(device)

		fmt.Println(request.Method, request.Header["User-Agent"])

	case http.MethodPost:
		var newData models.DeviceData

		if err := json.NewDecoder(request.Body).Decode(&newData); err != nil {
			http.Error(writer, "invalid JSON", http.StatusBadRequest)
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		device = newData

		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(device)

		fmt.Println(request.Method, request.Header["User-Agent"], device)

	default:
		http.Error(writer, "invalid method", http.StatusMethodNotAllowed)
	}
}

func RouteDeviceDataHandler() {
	http.HandleFunc("/api/user/device", deviceData)
}
