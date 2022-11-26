package ptypecontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (ptc *ProductTypeController) GetProductTypesPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		brands := ptc.ProductTypeService.GetProductTypesPaginated(c)
		c.JSON(http.StatusOK, &brands)
	}
}

func (ptc *ProductTypeController) GetProductTypeById() gin.HandlerFunc {
	return func(c *gin.Context) {
		brandID := c.Param("id")
		id, err := strconv.Atoi(brandID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid product type id"})
			return
		}
		productType := ptc.ProductTypeService.GetProductTypeById(uint(id))
		if productType == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no product type with this id"})
			return
		}
		c.JSON(http.StatusOK, &productType)
	}
}
