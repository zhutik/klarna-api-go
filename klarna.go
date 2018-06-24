package klarna

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
