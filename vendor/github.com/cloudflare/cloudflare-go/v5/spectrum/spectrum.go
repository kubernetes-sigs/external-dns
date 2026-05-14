// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// SpectrumService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpectrumService] method instead.
type SpectrumService struct {
	Options   []option.RequestOption
	Analytics *AnalyticsService
	Apps      *AppService
}

// NewSpectrumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpectrumService(opts ...option.RequestOption) (r *SpectrumService) {
	r = &SpectrumService{}
	r.Options = opts
	r.Analytics = NewAnalyticsService(opts...)
	r.Apps = NewAppService(opts...)
	return
}

// The name and type of DNS record for the Spectrum application.
type DNS struct {
	// The name of the DNS record associated with the application.
	Name string `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type DNSType `json:"type"`
	JSON dnsJSON `json:"-"`
}

// dnsJSON contains the JSON metadata for the struct [DNS]
type dnsJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the application.
type DNSType string

const (
	DNSTypeCNAME   DNSType = "CNAME"
	DNSTypeAddress DNSType = "ADDRESS"
)

func (r DNSType) IsKnown() bool {
	switch r {
	case DNSTypeCNAME, DNSTypeAddress:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type DNSParam struct {
	// The name of the DNS record associated with the application.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The type of DNS record associated with the application.
	Type param.Field[DNSType] `json:"type"`
}

func (r DNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The anycast edge IP configuration for the hostname of this application.
type EdgeIPs struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity EdgeIPsConnectivity `json:"connectivity"`
	// This field can have the runtime type of [[]string].
	IPs interface{} `json:"ips"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type  EdgeIPsType `json:"type"`
	JSON  edgeIPsJSON `json:"-"`
	union EdgeIPsUnion
}

// edgeIPsJSON contains the JSON metadata for the struct [EdgeIPs]
type edgeIPsJSON struct {
	Connectivity apijson.Field
	IPs          apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r edgeIPsJSON) RawJSON() string {
	return r.raw
}

func (r *EdgeIPs) UnmarshalJSON(data []byte) (err error) {
	*r = EdgeIPs{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [EdgeIPsUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [EdgeIPsDynamic], [EdgeIPsStatic].
func (r EdgeIPs) AsUnion() EdgeIPsUnion {
	return r.union
}

// The anycast edge IP configuration for the hostname of this application.
//
// Union satisfied by [EdgeIPsDynamic] or [EdgeIPsStatic].
type EdgeIPsUnion interface {
	implementsEdgeIPs()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EdgeIPsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EdgeIPsDynamic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(EdgeIPsStatic{}),
		},
	)
}

type EdgeIPsDynamic struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity EdgeIPsDynamicConnectivity `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type EdgeIPsDynamicType `json:"type"`
	JSON edgeIPsDynamicJSON `json:"-"`
}

// edgeIPsDynamicJSON contains the JSON metadata for the struct [EdgeIPsDynamic]
type edgeIPsDynamicJSON struct {
	Connectivity apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EdgeIPsDynamic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeIPsDynamicJSON) RawJSON() string {
	return r.raw
}

func (r EdgeIPsDynamic) implementsEdgeIPs() {}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type EdgeIPsDynamicConnectivity string

const (
	EdgeIPsDynamicConnectivityAll  EdgeIPsDynamicConnectivity = "all"
	EdgeIPsDynamicConnectivityIPV4 EdgeIPsDynamicConnectivity = "ipv4"
	EdgeIPsDynamicConnectivityIPV6 EdgeIPsDynamicConnectivity = "ipv6"
)

func (r EdgeIPsDynamicConnectivity) IsKnown() bool {
	switch r {
	case EdgeIPsDynamicConnectivityAll, EdgeIPsDynamicConnectivityIPV4, EdgeIPsDynamicConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type EdgeIPsDynamicType string

const (
	EdgeIPsDynamicTypeDynamic EdgeIPsDynamicType = "dynamic"
)

func (r EdgeIPsDynamicType) IsKnown() bool {
	switch r {
	case EdgeIPsDynamicTypeDynamic:
		return true
	}
	return false
}

type EdgeIPsStatic struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs []string `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type EdgeIPsStaticType `json:"type"`
	JSON edgeIPsStaticJSON `json:"-"`
}

// edgeIPsStaticJSON contains the JSON metadata for the struct [EdgeIPsStatic]
type edgeIPsStaticJSON struct {
	IPs         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EdgeIPsStatic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r edgeIPsStaticJSON) RawJSON() string {
	return r.raw
}

func (r EdgeIPsStatic) implementsEdgeIPs() {}

// The type of edge IP configuration specified. Statically allocated edge IPs use
// customer IPs in accordance with the ips array you specify. Only valid with
// ADDRESS DNS names.
type EdgeIPsStaticType string

const (
	EdgeIPsStaticTypeStatic EdgeIPsStaticType = "static"
)

func (r EdgeIPsStaticType) IsKnown() bool {
	switch r {
	case EdgeIPsStaticTypeStatic:
		return true
	}
	return false
}

// The IP versions supported for inbound connections on Spectrum anycast IPs.
type EdgeIPsConnectivity string

const (
	EdgeIPsConnectivityAll  EdgeIPsConnectivity = "all"
	EdgeIPsConnectivityIPV4 EdgeIPsConnectivity = "ipv4"
	EdgeIPsConnectivityIPV6 EdgeIPsConnectivity = "ipv6"
)

func (r EdgeIPsConnectivity) IsKnown() bool {
	switch r {
	case EdgeIPsConnectivityAll, EdgeIPsConnectivityIPV4, EdgeIPsConnectivityIPV6:
		return true
	}
	return false
}

// The type of edge IP configuration specified. Dynamically allocated edge IPs use
// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
// with CNAME DNS names.
type EdgeIPsType string

const (
	EdgeIPsTypeDynamic EdgeIPsType = "dynamic"
	EdgeIPsTypeStatic  EdgeIPsType = "static"
)

func (r EdgeIPsType) IsKnown() bool {
	switch r {
	case EdgeIPsTypeDynamic, EdgeIPsTypeStatic:
		return true
	}
	return false
}

// The anycast edge IP configuration for the hostname of this application.
type EdgeIPsParam struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[EdgeIPsConnectivity] `json:"connectivity"`
	IPs          param.Field[interface{}]         `json:"ips"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[EdgeIPsType] `json:"type"`
}

func (r EdgeIPsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EdgeIPsParam) implementsEdgeIPsUnionParam() {}

// The anycast edge IP configuration for the hostname of this application.
//
// Satisfied by [spectrum.EdgeIPsDynamicParam], [spectrum.EdgeIPsStaticParam],
// [EdgeIPsParam].
type EdgeIPsUnionParam interface {
	implementsEdgeIPsUnionParam()
}

type EdgeIPsDynamicParam struct {
	// The IP versions supported for inbound connections on Spectrum anycast IPs.
	Connectivity param.Field[EdgeIPsDynamicConnectivity] `json:"connectivity"`
	// The type of edge IP configuration specified. Dynamically allocated edge IPs use
	// Spectrum anycast IPs in accordance with the connectivity you specify. Only valid
	// with CNAME DNS names.
	Type param.Field[EdgeIPsDynamicType] `json:"type"`
}

func (r EdgeIPsDynamicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EdgeIPsDynamicParam) implementsEdgeIPsUnionParam() {}

type EdgeIPsStaticParam struct {
	// The array of customer owned IPs we broadcast via anycast for this hostname and
	// application.
	IPs param.Field[[]string] `json:"ips"`
	// The type of edge IP configuration specified. Statically allocated edge IPs use
	// customer IPs in accordance with the ips array you specify. Only valid with
	// ADDRESS DNS names.
	Type param.Field[EdgeIPsStaticType] `json:"type"`
}

func (r EdgeIPsStaticParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r EdgeIPsStaticParam) implementsEdgeIPsUnionParam() {}

// The name and type of DNS record for the Spectrum application.
type OriginDNS struct {
	// The name of the DNS record associated with the origin.
	Name string `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL int64 `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type OriginDNSType `json:"type"`
	JSON originDNSJSON `json:"-"`
}

// originDNSJSON contains the JSON metadata for the struct [OriginDNS]
type originDNSJSON struct {
	Name        apijson.Field
	TTL         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originDNSJSON) RawJSON() string {
	return r.raw
}

// The type of DNS record associated with the origin. "" is used to specify a
// combination of A/AAAA records.
type OriginDNSType string

const (
	OriginDNSTypeEmpty OriginDNSType = ""
	OriginDNSTypeA     OriginDNSType = "A"
	OriginDNSTypeAAAA  OriginDNSType = "AAAA"
	OriginDNSTypeSRV   OriginDNSType = "SRV"
)

func (r OriginDNSType) IsKnown() bool {
	switch r {
	case OriginDNSTypeEmpty, OriginDNSTypeA, OriginDNSTypeAAAA, OriginDNSTypeSRV:
		return true
	}
	return false
}

// The name and type of DNS record for the Spectrum application.
type OriginDNSParam struct {
	// The name of the DNS record associated with the origin.
	Name param.Field[string] `json:"name" format:"hostname"`
	// The TTL of our resolution of your DNS record in seconds.
	TTL param.Field[int64] `json:"ttl"`
	// The type of DNS record associated with the origin. "" is used to specify a
	// combination of A/AAAA records.
	Type param.Field[OriginDNSType] `json:"type"`
}

func (r OriginDNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Union satisfied by [shared.UnionInt] or [shared.UnionString].
type OriginPortUnion interface {
	ImplementsOriginPortUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OriginPortUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination port at the origin. Only specified in conjunction with
// origin_dns. May use an integer to specify a single origin port, for example
// `1000`, or a string to specify a range of origin ports, for example
// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
// range must match the number of ports specified in the "protocol" field.
//
// Satisfied by [shared.UnionInt], [shared.UnionString].
type OriginPortUnionParam interface {
	ImplementsOriginPortUnionParam()
}
