package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/ipam"
)

// IPAMService handles the 'ipam' endpoint.
type IPAMService service

// ListAddrs returns a list of all root addresses (i.e. Parent == 0) in all
// networks.
//
// NS1 API docs: https://ns1.com/api#getview-a-list-of-root-addresses
func (s *IPAMService) ListAddrs() ([]ipam.Address, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "ipam/address", nil)
	if err != nil {
		return nil, nil, err
	}

	addrs := []ipam.Address{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &addrs, s.nextAddrs)
	} else {
		resp, err = s.client.Do(req, &addrs)
	}
	if err != nil {
		return nil, resp, err
	}

	return addrs, resp, nil
}

// GetSubnet returns the subnet corresponding to the provided address ID.
//
// NS1 API docs: https://ns1.com/api#getview-a-subnet
func (s *IPAMService) GetSubnet(addrID int) (*ipam.Address, *http.Response, error) {
	reqPath := fmt.Sprintf("ipam/address/%d", addrID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	addr := &ipam.Address{}
	var resp *http.Response
	resp, err = s.client.Do(req, addr)
	if err != nil {
		return nil, resp, err
	}

	return addr, resp, nil
}

// GetChildren requests a list of all child addresses (or subnets) for the
// specified IP address.
//
// NS1 API docs: https://ns1.com/api#getview-address-children
func (s *IPAMService) GetChildren(addrID int) ([]*ipam.Address, *http.Response, error) {
	reqPath := fmt.Sprintf("ipam/address/%d/children", addrID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	addrs := []*ipam.Address{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &addrs, s.nextAddrs)
	} else {
		resp, err = s.client.Do(req, &addrs)
	}
	if err != nil {
		return nil, resp, err
	}

	return addrs, resp, nil
}

// GetParent fetches the addresses parent.
//
// NS1 API docs: https://ns1.com/api#getview-address-parent
func (s *IPAMService) GetParent(addrID int) (*ipam.Address, *http.Response, error) {
	reqPath := fmt.Sprintf("ipam/address/%d/parent", addrID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	addr := &ipam.Address{}
	var resp *http.Response
	resp, err = s.client.Do(req, addr)
	if err != nil {
		return nil, resp, err
	}

	return addr, resp, nil
}

// CreateSubnet creates an address or subnet.
// The Prefix and Network fields are required.
//
// NS1 API docs: https://ns1.com/api#putcreate-a-subnet
func (s *IPAMService) CreateSubnet(addr *ipam.Address) (*ipam.Address, *http.Response, error) {
	switch {
	case addr.Prefix == "":
		return nil, nil, errors.New("the Prefix field is required")
	case addr.Network == 0:
		return nil, nil, errors.New("the Network field is required")
	}

	req, err := s.client.NewRequest(http.MethodPut, "ipam/address", addr)
	if err != nil {
		return nil, nil, err
	}

	respAddr := &ipam.Address{}
	var resp *http.Response
	resp, err = s.client.Do(req, respAddr)
	if err != nil {
		return nil, resp, err
	}

	return respAddr, resp, nil
}

// EditSubnet updates an existing subnet.
// The ID field is required.
// Parent is whether or not to include the parent in the parent field.
//
// NS1 API docs: https://ns1.com/api#postedit-a-subnet
func (s *IPAMService) EditSubnet(addr *ipam.Address, parent bool) (newAddr, parentAddr *ipam.Address, resp *http.Response, err error) {
	if addr.ID == 0 {
		return nil, nil, nil, errors.New("the ID field is required")
	}

	reqPath := fmt.Sprintf("ipam/address/%d", addr.ID)
	req, err := s.client.NewRequest(http.MethodPost, reqPath, addr)
	if err != nil {
		return nil, nil, nil, err
	}
	if parent {
		q := req.URL.Query()
		q.Add("parent", "true")
		req.URL.RawQuery = q.Encode()
	}

	data := struct {
		ipam.Address
		Parent ipam.Address `json:"parent"`
	}{}
	resp, err = s.client.Do(req, &data)
	if err != nil {
		return nil, nil, resp, err
	}

	if parent {
		return &data.Address, &data.Parent, resp, nil
	}
	return &data.Address, nil, resp, nil
}

// SplitSubnet splits a block of unassigned IP space into equal pieces.
// This will not function with ranges or individual hosts. Normal breaking out
// of a subnet is done with the standard PUT route. (Eg. root address is a /24
// and request for /29s will break it into 32 /29s)
//
//   - Only planned subnets can be split
//   - Name and description will be unset on children
//   - KVPS and options will be copied; tags will be inherited
//
// NS1 API docs: https://ns1.com/api#postsplit-a-subnet
func (s *IPAMService) SplitSubnet(id, prefix int) (rootAddr int, prefixIDs []int, resp *http.Response, err error) {
	reqPath := fmt.Sprintf("ipam/address/%d/split", id)
	req, err := s.client.NewRequest(http.MethodPost, reqPath, struct {
		Prefix int `json:"prefix"`
	}{
		Prefix: prefix,
	})
	if err != nil {
		return 0, nil, nil, err
	}

	data := &struct {
		RootAddr  int   `json:"root_address_id"`
		PrefixIDs []int `json:"prefix_ids"`
	}{}
	resp, err = s.client.Do(req, &data)
	return data.RootAddr, data.PrefixIDs, resp, err
}

// MergeSubnet merges several subnets together.
//
// NS1 API docs: https://ns1.com/api#postmerge-a-subnet
func (s *IPAMService) MergeSubnet(rootID, mergeID int) (*ipam.Address, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "ipam/address/merge", struct {
		Root  int `json:"root_address_id"`
		Merge int `json:"merged_address_id"`
	}{
		Root:  rootID,
		Merge: mergeID,
	})
	if err != nil {
		return nil, nil, err
	}

	addr := &ipam.Address{}
	resp, err := s.client.Do(req, &addr)
	return addr, resp, err
}

// DeleteSubnet removes a subnet entirely.
//
// NS1 API docs: https://ns1.com/api#deletedelete-a-subnet
func (s *IPAMService) DeleteSubnet(id int) (*http.Response, error) {
	reqPath := fmt.Sprintf("ipam/address/%d", id)
	req, err := s.client.NewRequest(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// nextAddrs is a pagination helper than gets and appends another list of
// addresses to the passed list.
func (s *IPAMService) nextAddrs(v *interface{}, uri string) (*http.Response, error) {
	addrs := []*ipam.Address{}
	resp, err := s.client.getURI(&addrs, uri)
	if err != nil {
		return resp, err
	}
	addrList, ok := (*v).(*[]*ipam.Address)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *[]ipam.Address, got: %T", v,
		)
	}
	*addrList = append(*addrList, addrs...)
	return resp, nil
}
