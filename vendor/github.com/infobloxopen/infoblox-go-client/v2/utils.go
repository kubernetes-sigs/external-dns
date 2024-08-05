package ibclient

import (
	"fmt"
	"net"
	"net/url"
	"regexp"
<<<<<<< HEAD
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"strings"

	"golang.org/x/net/idna"
)

type NotFoundError struct {
	msg string
}

func (e *NotFoundError) Error() string {
	return e.msg
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{msg: msg}
}

type GenericObj interface {
	ObjectType() string
	ReturnFields() []string
	EaSearch() EASearch
	SetReturnFields([]string)
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
		Name: &m[1],
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

func BuildNetworkContainerFromRef(ref string) (*NetworkContainer, error) {
	// networkcontainer/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:89.0.0.0/24/global_view
	r := regexp.MustCompile(`networkcontainer/\w+:(\d+\.\d+\.\d+\.\d+/\d+)/(.+)`)
	m := r.FindStringSubmatch(ref)

	if m == nil {
		return nil, fmt.Errorf("CIDR format not matched")
	}

	newNet := NewNetworkContainer(m[2], m[1], false, "", nil)
	newNet.Ref = ref
	return newNet, nil
}

func BuildIPv6NetworkContainerFromRef(ref string) (*NetworkContainer, error) {
	// ipv6networkcontainer/ZG5zLm5ldHdvcmskODkuMC4wLjAvMjQvMjU:2001%3Adb8%3Aabcd%3A0012%3A%3A0/64/global_view
	r := regexp.MustCompile(`ipv6networkcontainer/[^:]+:(([^\/]+)\/\d+)\/(.+)`)
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

	newNet := NewNetworkContainer(m[3], cidr, true, "", nil)
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

const dnsLabelFormat = "[a-z0-9]+(([a-z0-9-]*[a-z0-9]+))?"

// ValidateDomainName return an error if the domain name does not conform to standards.
// The domain name may be in Unicode format (internationalized domain name)
func ValidateDomainName(name string) error {
	domainRegexpTemplate := fmt.Sprintf("^(?i)%s(\\.%s)*\\.?$", dnsLabelFormat, dnsLabelFormat)
	domainRegexp := regexp.MustCompile(domainRegexpTemplate)

	_, err := idna.ToASCII(name)
	if err != nil {
		return err
	}

	if !domainRegexp.MatchString(name) {
		return fmt.Errorf("the name '%s' is not a valid domain name", name)
	}

	return nil
}

// ValidateSrvRecName return an error if the record's name does not conform to standards.
func ValidateSrvRecName(name string) error {
	const protoLabelFormat = "[a-z0-9]+"

	const errorMsgFormat = "SRV-record's name '%s' does not conform to standards"
	var (
		srvNamePartRegExp  = regexp.MustCompile(fmt.Sprintf("^_%s", dnsLabelFormat))
		srvProtoPartRegExp = regexp.MustCompile(fmt.Sprintf("^_%s", protoLabelFormat))
	)

	nameParts := strings.SplitN(name, ".", 3)
	if len(nameParts) != 3 {
		return fmt.Errorf(errorMsgFormat, name)
	}
	if !srvNamePartRegExp.MatchString(nameParts[0]) {
		return fmt.Errorf(errorMsgFormat, name)
	}
	if !srvProtoPartRegExp.MatchString(nameParts[1]) {
		return fmt.Errorf(errorMsgFormat, name)
	}
	if err := ValidateDomainName(nameParts[2]); err != nil {
		return err
	}

	return nil
}

func CheckIntRange(name string, value int, min int, max int) error {
	if value < min || value > max {
		return fmt.Errorf("'%s' must be integer and must be in the range from 0 to 65535 inclusively", name)
	}

	return nil
}

func ValidateMultiValue(v string) ([]string, bool) {
	res := strings.Split(v, ",")
	if len(res) > 1 {
		return res, true
	} else {
		return nil, false
	}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
