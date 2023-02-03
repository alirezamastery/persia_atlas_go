package apcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	apsrvc "persia_atlas/server/services/actual_product"
)

type ActualProductController struct {
	ActualProductService apsrvc.ActualProductService
	DB                   *gorm.DB
}

func (apc *ActualProductController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(apc.DB))
	{
		routes.POST("/actual-products/", apc.CreateActualProduct())
		routes.GET("/actual-products/", apc.GetActualProductPaginated())
		routes.GET("/actual-products/:id/", apc.GetActualProductById())
		routes.PATCH("/actual-products/:id/", apc.UpdateActualProduct())
	}
}

func NewActualProductController(
	aps apsrvc.ActualProductService,
	db *gorm.DB,
) *ActualProductController {
	return &ActualProductController{
		ActualProductService: aps,
		DB:                   db,
	}
}
