package routes

import (
	"fmt"
	"net/http"
	"os"
)

func SetupServer(serverPort string) {
	RouteDeviceDataHandler()

	fmt.Printf("server running on port %s\n", serverPort)
	http.ListenAndServe(":"+serverPort, nil)
}

func StartServer() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	SetupServer(port)
}