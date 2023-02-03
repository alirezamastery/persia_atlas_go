package productcontroller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/middlewares"
	ps "persia_atlas/server/services/product"
)

type ProductController struct {
	ProductService ps.ProductService
	DB             *gorm.DB
}

func (pc *ProductController) RegisterRoutes(r *gin.Engine) {
	routes := r.Group("/api/products")
	routes.Use(middlewares.RequireAuth(pc.DB))
	{
		routes.POST("/products/", pc.CreateProduct())
		routes.GET("/products/", pc.GetProductsPaginated())
		routes.GET("/products/:id/", pc.GetProductById())
		routes.PATCH("/products/:id/", pc.UpdateProduct())
	}
}

func NewProductController(ps ps.ProductService, db *gorm.DB) *ProductController {
	return &ProductController{
		ProductService: ps,
		DB:             db,
	}
}
