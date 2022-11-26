package ptypecontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

type UpdateProductTypeRequest struct {
	Title        string `json:"title"`
	SelectorType uint   `json:"selector_type"`
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

		var pt models.ProductType
		if result := ptc.DB.First(&pt, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		ptc.DB.Model(&pt).Updates(models.ProductType{
			Title:          body.Title,
			SelectorTypeID: body.SelectorType,
		})

		c.JSON(http.StatusOK, &pt)
	}
}
