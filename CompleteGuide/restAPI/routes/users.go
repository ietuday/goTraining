package routes

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"example.com/rest-api/auth"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	// Validate JSON + binding tags (required/email/min=6 etc.)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid signup payload"})
		return
	}

	err := user.Save()
	if err != nil {
		// SQLite unique constraint error usually contains "unique"
		if strings.Contains(strings.ToLower(err.Error()), "unique") {
			c.JSON(http.StatusConflict, gin.H{"message": "email already exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "signup successful",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}

func login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}

	err := user.ValidateCredential()
	if err != nil {
		// invalid email/password
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
			return
		}

		// other DB errors
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not login"})
		return
	}

	token, err := auth.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not generate token"})
		return
	}

	// success
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})

}
