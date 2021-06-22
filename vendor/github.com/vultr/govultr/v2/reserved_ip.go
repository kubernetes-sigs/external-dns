package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const ripPath = "/v2/reserved-ips"

// ReservedIPService is the interface to interact with the reserved IP endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/reserved-ip
type ReservedIPService interface {
	Create(ctx context.Context, ripCreate *ReservedIPReq) (*ReservedIP, error)
	Get(ctx context.Context, id string) (*ReservedIP, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, options *ListOptions) ([]ReservedIP, *Meta, error)

	Convert(ctx context.Context, ripConvert *ReservedIPConvertReq) (*ReservedIP, error)
	Attach(ctx context.Context, id, instance string) error
	Detach(ctx context.Context, id string) error
}

// ReservedIPServiceHandler handles interaction with the reserved IP methods for the Vultr API
type ReservedIPServiceHandler struct {
	client *Client
}

// ReservedIP represents an reserved IP on Vultr
type ReservedIP struct {
	ID         string `json:"id"`
	Region     string `json:"region"`
	IPType     string `json:"ip_type"`
	Subnet     string `json:"subnet"`
	SubnetSize int    `json:"subnet_size"`
	Label      string `json:"label"`
	InstanceID string `json:"instance_id"`
}

// ReservedIPReq represents the parameters for creating a new Reserved IP on Vultr
type ReservedIPReq struct {
	Region     string `json:"region,omitempty"`
	IPType     string `json:"ip_type,omitempty"`
	IPAddress  string `json:"ip_address,omitempty"`
	Label      string `json:"label,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
}

type reservedIPsBase struct {
	ReservedIPs []ReservedIP `json:"reserved_ips"`
	Meta        *Meta        `json:"meta"`
}

type reservedIPBase struct {
	ReservedIP *ReservedIP `json:"reserved_ip"`
}

// ReservedIPConvertReq is the struct used for create and update calls.
type ReservedIPConvertReq struct {
	IPAddress string `json:"ip_address,omitempty"`
	Label     string `json:"label,omitempty"`
}

// Create adds the specified reserved IP to your Vultr account
func (r *ReservedIPServiceHandler) Create(ctx context.Context, ripCreate *ReservedIPReq) (*ReservedIP, error) {
	req, err := r.client.NewRequest(ctx, http.MethodPost, ripPath, ripCreate)
	if err != nil {
		return nil, err
	}

	rip := new(reservedIPBase)
	if err = r.client.DoWithContext(ctx, req, rip); err != nil {
		return nil, err
	}

	return rip.ReservedIP, nil
}

// Get gets the reserved IP associated with provided ID
func (r *ReservedIPServiceHandler) Get(ctx context.Context, id string) (*ReservedIP, error) {
	uri := fmt.Sprintf("%s/%s", ripPath, id)
	req, err := r.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	rip := new(reservedIPBase)
	if err = r.client.DoWithContext(ctx, req, rip); err != nil {
		return nil, err
	}

	return rip.ReservedIP, nil
}

// Delete removes the specified reserved IP from your Vultr account
func (r *ReservedIPServiceHandler) Delete(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s", ripPath, id)
	req, err := r.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return r.client.DoWithContext(ctx, req, nil)
}

// List lists all the reserved IPs associated with your Vultr account
func (r *ReservedIPServiceHandler) List(ctx context.Context, options *ListOptions) ([]ReservedIP, *Meta, error) {
	req, err := r.client.NewRequest(ctx, http.MethodGet, ripPath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	ips := new(reservedIPsBase)
	if err = r.client.DoWithContext(ctx, req, ips); err != nil {
		return nil, nil, err
	}

	return ips.ReservedIPs, ips.Meta, nil
}

// Convert an existing IP on a subscription to a reserved IP.
func (r *ReservedIPServiceHandler) Convert(ctx context.Context, ripConvert *ReservedIPConvertReq) (*ReservedIP, error) {
	uri := fmt.Sprintf("%s/convert", ripPath)
	req, err := r.client.NewRequest(ctx, http.MethodPost, uri, ripConvert)

	if err != nil {
		return nil, err
	}

	rip := new(reservedIPBase)
	if err = r.client.DoWithContext(ctx, req, rip); err != nil {
		return nil, err
	}

	return rip.ReservedIP, nil
}

// Attach a reserved IP to an existing subscription
func (r *ReservedIPServiceHandler) Attach(ctx context.Context, id, instance string) error {
	uri := fmt.Sprintf("%s/%s/attach", ripPath, id)
	reqBody := RequestBody{"instance_id": instance}
	req, err := r.client.NewRequest(ctx, http.MethodPost, uri, reqBody)
	if err != nil {
		return err
	}

	return r.client.DoWithContext(ctx, req, nil)
}

// Detach a reserved IP from an existing subscription.
func (r *ReservedIPServiceHandler) Detach(ctx context.Context, id string) error {
	uri := fmt.Sprintf("%s/%s/detach", ripPath, id)
	req, err := r.client.NewRequest(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return err
	}

	return r.client.DoWithContext(ctx, req, nil)
}
