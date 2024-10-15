package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/modals"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	// c.JSON serializes the given struct as JSON and returns it to the client
	events, err := modals.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, events)
}

func saveEvent(c *gin.Context) {
	var event modals.Event
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
	event, err := modals.GetEventByID(id)
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
	var event modals.Event
	event, err = modals.GetEventByID(id)
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
	var event modals.Event
	event, err = modals.GetEventByID(id)
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
