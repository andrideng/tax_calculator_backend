package apis

import (
	"github.com/andrideng/tax-calculator/app"
	"github.com/andrideng/tax-calculator/models"
	"github.com/go-ozzo/ozzo-routing"
)

type (
	// billService specifies the interface fo the bill services needed by billResource
	billService interface {
		List(rs app.RequestScope) ([]models.TaxBill, error)
		Create(rs app.RequestScope, model *models.Bill) (*models.Bill, error)
	}

	// billResource defines the handlers for the APIs.
	billResource struct {
		service billService
	}
)

// ServeBillResource sets up the routing of the bill endpoints and the corresponding handlers.
func ServeBillResource(rg *routing.RouteGroup, service billService) {
	r := &billResource{service}
	rg.Get("/bills", r.list)
	rg.Post("/bills", r.create)
}

func (r *billResource) list(c *routing.Context) error {
	rs := app.GetRequestScope(c)
	bills, err := r.service.List(rs)
	if err != nil {
		return err
	}
	expectedResultBill := models.TotalResponseBill{}
	responseBill := []models.ResponseBill{}
	for _, bill := range bills {
		amount := bill.Price + bill.ConvertedTax.ComputedTax
		responseBill = append(responseBill, models.ResponseBill{
			Name:       bill.Name,
			TaxCode:    bill.TaxCode,
			Type:       bill.Tax.Type,
			Refundable: bill.Tax.Refundable,
			Price:      bill.Price,
			Tax:        bill.ConvertedTax.ComputedTax,
			Amount:     amount,
		})
		expectedResultBill.PriceSubTotal += bill.Price
		expectedResultBill.TaxSubTotal += bill.ConvertedTax.ComputedTax
		expectedResultBill.GrandTotal += amount
	}
	expectedResultBill.Bills = responseBill
	return c.Write(expectedResultBill)
}

func (r *billResource) create(c *routing.Context) error {
	var model models.Bill
	rs := app.GetRequestScope(c)
	if err := c.Read(&model); err != nil {
		return err
	}
	response, err := r.service.Create(rs, &model)
	if err != nil {
		return err
	}
	return c.Write(response)
}
