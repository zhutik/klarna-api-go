package klarna

import "net/http"

// Klarna - base structure with configuration options
//
//        - Credentials instance of API creditials to connect to Adyen API
//        - client is HTTP client instance
type Klarna struct {
	Credentials apiCredentials

	client *http.Client
}
