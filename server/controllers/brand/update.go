package brandcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"persia_atlas/server/models"
	"strconv"
)

type UpdateBrandRequest struct {
	Title string `json:"title"`
}

func (bc *BrandController) UpdateBrand() gin.HandlerFunc {
	return func(c *gin.Context) {
		brandID := c.Param("id")
		id, err := strconv.Atoi(brandID)
		if err != nil || id < 1 {
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid brand id"})
			return
		}

		body := UpdateBrandRequest{}
		if err := c.BindJSON(&body); err != nil {
			fmt.Print("error:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var brand models.Brand
		if result := bc.DB.First(&brand, id); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}

		bc.DB.Model(&brand).Updates(models.Brand{
			Title: body.Title,
		})

		c.JSON(http.StatusOK, &brand)
	}
}
