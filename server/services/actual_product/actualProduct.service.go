package apsrvc

import (
	"github.com/gin-gonic/gin"
	"persia_atlas/server/models"
	"persia_atlas/server/pagination"
)

type ActualProductService interface {
	CreateActualProduct(ap *models.ActualProduct) error
	GetActualProductsPaginated(c *gin.Context) *pagination.PaginatedResponse
	GetActualProductById(id uint) *models.ActualProduct
}
