package klarna

import (
	"encoding/json"
)

// PaymentGateway - Klarna Payment API gateway
type PaymentGateway struct {
	*Klarna
}

const sessionCreateURL = "/payments/v1/sessions"

// CreateSession - Create a new credit session
//
// Link - https://developers.klarna.com/api/#payments-api-create-a-new-credit-session
func (k *PaymentGateway) CreateSession(req *PaymentOrder) (*PaymentOrderResponse, error) {
	uri := k.createURL(sessionCreateURL)

	res, err := k.post(uri, req)

	if nil != err {
		return nil, err
	}

	var p PaymentOrderResponse
	if err := json.Unmarshal(res.Body, &p); err != nil {
		return nil, err
	}

	return &p, nil
}
