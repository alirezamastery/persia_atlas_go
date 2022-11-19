package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"persia_atlas/server/auth"
	"persia_atlas/server/models"
)

func RequireAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := auth.ParseToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
			var user models.User
			db.Preload("Profile").First(&user, "id = ?", claims["user_id"])
			if user.ID == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func WsAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		protocols := c.Request.Header.Get("Sec-Websocket-Protocol")
		parts := strings.Split(protocols, ", ")
		if len(parts) > 0 {
			tokenString := parts[0]
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("API_SECRET")), nil
			})
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if float64(time.Now().Unix()) > claims["exp"].(float64) {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
				var user models.User
				db.Preload("Profile").First(&user, "id = ?", claims["user_id"])
				if user.ID == 0 {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
				c.Set("user", user)
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
