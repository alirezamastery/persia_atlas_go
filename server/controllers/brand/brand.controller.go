package brandcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	"persia_atlas/server/services/brand"
)

type BrandController struct {
	BrandService brandservice.BrandService
	DB           *gorm.DB
}

func (bc *BrandController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(bc.DB))
	{
		routes.POST("/brands", bc.CreateBrand())
		routes.GET("/brands", bc.GetBrands())
	}
}

func NewBrandController(bs brandservice.BrandService, db *gorm.DB) *BrandController {
	return &BrandController{
		BrandService: bs,
		DB:           db,
	}
}
