package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routers
	InitializeRouters(router)

	// Run the server
	router.Run(":8080")
}
