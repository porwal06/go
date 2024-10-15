package routes

import (
	"net/http"

	"example.com/rest-api/modals"
	"github.com/gin-gonic/gin"
)

// User struct with id, email, password
func registerUser(context *gin.Context) {
	var user modals.User

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
	var user modals.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "user": user})
		return
	}
	err = user.Login()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login user", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
