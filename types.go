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
