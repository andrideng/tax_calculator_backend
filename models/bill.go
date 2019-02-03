package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Bill represent model bills in database
type Bill struct {
	ID        int64     `json:"id" db:"pk,id"`
	Name      string    `json:"name" db:"name"`
	TaxCode   int64     `json:"tax_code" db:"tax_code"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// TableName represent table name in database
func (m Bill) TableName() string {
	return "bills"
}

// Validate validates the Bill fields.
func (m Bill) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
	)
}
