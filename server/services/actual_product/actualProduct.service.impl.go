package apsrvc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type ActualProductServiceImpl struct {
	db *gorm.DB
}

func NewActualProductService(db *gorm.DB) ActualProductService {
	return &ActualProductServiceImpl{
		db: db,
	}
}

func (aps ActualProductServiceImpl) CreateActualProduct(ap *models.ActualProduct) error {
	res := aps.db.Create(ap)
	return res.Error
}

func (aps ActualProductServiceImpl) GetActualProductsPaginated(c *gin.Context) *pagination.PaginatedResponse {
	var actualProducts []models.ActualProduct
	var scopes []func(db *gorm.DB) *gorm.DB

	if titleQuery := c.Query("title"); titleQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("search ILIKE ?", "%"+titleQuery+"%")
		})
	}

	paginator := pagination.NewPageNumberPaginator(c, aps.db, scopes, &actualProducts)
	response := paginator.GetPaginatedResponse()

	return response
}

func (aps ActualProductServiceImpl) GetActualProductById(id uint) *models.ActualProduct {
	var ap models.ActualProduct
	aps.db.Joins("Brand").First(&ap, id)
	if ap.ID == 0 {
		return nil
	}
	return &ap
}
