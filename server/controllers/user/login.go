package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"persia_atlas/server/auth"
	"persia_atlas/server/models"
)

type LoginRequestBody struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := LoginRequestBody{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "mobile and password required"})
			return
		}

		var user models.User
		h.DB.First(&user, "mobile = ?", body.Mobile)
		if user.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
			return
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
			return
		}

		tokenPair, err := auth.GenerateToken(1)
		if err != nil {
			fmt.Println("t e", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
			return
		}
		c.JSON(http.StatusCreated, &tokenPair)
	}
}
