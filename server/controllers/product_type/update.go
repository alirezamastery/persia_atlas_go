package ptypecontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

type UpdateProductTypeRequest struct {
	Title        *string `json:"title,omitempty" binding:"omitempty"`
	SelectorType *uint   `json:"selector_type,omitempty" binding:"omitempty"`
}

func (ptc *ProductTypeController) UpdateProductType() gin.HandlerFunc {
	return func(c *gin.Context) {
		ptID := c.Param("id")
		id, err := strconv.Atoi(ptID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid product type id"})
			return
		}

		body := UpdateProductTypeRequest{}
		if err := c.BindJSON(&body); err != nil {
			fmt.Print("error:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("body:", body)

		var pt models.ProductType
		if result := ptc.DB.First(&pt, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		payload := make(map[string]any)
		if body.Title != nil {
			payload["Title"] = *body.Title
		}
		if body.SelectorType != nil {
			selectorType := models.VariantSelectorType{}
			ptc.DB.First(&selectorType, body.SelectorType)
			if selectorType.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "selector type not found"})
				return
			}
			payload["SelectorTypeID"] = *body.SelectorType
		}

		ptc.DB.Model(&pt).Updates(payload)

		c.JSON(http.StatusOK, &pt)
	}
}
