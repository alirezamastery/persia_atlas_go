package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
)

func (h handler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get("user")
		if !user.(models.User).IsAdmin {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var userList []models.User
		h.DB.Preload("Profile").Find(&userList)
		c.JSON(http.StatusOK, &userList)
	}
}
