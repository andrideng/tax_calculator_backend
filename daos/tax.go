package daos

import (
	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
	"github.com/go-ozzo/ozzo-dbx"
)

// TaxDAO presists tax data in database
type TaxDAO struct{}

// NewTaxDAO creates a new TaxDAO
func NewTaxDAO() *TaxDAO {
	return &TaxDAO{}
}

// Get reads the tax with the specified tax_code from the database.
func (dao *TaxDAO) Get(rs app.RequestScope, taxCode int64) (*models.Tax, error) {
	var tax models.Tax
	err := rs.Tx().Select().Where(
		dbx.HashExp{"tax_code": taxCode},
	).One(&tax)
	return &tax, err
}
