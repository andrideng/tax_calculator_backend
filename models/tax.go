package models

type (
	// Tax represent taxs table model in database
	Tax struct {
		ID         int64  `json:"id" db:"pk,id"`
		TaxCode    int64  `json:"tax_code" db:"tax_code"`
		Type       string `json:"type" db:"type"`
		Condition  string `json:"condition" db:"condition"`
		Formula    string `json:"formula" db:"formula"`
		Refundable string `json:"refundable" db:"refundable"`
		CreatedAt  string `json:"created_at" db:"created_at"`
		UpdatedAt  string `json:"updated_at" db:"updated_at"`
	}

	// ConvertedTax ...
	ConvertedTax struct {
		ComputedTax float64 `json:"computed_tax"`
	}
)

// TableName represent table name in database
func (m Tax) TableName() string {
	return "taxs"
}
