package main

import (
	"hexagonal-news-api/adapter/input/routes"
	"hexagonal-news-api/configuration/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger.Info("About to start application")
	router := gin.Default()

	routes.InitRoutes(router)

	if err := router.Run(":3333"); err != nil {
		logger.Error("Error trying to init application on port 3333", err)
		return
	}
}
