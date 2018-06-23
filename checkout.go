package klarna

// Address - Customer's Billing or Shipping address information
//
// Link - https://developer.klarna.com/api/#checkout-api-create-a-new-order
type Address struct {
	OrganizationName string `json:"organization_name,omitempty"`
	Reference        string `json:"reference,omitempty"`
	Attention        string `json:"attention,omitempty"`
	GivenName        string `json:"given_name,omitempty"`
	FamilyName       string `json:"family_name,omitempty"`
	Email            string `json:"email,omitempty"`
	Title            string `json:"title,omitempty"`
	StreetAddress    string `json:"street_address,omitempty"`
	StreetAddress2   string `json:"street_address2,omitempty"`
	StreetName       string `json:"street_name,omitempty"`     // Street name. Only applicable in DE/AT/NL. Do not combine with street_address. See streetNumber.
	StreenNumber     string `json:"streen_number,omitempty"`   // Street number. Only applicable in DE/AT/NL. Do not combine with street_address. See streetName.
	HouseExtension   string `json:"house_extension,omitempty"` // Only applicable in NL
	PostalCode       string `json:"postal_code,omitempty"`
	City             string `json:"city,omitempty"`
	Region           string `json:"region,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Country          string `json:"country,omitempty"`
}
