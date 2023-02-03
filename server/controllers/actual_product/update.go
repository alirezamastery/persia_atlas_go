package apcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

type UpdateActualProductRequest struct {
	Title     *string `json:"title" binding:"omitempty"`
	PriceStep *uint   `json:"price_step" binding:"omitempty"`
	Brand     *uint   `json:"brand" binding:"omitempty,validate-brand"`
}

func (apc *ActualProductController) UpdateActualProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		apID := c.Param("id")
		id, err := strconv.Atoi(apID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid actual product id"})
			return
		}

		body := UpdateActualProductRequest{}
		if err := c.BindJSON(&body); err != nil {
			fmt.Print("error:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("body:", body)

		var ap models.ActualProduct
		if result := apc.DB.First(&ap, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		payload := make(map[string]any)
		if body.Title != nil {
			payload["title"] = *body.Title
		}
		if body.PriceStep != nil {
			payload["price_step"] = *body.PriceStep
		}
		if body.Brand != nil {
			payload["brand_id"] = *body.Brand
		}

		apc.DB.Model(&ap).Updates(payload)
		updatedAP := apc.ActualProductService.GetActualProductById(ap.ID)

		c.JSON(http.StatusOK, &updatedAP)
	}
}
