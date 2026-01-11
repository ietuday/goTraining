package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || eventID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	uidAny, ok := c.Get("userId") // set by Authenticate middleware
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userID := uidAny.(int64)

	err = models.RegisterForEvent(eventID, userID)
	if err != nil {
		if err.Error() == "already registered" {
			c.JSON(http.StatusConflict, gin.H{"message": "already registered"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not register"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered successfully"})
}


func unregisterForEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || eventID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	uidAny, ok := c.Get("userId") // set by Authenticate middleware
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}
	userID := uidAny.(int64)

	rows, err := models.UnregisterForEvent(eventID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not unregister"})
		return
	}

	if rows == 0 {
		// user was not registered (or event doesn't exist)
		c.JSON(http.StatusNotFound, gin.H{"message": "registration not found"})
		return
	}

	c.Status(http.StatusNoContent)
}


func getEventRegistrations(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || eventID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	users, err := models.GetRegistrationsByEventID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch registrations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"eventId":        eventID,
		"registrations":  users,
		"count":          len(users),
	})
}
