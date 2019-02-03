package models

import "time"

// Tax represent taxs table model in database
type Tax struct {
	TaxCode    int64     `json:"tax_code" db:"tax_code"`
	Type       string    `json:"type" db:"type"`
	Condition  string    `json:"condition" db:"condition"`
	Formula    string    `json:"formula" db:"formula"`
	Refundable string    `json:"refundable" db:"refundable"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// TableName represent table name in database
func (m Tax) TableName() string {
	return "taxs"
}
