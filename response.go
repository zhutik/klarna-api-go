package klarna

import (
	"fmt"
	"net/http"
)

// APIError - handle error (non 200 status) errors from Klarna API
type APIError struct {
	ErrorCode     string   `json:"error_code"`
	ErrorMessages []string `json:"error_messages"`
	CorrelationID string   `json:"correlation_id"`
}

// Response - Adyen API response structure
type Response struct {
	*http.Response
	Body []byte
}

// Error - error interface for APIError
func (e APIError) Error() string {
	return fmt.Sprintf("[%s]: %s [%s]", e.ErrorCode, e.ErrorMessages, e.CorrelationID)
}
