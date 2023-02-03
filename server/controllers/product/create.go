package productcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"persia_atlas/server/models"
)

func (pc *ProductController) CreateProduct() gin.HandlerFunc {
	type ProductCreateRequest struct {
		Title    string `json:"title" binding:"required"`
		DKP      string `json:"dkp" binding:"required"`
		IsActive bool   `json:"is_active" binding:"required"`
		Type     uint   `json:"type" binding:"required,validate-product-type"`
	}

	var validateProductType validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(uint)
		if ok {
			var pt models.ProductType
			pc.DB.First(&pt, value)
			if pt.ID == 0 {
				return false
			}
		}
		return true
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("validate-product-type", validateProductType, false)
		if err != nil {
			log.Fatalln("error in registering validator:", err.Error())
			return nil
		}
	}

	return func(c *gin.Context) {
		body := ProductCreateRequest{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid payload: %s", err.Error())},
			)
			return
		}
		product := models.Product{
			Title:    body.Title,
			DKP:      body.DKP,
			IsActive: body.IsActive,
			TypeID:   body.Type,
		}
		err := pc.ProductService.CreateProduct(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &product)
	}
}
