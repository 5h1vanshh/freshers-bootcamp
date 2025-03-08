package main

import (
	config "freshers-bootcamp/databaseConnect"
	"freshers-bootcamp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	routes.ProductRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
