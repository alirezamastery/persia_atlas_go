package variantcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	vs "persia_atlas/server/services/variant"
)

type VariantController struct {
	VariantService vs.VariantService
	DB             *gorm.DB
}

func (vc *VariantController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(vc.DB))
	{
		routes.POST("/variants/", vc.CreateVariant())
		routes.GET("/variants/", vc.GetVariantsPaginated())
		routes.GET("/variants/:id/", vc.GetVariantById())
		routes.PATCH("/variants/:id/", vc.UpdateVariant())
	}
}

func NewVariantController(vs vs.VariantService, db *gorm.DB) *VariantController {
	return &VariantController{
		VariantService: vs,
		DB:             db,
	}
}
