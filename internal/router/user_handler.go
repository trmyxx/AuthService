package router

import (
	"net/http"

	"github.com/trmyxx/auth/model"

	"github.com/gin-gonic/gin"
)

func (router *Router) postUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := router.service.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": savedUser.ID})
}
