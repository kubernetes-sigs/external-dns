package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// FeedPtr represents the dynamic metadata value in which a feed is providing the value.
type FeedPtr struct {
	FeedID string `json:"feed,omitempty"`
}

// PulsarMeta is currently only used for validation
type PulsarMeta struct {
	JobID     string  `json:"job_id,omitempty"`
	Bias      string  `json:"bias,omitempty"`
	A5MCutoff float64 `json:"a5m_cutoff,omitempty"`
}

// Meta contains information on an entity's metadata table. Metadata key/value
// pairs are used by a record's filter pipeline during a dns query.
// All values can be a feed id as well, indicating real-time updates of these values.
// Structure/Precedence of metadata tables:
//   - Record
//   - Meta <- lowest precedence in filter
//   - Region(s)
//   - Meta <- middle precedence in filter chain
//   - ...
//   - Answer(s)
//   - Meta <- highest precedence in filter chain
//   - ...
//   - ...
type Meta struct {
	// STATUS

	// Indicates whether or not entity is considered 'up'
	// bool or FeedPtr.
	Up interface{} `json:"up,omitempty"`

	// Indicates the number of active connections.
	// Values must be positive.
	// int or FeedPtr.
	Connections interface{} `json:"connections,omitempty"`

	// Indicates the number of active requests (HTTP or otherwise).
	// Values must be positive.
	// int or FeedPtr.
	Requests interface{} `json:"requests,omitempty"`

	// Indicates the "load average".
	// Values must be positive, and will be rounded to the nearest tenth.
	// float64 or FeedPtr.
	LoadAvg interface{} `json:"loadavg,omitempty"`

	// The Job ID of a Pulsar telemetry gathering job and associated metadata.
	// list of PulsarMeta
	Pulsar interface{} `json:"pulsar,omitempty"`

	// GEOGRAPHICAL

	// Must be between -180.0 and +180.0 where negative
	// indicates South and positive indicates North.
	// e.g., the longitude of the datacenter where a server resides.
	// float64 or FeedPtr.
	Latitude interface{} `json:"latitude,omitempty"`

	// Must be between -180.0 and +180.0 where negative
	// indicates West and positive indicates East.
	// e.g., the longitude of the datacenter where a server resides.
	// float64 or FeedPtr.
	Longitude interface{} `json:"longitude,omitempty"`

	// Valid geographic regions are: 'US-EAST', 'US-CENTRAL', 'US-WEST',
	// 'EUROPE', 'ASIAPAC', 'SOUTH-AMERICA', 'AFRICA'.
	// e.g., the rough geographic location of the Datacenter where a server resides.
	// []string or FeedPtr.
	Georegion interface{} `json:"georegion,omitempty"`

	// Countr(ies) must be specified as ISO3166 2-character country code(s).
	// []string or FeedPtr.
	Country interface{} `json:"country,omitempty"`

	// State(s) must be specified as standard 2-character state code(s).
	// []string or FeedPtr.
	USState interface{} `json:"us_state,omitempty"`

	// Canadian Province(s) must be specified as standard 2-character province
	// code(s).
	// []string or FeedPtr.
	CAProvince interface{} `json:"ca_province,omitempty"`

	// INFORMATIONAL

	// Notes to indicate any necessary details for operators.
	// Up to 256 characters in length.
	// string or FeedPtr.
	Note interface{} `json:"note,omitempty"`

	// NETWORK

	// IP (v4 and v6) prefixes in CIDR format ("a.b.c.d/mask").
	// May include up to 1000 prefixes.
	// e.g., "1.2.3.4/24"
	// []string or FeedPtr.
	IPPrefixes interface{} `json:"ip_prefixes,omitempty"`

	// Autonomous System (AS) number(s).
	// May include up to 1000 AS numbers.
	// []string or FeedPtr.
	ASN interface{} `json:"asn,omitempty"`

	// TRAFFIC

	// Indicates the "priority tier".
	// Lower values indicate higher priority.
	// Values must be positive.
	// int or FeedPtr.
	Priority interface{} `json:"priority,omitempty"`

	// Indicates a weight.
	// Filters that use weights normalize them.
	// Any positive values are allowed.
	// Values between 0 and 100 are recommended for simplicity's sake.
	// float64 or FeedPtr.
	Weight interface{} `json:"weight,omitempty"`

	// Indicates a cost.
	// Filters that use costs normalize them.
	// Any positive values are allowed.
	// float64 or FeedPtr.
	Cost interface{} `json:"cost,omitempty"`

	// Indicates a "low watermark" to use for load shedding.
	// The value should depend on the metric used to determine
	// load (e.g., loadavg, connections, etc).
	// int or FeedPtr.
	LowWatermark interface{} `json:"low_watermark,omitempty"`

	// Indicates a "high watermark" to use for load shedding.
	// The value should depend on the metric used to determine
	// load (e.g., loadavg, connections, etc).
	// int or FeedPtr.
	HighWatermark interface{} `json:"high_watermark,omitempty"`

	// subdivisions must follow the ISO-3166-2 code for a country and subdivisions
	// map[string]interface{} or FeedPtr.
	Subdivisions interface{} `json:"subdivisions,omitempty"`

	AdditionalMetadata interface{} `json:"additional_metadata,omitempty"`
}

// StringMap returns a map[string]interface{} representation of metadata (for use with terraform in nested structures)
func (meta *Meta) StringMap() map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.Indirect(reflect.ValueOf(meta))
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fv := v.Field(i)
		if fv.IsNil() {
			continue
		}
		tag := f.Tag.Get("json")

		tag = strings.Split(tag, ",")[0]

		m[tag] = FormatInterface(fv.Interface())
	}
	return m
}

// FormatInterface takes an interface of types: string, bool, int, float64, []string, map[string]interface{} and FeedPtr, and returns a string representation of said interface
func FormatInterface(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case bool:
		if v {
			return "1"
		}
		return "0"
	case int:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []string:
		return strings.Join(v, ",")
	case []interface{}:
		slc := make([]string, 0)
		for _, s := range v {
			switch ss := s.(type) {
			// Pulsar
			case map[string]interface{}:
				data, _ := json.Marshal(v)
				return string(data)
			case string:
				slc = append(slc, ss)
			// The ASN field specifically is returned from the API as an integer,
			// which Go treats as a float64 when it parses the json,
			// so this is to account for that field.
			case float64:
				slc = append(slc, strconv.FormatFloat(ss, 'f', -1, 64))
			}
		}
		return strings.Join(slc, ",")
	case map[string]interface{}:
		// Required for Terraform workaround to allow users to submit raw json of feed pointer
		// as value for metadata.  See https://github.com/terraform-providers/terraform-provider-ns1/issues/35
		if val, ok := v["feed"].(string); ok {
			feedPtr := FeedPtr{FeedID: val}
			data, _ := json.Marshal(feedPtr)
			return string(data)
		}
		data, _ := json.Marshal(v)
		return string(data)
	case FeedPtr:
		data, _ := json.Marshal(v)
		return string(data)
	default:
		panic(fmt.Sprintf("expected v to be convertible to a string, got: %+v, %T", v, v))
	}
}

// ParseType returns an interface containing a string, bool, int, float64, []string, or FeedPtr
// float64 values with no decimal may be returned as integers, but that should be ok because the api won't know the difference
// when it's json encoded
func ParseType(s string) interface{} {
	slc := strings.Split(s, ",")
	if len(slc) > 1 {
		sort.Strings(slc)
		return slc
	}

	feedptr := FeedPtr{}
	err := json.Unmarshal([]byte(s), &feedptr)
	if err == nil {
		return feedptr
	}

	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		if !isIntegral(f) {
			return f
		}
		return int(f)
	}

	return s
}

func isIntegral(f float64) bool {
	return f == float64(int(f))
}

// MetaFromMap creates a *Meta and uses reflection to set fields from a map. This will panic if a value for a key is not a string.
// This it to ensure compatibility with terraform
func MetaFromMap(m map[string]interface{}) *Meta {
	meta := &Meta{}
	mv := reflect.Indirect(reflect.ValueOf(meta))
	mt := mv.Type()
	for k, v := range m {
		name := ToCamel(k)
		switch name {
		case "UsState":
			name = "USState"
		case "Loadavg":
			name = "LoadAvg"
		case "CaProvince":
			name = "CAProvince"
		case "IpPrefixes":
			name = "IPPrefixes"
		case "Asn":
			name = "ASN"
		}
		if _, ok := mt.FieldByName(name); ok {
			fv := mv.FieldByName(name)
			switch name {
			case "Up":
				if v.(string) == "1" || strings.ToLower(v.(string)) == "true" {
					fv.Set(reflect.ValueOf(true))
				} else if v.(string) == "0" || strings.ToLower(v.(string)) == "false" {
					fv.Set(reflect.ValueOf(false))
				} else {
					fv.Set(reflect.ValueOf(ParseType(v.(string))))
				}
			case "ASN":
				// If there is only one ASN, it should still be treated as a string.-
				// otherwise this gets parsed into a float64 and breaks stuff.
				i := strings.Index(v.(string), ",")
				if i == -1 {
					fv.Set(reflect.ValueOf(v.(string)))
				} else {
					fv.Set(reflect.ValueOf(ParseType(v.(string))))
				}
			case "Pulsar":
				var pulsars []map[string]interface{}
				if err := json.Unmarshal([]byte(v.(string)), &pulsars); err == nil {
					fv.Set(reflect.ValueOf(pulsars))
				}
			case "Subdivisions":
				switch v.(type) {
				case string:
					var subMap map[string]interface{}
					json.Unmarshal([]byte(v.(string)), &subMap)
					fv.Set(reflect.ValueOf(subMap))
				case map[string]interface{}:
					fv.Set(reflect.ValueOf(v.(map[string]interface{})))
				}
			case "Note":
				// If it's a Note, just pass the string without any type of parse.
				fv.Set(reflect.ValueOf(v.(string)))
			case "AdditionalMetadata":
				var additional []map[string]interface{}
				if err := json.Unmarshal([]byte(v.(string)), &additional); err == nil {
					fv.Set(reflect.ValueOf(additional))
				}
			default:
				fv.Set(reflect.ValueOf(ParseType(v.(string))))
			}
		}
	}
	return meta
}

// metaValidation is a validation struct for a metadata field.
// It contains the kinds of types that the field can be, and a list of check functions that will run on the field
type metaValidation struct {
	kinds      []reflect.Kind
	checkFuncs []func(v reflect.Value) error
}

// validateLatLong makes sure that the given lat/long is within the range 180.0 to -180.0
func validateLatLong(v reflect.Value) error {
	if v.Kind() == reflect.Float64 {
		f := v.Interface().(float64)
		if f < -180.0 || f > 180.0 {
			return fmt.Errorf("latitude/longitude values must be between -180.0 and 180.0, got %f", f)
		}
	}
	return nil
}

// validateCidr makes sure that the given string is a valid cidr
func validateCidr(v reflect.Value) error {
	if v.Kind() == reflect.String {
		s := v.Interface().(string)
		_, _, err := net.ParseCIDR(s)
		if err != nil {
			return err
		}
	}
	if v.Kind() == reflect.Slice {
		if slc, ok := v.Interface().([]string); ok {
			for _, s := range slc {
				_, _, err := net.ParseCIDR(s)
				if err != nil {
					return fmt.Errorf("%s is not a valid CIDR block", s)
				}
			}
			return nil
		}
		slc := v.Interface().([]interface{})
		for _, s := range slc {
			_, _, err := net.ParseCIDR(s.(string))
			if err != nil {
				return fmt.Errorf("%s is not a valid CIDR block", s.(string))
			}
		}
	}
	return nil
}

// validatePositiveNumber makes sure that the given number (float or int) is positive
func validatePositiveNumber(fieldName string, v reflect.Value) error {
	i := 0
	if v.Kind() == reflect.Int {
		i = v.Interface().(int)

	}

	if v.Kind() == reflect.Float64 {
		i = int(v.Interface().(float64))
	}

	if i < 0 {
		return fmt.Errorf("%s must be a positive number, was %+v", fieldName, v.Interface())
	}

	return nil
}

// geoMap is a map of all of the georegions
var geoMap = map[string]struct{}{
	"US-EAST": {}, "US-CENTRAL": {}, "US-WEST": {},
	"EUROPE": {}, "ASIAPAC": {}, "SOUTH-AMERICA": {}, "AFRICA": {},
}

// geoKeyString returns a string representation of all of the georegions
func geoKeyString() string {
	length := 0
	slc := make([]string, 0)
	for k := range geoMap {
		slc = append(slc, k)
		length += len(k) + 1
	}
	sort.Strings(slc)

	b := bytes.NewBuffer(make([]byte, 0, length-1))

	for _, k := range slc {
		b.WriteString(k + ",")
	}

	return strings.TrimRight(b.String(), ",")
}

// validateGeoregion makes sure that the given georegion is correct
func validateGeoregion(v reflect.Value) error {
	if v.Kind() == reflect.String {
		s := v.String()
		if _, ok := geoMap[s]; !ok {
			return fmt.Errorf("georegion must be one or more of %s, found %s", geoKeyString(), s)
		}
	}

	if v.Kind() == reflect.Slice {
		if slc, ok := v.Interface().([]string); ok {
			for _, s := range slc {
				if _, ok := geoMap[s]; !ok {
					return fmt.Errorf("georegion must be one or more of %s, found %s", geoKeyString(), s)
				}
			}
			return nil
		}
		slc := v.Interface().([]interface{})
		for _, s := range slc {
			if _, ok := geoMap[s.(string)]; !ok {
				return fmt.Errorf("georegion must be one or more of %s, found %s", geoKeyString(), s)
			}
		}
	}
	return nil
}

// validateCountryStateProvince makes sure that the given field only has two characters
func validateCountryStateProvince(v reflect.Value) error {
	if v.Kind() == reflect.String {
		s := v.String()
		if len(s) != 2 {
			return fmt.Errorf("country/state/province codes must be 2 digits as specified in ISO3166/ISO3166-2, got: %s", s)
		}
	}

	if v.Kind() == reflect.Slice {
		if slc, ok := v.Interface().([]string); ok {
			for _, s := range slc {
				if len(s) != 2 {
					return fmt.Errorf("country/state/province codes must be 2 digits as specified in ISO3166/ISO3166-2, got: %s", s)
				}
			}
			return nil
		}
		slc := v.Interface().([]interface{})
		for _, s := range slc {
			if len(s.(string)) != 2 {
				return fmt.Errorf("country/state/province codes must be 2 digits as specified in ISO3166/ISO3166-2, got: %s", s)
			}
		}
	}
	return nil
}

// validateNoteLength validates that a note's length is less than 256 characters
func validateNoteLength(v reflect.Value) error {
	if v.Kind() == reflect.String {
		s := v.String()
		if len(s) > 256 {
			return fmt.Errorf("note length must be less than 256 characters, was %d", len(s))
		}
	}
	return nil
}

func validatePulsar(v reflect.Value) error {
	var pulsars []*PulsarMeta

	switch v.Kind() {
	case reflect.Slice:
		// Slice from API
		bs, err := json.Marshal(v.Interface())
		if err != nil {
			return fmt.Errorf("pulsar: unexpected value: `%v`", v.Interface())
		}
		if err := json.Unmarshal(bs, &pulsars); err != nil {
			return fmt.Errorf("pulsar: invalid value: `%v`", v.Interface())
		}
	case reflect.String:
		// String from terraform
		if err := json.Unmarshal([]byte(v.String()), &pulsars); err != nil {
			return fmt.Errorf("pulsar: invalid value: `%v`", v.String())
		}
	}

	for _, p := range pulsars {
		if p.JobID == "" {
			return fmt.Errorf("pulsar Job ID is required")
		}
	}
	return nil
}

func validateAdditionalMetadata(v reflect.Value) error {
	// API expects additional_metadata to be array of length 1
	if v.Len() > 1 {
		return fmt.Errorf("unexpected length of `%d`, expected 1", v.Len())
	}

	return nil
}

// checkFuncs is shorthand for returning a slice of functions that take a reflect.Value and return an error
func checkFuncs(f ...func(v reflect.Value) error) []func(v reflect.Value) error {
	return f
}

// kinds is shorthand for returning a slice of reflect.Kind
func kinds(k ...reflect.Kind) []reflect.Kind {
	return k
}

// validationMap is a map of meta fields to validation types and functions
var validationMap = map[string]metaValidation{
	"Up": {kinds(reflect.Bool), nil},
	"Connections": {kinds(reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("Connections", v)
		})},
	"Requests": {kinds(reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("Requests", v)
		})},
	"LoadAvg": {kinds(reflect.Float64, reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("LoadAvg", v)
		})},
	"Pulsar":     {kinds(reflect.String, reflect.Slice), checkFuncs(validatePulsar)},
	"Latitude":   {kinds(reflect.Float64, reflect.Int), checkFuncs(validateLatLong)},
	"Longitude":  {kinds(reflect.Float64, reflect.Int), checkFuncs(validateLatLong)},
	"Georegion":  {kinds(reflect.String, reflect.Slice), checkFuncs(validateGeoregion)},
	"Country":    {kinds(reflect.String, reflect.Slice), checkFuncs(validateCountryStateProvince)},
	"USState":    {kinds(reflect.String, reflect.Slice), checkFuncs(validateCountryStateProvince)},
	"CAProvince": {kinds(reflect.String, reflect.Slice), checkFuncs(validateCountryStateProvince)},
	"Note":       {kinds(reflect.String), checkFuncs(validateNoteLength)},
	"IPPrefixes": {kinds(reflect.String, reflect.Slice), checkFuncs(validateCidr)},
	"ASN":        {kinds(reflect.String, reflect.Slice), nil},
	"Priority": {kinds(reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("Priority", v)
		})},
	"Weight": {kinds(reflect.Float64, reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("Weight", v)
		})},
	"Cost": {kinds(reflect.Float64, reflect.Int), checkFuncs(
		func(v reflect.Value) error {
			return validatePositiveNumber("Cost", v)
		})},
	"LowWatermark":       {kinds(reflect.Int), nil},
	"HighWatermark":      {kinds(reflect.Int), nil},
	"Subdivisions":       {kinds(reflect.String, reflect.Map), nil},
	"AdditionalMetadata": {kinds(reflect.String, reflect.Slice), checkFuncs(validateAdditionalMetadata)},
}

// validate takes a field name, a reflect value, and metaValidation and validates the given field
func validate(name string, v reflect.Value, m metaValidation) (errs []error) {

	check := true
	// if this is a FeedPtr or a *FeedPtr then we're ok, skip checking the rest of the types
	if v.Kind() == reflect.Struct || v.Kind() == reflect.Invalid {
		check = false
	}

	if check {
		match := false
		for _, k := range m.kinds {
			if k == v.Kind() {
				match = true
			}
		}

		if !match {
			errs = append(errs, fmt.Errorf("found type mismatch for meta field '%s'. expected %+v, got: %+v", name, m.kinds, v.Kind()))
		}

		for _, f := range m.checkFuncs {
			err := f(v)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if v.Kind() == reflect.Struct {
		if _, ok := v.Interface().(FeedPtr); !ok {
			errs = append(errs, fmt.Errorf("if a meta field is a struct, it must be a FeedPtr, got: %s", v.Type()))
		}
	}

	return
}

// Validate validates metadata fields and returns a list of errors if any are found
func (meta *Meta) Validate() (errs []error) {
	mv := reflect.Indirect(reflect.ValueOf(meta))
	mt := mv.Type()
	for i := 0; i < mt.NumField(); i++ {
		fv := mt.Field(i)
		err := validate(fv.Name, mv.Field(i).Elem(), validationMap[fv.Name])
		if err != nil {
			errs = append(errs, err...)
		}
	}

	return errs
}
