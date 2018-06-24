package klarna

// ExpireDate - expiration data respresentation for Klarna API
type ExpireDate struct {
	Nano         int `json:"nano"`
	EpochSeconds int `json:"epoch_seconds"`
}

// Attachment - information about attachment
type Attachment struct {
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
}

// ProductIdentifier - Product Identifiers Object
//
// Link - https://developers.klarna.com/api/#payments-api__create-a-new-credit-session__order_lines__product_identifiers
type ProductIdentifier struct {
	CategoryPath           string `json:"category_path,omitempty"`
	GlobalTradeItemNumber  string `json:"global_trade_item_number,omitempty"`
	ManufacturerPartNumber string `json:"manufacturer_part_number,omitempty"`
	Brand                  string `json:"brand,omitempty"`
}

// LineItem - Order Line Object. Represents order items, costs, discount and any additional fees
//
// Link - https://developers.klarna.com/api/#payments-api__create-a-new-credit-session__order_lines
type LineItem struct {
	Type                string             `json:"type,omitempty"`
	Reference           string             `json:"reference,omitempty"`
	Name                string             `json:"name"`
	Quantity            int                `json:"quantity"`
	QuantityUnit        string             `json:"quantity_unit,omitempty"`
	UnitPrice           int                `json:"unit_price"`
	TaxRate             int                `json:"tax_rate"`
	TotalAmount         int                `json:"total_amount"`
	TotalDiscountAmount int                `json:"total_discount_amount,omitempty"`
	TotalTaxAmount      int                `json:"total_tax_amount"`
	MerchantData        string             `json:"merchant_data,omitempty"`
	ProductURL          string             `json:"product_url,omitempty"`
	ImageURL            string             `json:"image_url,omitempty"`
	ProductIdentifiers  *ProductIdentifier `json:"product_identifiers,omitempty"`
}
