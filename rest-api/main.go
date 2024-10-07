package main

import (
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/modules"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default() // Default returns a gin router with default middleware
	// GET is a shortcut for router.Handle("GET", path, handle)
	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent) //"For dynamic param use ':' and some identifier"
	r.POST("/events", saveEvent)
	r.PUT("/events/:id", updateEvent)
	r.DELETE("/events/:id", deleteEvent)

	// Make 8080 port free, stop apache server or anyother http server if already running.
	r.Run(":9001") // listen and serve on 0.0.0.0:8080

}

func getEvents(c *gin.Context) {
	// c.JSON serializes the given struct as JSON and returns it to the client
	events, err := modules.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, events)
}

func saveEvent(c *gin.Context) {
	var event modules.Event
	// c.BindJSON deserializes the JSON body and stores it in the value pointed to by event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry couldn't save event, please try again later", "error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event successfully created", "events": event})
}

func getEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event, err := modules.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var event modules.Event
	event, err = modules.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	event.ID = id
	err = event.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry couldn't update event, please try again later", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event successfully updated", "events": event})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var event modules.Event
	event, err = modules.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Sorry couldn't delete event, please try again later", "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted"})
}
