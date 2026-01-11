package middlewares

import (
	"database/sql"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

func RequireEventOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get logged-in userId from context (set by Authenticate middleware)
		uidAny, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			c.Abort()
			return
		}
		userID, ok := uidAny.(int64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			c.Abort()
			return
		}

		// parse event id from path
		eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil || eventID <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
			c.Abort()
			return
		}

		// verify ownership
		var ownerID sql.NullInt64
		err = db.DB.QueryRow(`SELECT user_id FROM events WHERE id = ?`, eventID).Scan(&ownerID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"message": "event not found"})
				c.Abort()
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not verify event owner"})
			c.Abort()
			return
		}

		if !ownerID.Valid || ownerID.Int64 != userID {
			c.JSON(http.StatusForbidden, gin.H{"message": "Not allowed"})
			c.Abort()
			return
		}

		// ok: user owns this event
		c.Next()
	}
}
