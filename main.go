package main

import (
	"github.com/gin-gonic/gin"
	"github.com/whary11/restapi/controllers"
)

func main() {

	route := gin.Default()
	route.GET("/get_user/:id", controllers.HomeHandler)

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
