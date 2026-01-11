package routes

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil || eventID <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	event, err := models.GetEventByID(eventID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			context.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return

	}

	context.JSON(http.StatusOK, event)

}

func createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data."})
		return
	}

	// Get userId set by middleware
	uid, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	// middleware stored int64
	event.UserID = int(uid.(int64))

	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create events. Try again later.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}

func updateEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || eventID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	var updatedEvent models.Event
	if err := c.ShouldBindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventID

	rows, err := updatedEvent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again later."})
		return
	}

	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || eventID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	e := models.Event{ID: eventID}

	rows, err := e.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete event"})
		return
	}
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
		return
	}

	c.Status(http.StatusNoContent)
}



