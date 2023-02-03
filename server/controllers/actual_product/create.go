package apcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"persia_atlas/server/models"
)

func (apc *ActualProductController) CreateActualProduct() gin.HandlerFunc {
	type ActualProductRequest struct {
		Title     string `json:"title" binding:"required"`
		PriceStep uint   `json:"price_step" binding:"required"`
		Brand     uint   `json:"brand" binding:"required,validate-brand"`
	}

	var validateBrand validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(uint)
		fmt.Println("validate brand:", value)
		if ok {
			var brand models.Brand
			apc.DB.First(&brand, value)
			if brand.ID == 0 {
				return false
			}
		}
		return true
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("validate-brand", validateBrand, false)
		if err != nil {
			log.Fatalln("error in registering validator:", err.Error())
			return nil
		}
	}

	return func(c *gin.Context) {
		body := ActualProductRequest{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid payload: %s", err.Error())},
			)
			return
		}
		ap := models.ActualProduct{
			Title:     body.Title,
			PriceStep: body.PriceStep,
			BrandID:   body.Brand,
		}
		err := apc.ActualProductService.CreateActualProduct(&ap)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, &ap)
	}
}
