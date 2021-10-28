package ibclient

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
)

type NotFoundError struct {
	msg string
}

func (e *NotFoundError) Error() string {
	return "not found"
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{msg: msg}
}

func BuildNetworkViewFromRef(ref string) *NetworkView {
	// networkview/ZG5zLm5ldHdvcmtfdmlldyQyMw:global_view/false
	r := regexp.MustCompile(`networkview/\w+:([^/]+)/\w+`)
	m := r.FindStringSubmatch(ref)

	if m == nil {
		return nil
	}

	return &NetworkView{
		Ref:  ref,
		Name: m[1],
	}
}

func BuildNetworkFromRef(ref string) (*Network, error) {
	// network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:89.0.0.0/24/global_view
	r := regexp.MustCompile(`network/\w+:(\d+\.\d+\.\d+\.\d+/\d+)/(.+)`)
	m := r.FindStringSubmatch(ref)

	if m == nil {
		return nil, fmt.Errorf("CIDR format not matched")
	}

	newNet := NewNetwork(m[2], m[1], false, "", nil)
	newNet.Ref = ref
	return newNet, nil
}

func GetIPAddressFromRef(ref string) string {
	// fixedaddress/ZG5zLmJpbmRfY25h:12.0.10.1/external
	r := regexp.MustCompile(`fixedaddress/\w+:(\d+\.\d+\.\d+\.\d+)/.+`)
	m := r.FindStringSubmatch(ref)

	if m != nil {
		return m[1]
	}
	return ""
}

// validation  for match_client
func validateMatchClient(value string) bool {
	matchClientList := [5]string{
		"MAC_ADDRESS",
		"CLIENT_ID",
		"RESERVED",
		"CIRCUIT_ID",
		"REMOTE_ID"}

	for _, val := range matchClientList {
		if val == value {
			return true
		}
	}
	return false
}

func BuildIPv6NetworkFromRef(ref string) (*Network, error) {
	// ipv6network/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:2001%3Adb8%3Aabcd%3A0012%3A%3A0/64/global_view
	r := regexp.MustCompile(`ipv6network/[^:]+:(([^\/]+)\/\d+)\/(.+)`)
	m := r.FindStringSubmatch(ref)

	if m == nil {
		return nil, fmt.Errorf("CIDR format not matched")
	}

	cidr, err := url.QueryUnescape(m[1])
	if err != nil {
		return nil, fmt.Errorf(
			"cannot extract network CIDR information from the reference '%s': %s",
			ref, err.Error())
	}

	if _, _, err = net.ParseCIDR(cidr); err != nil {
		return nil, fmt.Errorf("CIDR format not matched")
	}

	newNet := NewNetwork(m[3], cidr, true, "", nil)
	newNet.Ref = ref

	return newNet, nil
}
