package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/users")
	{
		routes.POST("/signup", h.Signup())
		routes.POST("/login", h.Login())
		routes.PATCH("/profile", middlewares.RequireAuth(db), h.UpdateProfile())
		routes.GET("/users", middlewares.RequireAuth(db), h.GetAllUsers())
	}
}
