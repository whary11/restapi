package main

import (
	"fmt"
	"restapi/utils"

	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "5000"
	route := gin.Default()
	route.GET("/get_user/:id", controllers.HomeHandler)
	fmt.Println("Corriendo en esta ip: " + utils.GetOutboundIP() + ":" + port)
	route.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
