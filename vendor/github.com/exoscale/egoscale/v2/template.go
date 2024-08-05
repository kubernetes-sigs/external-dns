package v2

import (
	"context"
<<<<<<< HEAD
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// Template represents a Compute instance template.
type Template struct {
	BootMode        *string
	Build           *string
	Checksum        *string `req-for:"create"`
	CreatedAt       *time.Time
	DefaultUser     *string
	Description     *string
	Family          *string
	ID              *string `req-for:"delete"`
	Name            *string `req-for:"create"`
	PasswordEnabled *bool   `req-for:"create"`
	SSHKeyEnabled   *bool   `req-for:"create"`
	Size            *int64
	URL             *string `req-for:"create"`
	Version         *string
	Visibility      *string
}

func templateFromAPI(t *oapi.Template) *Template {
	return &Template{
		BootMode:        (*string)(t.BootMode),
		Build:           t.Build,
		Checksum:        t.Checksum,
		CreatedAt:       t.CreatedAt,
		DefaultUser:     t.DefaultUser,
		Description:     t.Description,
		Family:          t.Family,
		ID:              t.Id,
		Name:            t.Name,
		PasswordEnabled: t.PasswordEnabled,
		SSHKeyEnabled:   t.SshKeyEnabled,
		Size:            t.Size,
		URL:             t.Url,
		Version:         t.Version,
		Visibility:      (*string)(t.Visibility),
	}
}

// DeleteTemplate deletes a Template.
func (c *Client) DeleteTemplate(ctx context.Context, zone string, template *Template) error {
	if err := validateOperationParams(template, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteTemplateWithResponse(apiv2.WithZone(ctx, zone), *template.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// GetTemplate returns the Template corresponding to the specified ID.
func (c *Client) GetTemplate(ctx context.Context, zone, id string) (*Template, error) {
	resp, err := c.GetTemplateWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return templateFromAPI(resp.JSON200), nil
}

// ListTemplates returns the list of existing Templates.
func (c *Client) ListTemplates(ctx context.Context, zone, visibility, family string) ([]*Template, error) {
	list := make([]*Template, 0)

	resp, err := c.ListTemplatesWithResponse(apiv2.WithZone(ctx, zone), &oapi.ListTemplatesParams{
		Visibility: (*oapi.ListTemplatesParamsVisibility)(&visibility),
		Family: func() *string {
			if family != "" {
				return &family
			}
			return nil
		}(),
	})
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Templates != nil {
		for i := range *resp.JSON200.Templates {
			list = append(list, templateFromAPI(&(*resp.JSON200.Templates)[i]))
		}
	}

	return list, nil
}

// RegisterTemplate registers a new Template.
func (c *Client) RegisterTemplate(ctx context.Context, zone string, template *Template) (*Template, error) {
	if err := validateOperationParams(template, "create"); err != nil {
		return nil, err
	}

	resp, err := c.RegisterTemplateWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.RegisterTemplateJSONRequestBody{
			BootMode:        (*oapi.RegisterTemplateJSONBodyBootMode)(template.BootMode),
			Checksum:        *template.Checksum,
			DefaultUser:     template.DefaultUser,
			Description:     template.Description,
			Name:            *template.Name,
			PasswordEnabled: *template.PasswordEnabled,
			SshKeyEnabled:   *template.SSHKeyEnabled,
			Url:             *template.URL,
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetTemplate(ctx, zone, *res.(*oapi.Reference).Id)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"errors"
	"sort"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// Template represents a Compute instance template.
type Template struct {
	BootMode        *string
	Build           *string
	Checksum        *string `req-for:"create"`
	CreatedAt       *time.Time
	DefaultUser     *string
	Description     *string
	Family          *string
	ID              *string `req-for:"update,delete"`
	Maintainer      *string
	Name            *string `req-for:"create"`
	PasswordEnabled *bool   `req-for:"create"`
	SSHKeyEnabled   *bool   `req-for:"create"`
	Size            *int64
	URL             *string `req-for:"create"`
	Version         *string
	Visibility      *string
	Zone            *string
}

// Implement sort.Interface for []*Template by CreatedAt or by Nane
type Templates []*Template

func (s Templates) Len() int      { return len(s) }
func (s Templates) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Templates }

func (s ByName) Less(i, j int) bool { return *s.Templates[i].Name < *s.Templates[j].Name }

type ByCreatedAt struct{ Templates }

func (s ByCreatedAt) Less(i, j int) bool {
	return s.Templates[i].CreatedAt.Before(*s.Templates[j].CreatedAt)
}

// ListTemplatesOpt represents an ListTemplates operation option.
type ListTemplatesOpt func(params *oapi.ListTemplatesParams)

// ListTemplatesWithFamily sets a family filter to list Templates with.
func ListTemplatesWithFamily(v string) ListTemplatesOpt {
	return func(p *oapi.ListTemplatesParams) {
		if v != "" {
			p.Family = &v
		}
	}
}

// ListTemplatesWithVisibility sets a visibility filter to list Templates with (default: "public").
func ListTemplatesWithVisibility(v string) ListTemplatesOpt {
	return func(p *oapi.ListTemplatesParams) {
		if v != "" {
			p.Visibility = (*oapi.ListTemplatesParamsVisibility)(&v)
		}
	}
}

func templateFromAPI(t *oapi.Template, zone string) *Template {
	return &Template{
		BootMode:        (*string)(t.BootMode),
		Build:           t.Build,
		Checksum:        t.Checksum,
		CreatedAt:       t.CreatedAt,
		DefaultUser:     t.DefaultUser,
		Description:     t.Description,
		Family:          t.Family,
		ID:              t.Id,
		Maintainer:      t.Maintainer,
		Name:            t.Name,
		PasswordEnabled: t.PasswordEnabled,
		SSHKeyEnabled:   t.SshKeyEnabled,
		Size:            t.Size,
		URL:             t.Url,
		Version:         t.Version,
		Visibility:      (*string)(t.Visibility),
		Zone:            &zone,
	}
}

// CopyTemplate copies a Template to a different Exoscale zone.
func (c *Client) CopyTemplate(ctx context.Context, zone string, template *Template, dstZone string) (*Template, error) {
	if err := validateOperationParams(template, "update"); err != nil {
		return nil, err
	}

	resp, err := c.CopyTemplateWithResponse(
		apiv2.WithZone(ctx, zone),
		*template.ID,
		oapi.CopyTemplateJSONRequestBody{TargetZone: oapi.Zone{Name: (*oapi.ZoneName)(&dstZone)}},
	)
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetTemplate(ctx, dstZone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteTemplate deletes a Template.
func (c *Client) DeleteTemplate(ctx context.Context, zone string, template *Template) error {
	if err := validateOperationParams(template, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteTemplateWithResponse(apiv2.WithZone(ctx, zone), *template.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// GetTemplate returns the Template corresponding to the specified ID.
func (c *Client) GetTemplate(ctx context.Context, zone, id string) (*Template, error) {
	resp, err := c.GetTemplateWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return templateFromAPI(resp.JSON200, zone), nil
}

// GetTemplateByName returns the newest Template corresponding to the specified Name.
func (c *Client) GetTemplateByName(ctx context.Context, zone string, templateName string, visibility string) (*Template, error) {
	templates, err := c.ListTemplates(ctx, zone, ListTemplatesWithVisibility(visibility))
	if err != nil {
		return nil, err
	}
	// Newest first (multiple private templates can have the same name)
	sort.Sort(sort.Reverse(ByCreatedAt{templates}))
	for _, template := range templates {
		if *template.Name == templateName {
			return template, nil
		}
	}

	return nil, apiv2.ErrNotFound
}

// FindTemplate attempts to find a template by name or ID.
// In case the identifier is a name and multiple resources match, the newest template is returned.
func (c *Client) FindTemplate(ctx context.Context, zone, x string, visibilty string) (*Template, error) {
	// Check if template is referenced by ID.
	template, err := c.GetTemplate(ctx, zone, x)
	if err != nil {
		// ErrInvalidRequest when we pass an invalid id (should be a UUID)
		// We can check if template is referenced by Name.
		if errors.Is(err, apiv2.ErrInvalidRequest) || errors.Is(err, apiv2.ErrNotFound) {
			template, err = c.GetTemplateByName(ctx, zone, x, visibilty)
			if err == nil {
				return template, nil
			}
		}
		return nil, err
	}
	return template, nil
}

// ListTemplates returns the list of existing Templates.
func (c *Client) ListTemplates(ctx context.Context, zone string, opts ...ListTemplatesOpt) ([]*Template, error) {
	list := make([]*Template, 0)

	defaultVisibility := oapi.TemplateVisibilityPublic
	params := oapi.ListTemplatesParams{
		Visibility: (*oapi.ListTemplatesParamsVisibility)(&defaultVisibility),
	}

	for _, opt := range opts {
		opt(&params)
	}

	resp, err := c.ListTemplatesWithResponse(apiv2.WithZone(ctx, zone), &params)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.Templates != nil {
		for i := range *resp.JSON200.Templates {
			list = append(list, templateFromAPI(&(*resp.JSON200.Templates)[i], zone))
		}
	}

	return list, nil
}

// RegisterTemplate registers a new Template.
func (c *Client) RegisterTemplate(ctx context.Context, zone string, template *Template) (*Template, error) {
	if err := validateOperationParams(template, "create"); err != nil {
		return nil, err
	}

	resp, err := c.RegisterTemplateWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.RegisterTemplateJSONRequestBody{
			BootMode:        (*oapi.RegisterTemplateJSONBodyBootMode)(template.BootMode),
			Build:           template.Build,
			Checksum:        *template.Checksum,
			DefaultUser:     template.DefaultUser,
			Description:     template.Description,
			Maintainer:      template.Maintainer,
			Name:            *template.Name,
			PasswordEnabled: *template.PasswordEnabled,
			SshKeyEnabled:   *template.SSHKeyEnabled,
			Url:             *template.URL,
			Version:         template.Version,
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetTemplate(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// UpdateTemplate updates a Template.
func (c *Client) UpdateTemplate(ctx context.Context, zone string, template *Template) error {
	if err := validateOperationParams(template, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateTemplateWithResponse(
		apiv2.WithZone(ctx, zone),
		*template.ID,
		oapi.UpdateTemplateJSONRequestBody{
			Description: oapi.NilableString(template.Description),
			Name:        template.Name,
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
