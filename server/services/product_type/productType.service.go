package ptypesrvc

import (
	"github.com/gin-gonic/gin"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type ProductTypeService interface {
	CreateProductType(brand *models.ProductType) error
	GetProductTypesPaginated(c *gin.Context) *pagination.PaginatedResponse
	GetProductTypeById(id uint) *models.ProductType
}
