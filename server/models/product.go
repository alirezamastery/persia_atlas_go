package models

type Brand struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size:255" json:"title"`

	ActualProducts []ActualProduct `gorm:"constraint:OnDelete:CASCADE;"`
}

type ActualProduct struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"size:255" json:"title"`
	PriceStep uint   `gorm:"default:500" json:"price_step"`
	BrandID   uint
	Brand     Brand `json:"brand"`
	Variants  []Variant
}

type Product struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	DKP      string `gorm:"size:255;unique" json:"dkp"`
	Title    string `gorm:"size:255" json:"title"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	TypeID   uint
	Type     ProductType `json:"type"`
	Variants []Variant
}

type ProductType struct {
	ID             uint                `gorm:"primaryKey" json:"id"`
	Title          string              `gorm:"size:255" json:"title"`
	Products       []Product           `gorm:"foreignKey:TypeID"`
	SelectorTypeID uint                `json:"selector_type_id"`
	SelectorType   VariantSelectorType `json:"selector_type"`
}

type VariantSelectorType struct {
	ID               uint              `gorm:"primaryKey" json:"id"`
	Title            string            `gorm:"size:255" json:"title"`
	ProductTypes     []ProductType     `gorm:"foreignKey:SelectorTypeID"`
	VariantSelectors []VariantSelector `gorm:"foreignKey:SelectorTypeID"`
}

type VariantSelector struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	DigikalaId     uint    `gorm:"unique" json:"digikala_id"`
	Value          string  `gorm:"size:255;unique" json:"value"`
	ExtraInfo      *string `gorm:"size:255" json:"extra_info"`
	SelectorTypeID uint
	SelectorType   VariantSelectorType `json:"selector_type"`
	Variants       []Variant           `gorm:"foreignKey:SelectorID"`
}

type Variant struct {
	ID              uint  `gorm:"primaryKey" json:"id"`
	DKPC            uint  `gorm:"unique" json:"dkpc"`
	PriceMin        uint  `json:"price_min"`
	StopLoss        *uint `gorm:"default:0" json:"stop_loss"`
	IsActive        *bool `gorm:"default:true" json:"is_active"`
	HasCompetition  bool  `gorm:"default:true" json:"has_competition"`
	ProductID       uint
	Product         Product
	SelectorID      uint
	Selector        VariantSelector
	ActualProductID uint
	ActualProduct   ActualProduct
}
