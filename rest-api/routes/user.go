package routes

import (
	"net/http"

	"example.com/rest-api/modules"
	"github.com/gin-gonic/gin"
)

// User struct with id, email, password
func registerUser(context *gin.Context) {
	var user modules.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Store user in database
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func loginUser(context *gin.Context) {
	var user modules.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = user.Login()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
