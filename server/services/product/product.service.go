package productsrvc

import (
	"github.com/gin-gonic/gin"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type ProductService interface {
	CreateProduct(p *models.Product) error
	GetProductsPaginated(c *gin.Context) *pagination.PaginatedResponse
	GetProductById(id uint) *models.Product
}
