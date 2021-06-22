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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageID    string `json:"image_id"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageID    string `json:"image_id"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageID    string `json:"image_id"`
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	Type       string `json:"type"`
	Vendor     string `json:"vendor"`
	ImageID    string `json:"image_id"`
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
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
