package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type DeviceData struct {
	Name        string  `json:"deviceName"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

var (
	device DeviceData
	mutex  sync.RWMutex
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
		var newData DeviceData

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

func checkFileExistance(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		return false, fmt.Errorf("%s does not exist", path)
	}

	return true, nil
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/api/user/device", deviceData)

	http.ListenAndServe(":"+port, nil)
	fmt.Printf("server running on port %s\n", port)
}