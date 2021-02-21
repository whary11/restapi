package main

import (
	"fmt"
	"restapi/utils"

	"restapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.GET("/get_user/:id", controllers.HomeHandler)
	fmt.Println(utils.GetOutboundIP())
	route.Run(":5000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
