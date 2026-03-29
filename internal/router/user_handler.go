package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trmyxx/AuthService/internal/model"
)

func (router *Router) Signup(c *gin.Context) {
	//Get email/pass off req body
	var body model.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	err := router.service.Signup(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (router *Router) Login(c *gin.Context) {
	//Get email/pass off req body
	var body model.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	tokenString, err := router.service.Login(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Autorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func (router *Router) Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
