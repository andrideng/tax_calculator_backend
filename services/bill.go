package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
)

// billDAO specifies the interface of the bill DAO needed by BillService.
type billDAO interface {
	// Get return the bill with the specified bill ID.
	Get(rs app.RequestScope, id int64) (*models.Bill, error)
	// List returns the list of bill.
	List(rs app.RequestScope) ([]models.TaxBill, error)
	// Create saves a new bill in the storage.
	Create(rs app.RequestScope, artist *models.Bill) error
}

// BillService provides services related with bill.
type BillService struct {
	dao    billDAO
	taxDAO taxDAO
}

// NewBillService creates a new BillService with the given billDAO.
func NewBillService(dao billDAO, taxDAO taxDAO) *BillService {
	return &BillService{dao, taxDAO}
}

// Get return the bill with the specified bill ID.
func (s *BillService) Get(rs app.RequestScope, id int64) (*models.Bill, error) {
	return s.dao.Get(rs, id)
}

// List return the list of bills.
func (s *BillService) List(rs app.RequestScope) ([]models.TaxBill, error) {
	bills, err := s.dao.List(rs)
	if err != nil {
		return nil, err
	}
	for idx, bill := range bills {
		taxCode := bill.TaxCode
		// - get tax dao data
		tax, err := s.taxDAO.Get(rs, taxCode)
		if err != nil {
			return nil, err
		}

		condition := tax.Condition
		parseFormula := false
		if condition == "" {
			parseFormula = true
		} else {
			// - parse condition
			str := strings.Replace(tax.Condition, "{price}", strconv.Itoa(int(bill.Price)), -1)
			fmt.Println(str)
			exp, err := govaluate.NewEvaluableExpression(str)
			if err != nil {
				return nil, err
			}
			result, err := exp.Evaluate(nil)
			if err != nil {
				return nil, err
			}
			parseFormula = result.(bool)
		}

		if parseFormula {
			// - parse formula
			str := strings.Replace(tax.Formula, "{price}", strconv.Itoa(int(bill.Price)), -1)
			exp, err := govaluate.NewEvaluableExpression(str)
			if err != nil {
				return nil, err
			}
			result, err := exp.Evaluate(nil)
			if err != nil {
				return nil, err
			}
			bills[idx].ConvertedTax.ComputedTax = result.(float64)
		}
		// - set biil tax value
		bills[idx].Tax = *tax
	}

	return bills, nil
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
