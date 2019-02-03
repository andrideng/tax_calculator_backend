package services

import (
	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
)

// billDAO specifies the interface of the bill DAO needed by BillService.
type billDAO interface {
	// Get return the bill with the specified bill ID.
	Get(rs app.RequestScope, id int64) (*models.Bill, error)
	// List returns the list of bill.
	List(rs app.RequestScope) ([]models.Bill, error)
	// Create saves a new bill in the storage.
	Create(rs app.RequestScope, artist *models.Bill) error
}

// BillService provides services related with bill.
type BillService struct {
	dao billDAO
}

// NewBillService creates a new BillService with the given billDAO.
func NewBillService(dao billDAO) *BillService {
	return &BillService{dao}
}

// Get return the bill with the specified bill ID.
func (s *BillService) Get(rs app.RequestScope, id int64) (*models.Bill, error) {
	return s.dao.Get(rs, id)
}

// List return the list of bills.
func (s *BillService) List(rs app.RequestScope) ([]models.Bill, error) {
	return s.dao.List(rs)
}

// Create creates a new bill.
func (s *BillService) Create(rs app.RequestScope, model *models.Bill) (*models.Bill, error) {
	if err := model.Validate(); err != nil {
		return nil, err
	}
	if err := s.dao.Create(rs, model); err != nil {
		return nil, err
	}
	return s.dao.Get(rs, model.ID)
}
