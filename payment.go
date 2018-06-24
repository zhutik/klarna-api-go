package klarna

// PaymentOrder - used to create Payment session for new order
//
// https://developers.klarna.com/api/#payments-api-create-a-new-credit-session
type PaymentOrder struct {
	Design                 string               `json:"design,omitempty"`
	PurchaseCountry        string               `json:"purchase_country"`
	PurchaseCurrency       string               `json:"purchase_currency"`
	Locale                 string               `json:"locale"`
	BillingAddress         *Address             `json:"billing_address,omitempty"`
	ShippingAddress        *Address             `json:"shipping_address,omitempty"`
	OrderAmount            int                  `json:"order_amount"`
	OrderTaxAmount         int                  `json:"order_tax_amount"`
	OrderLines             []*LineItem          `json:"order_lines"`
	Customer               *CustomerInformation `json:"customer,omitempty"`
	MerchantURLS           *MerchantURLS        `json:"merchant_urls,omitempty"`
	MerchantReference1     string               `json:"merchant_reference1,omitempty"`
	MerchantReference2     string               `json:"merchant_reference2,omitempty"`
	MerchantData           string               `json:"merchant_data,omitempty"`
	Options                *Options             `json:"options,omitempty"`
	Attachment             *Attachment          `json:"attachment,omitempty"`
	CustomPaymentMethodIDS []string             `json:"custom_payment_method_ids,omitempty"`
	Status                 string               `json:"status,omitempty"` // https://developers.klarna.com/api/#payments-api__create-a-new-credit-session__status
	ClientToken            string               `json:"client_token,omitempty"`
	ExpiresAt              *ExpireDate          `json:"expires_at,omitempty"`
	AcquiringChannel       string               `json:"acquiring_channel,omitempty"`
}

// PaymentOrderResponse - Payment session API response
//
// Link - https://developer.klarna.com/api/#payments-api__merchant_session__session_id
type PaymentOrderResponse struct {
	SessionID               string                   `json:"session_id"`
	ClientID                string                   `json:"client_token"`
	PaymentMethodCategories []*PaymentMethodCategory `json:"payment_method_categories,omitempty"`
}

// PaymentMethodCategory - Available payment method categories
type PaymentMethodCategory struct {
	Identifier string   `json:"identifier"`
	Name       string   `json:"name"`
	AssetURLS  AssetURL `json:"asset_urls"`
}

// AssetURL - Asset Urls Object
type AssetURL struct {
	Descriptive string `json:"descriptive"`
	Standard    string `json:"standard"`
}

// MerchantURLS - The merchant urls structure
type MerchantURLS struct {
	Confirmation string `json:"confirmation"`
	Notification string `json:"notification,omitempty"`
	Push         string `json:"push,omitempty"`
}

// CustomerInformation - Information about the liable customer of the order
//
// Link - https://developers.klarna.com/api/#payments-api__create-a-new-credit-session__customer
type CustomerInformation struct {
	DateOfBirth                  string `json:"date_of_birth,omitempty"`
	Gender                       string `json:"gender,omitempty"`
	LastFourSSN                  string `json:"last_four_ssn,omitempty"`
	NationalIdentificationNumber string `json:"national_identification_number,omitempty"`
	Type                         string `json:"type,omitempty"`
	VatID                        string `json:"vat_id,omitempty"`
	OrganizationRegistrationID   string `json:"organization_registration_id,omitempty"`
	OrganizationEntityType       string `json:"organization_entity_type,omitempty"` // https://developers.klarna.com/api/#payments-api__create-a-new-credit-sessioncustomer__organization_entity_type
}

// Options - type Options for this purchase
type Options struct {
	ColorButton            string `json:"color_button,omitempty"`
	ColorButtonText        string `json:"color_button_text,omitempty"`
	ColorCheckbox          string `json:"color_checkbox,omitempty"`
	ColorCheckboxCheckmark string `json:"color_checkbox_checkmark,omitempty"`
	ColorHeader            string `json:"color_header,omitempty"`
	ColorLink              string `json:"color_link,omitempty"`
	ColorBorder            string `json:"color_border,omitempty"`
	ColorBorderSelected    string `json:"color_border_selected,omitempty"`
	ColorText              string `json:"color_text,omitempty"`
	ColorDetails           string `json:"color_details,omitempty"`
	ColorTextSecondary     string `json:"color_text_secondary,omitempty"`
	RadiusBorder           string `json:"radius_border,omitempty"`
}
