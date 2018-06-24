package klarna

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// LibraryVersion - current library version
const LibraryVersion = "0.1"

// Klarna - base structure with configuration options
//
//        - Credentials instance of API creditials to connect to Adyen API
//        - client is HTTP client instance
type Klarna struct {
	Credentials apiCredentials

	client *http.Client
}

// New - creates new Klarna instance
//
// Description:
//
//     - env - Environment for next API calls
//     - username - API username (UID)
//     - password - API password
//     - opts - an optional collection of functions that allow you to tweak configurations.
func New(env Environment, username, password string, opts ...Option) *Klarna {
	credentials := makeCredentials(env, username, password)
	return NewWithCredentials(credentials, opts...)
}

// NewWithCredentials - create new Klarna instance with pre-configured credentials.
//
// Description:
//
//     - credentials - configured apiCredentials to use when interacting with Adyen.
//     - opts - an optional collection of functions that allow you to tweak configurations.
//
func NewWithCredentials(creds apiCredentials, opts ...Option) *Klarna {
	k := Klarna{
		Credentials: creds,
		client:      &http.Client{},
	}

	if opts != nil {
		for _, opt := range opts {
			opt(&k)
		}
	}

	return &k
}

// Option allows for custom configuration overrides.
type Option func(*Klarna)

// WithTimeout allows for a custom timeout to be provided to the underlying
// HTTP client that's used to communicate with Adyen.
func WithTimeout(d time.Duration) func(*Klarna) {
	return func(a *Klarna) {
		a.client.Timeout = d
	}
}

// Payment - returns PaymentGateway
func (k *Klarna) Payment() *PaymentGateway {
	return &PaymentGateway{k}
}

// perform POST request to the Klarna API
//
// Should be used by gateway instances directly
func (k *Klarna) post(path string, requestEntity interface{}) (*http.Response, error) {
	return k.execute(http.MethodPost, path, requestEntity)
}

// perform HTTP request to a given URL with requestEntity payload
// adds HTTP Basic Authentication based on provided API credentials
// handles API error response and provide human readable errors
func (k *Klarna) execute(method string, url string, requestEntity interface{}) (*http.Response, error) {
	var reader io.Reader
	if nil != requestEntity {
		bytesBody, err := json.Marshal(requestEntity)
		if nil != err {
			return nil, err
		}
		reader = bytes.NewReader(bytesBody)
	}

	req, err := http.NewRequest(method, url, reader)
	if nil != err {
		return nil, err
	}

	req.SetBasicAuth(k.Credentials.Username, k.Credentials.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", fmt.Sprintf("Klarna Go %s", LibraryVersion))

	res, err := k.client.Do(req)
	if nil != err {
		return nil, err
	}

	defer func() {
		if cerr := res.Body.Close(); cerr != nil {
			err = cerr
		}
	}()

	err = k.handleHTTPError(res)
	if nil != err {
		return nil, err
	}

	return res, nil
}

// handleHTTPError - method to handle Klarna API error response
//
// Link - https://developers.klarna.com/api/#errors
func (k *Klarna) handleHTTPError(res *http.Response) error {
	if res.StatusCode > 299 {
		// handle errors from Klarna API
	}

	return nil
}

// createURL - creates full URI to perform an API request
func (k *Klarna) createURL(resource string) string {
	uri := fmt.Sprintf(
		"%s%s",
		strings.TrimRight(k.Credentials.Env.apiURL, "/"),
		resource,
	)

	return uri
}
