package variantcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

func (vc *VariantController) UpdateVariant() gin.HandlerFunc {
	type UpdateVariantRequest struct {
		DKPC            *uint `json:"dkpc" binding:"omitempty"`
		PriceMin        *uint `json:"price_min" binding:"omitempty"`
		StopLoss        *uint `json:"stop_loss" binding:"omitempty"`
		IsActive        *bool `json:"is_active" binding:"omitempty"`
		ProductID       *uint `json:"product" binding:"omitempty,validate-product-id"`
		SelectorID      *uint `json:"selector" binding:"omitempty,validate-variant-selector-id"`
		ActualProductID *uint `json:"actual_product" binding:"omitempty,validate-actual-product-id"`
	}

	return func(c *gin.Context) {
		variantID := c.Param("id")
		id, err := strconv.Atoi(variantID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid variant id"})
			return
		}

		body := UpdateVariantRequest{}
		if err := c.BindJSON(&body); err != nil {
			fmt.Print("error:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("body:", body)

		var variant models.Variant
		if result := vc.DB.First(&variant, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		payload := make(map[string]any)
		if body.DKPC != nil {
			payload["dkpc"] = *body.DKPC
		}
		if body.PriceMin != nil {
			payload["price_min"] = *body.PriceMin
		}
		if body.StopLoss != nil {
			payload["stop_loss"] = *body.StopLoss
		}
		if body.IsActive != nil {
			payload["is_active"] = *body.IsActive
		}
		if body.ProductID != nil {
			payload["product_id"] = *body.ProductID
		}
		if body.SelectorID != nil {
			payload["selector_id"] = *body.SelectorID
		}
		if body.ActualProductID != nil {
			payload["actual_product_id"] = *body.ActualProductID
		}

		vc.DB.Model(&variant).Updates(payload)
		updatedProduct := vc.VariantService.GetVariantById(variant.ID)

		c.JSON(http.StatusOK, &updatedProduct)
	}
}
