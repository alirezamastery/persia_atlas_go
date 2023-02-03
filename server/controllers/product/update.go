package productcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

func (pc *ProductController) UpdateProduct() gin.HandlerFunc {
	type UpdateProductRequest struct {
		Title    *string `json:"title" binding:"omitempty"`
		DKP      *string `json:"dkp" binding:"omitempty"`
		IsActive *bool   `json:"is_active" binding:"omitempty"`
		Type     *uint   `json:"type" binding:"omitempty,validate-product-type"`
	}

	return func(c *gin.Context) {
		productID := c.Param("id")
		id, err := strconv.Atoi(productID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid product id"})
			return
		}

		body := UpdateProductRequest{}
		if err := c.BindJSON(&body); err != nil {
			fmt.Print("error:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("body:", body)

		var product models.Product
		if result := pc.DB.First(&product, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		payload := make(map[string]any)
		if body.Title != nil {
			payload["title"] = *body.Title
		}
		if body.DKP != nil {
			payload["dkp"] = *body.DKP
		}
		if body.IsActive != nil {
			payload["is_active"] = *body.IsActive
		}
		if body.Type != nil {
			payload["type_id"] = *body.Type
		}

		pc.DB.Model(&product).Updates(payload)
		updatedProduct := pc.ProductService.GetProductById(product.ID)

		c.JSON(http.StatusOK, &updatedProduct)
	}
}
