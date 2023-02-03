package variantcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (vc *VariantController) GetVariantsPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		variants := vc.VariantService.GetVariantsPaginated(c)
		c.JSON(http.StatusOK, &variants)
	}
}

func (vc *VariantController) GetVariantById() gin.HandlerFunc {
	return func(c *gin.Context) {
		variantID := c.Param("id")
		id, err := strconv.Atoi(variantID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid variant id"})
			return
		}
		p := vc.VariantService.GetVariantById(uint(id))
		if p == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no variant with this id"})
			return
		}
		c.JSON(http.StatusOK, &p)
	}
}
