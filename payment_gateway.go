package klarna

import (
	"bytes"
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

	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(res.Body); err != nil {
		return nil, err
	}

	var p PaymentOrderResponse
	if err := json.Unmarshal(buf.Bytes(), &p); err != nil {
		return nil, err
	}

	return &p, nil
}
