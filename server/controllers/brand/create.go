package brandcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
)

func (bc *BrandController) CreateBrand() gin.HandlerFunc {
	type BrandCreateRequest struct {
		Title string `json:"title" binding:"required"`
	}
	return func(c *gin.Context) {
		body := BrandCreateRequest{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid payload: %s", err.Error())},
			)
			return
		}
		b := models.Brand{Title: body.Title}
		err := bc.BrandService.CreateBrand(&b)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &b)
	}
}
