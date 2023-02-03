package productcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (pc *ProductController) GetProductsPaginated() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := pc.ProductService.GetProductsPaginated(c)
		c.JSON(http.StatusOK, &products)
	}
}

func (pc *ProductController) GetProductById() gin.HandlerFunc {
	return func(c *gin.Context) {
		productID := c.Param("id")
		id, err := strconv.Atoi(productID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid product id"})
			return
		}
		p := pc.ProductService.GetProductById(uint(id))
		if p == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "no product with this id"})
			return
		}
		c.JSON(http.StatusOK, &p)
	}
}
