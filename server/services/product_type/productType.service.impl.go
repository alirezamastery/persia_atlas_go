package ptypesrvc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

func NewProductTypeService(db *gorm.DB) ProductTypeService {
	return &ProductTypeServiceImpl{
		db: db,
	}
}

type ProductTypeServiceImpl struct {
	db *gorm.DB
	//ctx context.Context
}

func (bs *ProductTypeServiceImpl) CreateProductType(pt *models.ProductType) error {
	res := bs.db.Create(pt)
	return res.Error
}

func (bs *ProductTypeServiceImpl) GetProductTypesPaginated(c *gin.Context) *pagination.PaginatedResponse {
	var productTypes []models.ProductType

	var scopes []func(*gorm.DB) *gorm.DB

	if titleQuery := c.Query("title"); titleQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ?", "%"+titleQuery+"%")
		})
	}

	paginator := pagination.NewPageNumberPaginator(c, bs.db, scopes, &productTypes)
	response := paginator.GetPaginatedResponse()

	return response
}

func (bs *ProductTypeServiceImpl) GetProductTypeById(id uint) *models.ProductType {
	var pt models.ProductType
	bs.db.First(&pt, id)
	if pt.ID == 0 {
		return nil
	}
	return &pt
}
