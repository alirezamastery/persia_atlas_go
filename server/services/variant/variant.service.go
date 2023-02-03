package variantsrvc

import (
	"github.com/gin-gonic/gin"
	"persia_atlas/server/models"
	"persia_atlas/server/models/network"
	"persia_atlas/server/pagination"
)

type VariantService interface {
	CreateVariant(p *models.Variant) error
	GetVariantsPaginated(c *gin.Context) *pagination.PaginatedResponse
	GetVariantById(id uint) *network.VariantSerializer
}
