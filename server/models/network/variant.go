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
