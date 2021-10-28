package godo

import (
	"context"
	"net/http"
)

// SizesService is an interface for interfacing with the size
// endpoints of the DigitalOcean API
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes
type SizesService interface {
	List(context.Context, *ListOptions) ([]Size, *Response, error)
}

// SizesServiceOp handles communication with the size related methods of the
// DigitalOcean API.
type SizesServiceOp struct {
	client *Client
}

var _ SizesService = &SizesServiceOp{}

// Size represents a DigitalOcean Size
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
	Description  string   `json:"description,omitempty"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#sizes
||||||| parent of 5ce8c7613 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#sizes
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes
>>>>>>> 5ce8c7613 (update vendored files)
type SizesService interface {
	List(context.Context, *ListOptions) ([]Size, *Response, error)
}

// SizesServiceOp handles communication with the size related methods of the
// DigitalOcean API.
type SizesServiceOp struct {
	client *Client
}

var _ SizesService = &SizesServiceOp{}

// Size represents a DigitalOcean Size
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	Description  string   `json:"description,omitempty"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#sizes
||||||| parent of 6b7ce455e (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#sizes
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes
>>>>>>> 6b7ce455e (update vendored files)
type SizesService interface {
	List(context.Context, *ListOptions) ([]Size, *Response, error)
}

// SizesServiceOp handles communication with the size related methods of the
// DigitalOcean API.
type SizesServiceOp struct {
	client *Client
}

var _ SizesService = &SizesServiceOp{}

// Size represents a DigitalOcean Size
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	Description  string   `json:"description,omitempty"`
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#sizes
||||||| parent of 4d7e5ad26 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#sizes
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes
>>>>>>> 4d7e5ad26 (update vendored files)
type SizesService interface {
	List(context.Context, *ListOptions) ([]Size, *Response, error)
}

// SizesServiceOp handles communication with the size related methods of the
// DigitalOcean API.
type SizesServiceOp struct {
	client *Client
}

var _ SizesService = &SizesServiceOp{}

// Size represents a DigitalOcean Size
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	PriceMonthly float64  `json:"price_monthly,omitempty"`
	PriceHourly  float64  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	Available    bool     `json:"available,omitempty"`
	Transfer     float64  `json:"transfer,omitempty"`
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	Description  string   `json:"description,omitempty"`
>>>>>>> 4d7e5ad26 (update vendored files)
}

func (s Size) String() string {
	return Stringify(s)
}

type sizesRoot struct {
	Sizes []Size
	Links *Links `json:"links"`
	Meta  *Meta  `json:"meta"`
}

// List all images
func (s *SizesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Size, *Response, error) {
	path := "v2/sizes"
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(sizesRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	if l := root.Links; l != nil {
		resp.Links = l
	}
	if m := root.Meta; m != nil {
		resp.Meta = m
	}

	return root.Sizes, resp, err
}
