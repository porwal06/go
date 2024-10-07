package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/modules"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default() // Default returns a gin router with default middleware
	// GET is a shortcut for router.Handle("GET", path, handle)
	r.GET("/events", getEvents)
	r.POST("/events", saveEvent)

	// Make 8080 port free, stop apache server or anyother http server if already running.
	r.Run(":9001") // listen and serve on 0.0.0.0:8080

}

func getEvents(c *gin.Context) {
	// c.JSON serializes the given struct as JSON and returns it to the client
	c.JSON(http.StatusOK, modules.GetAllEvents())
}

func saveEvent(c *gin.Context) {
	var event modules.Event
	// c.BindJSON deserializes the JSON body and stores it in the value pointed to by event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event successfully created", "event": event})
}
