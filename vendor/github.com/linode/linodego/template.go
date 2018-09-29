// +build ignore

package linodego

/*
 - replace "Template" with "NameOfResource"
 - replace "template" with "nameOfResource"
 - When updating Template structs,
   - use pointers where ever null'able would have a different meaning if the wrapper
	 supplied "" or 0 instead
 - Add "NameOfResource" to client.go, resources.go, pagination.go
*/

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty"
)

// Template represents a Template object
type Template struct {
	ID int
	// UpdatedStr string `json:"updated"`
	// Updated *time.Time `json:"-"`
}

type TemplateCreateOptions struct {
}

type TemplateUpdateOptions struct {
}

func (i Template) GetCreateOptions() (o TemplateCreateOptions) {
	// o.Label = i.Label
	// o.Description = copyString(o.Description)
	return
}

func (i Template) GetUpdateOptions() (o TemplateCreateOptions) {
	// o.Label = i.Label
	// o.Description = copyString(o.Description)
	return
}

// TemplatesPagedResponse represents a paginated Template API response
type TemplatesPagedResponse struct {
	*PageOptions
	Data []*Template
}

// endpoint gets the endpoint URL for Template
func (TemplatesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Templates.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends Templates when processing paginated Template responses
func (resp *TemplatesPagedResponse) appendData(r *TemplatesPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Template
func (TemplatesPagedResponse) setResult(r *resty.Request) {
	r.SetResult(TemplatesPagedResponse{})
}

// ListTemplates lists Templates
func (c *Client) ListTemplates(opts *ListOptions) ([]*Template, error) {
	response := TemplatesPagedResponse{}
	err := c.listHelper(&response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Template) fixDates() *Template {
	// v.Created, _ = parseDates(v.CreatedStr)
	// v.Updated, _ = parseDates(v.UpdatedStr)
	return v
}

// GetTemplate gets the template with the provided ID
func (c *Client) GetTemplate(id int) (*Template, error) {
	e, err := c.Templates.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R().SetResult(&Template{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*Template).fixDates(), nil
}

// CreateTemplate creates a Template
func (c *Client) CreateTemplate(ctx context.Context, createOpts TemplateCreateOptions) (*Template, error) {
	var body string
	e, err := c.Templates.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&Template{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}
	return r.Result().(*Template).fixDates(), nil
}

// UpdateTemplate updates the Template with the specified id
func (c *Client) UpdateTemplate(ctx context.Context, id int, updateOpts TemplateUpdateOptions) (*Template, error) {
	var body string
	e, err := c.Templates.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	req := c.R(ctx).SetResult(&Template{})

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
	return r.Result().(*Template).fixDates(), nil
}

// DeleteTemplate deletes the Template with the specified id
func (c *Client) DeleteTemplate(ctx context.Context, id int) error {
	e, err := c.Templates.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	if _, err := coupleAPIErrors(c.R(ctx).Delete(e)); err != nil {
		return err
	}

	return nil
}
