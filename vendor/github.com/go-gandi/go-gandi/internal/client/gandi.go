package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
)

const (
	gandiEndpoint = "https://api.gandi.net/v5/"
)

// Gandi is the handle used to interact with the Gandi API
type Gandi struct {
	apikey    string
	endpoint  string
	sharingID string
	debug     bool
	dryRun    bool
}

// New instantiates a new Gandi client
func New(apikey string, sharingID string, debug bool, dryRun bool) *Gandi {
	return &Gandi{apikey: apikey, endpoint: gandiEndpoint, sharingID: sharingID, debug: debug, dryRun: dryRun}
}

// SetEndpoint sets the URL to the endpoint. It takes a string defining the subpath under https://api.gandi.net/v5/
func (g *Gandi) SetEndpoint(endpoint string) {
	g.endpoint = gandiEndpoint + endpoint
}

// Get issues a GET request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Get(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodGet, path, params, recipient)
}

// Post issues a POST request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Post(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPost, path, params, recipient)
}

// Patch issues a PATCH request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Patch(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPatch, path, params, recipient)
}

// Delete issues a DELETE request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Delete(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodDelete, path, params, recipient)
}

// Put issues a PUT request. It takes a subpath rooted in the endpoint. Response data is written to the recipient.
// Returns the response headers and any error
func (g *Gandi) Put(path string, params, recipient interface{}) (http.Header, error) {
	return g.askGandi(http.MethodPut, path, params, recipient)
}

func (g *Gandi) askGandi(method, path string, params, recipient interface{}) (http.Header, error) {
	resp, err := g.doAskGandi(method, path, params, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(recipient)
	return resp.Header, nil
}

// GetBytes issues a GET request but does not attempt to parse any response into JSON.
// It returns the response headers, a byteslice of the response, and any error
func (g *Gandi) GetBytes(path string, params interface{}) (http.Header, []byte, error) {
	headers := [][2]string{
		{"Accept", "text/plain"},
	}
	resp, err := g.doAskGandi(http.MethodGet, path, params, headers)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	return resp.Header, content, err
}

func (g *Gandi) doAskGandi(method, path string, p interface{}, extraHeaders [][2]string) (*http.Response, error) {
	var (
		err error
		req *http.Request
	)
	params, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	suffix := ""
	if len(g.sharingID) != 0 {
		suffix += "?sharing_id=" + g.sharingID
	}
	if params != nil && string(params) != "null" {
		req, err = http.NewRequest(method, g.endpoint+path+suffix, bytes.NewReader(params))
	} else {
		req, err = http.NewRequest(method, g.endpoint+path+suffix, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Apikey "+g.apikey)
	req.Header.Add("Content-Type", "application/json")
	if g.dryRun {
		req.Header.Add("Dry-Run", "1")
	}
	for _, header := range extraHeaders {
		req.Header.Add(header[0], header[1])
	}
	if g.debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Println("=======================================\nREQUEST:")
		fmt.Println(string(dump))
	}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}
	if g.debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Println("=======================================\nRESPONSE:")
		fmt.Println(string(dump))
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		defer resp.Body.Close()
		var message StandardResponse
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		decoder.Decode(&message)
		if message.Message != "" {
			err = fmt.Errorf("%d: %s", resp.StatusCode, message.Message)
		} else if len(message.Errors) > 0 {
			var errors []string
			for _, oneError := range message.Errors {
				errors = append(errors, fmt.Sprintf("%s: %s", oneError.Name, oneError.Description))
			}
			err = fmt.Errorf(strings.Join(errors, ", "))
		} else {
			err = fmt.Errorf("%d", resp.StatusCode)

		}
	}
	return resp, err
}

// StandardResponse is a standard response
type StandardResponse struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	UUID    string          `json:"uuid,omitempty"`
	Object  string          `json:"object,omitempty"`
	Cause   string          `json:"cause,omitempty"`
	Status  string          `json:"status,omitempty"`
	Errors  []StandardError `json:"errors,omitempty"`
}

// StandardError is embedded in a standard error
type StandardError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
