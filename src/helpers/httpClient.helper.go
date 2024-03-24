package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type S_Client struct {
	baseURL     string
	bearerToken string
}

func NewClient(baseURL, bearerToken string) *S_Client {
	return &S_Client{baseURL: baseURL, bearerToken: bearerToken}
}

func (c *S_Client) Get(endpoint string, queryParams map[string]string) ([]byte, error) {
	fullUrl := c.baseURL + endpoint

	// Add query parameters to the URL
	if len(queryParams) > 0 {
		query := url.Values{}
		for key, value := range queryParams {
			query.Set(key, value)
		}
		fullUrl += "?" + query.Encode()
	}

	// Create a new GET request
	request, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Bearer token header if provided
	if c.bearerToken != "" {
		request.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}

	// Send the GET request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}

// Delete sends a DELETE request with optional query parameters
func (c *S_Client) Delete(endpoint string, queryParams map[string]string) ([]byte, error) {
	fullUrl := c.baseURL + endpoint

	// Add query parameters to the URL
	if len(queryParams) > 0 {
		query := url.Values{}
		for key, value := range queryParams {
			query.Set(key, value)
		}
		fullUrl += "?" + query.Encode()
	}

	// Create a new DELETE request
	request, err := http.NewRequest("DELETE", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Bearer token header if provided
	if c.bearerToken != "" {
		request.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}

	// Send the DELETE request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}

// Post sends a POST request with the provided payload
func (c *S_Client) Post(endpoint string, payload interface{}) ([]byte, error) {
	return c.doRequest("POST", endpoint, payload)
}

// Put sends a PUT request with the provided payload
func (c *S_Client) Put(endpoint string, payload interface{}) ([]byte, error) {
	return c.doRequest("PUT", endpoint, payload)
}

// doRequest sends an HTTP request with the specified method and payload
func (c *S_Client) doRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := c.baseURL + endpoint

	// Convert payload to JSON if provided
	var jsonPayload []byte
	if payload != nil {
		var err error
		jsonPayload, err = json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("error marshaling JSON: %w", err)
		}
	}

	// Create a new request with the specified method and URL
	var request *http.Request
	var err error
	if jsonPayload != nil {
		request, err = http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set the Content-Type header to application/json if payload is provided
	if jsonPayload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	// Set the Bearer token header if provided
	if c.bearerToken != "" {
		request.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return body, nil
}
