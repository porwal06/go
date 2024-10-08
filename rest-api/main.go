package main

import (
	"example.com/rest-api/db"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default() // Default returns a gin router with default middleware
	routes.RouteList(r)

	// Make 8080 port free, stop apache server or anyother http server if already running.
	r.Run(":9001") // listen and serve on 0.0.0.0:8080

}
