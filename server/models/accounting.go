package models

import "time"

type CostType struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `gorm:"size:255" json:"title"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"updated_at"`
	Costs      []Cost    `gorm:"foreignKey:TypeID"`
}

type Cost struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	TypeID      uint
	Title       string    `gorm:"size:255" json:"title"`
	Amount      uint      `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `gorm:"type:text;default:''" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"updated_at"`
}

type Income struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Amount      uint      `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `gorm:"type:text;default:''" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"updated_at"`
}

type ProductCost struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Amount      uint      `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `gorm:"type:text;default:''" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"updated_at"`
}

type Invoice struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	StartDatePersian time.Time `json:"start_date_persian"`
	EndDatePersian   time.Time `json:"end_date_persian"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	InvoiceItems     []InvoiceItem
}

type InvoiceItem struct {
	ID             uint `gorm:"primaryKey" json:"id"`
	InvoiceID      uint
	RowNumber      int       `json:"row_number"`
	Code           uint      `json:"code"`
	DatePersian    time.Time `json:"date_persian"`
	Date           time.Time `json:"date"`
	DKPC           uint      `json:"dkpc"`
	VariantTitle   string    `gorm:"size:255" json:"variant_title"`
	OrderId        uint      `json:"order_id"`
	Serial         string    `gorm:"size:255" json:"serial"`
	Credit         uint      `json:"credit"`
	Debit          uint      `json:"debit"`
	CreditDiscount uint      `json:"credit_discount"`
	DebitDiscount  uint      `json:"debit_discount"`
	CreditFinal    uint      `json:"credit_final"`
	DebitFinal     uint      `json:"debit_final"`
	Description    string    `gorm:"type:text;default:''" json:"description"`
	Calculated     bool      `gorm:"default:false" json:"Calculated"`
}
