package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, nil)
	fmt.Printf("server running on port %s\n", port)
}
