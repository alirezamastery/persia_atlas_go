package brandsrvc

import (
	"github.com/gin-gonic/gin"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type BrandService interface {
	CreateBrand(brand *models.Brand) error
	GetBrandsPaginated(c *gin.Context) *pagination.PaginatedResponse
	GetBrandById(id uint) *models.Brand
	//UpdateBrand(id uint) *models.Brand
}
