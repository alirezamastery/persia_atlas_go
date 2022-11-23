package brandcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (bc *BrandController) GetBrands() gin.HandlerFunc {
	return func(c *gin.Context) {
		brands := bc.BrandService.GetBrandsPaginated(c)
		c.JSON(http.StatusOK, &brands)
	}
}
