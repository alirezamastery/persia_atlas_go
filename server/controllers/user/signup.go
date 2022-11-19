package user

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"persia_atlas/server/models"
)

type SignupRequestBody struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h handler) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		body := SignupRequestBody{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		var user models.User
		h.DB.First(&user, "mobile = ?", body.Mobile)
		if user.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with this mobile already exists"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash"})
			return
		}
		user = models.User{Mobile: body.Mobile, Password: string(hash)}
		result := h.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}

		c.JSON(http.StatusOK, &user)
	}
}
