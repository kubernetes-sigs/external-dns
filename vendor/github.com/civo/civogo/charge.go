package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// Charge represents a Civo resource with number of hours within the specified billing period
type Charge struct {
	Code          string    `json:"code"`
	Label         string    `json:"label"`
	From          time.Time `json:"from"`
	To            time.Time `json:"to"`
	NumHours      int       `json:"num_hours"`
	SizeGigabytes int       `json:"size_gb"`
}

// ListCharges returns all charges for the calling API account
func (c *Client) ListCharges(from, to time.Time) ([]Charge, error) {
	url := "/v2/charges"
	url = url + fmt.Sprintf("?from=%s&to=%s", from.Format(time.RFC3339), to.Format(time.RFC3339))

	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	charges := make([]Charge, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&charges); err != nil {
		return nil, err
	}

	return charges, nil
}
