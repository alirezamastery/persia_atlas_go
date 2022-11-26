package ptypecontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
)

func (ptc *ProductTypeController) CreateProductType() gin.HandlerFunc {
	type ProductTypeRequest struct {
		Title        string `json:"title" binding:"required"`
		SelectorType uint   `json:"selector_type" binding:"required"`
	}
	return func(c *gin.Context) {
		body := ProductTypeRequest{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid payload: %s", err.Error())},
			)
			return
		}
		pt := models.ProductType{
			Title:          body.Title,
			SelectorTypeID: body.SelectorType,
		}
		err := ptc.ProductTypeService.CreateProductType(&pt)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &pt)
	}
}
