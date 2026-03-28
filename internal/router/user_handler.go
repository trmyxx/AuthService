package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *Router) postUser(c *gin.Context) {
	//Get email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	savedUser, err := router.service.Signup(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": savedUser.ID})
}
