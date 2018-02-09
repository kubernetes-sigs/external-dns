package dynect

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	DynAPIPrefix = "https://api.dynect.net/REST"
)

var (
	PollingInterval  = 1 * time.Second
	ErrPromotedToJob = errors.New("promoted to job")
	ErrRateLimited   = errors.New("too many requests")
)

// handleJobRedirect overrides the net/http.DefaultClient's redirection policy
// function.
//
// This function will set the Content-Type, and Auth-Token headers, so that we
// don't get an error back from the API.
func handleJobRedirect(req *http.Request, via []*http.Request) error {
	// Set the Content-Type header.
	req.Header.Set("Content-Type", "application/json")

	// Now, try and divine the Auth-Token header's value from previous
	// requests.
	for _, r := range via {
		if authHdr := r.Header.Get("Auth-Token"); authHdr != "" {
			req.Header.Set("Auth-Token", authHdr)
			return nil
		}
	}
	return fmt.Errorf("failed to set Auth-Token header from previous requests")
}

// A client for use with DynECT's REST API.
type Client struct {
	Token        string
	CustomerName string
	Transport    http.RoundTripper
	verbose      bool
}

// Creates a new Httpclient.
func NewClient(customerName string) *Client {
	return &Client{
		CustomerName: customerName,
		Transport:    &http.Transport{Proxy: http.ProxyFromEnvironment},
	}
}

// Sets the transport for the client.
func (c *Client) SetTransport(t http.RoundTripper) {
	c.Transport = t
}

// Enable, or disable verbose output from the client.
//
// This will enable (or disable) logging messages that explain what the client
// is about to do, like the endpoint it is about to make a request to. If the
// request fails with an unexpected HTTP response code, then the response body
// will be logged out, as well.
func (c *Client) Verbose(p bool) {
	c.verbose = p
}

// Establishes a new session with the DynECT API.
func (c *Client) Login(username, password string) error {
	var req = LoginBlock{
		Username:     username,
		Password:     password,
		CustomerName: c.CustomerName}

	var resp LoginResponse

	err := c.Do("POST", "Session", req, &resp)
	if err != nil {
		return err
	}

	c.Token = resp.Data.Token
	return nil
}

func (c *Client) LoggedIn() bool {
	return len(c.Token) > 0
}

func (c *Client) Logout() error {
	return c.Do("DELETE", "Session", nil, nil)
}

// newRequest creates a new *http.Request, and sets the following headers:
// <ul>
// <li>Auth-Token</li>
// <li>Content-Type</li>
// </ul>
func (c *Client) newRequest(method, urlStr string, data []byte) (*http.Request, error) {
	var r *http.Request
	var err error

	if data != nil {
		r, err = http.NewRequest(method, urlStr, bytes.NewReader(data))
	} else {
		r, err = http.NewRequest(method, urlStr, nil)
	}

	r.Header.Set("Auth-Token", c.Token)
	r.Header.Set("Content-Type", "application/json")

	return r, err
}

func (c *Client) Do(method, endpoint string, requestData, responseData interface{}) error {
	// Throw an error if the user tries to make a request if the client is
	// logged out/unauthenticated, but make an exemption for when the
	// caller is trying to log in.
	if !c.LoggedIn() && method != "POST" && endpoint != "Session" {
		return errors.New("Will not perform request; client is closed")
	}

	var err error

	// Marshal the request data into a byte slice.
	if c.verbose {
		log.Println("dynect: marshaling request data")
	}
	var js []byte
	if requestData != nil {
		js, err = json.Marshal(requestData)
	} else {
		js = []byte("")
	}
	if err != nil {
		return err
	}

	urlStr := fmt.Sprintf("%s/%s", DynAPIPrefix, endpoint)

	// Create a new http.Request.
	req, err := c.newRequest(method, urlStr, js)
	if err != nil {
		return err
	}

	if c.verbose {
		log.Printf("Making %s request to %q", method, urlStr)
	}

	var resp *http.Response
	resp, err = c.Transport.RoundTrip(req)

	if err != nil {
		if c.verbose {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			log.Printf("%s", string(respBody))
		}
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		if resp.ContentLength == 0 {
			// Zero-length content body?
			log.Println("dynect: warning: zero-length response body; skipping decoding of response")
			return nil
		}

		//dec := json.NewDecoder(resp.Body)
		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Could not read response body")
		}
		if err := json.Unmarshal(text, &responseData); err != nil {
			return fmt.Errorf("Error unmarshalling response:", err)
		}

		return nil

	case 307:
		// Handle the temporary redirect, which should point to a
		// /REST/Jobs endpoint.
		loc := resp.Header.Get("Location")
		log.Println("dynect: request is taking too long to complete: redirecting to", loc)

		// Going in to this blind, the documentation says that it will
		// return a URI when promoting a long-running request to a
		// job.
		//
		// Since a URL is technically a URI, we should do some checks
		// on the returned URI to sanitize it, and make sure that it is
		// in the format we would like it to be.
		if strings.HasPrefix(loc, "/REST/") {
			loc = strings.TrimLeft(loc, "/REST/")
		}
		if !strings.HasPrefix(loc, DynAPIPrefix) {
			loc = fmt.Sprintf("%s/%s", DynAPIPrefix, loc)
		}

		log.Println("Fetching location:", loc)

		// Generate a new request.
		req, err := c.newRequest("GET", loc, nil)
		if err != nil {
			return err
		}

		var jobData JobData

		// Poll the API endpoint, until we get a response back.
		for {
			select {
			case <-time.After(PollingInterval):
				resp, err := c.Transport.RoundTrip(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				text, err := ioutil.ReadAll(resp.Body)
				//log.Println(string(text))
				if err != nil {
					return fmt.Errorf("Could not read response body:", err)
				}
				if err := json.Unmarshal(text, &jobData); err != nil {
					return fmt.Errorf("failed to decode job response body:", err)
				}

				// Check to see the status of the job.
				//
				// If it is "incomplete", loop around again.
				//
				// Should the job's status be "success", then
				// return the data, business-as-usual.
				//
				// TODO(nesv): Figure out what to do in the
				// event of a "failure" job status.

				switch jobData.Status {
				case "incomplete":
					continue
				case "success":
					if err := json.Unmarshal(text, &responseData); err != nil {
						return fmt.Errorf("failed to decode response body:", err)
					}
					return nil
				case "failure":
					return fmt.Errorf("request failed: %v", jobData.Messages)
				}
			}
		}

		return nil

	case 429:
		return ErrRateLimited
	}

	// If we got here, this means that the client does not know how to
	// interpret the response, and it should just error out.
	reason, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read in response body")
	}
	return fmt.Errorf("server responded with %v: %v",
		resp.Status,
		string(reason))
}
