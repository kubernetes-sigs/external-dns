package dhcp

// Option encapsulates DHCPv4 and DHCPv6 option information.
//
// Serialized:
//
//	'{"name": "dhcpv4/routers", "value": ["127.0.0.1"]}'
//	'{"name": "dhcpv4/boot-file-name", "value": "/bootfilename"}'
//	'{"name": "dhcpv6/dns-servers", "value": ["2001:db8::cafe"]}'
type Option struct {
	Name        string      `json:"name"`
	Value       interface{} `json:"value"`
	AlwaysSend  *bool       `json:"always_send,omitempty"`
	Encapsulate *string     `json:"encapsulate,omitempty"`
}

// OptionSet is a convenience type for marshalling an array of options to and from a JSON field.
type OptionSet []Option

// ItemType defines the type of the item as a string
type ItemType string

// List of available types
const (
	ItemTypeBinary      = "binary"
	ItemTypeBoolean     = "boolean"
	ItemTypeEmpty       = "empty"
	ItemTypeFQDN        = "fqdn"
	ItemTypeInt16       = "int16"
	ItemTypeInt32       = "int32"
	ItemTypeInt8        = "int8"
	ItemTypeIPv4Address = "ipv4_address"
	ItemTypeIPv6Address = "ipv6_address"
	ItemTypeIPv6Prefix  = "ipv6_prefix"
	ItemTypePSID        = "psid"
	ItemTypeString      = "string"
	ItemTypeTuple       = "tuple"
	ItemTypeUint16      = "uint16"
	ItemTypeUint32      = "uint32"
	ItemTypeUint8       = "uint8"
)

// OptionDefSchemaItems is generated from https://apispec.ns1.com/v1/dhcp/schemas/option-definition-fields.json#/properties/schema/properties/fields/items
type OptionDefSchemaItems struct {
	Name string   `json:"name"`
	Type ItemType `json:"type"`
}

// SchemaType defines the type of the schema as a string
type SchemaType string

// List of available types
const (
	SchemaTypeArray       SchemaType = "array"
	SchemaTypeBinary                 = "binary"
	SchemaTypeBoolean                = "boolean"
	SchemaTypeEmpty                  = "empty"
	SchemaTypeFQDN                   = "fqdn"
	SchemaTypeInt16                  = "int16"
	SchemaTypeInt32                  = "int32"
	SchemaTypeInt8                   = "int8"
	SchemaTypeIPv4Address            = "ipv4_address"
	SchemaTypeIPv6Address            = "ipv6_address"
	SchemaTypeIPv6Prefix             = "ipv6_prefix"
	SchemaTypePSID                   = "psid"
	SchemaTypeRecord                 = "record"
	SchemaTypeString                 = "string"
	SchemaTypeTuple                  = "tuple"
	SchemaTypeUint16                 = "uint16"
	SchemaTypeUint32                 = "uint32"
	SchemaTypeUint8                  = "uint8"
)

// OptionDefSchema is a schema of the option def which describes the value that option can hold
type OptionDefSchema struct {
	Fields             []OptionDefSchemaItems `json:"fields"`
	Items              *string                `json:"items,omitempty"`
	MultipleFinalValue *bool                  `json:"multiple_final_value,omitempty"`
	Type               SchemaType             `json:"type"`
}

// OptionDef configures a custom option definition
// https://ftp.isc.org/isc/kea/1.4.0/doc/kea-guide.html#dhcp4-custom-options
type OptionDef struct {
	Space        *string         `json:"space,omitempty"`
	FriendlyName string          `json:"friendly_name"`
	Description  string          `json:"description"`
	Code         int             `json:"code"`
	Encapsulate  *string         `json:"encapsulate,omitempty"`
	Schema       OptionDefSchema `json:"schema"`
}
