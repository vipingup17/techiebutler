package main

import (
	"techiebutler/configs"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin router
	router := gin.Default()

	// Connect to the database
	configs.InitDB()

	// Initialise routes
	InitRoutes(router)

	// Start the server
	router.Run(":8080")
}
