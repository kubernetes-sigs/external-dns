package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty"
)

// Stackscript represents a Linode StackScript
type Stackscript struct {
	CreatedStr string `json:"created"`
	UpdatedStr string `json:"updated"`

	ID                int
	Username          string
	Label             string
	Description       string
	Images            []string
	DeploymentsTotal  int
	DeploymentsActive int
	IsPublic          bool
	Created           *time.Time `json:"-"`
	Updated           *time.Time `json:"-"`
	RevNote           string
	Script            string
	UserDefinedFields *map[string]string
	UserGravatarID    string
}

type StackscriptCreateOptions struct {
	Label       string   `json:"label"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	IsPublic    bool     `json:"is_public"`
	RevNote     string   `json:"rev_note"`
	Script      string   `json:"script"`
}

type StackscriptUpdateOptions StackscriptCreateOptions

func (i Stackscript) GetCreateOptions() StackscriptCreateOptions {
	return StackscriptCreateOptions{
		Label:       i.Label,
		Description: i.Description,
		Images:      i.Images,
		IsPublic:    i.IsPublic,
		RevNote:     i.RevNote,
		Script:      i.Script,
	}
}

func (i Stackscript) GetUpdateOptions() StackscriptUpdateOptions {
	return StackscriptUpdateOptions{
		Label:       i.Label,
		Description: i.Description,
		Images:      i.Images,
		IsPublic:    i.IsPublic,
		RevNote:     i.RevNote,
		Script:      i.Script,
	}
}

// StackscriptsPagedResponse represents a paginated Stackscript API response
type StackscriptsPagedResponse struct {
	*PageOptions
	Data []*Stackscript
}

// endpoint gets the endpoint URL for Stackscript
func (StackscriptsPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.StackScripts.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends Stackscripts when processing paginated Stackscript responses
func (resp *StackscriptsPagedResponse) appendData(r *StackscriptsPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Stackscript
func (StackscriptsPagedResponse) setResult(r *resty.Request) {
	r.SetResult(StackscriptsPagedResponse{})
}

// ListStackscripts lists Stackscripts
func (c *Client) ListStackscripts(ctx context.Context, opts *ListOptions) ([]*Stackscript, error) {
	response := StackscriptsPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Stackscript) fixDates() *Stackscript {
	v.Created, _ = parseDates(v.CreatedStr)
	v.Updated, _ = parseDates(v.UpdatedStr)
	return v
}

// GetStackscript gets the Stackscript with the provided ID
func (c *Client) GetStackscript(ctx context.Context, id int) (*Stackscript, error) {
	e, err := c.StackScripts.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := c.R(ctx).SetResult(&Stackscript{}).Get(e)
	if err != nil {
		return nil, err
	}
	return r.Result().(*Stackscript).fixDates(), nil
}

// CreateStackscript creates a StackScript
func (c *Client) CreateStackscript(ctx context.Context, createOpts StackscriptCreateOptions) (*Stackscript, error) {
	var body string
	e, err := c.StackScripts.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Stackscript{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*Stackscript).fixDates(), nil
}

// UpdateStackscript updates the StackScript with the specified id
func (c *Client) UpdateStackscript(ctx context.Context, id int, updateOpts StackscriptUpdateOptions) (*Stackscript, error) {
	var body string
	e, err := c.StackScripts.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&Stackscript{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*Stackscript).fixDates(), nil
}

// DeleteStackscript deletes the StackScript with the specified id
func (c *Client) DeleteStackscript(ctx context.Context, id int) error {
	e, err := c.StackScripts.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	if _, err := coupleAPIErrors(c.R(ctx).Delete(e)); err != nil {
		return err
	}

	return nil
}
