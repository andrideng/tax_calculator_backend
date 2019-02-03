package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	// TaxBill ...
	TaxBill struct {
		ID           int64        `json:"id" db:"pk,id"`
		Name         string       `json:"name" db:"name"`
		TaxCode      int64        `json:"tax_code" db:"tax_code"`
		Price        float64      `json:"price" db:"price"`
		CreatedAt    string       `json:"created_at" db:"created_at"`
		UpdatedAt    string       `json:"updated_at" db:"updated_at"`
		Tax          Tax          `json:"tax"`
		ConvertedTax ConvertedTax `json:"converted_tax"`
	}
	// Bill represent model bills in database
	Bill struct {
		ID        int64   `json:"id" db:"pk,id"`
		Name      string  `json:"name" db:"name"`
		TaxCode   int64   `json:"tax_code" db:"tax_code"`
		Price     float64 `json:"price" db:"price"`
		CreatedAt string  `json:"created_at" db:"created_at"`
		UpdatedAt string  `json:"updated_at" db:"updated_at"`
	}

	// ResponseBill expected bill response to user.
	ResponseBill struct {
		Name       string  `json:"name"`
		TaxCode    int64   `json:"tax_code"`
		Type       string  `json:"type"`
		Refundable string  `json:"refundable"`
		Price      float64 `json:"price"`
		Tax        float64 `json:"tax"`
		Amount     float64 `json:"amount"`
	}

	// TotalResponseBill ...
	TotalResponseBill struct {
		Bills         []ResponseBill `json:"bills"`
		PriceSubTotal float64        `json:"price_sub_total"`
		TaxSubTotal   float64        `json:"tax_sub_total"`
		GrandTotal    float64        `json:"grand_total"`
	}
)

// TableName represent table name in database
func (m Bill) TableName() string {
	return "bills"
}

// Validate validates the Bill fields.
func (m Bill) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 100)),
		validation.Field(&m.TaxCode, validation.Required, validation.Min(0)),
		validation.Field(&m.Price, validation.Required),
	)
}
