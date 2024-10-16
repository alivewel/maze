package main

import (
	"fmt"
	"os"

	"maze/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/solve-maze", handlers.SolveMazeHandler)
	router.POST("/generate-maze", handlers.GenerateMazeHandler)

	port := os.Getenv("SERVERPORT")
	if port == "" {
		port = "8080"
	}

	url := os.Getenv("URL")
	if url == "" {
		url = "localhost"
	}

	fmt.Printf("Сервер запущен на %s:%s\n", url, port)
	router.Run(url + ":" + port)
}
