package ptypecontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	ptypesrvc "persia_atlas/server/services/product_type"
)

type ProductTypeController struct {
	ProductTypeService ptypesrvc.ProductTypeService
	DB                 *gorm.DB
}

func (ptc *ProductTypeController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(ptc.DB))
	{
		routes.POST("/product-types", ptc.CreateProductType())
		routes.GET("/product-types", ptc.GetProductTypesPaginated())
		routes.GET("/product-types/:id", ptc.GetProductTypeById())
		routes.PATCH("/product-types/:id", ptc.UpdateProductType())
	}
}

func NewProductTypeController(
	pts ptypesrvc.ProductTypeService,
	db *gorm.DB,
) *ProductTypeController {
	return &ProductTypeController{
		ProductTypeService: pts,
		DB:                 db,
	}
}
