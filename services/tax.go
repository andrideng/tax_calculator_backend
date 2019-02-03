package services

import (
	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
)

// taxDAO specifies the interface of the tax DAO needed by TaxService.
type taxDAO interface {
	// Get returns the tax with the specified tax_code.
	Get(rs app.RequestScope, taxCode int64) (*models.Tax, error)
}

// TaxService provides services related with tax.
type TaxService struct {
	dao taxDAO
}

// NewTaxService creates a new TaxService with the given tax DAO.
func NewTaxService(dao taxDAO) *TaxService {
	return &TaxService{dao}
}

// Get return the tax with the specified the tax_code.
func (s *TaxService) Get(rs app.RequestScope, taxCode int64) (*models.Tax, error) {
	return s.dao.Get(rs, taxCode)
}
