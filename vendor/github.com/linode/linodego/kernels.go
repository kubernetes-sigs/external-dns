package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// LinodeKernel represents a Linode Instance kernel object
type LinodeKernel struct {
<<<<<<< HEAD
	ID           string `json:"id"`
	Label        string `json:"label"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Deprecated   bool   `json:"deprecated"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	Deprecated   bool   `json:"deprecated"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	Deprecated   bool   `json:"deprecated"`
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	Deprecated   bool   `json:"deprecated"`
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	KVM          bool   `json:"kvm"`
	XEN          bool   `json:"xen"`
	PVOPS        bool   `json:"pvops"`
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	ID           string `json:"id"`
	Label        string `json:"label"`
	Version      string `json:"version"`
	Architecture string `json:"architecture"`
	KVM          bool   `json:"kvm"`
	XEN          bool   `json:"xen"`
	PVOPS        bool   `json:"pvops"`
=======
	ID           string     `json:"id"`
	Label        string     `json:"label"`
	Version      string     `json:"version"`
	Architecture string     `json:"architecture"`
	Deprecated   bool       `json:"deprecated"`
	KVM          bool       `json:"kvm"`
	XEN          bool       `json:"xen"`
	PVOPS        bool       `json:"pvops"`
	Built        *time.Time `json:"-"`
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

// LinodeKernelsPagedResponse represents a Linode kernels API response for listing
type LinodeKernelsPagedResponse struct {
	*PageOptions
	Data []LinodeKernel `json:"data"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *LinodeKernel) UnmarshalJSON(b []byte) error {
	type Mask LinodeKernel

	p := struct {
		*Mask
		Built *parseabletime.ParseableTime `json:"built"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Built = (*time.Time)(p.Built)

	return nil
}

func (LinodeKernelsPagedResponse) endpoint(_ ...any) string {
	return "linode/kernels"
}

func (resp *LinodeKernelsPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(LinodeKernelsPagedResponse{}).Get(e))
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*LinodeKernelsPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListKernels lists linode kernels. This endpoint is cached by default.
func (c *Client) ListKernels(ctx context.Context, opts *ListOptions) ([]LinodeKernel, error) {
	response := LinodeKernelsPagedResponse{}

	endpoint, err := generateListCacheURL(response.endpoint(), opts)
	if err != nil {
		return nil, err
	}

	if result := c.getCachedResponse(endpoint); result != nil {
		return result.([]LinodeKernel), nil
	}

	err = c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(endpoint, response.Data, nil)

	return response.Data, nil
}

// GetKernel gets the kernel with the provided ID. This endpoint is cached by default.
func (c *Client) GetKernel(ctx context.Context, kernelID string) (*LinodeKernel, error) {
	kernelID = url.PathEscape(kernelID)
	e := fmt.Sprintf("linode/kernels/%s", kernelID)

	if result := c.getCachedResponse(e); result != nil {
		result := result.(LinodeKernel)
		return &result, nil
	}

	req := c.R(ctx).SetResult(&LinodeKernel{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}

	c.addCachedResponse(e, r.Result(), nil)

	return r.Result().(*LinodeKernel), nil
}
