// +build integration

package klarna

import (
	"fmt"
	"testing"
)

func TestCreateSession(t *testing.T) {
	instance := getInstance(TestEnvironment())

	req := &PaymentOrder{
		PurchaseCountry:  "DE",
		PurchaseCurrency: "EUR",
		Locale:           "de-DE",
		OrderAmount:      10,
		OrderTaxAmount:   0,
		OrderLines: []*LineItem{
			{
				Type:                "physical",
				Reference:           "19-402",
				Name:                "Battery Power Pack",
				Quantity:            1,
				UnitPrice:           10,
				TaxRate:             0,
				TotalAmount:         10,
				TotalDiscountAmount: 0,
				TotalTaxAmount:      0,
			},
		},
	}

	response, err := instance.Payment().CreateSession(req)

	if err != nil {
		t.Errorf("CreateSession response should be successful, error - %s", err)
	}

	fmt.Println(response.ClientID)
	fmt.Println(response.SessionID)

	assert(t, len(response.ClientID) > 0, "Response Client ID could not be empty")
	assert(t, len(response.SessionID) > 0, "Response SessionID could not be empty")
}
