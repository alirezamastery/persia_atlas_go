package brandservice

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type BrandServiceImpl struct {
	db *gorm.DB
	//ctx context.Context
}

func NewBrandService(db *gorm.DB) BrandService {
	return &BrandServiceImpl{
		db: db,
	}
}

func (bs *BrandServiceImpl) CreateBrand(brand *models.Brand) error {
	res := bs.db.Create(brand)
	return res.Error
}

func (bs *BrandServiceImpl) GetBrandsPaginated(c *gin.Context) *pagination.PaginatedResponse {
	var brands []models.Brand

	var scopes []func(*gorm.DB) *gorm.DB

	if titleQuery := c.Query("title"); titleQuery != "" {
		scopes = append(scopes, func(db *gorm.DB) *gorm.DB {
			return db.Where("title LIKE ?", "%"+titleQuery+"%")
		})
	}

	paginator := pagination.NewPageNumberPaginator(c, bs.db, scopes, &brands)
	response := paginator.GetPaginatedResponse()

	return response
}

func (bs *BrandServiceImpl) GetBrandById(id uint) *models.Brand {
	var brand models.Brand
	bs.db.First(&brand, id)
	if brand.ID == 0 {
		return nil
	}
	return &brand
}

func (bs *BrandServiceImpl) UpdateBrand(b *models.Brand, data map[string]any) *models.Brand {
	var brand models.Brand
	bs.db.First(&brand)
	if brand.ID == 0 {

	}
	return &brand
}
