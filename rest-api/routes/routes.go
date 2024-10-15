package routes

import "github.com/gin-gonic/gin"

func RouteList(r *gin.Engine) {
	// GET is a shortcut for router.Handle("GET", path, handle)
	r.GET("/events", getEvents)
	r.GET("/event/:id", getEvent) //"For dynamic param use ':' and some identifier"
	r.POST("/event", saveEvent)
	r.PUT("/event/:id", updateEvent)
	r.DELETE("/event/:id", deleteEvent)

	r.POST("/user/register", registerUser)
	r.POST("/user/login", loginUser)
}
