package variantcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"persia_atlas/server/models"
)

func (vc *VariantController) CreateVariant() gin.HandlerFunc {

	type CreateVariantRequest struct {
		DKPC            uint  `json:"dkpc" binding:"required"`
		PriceMin        uint  `json:"price_min" binding:"required"`
		StopLoss        *uint `json:"stop_loss" binding:"required"`
		IsActive        *bool `json:"is_active" binding:"required"`
		ProductID       uint  `json:"product" binding:"required,validate-product-id"`
		SelectorID      uint  `json:"selector" binding:"required,validate-variant-selector-id"`
		ActualProductID uint  `json:"actual_product" binding:"required,validate-actual-product-id"`
	}

	var validateProductID validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(uint)
		if ok {
			var p models.Product
			vc.DB.First(&p, value)
			if p.ID == 0 {
				return false
			}
		}
		return true
	}

	var validateVariantSelectorID validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(uint)
		if ok {
			var vs models.VariantSelector
			vc.DB.First(&vs, value)
			if vs.ID == 0 {
				return false
			}
		}
		return true
	}

	var validateActualProductID validator.Func = func(fl validator.FieldLevel) bool {
		value, ok := fl.Field().Interface().(uint)
		if ok {
			var ap models.ActualProduct
			vc.DB.First(&ap, value)
			if ap.ID == 0 {
				return false
			}
		}
		return true
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("validate-product-id", validateProductID, false)
		if err != nil {
			log.Fatalln("error in registering validateProductID:", err.Error())
			return nil
		}
		err = v.RegisterValidation("validate-variant-selector-id", validateVariantSelectorID, false)
		if err != nil {
			log.Fatalln("error in registering validateVariantSelectorID:", err.Error())
			return nil
		}
		err = v.RegisterValidation("validate-actual-product-id", validateActualProductID, false)
		if err != nil {
			log.Fatalln("error in registering validateActualProductID:", err.Error())
			return nil
		}
	}

	return func(c *gin.Context) {
		body := CreateVariantRequest{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("invalid payload: %s", err.Error())},
			)
			return
		}
		fmt.Println("body:", *body.IsActive)
		variant := models.Variant{
			DKPC:            body.DKPC,
			PriceMin:        body.PriceMin,
			StopLoss:        body.StopLoss,
			IsActive:        body.IsActive,
			ProductID:       body.ProductID,
			SelectorID:      body.SelectorID,
			ActualProductID: body.ActualProductID,
		}
		err := vc.VariantService.CreateVariant(&variant)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		variantSerialized := vc.VariantService.GetVariantById(variant.ID)
		c.JSON(http.StatusCreated, &variantSerialized)
	}
}
