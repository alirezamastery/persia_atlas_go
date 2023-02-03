package apcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (apc *ActualProductController) GetActualProductPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		actualProducts := apc.ActualProductService.GetActualProductsPaginated(c)
		c.JSON(http.StatusOK, &actualProducts)
	}
}

func (apc *ActualProductController) GetActualProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		brandID := c.Param("id")
		id, err := strconv.Atoi(brandID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid actual product id"})
			return
		}
		productType := apc.ActualProductService.GetActualProductById(uint(id))
		if productType == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no actual product with this id"})
			return
		}
		c.JSON(http.StatusOK, &productType)
	}
}
