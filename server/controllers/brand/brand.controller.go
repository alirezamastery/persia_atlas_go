package brandcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	bs "persia_atlas/server/services/brand"
)

type BrandController struct {
	BrandService bs.BrandService
	DB           *gorm.DB
}

func (bc *BrandController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(bc.DB))
	{
		routes.POST("/brands", bc.CreateBrand())
		routes.GET("/brands", bc.GetBrands())
		routes.GET("/brands/:id", bc.GetBrandById())
		routes.PATCH("/brands/:id", bc.UpdateBrand())
	}
}

func NewBrandController(bs bs.BrandService, db *gorm.DB) *BrandController {
	return &BrandController{
		BrandService: bs,
		DB:           db,
	}
}
