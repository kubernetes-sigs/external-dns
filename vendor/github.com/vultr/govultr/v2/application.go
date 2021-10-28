package govultr

import (
	"context"
	"net/http"

	"github.com/google/go-querystring/query"
)

// ApplicationService is the interface to interact with the Application endpoint on the Vultr API.
// Link : https://www.vultr.com/api/#tag/application
type ApplicationService interface {
	List(ctx context.Context, options *ListOptions) ([]Application, *Meta, error)
}

// ApplicationServiceHandler handles interaction with the application methods for the Vultr API.
type ApplicationServiceHandler struct {
	client *Client
}

// Application represents all available apps that can be used to deployed with vultr Instances.
type Application struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	DeployName string `json:"deploy_name"`
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageID    string `json:"image_id"`
}

type applicationBase struct {
	Applications []Application `json:"applications"`
	Meta         *Meta         `json:"meta"`
}

// List retrieves a list of available applications that can be launched when creating a Vultr instance
func (a *ApplicationServiceHandler) List(ctx context.Context, options *ListOptions) ([]Application, *Meta, error) {
	uri := "/v2/applications"

	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()
	apps := new(applicationBase)

	err = a.client.DoWithContext(ctx, req, apps)
	if err != nil {
		return nil, nil, err
	}

	return apps.Applications, apps.Meta, nil
}
