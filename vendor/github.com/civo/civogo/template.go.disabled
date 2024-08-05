package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Template represents a Template for launching instances from
type Template struct {
	ID               string `json:"id"`
	Code             string `json:"code"`
	Name             string `json:"name"`
	Region           string `json:"region"`
	AccountID        string `json:"account_id,omitempty"`
	ImageID          string `json:"image_id,omitempty"`
	VolumeID         string `json:"volume_id"`
	ShortDescription string `json:"short_description"`
	Description      string `json:"description"`
	DefaultUsername  string `json:"default_username"`
	CloudConfig      string `json:"cloud_config"`
}

// ListTemplates return all template in system
func (c *Client) ListTemplates() ([]Template, error) {
	resp, err := c.SendGetRequest("/v2/templates")
	if err != nil {
		return nil, decodeError(err)
	}

	templates := make([]Template, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&templates); err != nil {
		return nil, err
	}

	return templates, nil
}

// NewTemplate will create a new template for the current user
func (c *Client) NewTemplate(conf *Template) (*SimpleResponse, error) {
	if conf.ImageID == "" && conf.VolumeID == "" {
		return nil, errors.New("if image id is not present, volume id must be")
	}

	resp, err := c.SendPostRequest("/v2/templates", conf)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateTemplate will update a template for the current user
func (c *Client) UpdateTemplate(id string, conf *Template) (*Template, error) {
	if conf.ImageID == "" && conf.VolumeID == "" {
		return nil, errors.New("if image id is not present, volume id must be")
	}

	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/templates/%s", id), conf)
	if err != nil {
		return nil, decodeError(err)
	}

	template := &Template{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(template); err != nil {
		return nil, err
	}

	return template, nil
}

// GetTemplateByCode finds the Template for an account with the specified code
func (c *Client) GetTemplateByCode(code string) (*Template, error) {
	resp, err := c.ListTemplates()
	if err != nil {
		return nil, decodeError(err)
	}

	for _, template := range resp {
		if template.Code == code {
			return &template, nil
		}
	}

	return nil, errors.New("template not found")
}

// FindTemplate finds a template by either part of the ID or part of the code
func (c *Client) FindTemplate(search string) (*Template, error) {
	templateList, err := c.ListTemplates()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Template{}

	for _, value := range templateList {
		if value.Code == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Code, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// DeleteTemplate deletes requested template
func (c *Client) DeleteTemplate(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/templates/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
