package productsrvc

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type ProductServiceImpl struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &ProductServiceImpl{
		db: db,
	}
}

func (ps ProductServiceImpl) CreateProduct(p *models.Product) error {
	res := ps.db.Create(p)
	return res.Error
}

func (ps ProductServiceImpl) GetProductsPaginated(c *gin.Context) *pagination.PaginatedResponse {
	var products []models.ActualProduct
	var scopes []func(db *gorm.DB) *gorm.DB

	if searchQuery := c.Query("search"); searchQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.
				Where("title ILIKE ?", "%"+searchQuery+"%").
				Or("dkp = ?", searchQuery)
		})
	}
	if priceStepQuery := c.Query("price_step"); priceStepQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("price_step = ?", priceStepQuery)
		})
	}
	if isActiveQuery := c.Query("is_active"); isActiveQuery != "" {
		var isActive bool
		if isActiveQuery == "true" {
			isActive = true
		} else {
			isActive = false
		}
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", isActive)
		})
	}

	paginator := pagination.NewPageNumberPaginator(c, ps.db, scopes, &products)
	response := paginator.GetPaginatedResponse()

	return response
}

func (ps ProductServiceImpl) GetProductById(id uint) *models.Product {
	var p models.Product
	ps.db.Joins("Type").First(&p, id)
	if p.ID == 0 {
		return nil
	}
	return &p
}
