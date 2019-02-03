package daos

import (
	"time"

	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
)

// BillDAO presists bill data in database
type BillDAO struct{}

// NewBillDAO creates a new BillDAO
func NewBillDAO() *BillDAO {
	return &BillDAO{}
}

// Get reads the bill will the specified id from the database
func (dao *BillDAO) Get(rs app.RequestScope, id int64) (*models.Bill, error) {
	var bill models.Bill
	err := rs.Tx().Select().Model(id, &bill)
	return &bill, err
}

// List retrieves all bill record from database.
func (dao *BillDAO) List(rs app.RequestScope) ([]models.Bill, error) {
	bills := []models.Bill{}
	err := rs.Tx().Select().All(&bills)
	return bills, err
}

// Create saves a new bill record in the database.
// The Tax.Id field will be populated with an automatically generated ID upon successfull saving.
func (dao *BillDAO) Create(rs app.RequestScope, bill *models.Bill) error {
	bill.CreatedAt = time.Now()
	bill.UpdatedAt = time.Now()
	return rs.Tx().Model(bill).Insert()
}
