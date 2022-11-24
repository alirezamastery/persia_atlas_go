package brandcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (bc *BrandController) GetBrands() gin.HandlerFunc {
	return func(c *gin.Context) {
		brands := bc.BrandService.GetBrandsPaginated(c)
		c.JSON(http.StatusOK, &brands)
	}
}

func (bc *BrandController) GetBrandById() gin.HandlerFunc {
	return func(c *gin.Context) {
		brandID := c.Param("id")
		id, err := strconv.Atoi(brandID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid brand id"})
			return
		}
		brand := bc.BrandService.GetBrandById(uint(id))
		if brand == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no brand with this id"})
			return
		}
		c.JSON(http.StatusOK, &brand)
	}
}
