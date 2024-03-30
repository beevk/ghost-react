package main

import (
	"github.com/joho/godotenv"
	"log"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	apiPort := GetEnv("API_PORT", "3000")
	port := ":" + apiPort
	server := NewAPIServer(port)
	server.RunServer()
}
