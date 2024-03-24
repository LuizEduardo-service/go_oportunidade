package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// intitialize server
	router := gin.Default()

	//initialize routes
	initializeRoutes(router)

	//run the server
	router.Run(":8080")
}
