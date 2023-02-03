package network

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
