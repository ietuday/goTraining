package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Public
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Public read
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Protected write
	protected := server.Group("/")
	protected.Use(middlewares.Authenticate())
	protected.POST("/events", createEvent)
	protected.POST("/events/:id/register", registerForEvent)


	// Owner-only update/delete
	owner := protected.Group("/")
	owner.Use(middlewares.RequireEventOwner())
	owner.PUT("/events/:id", updateEvent)
	owner.DELETE("/events/:id", deleteEvent)
	owner.GET("/events/:id/registrations", getEventRegistrations)
	owner.DELETE("/events/:id/registrations", unregisterForEvent)

}
