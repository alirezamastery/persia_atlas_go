package network

import (
	"encoding/json"
	"errors"
	"fmt"
)

type VariantSerializer struct {
	ID              uint            `json:"id"`
	DKPC            uint            `json:"dkpc"`
	PriceMin        uint            `json:"price_min"`
	StopLoss        uint            `json:"stop_loss"`
	IsActive        bool            `json:"is_active"`
	HasCompetition  bool            `json:"has_competition"`
	ProductID       uint            `json:"-"`
	Product         product         `json:"product"`
	SelectorID      uint            `json:"-"`
	Selector        variantSelector `json:"selector"`
	ActualProductID uint            `json:"-"`
	ActualProduct   actualProduct   `json:"actual_product"`
}

type product struct {
	ID       uint        `json:"id"`
	DKP      string      `json:"dkp"`
	Title    string      `json:"title"`
	IsActive bool        `json:"is_active"`
	TypeID   uint        `json:"-"`
	Type     productType `json:"type"`
}

type productType struct {
	ID             uint                `json:"id"`
	Title          string              `json:"title"`
	SelectorTypeID uint                `json:"-"`
	SelectorType   variantSelectorType `json:"selector_type"`
}

type variantSelector struct {
	ID             uint                `json:"id"`
	DigikalaId     uint                `json:"digikala_id"`
	Value          string              `json:"value"`
	ExtraInfo      *string             `json:"extra_info"`
	SelectorTypeID uint                `json:"-"`
	SelectorType   variantSelectorType `json:"selector_type"`
}

type variantSelectorType struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type actualProduct struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	PriceStep uint   `json:"price_step"`
	BrandID   uint   `json:"-"`
	Brand     brand  `json:"brand"`
}

type brand struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type VariantSerializerNested struct {
	ID             uint `json:"id"`
	DKPC           uint `json:"dkpc"`
	PriceMin       uint `json:"price_min"`
	StopLoss       uint `json:"stop_loss"`
	IsActive       bool `json:"is_active"`
	HasCompetition bool `json:"has_competition"`
	ProductID      uint `json:"-"`
	Product        struct {
		ID       uint   `json:"id"`
		DKP      string `json:"dkp"`
		Title    string `json:"title"`
		IsActive bool   `json:"is_active"`
		TypeID   uint   `json:"-"`
		Type     struct {
			ID             uint   `json:"id"`
			Title          string `json:"title"`
			SelectorTypeID uint   `json:"-"`
			SelectorType   struct {
				ID    uint   `json:"id"`
				Title string `json:"title"`
			} `gorm:"embedded" json:"selector_type"`
		} `gorm:"embedded" json:"type"`
	} `gorm:"embedded" json:"product"`
	SelectorID uint `json:"-"`
	Selector   struct {
		ID             uint    `json:"id"`
		DigikalaId     uint    `json:"digikala_id"`
		Value          string  `json:"value"`
		ExtraInfo      *string `json:"extra_info"`
		SelectorTypeID uint    `json:"-"`
		SelectorType   struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
		} `gorm:"embedded" json:"selector_type"`
	} `gorm:"embedded" json:"selector"`
	ActualProductID uint `json:"-"`
	ActualProduct   struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		PriceStep uint   `json:"price_step"`
		BrandID   uint   `json:"-"`
		Brand     struct {
			ID    uint   `json:"id"`
			Title string `json:"title"`
		} `gorm:"embedded" json:"brand"`
	} `gorm:"embedded" json:"actual_product"`
}

func (v *variantSelector) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	fmt.Println("the scanner:", bytes)
	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	fmt.Println("the scanner result:", bytes)
	*v = variantSelector{}
	return err
}

type VariantScanner struct {
	Id                          uint
	Dkpc                        uint
	Product_id                  uint
	Stop_loss                   uint
	Price_min                   uint
	Is_active                   bool
	Has_competition             bool
	Product__id                 uint
	Product__dkp                string
	Product__title              string
	Product__is_active          bool
	Product__Type__id           uint
	Product__Type__title        string
	SelectorType__id            uint
	SelectorType__title         string
	Selector__id                uint
	Selector__digikala_id       uint
	Selector__value             string
	Selector__extra_info        *string
	ActualProduct__id           uint
	ActualProduct__title        string
	ActualProduct__price_step   uint
	ActualProduct__Brand__id    uint
	ActualProduct__Brand__title string
}

func (s *VariantScanner) Serialize() VariantSerializer {
	return VariantSerializer{
		ID:             s.Id,
		DKPC:           s.Dkpc,
		PriceMin:       s.Price_min,
		StopLoss:       s.Stop_loss,
		IsActive:       s.Is_active,
		HasCompetition: s.Has_competition,
		ProductID:      s.Product_id,
		Product: product{
			ID:       s.Product__id,
			DKP:      s.Product__dkp,
			Title:    s.Product__title,
			IsActive: s.Product__is_active,
			TypeID:   s.Product__Type__id,
			Type: productType{
				ID:             s.Product__Type__id,
				Title:          s.Product__Type__title,
				SelectorTypeID: s.SelectorType__id,
				SelectorType: variantSelectorType{
					ID:    s.SelectorType__id,
					Title: s.SelectorType__title,
				},
			},
		},
		SelectorID: s.Selector__id,
		Selector: variantSelector{
			ID:             s.Selector__id,
			DigikalaId:     s.Selector__digikala_id,
			Value:          s.Selector__value,
			ExtraInfo:      s.Selector__extra_info,
			SelectorTypeID: s.SelectorType__id,
			SelectorType: variantSelectorType{
				ID:    s.SelectorType__id,
				Title: s.SelectorType__title,
			},
		},
		ActualProductID: s.ActualProduct__id,
		ActualProduct: actualProduct{
			ID:        s.ActualProduct__id,
			Title:     s.ActualProduct__title,
			PriceStep: s.ActualProduct__price_step,
			BrandID:   s.ActualProduct__Brand__id,
			Brand: brand{
				ID:    s.ActualProduct__Brand__id,
				Title: s.ActualProduct__Brand__title,
			},
		},
	}
}

type VariantScannerGoConvention struct {
	Id                     uint
	Dkpc                   uint
	ProductId              uint
	StopLoss               uint
	PriceMin               uint
	IsActive               bool
	HasCompetition         bool
	ProductDkp             string
	ProductTitle           string
	ProductIsActive        bool
	ProductTypeId          uint
	ProductTypeTitle       string
	SelectorTypeId         uint
	SelectorTypeTitle      string
	SelectorId             uint
	SelectorDigikalaId     uint
	SelectorValue          string
	SelectorExtraInfo      *string
	ActualProductId        uint
	ActualProductTitle     string
	ActualProductPriceStep uint
	BrandId                uint
	BrandTitle             string
}

func (s *VariantScannerGoConvention) Serialize() VariantSerializer {
	return VariantSerializer{
		ID:             s.Id,
		DKPC:           s.Dkpc,
		PriceMin:       s.PriceMin,
		StopLoss:       s.StopLoss,
		IsActive:       s.IsActive,
		HasCompetition: s.HasCompetition,
		ProductID:      s.ProductId,
		Product: product{
			ID:       s.ProductId,
			DKP:      s.ProductDkp,
			Title:    s.ProductTitle,
			IsActive: s.ProductIsActive,
			TypeID:   s.ProductTypeId,
			Type: productType{
				ID:             s.ProductTypeId,
				Title:          s.ProductTypeTitle,
				SelectorTypeID: s.SelectorTypeId,
				SelectorType: variantSelectorType{
					ID:    s.SelectorTypeId,
					Title: s.SelectorTypeTitle,
				},
			},
		},
		SelectorID: s.SelectorId,
		Selector: variantSelector{
			ID:             s.SelectorId,
			DigikalaId:     s.SelectorDigikalaId,
			Value:          s.SelectorValue,
			ExtraInfo:      s.SelectorExtraInfo,
			SelectorTypeID: s.SelectorTypeId,
			SelectorType: variantSelectorType{
				ID:    s.SelectorTypeId,
				Title: s.SelectorTypeTitle,
			},
		},
		ActualProductID: s.ActualProductId,
		ActualProduct: actualProduct{
			ID:        s.ActualProductId,
			Title:     s.ActualProductTitle,
			PriceStep: s.ActualProductPriceStep,
			BrandID:   s.BrandId,
			Brand: brand{
				ID:    s.BrandId,
				Title: s.BrandTitle,
			},
		},
	}
}
