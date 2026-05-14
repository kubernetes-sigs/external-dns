package attributevalue

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// An Unmarshaler is an interface to provide custom unmarshaling of
// AttributeValues. Use this to provide custom logic determining
// how AttributeValues should be unmarshaled.
//
//			type ExampleUnmarshaler struct {
//				Value int
//			}
//
//			func (u *ExampleUnmarshaler) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
//	         avN, ok := av.(*types.AttributeValueMemberN)
//				if !ok {
//					return nil
//				}
//
//				n, err := strconv.ParseInt(avN.Value, 10, 0)
//				if err != nil {
//					return err
//				}
//
//				u.Value = int(n)
//				return nil
//			}
type Unmarshaler interface {
	UnmarshalDynamoDBAttributeValue(types.AttributeValue) error
}

// Unmarshal will unmarshal AttributeValues to Go value types.
// Both generic interface{} and concrete types are valid unmarshal
// destination types.
//
// Unmarshal will allocate maps, slices, and pointers as needed to
// unmarshal the AttributeValue into the provided type value.
//
// When unmarshaling AttributeValues into structs Unmarshal matches
// the field names of the struct to the AttributeValue Map keys.
// Initially it will look for exact field name matching, but will
// fall back to case insensitive if not exact match is found.
//
// With the exception of omitempty, omitemptyelem, binaryset, numberset
// and stringset all struct tags used by Marshal are also used by
// Unmarshal.
//
// When decoding AttributeValues to interfaces Unmarshal will use the
// following types.
//
//	[]byte,                 AV Binary (B)
//	[][]byte,               AV Binary Set (BS)
//	bool,                   AV Boolean (BOOL)
//	[]interface{},          AV List (L)
//	map[string]interface{}, AV Map (M)
//	float64,                AV Number (N)
//	Number,                 AV Number (N) with UseNumber set
//	[]float64,              AV Number Set (NS)
//	[]Number,               AV Number Set (NS) with UseNumber set
//	string,                 AV String (S)
//	[]string,               AV String Set (SS)
//
// If the Decoder option, UseNumber is set numbers will be unmarshaled
// as Number values instead of float64. Use this to maintain the original
// string formating of the number as it was represented in the AttributeValue.
// In addition provides additional opportunities to parse the number
// string based on individual use cases.
//
// When unmarshaling any error that occurs will halt the unmarshal
// and return the error.
//
// The output value provided must be a non-nil pointer
func Unmarshal(av types.AttributeValue, out interface{}) error {
	return NewDecoder().Decode(av, out)
}

// UnmarshalWithOptions will unmarshal AttributeValues to Go value types.
// Both generic interface{} and concrete types are valid unmarshal
// destination types.
//
// Use the `optsFns` functional options to override the default configuration.
//
// UnmarshalWithOptions will allocate maps, slices, and pointers as needed to
// unmarshal the AttributeValue into the provided type value.
//
// When unmarshaling AttributeValues into structs Unmarshal matches
// the field names of the struct to the AttributeValue Map keys.
// Initially it will look for exact field name matching, but will
// fall back to case insensitive if not exact match is found.
//
// With the exception of omitempty, omitemptyelem, binaryset, numberset
// and stringset all struct tags used by Marshal are also used by
// UnmarshalWithOptions.
//
// When decoding AttributeValues to interfaces Unmarshal will use the
// following types.
//
//	[]byte,                 AV Binary (B)
//	[][]byte,               AV Binary Set (BS)
//	bool,                   AV Boolean (BOOL)
//	[]interface{},          AV List (L)
//	map[string]interface{}, AV Map (M)
//	float64,                AV Number (N)
//	Number,                 AV Number (N) with UseNumber set
//	[]float64,              AV Number Set (NS)
//	[]Number,               AV Number Set (NS) with UseNumber set
//	string,                 AV String (S)
//	[]string,               AV String Set (SS)
//
// If the Decoder option, UseNumber is set numbers will be unmarshaled
// as Number values instead of float64. Use this to maintain the original
// string formating of the number as it was represented in the AttributeValue.
// In addition provides additional opportunities to parse the number
// string based on individual use cases.
//
// When unmarshaling any error that occurs will halt the unmarshal
// and return the error.
//
// The output value provided must be a non-nil pointer
func UnmarshalWithOptions(av types.AttributeValue, out interface{}, optFns ...func(options *DecoderOptions)) error {
	return NewDecoder(optFns...).Decode(av, out)
}

// UnmarshalMap is an alias for Unmarshal which unmarshals from
// a map of AttributeValues.
//
// The output value provided must be a non-nil pointer
func UnmarshalMap(m map[string]types.AttributeValue, out interface{}) error {
	return NewDecoder().Decode(&types.AttributeValueMemberM{Value: m}, out)
}

// UnmarshalMapWithOptions is an alias for UnmarshalWithOptions which unmarshals from
// a map of AttributeValues.
//
// Use the `optsFns` functional options to override the default configuration.
//
// The output value provided must be a non-nil pointer
func UnmarshalMapWithOptions(m map[string]types.AttributeValue, out interface{}, optFns ...func(options *DecoderOptions)) error {
	return NewDecoder(optFns...).Decode(&types.AttributeValueMemberM{Value: m}, out)
}

// UnmarshalList is an alias for Unmarshal func which unmarshals
// a slice of AttributeValues.
//
// The output value provided must be a non-nil pointer
func UnmarshalList(l []types.AttributeValue, out interface{}) error {
	return NewDecoder().Decode(&types.AttributeValueMemberL{Value: l}, out)
}

// UnmarshalListWithOptions is an alias for UnmarshalWithOptions func which unmarshals
// a slice of AttributeValues.
//
// Use the `optsFns` functional options to override the default configuration.
//
// The output value provided must be a non-nil pointer
func UnmarshalListWithOptions(l []types.AttributeValue, out interface{}, optFns ...func(options *DecoderOptions)) error {
	return NewDecoder(optFns...).Decode(&types.AttributeValueMemberL{Value: l}, out)
}

// UnmarshalListOfMaps is an alias for Unmarshal func which unmarshals a
// slice of maps of attribute values.
//
// This is useful for when you need to unmarshal the Items from a Query API
// call.
//
// The output value provided must be a non-nil pointer
func UnmarshalListOfMaps(l []map[string]types.AttributeValue, out interface{}) error {
	items := make([]types.AttributeValue, len(l))
	for i, m := range l {
		items[i] = &types.AttributeValueMemberM{Value: m}
	}

	return UnmarshalList(items, out)
}

// UnmarshalListOfMapsWithOptions is an alias for UnmarshalWithOptions func which unmarshals a
// slice of maps of attribute values.
//
// Use the `optsFns` functional options to override the default configuration.
//
// This is useful for when you need to unmarshal the Items from a Query API
// call.
//
// The output value provided must be a non-nil pointer
func UnmarshalListOfMapsWithOptions(l []map[string]types.AttributeValue, out interface{}, optFns ...func(options *DecoderOptions)) error {
	items := make([]types.AttributeValue, len(l))
	for i, m := range l {
		items[i] = &types.AttributeValueMemberM{Value: m}
	}

	return UnmarshalListWithOptions(items, out, optFns...)
}

// DecodeTimeAttributes is the set of time decoding functions for different AttributeValues.
type DecodeTimeAttributes struct {
	// Will decode S attribute values and SS attribute value elements into time.Time
	//
	// Default string parsing format is time.RFC3339
	S func(string) (time.Time, error)
	// Will decode N attribute values and NS attribute value elements into time.Time
	//
	// Default number parsing format is seconds since January 1, 1970 UTC
	N func(string) (time.Time, error)
}

// DecoderOptions is a collection of options to configure how the decoder
// unmarshals the value.
type DecoderOptions struct {
	// Support other custom struct tag keys, such as `yaml`, `json`, or `toml`.
	// Note that values provided with a custom TagKey must also be supported
	// by the (un)marshalers in this package.
	//
	// Tag key `dynamodbav` will always be read, but if custom tag key
	// conflicts with `dynamodbav` the custom tag key value will be used.
	TagKey string

	// Instructs the decoder to decode AttributeValue Numbers as
	// Number type instead of float64 when the destination type
	// is interface{}. Similar to encoding/json.Number
	UseNumber bool

	// Contains the time decoding functions for different AttributeValues
	//
	// Default string parsing format is time.RFC3339
	// Default number parsing format is seconds since January 1, 1970 UTC
	DecodeTime DecodeTimeAttributes

	// When enabled, the decoder will use implementations of
	// encoding.TextUnmarshaler and encoding.BinaryUnmarshaler when present on
	// unmarshaling targets.
	//
	// If a target implements [Unmarshaler], encoding unmarshaler
	// implementations are ignored.
	//
	// If the attributevalue is a string, its underlying value will be used to
	// call UnmarshalText on the target. If the attributevalue is a binary, its
	// value will be used to call UnmarshalBinary.
	UseEncodingUnmarshalers bool

	// When enabled, the decoder will call Unmarshaler.UnmarshalDynamoDBAttributeValue
	// for each individual set item instead of the whole set at once.
	// See issue #2895.
	FixUnmarshalIndividualSetValues bool
}

// A Decoder provides unmarshaling AttributeValues to Go value types.
type Decoder struct {
	options DecoderOptions
}

// NewDecoder creates a new Decoder with default configuration. Use
// the `opts` functional options to override the default configuration.
func NewDecoder(optFns ...func(*DecoderOptions)) *Decoder {
	options := DecoderOptions{
		TagKey: defaultTagKey,
		DecodeTime: DecodeTimeAttributes{
			S: defaultDecodeTimeS,
			N: defaultDecodeTimeN,
		},
	}
	for _, fn := range optFns {
		fn(&options)
	}

	if options.DecodeTime.S == nil {
		options.DecodeTime.S = defaultDecodeTimeS
	}

	if options.DecodeTime.N == nil {
		options.DecodeTime.N = defaultDecodeTimeN
	}

	return &Decoder{
		options: options,
	}
}

// Decode will unmarshal an AttributeValue into a Go value type. An error
// will be return if the decoder is unable to unmarshal the AttributeValue
// to the provide Go value type.
//
// The output value provided must be a non-nil pointer
func (d *Decoder) Decode(av types.AttributeValue, out interface{}, opts ...func(*Decoder)) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() || !v.IsValid() {
		return &InvalidUnmarshalError{Type: reflect.TypeOf(out)}
	}

	return d.decode(av, v, tag{})
}

var stringInterfaceMapType = reflect.TypeOf(map[string]interface{}(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))
var byteSliceSliceType = reflect.TypeOf([][]byte(nil))
var timeType = reflect.TypeOf(time.Time{})

func (d *Decoder) decode(av types.AttributeValue, v reflect.Value, fieldTag tag) error {
	var u Unmarshaler
	_, isNull := av.(*types.AttributeValueMemberNULL)
	if av == nil || isNull {
		u, v = indirect[Unmarshaler](v, indirectOptions{decodeNull: true})
		if u != nil {
			return u.UnmarshalDynamoDBAttributeValue(av)
		}
		return d.decodeNull(v)
	}

	v0 := v
	u, v = indirect[Unmarshaler](v, indirectOptions{})
	if u != nil {
		return u.UnmarshalDynamoDBAttributeValue(av)
	}
	if d.options.UseEncodingUnmarshalers {
		if s, ok := av.(*types.AttributeValueMemberS); ok {
			if u, _ := indirect[encoding.TextUnmarshaler](v0, indirectOptions{}); u != nil {
				return u.UnmarshalText([]byte(s.Value))
			}
		}
		if b, ok := av.(*types.AttributeValueMemberB); ok {
			if u, _ := indirect[encoding.BinaryUnmarshaler](v0, indirectOptions{}); u != nil {
				return u.UnmarshalBinary(b.Value)
			}
		}
	}

	switch tv := av.(type) {
	case *types.AttributeValueMemberB:
		return d.decodeBinary(tv.Value, v)

	case *types.AttributeValueMemberBOOL:
		return d.decodeBool(tv.Value, v)

	case *types.AttributeValueMemberBS:
		return d.decodeBinarySet(tv.Value, v)

	case *types.AttributeValueMemberL:
		return d.decodeList(tv.Value, v)

	case *types.AttributeValueMemberM:
		return d.decodeMap(tv.Value, v)

	case *types.AttributeValueMemberN:
		return d.decodeNumber(tv.Value, v, fieldTag)

	case *types.AttributeValueMemberNS:
		return d.decodeNumberSet(tv.Value, v)

	case *types.AttributeValueMemberS:
		return d.decodeString(tv.Value, v, fieldTag)

	case *types.AttributeValueMemberSS:
		return d.decodeStringSet(tv.Value, v)

	default:
		return fmt.Errorf("unsupported AttributeValue type, %T", av)
	}
}

func (d *Decoder) decodeBinary(b []byte, v reflect.Value) error {
	if v.Kind() == reflect.Interface {
		buf := make([]byte, len(b))
		copy(buf, b)
		v.Set(reflect.ValueOf(buf))
		return nil
	}

	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return &UnmarshalTypeError{Value: "binary", Type: v.Type()}
	}

	if v.Type() == byteSliceType {
		// Optimization for []byte types
		if v.IsNil() || v.Cap() < len(b) {
			v.Set(reflect.MakeSlice(byteSliceType, len(b), len(b)))
		} else if v.Len() != len(b) {
			v.SetLen(len(b))
		}
		copy(v.Interface().([]byte), b)
		return nil
	}

	switch v.Type().Elem().Kind() {
	case reflect.Uint8:
		// Fallback to reflection copy for type aliased of []byte type
		if v.Kind() != reflect.Array && (v.IsNil() || v.Cap() < len(b)) {
			v.Set(reflect.MakeSlice(v.Type(), len(b), len(b)))
		} else if v.Len() != len(b) {
			v.SetLen(len(b))
		}
		for i := 0; i < len(b); i++ {
			v.Index(i).SetUint(uint64(b[i]))
		}
	default:
		if v.Kind() == reflect.Array && v.Type().Elem().Kind() == reflect.Uint8 {
			reflect.Copy(v, reflect.ValueOf(b))
			break
		}
		return &UnmarshalTypeError{Value: "binary", Type: v.Type()}
	}

	return nil
}

func (d *Decoder) decodeBool(b bool, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Bool, reflect.Interface:
		v.Set(reflect.ValueOf(b).Convert(v.Type()))

	default:
		return &UnmarshalTypeError{Value: "bool", Type: v.Type()}
	}

	return nil
}

func (d *Decoder) decodeBinarySet(bs [][]byte, v reflect.Value) error {
	var isArray bool

	switch v.Kind() {
	case reflect.Slice:
		// Make room for the slice elements if needed
		if v.IsNil() || v.Cap() < len(bs) {
			// What about if ignoring nil/empty values?
			v.Set(reflect.MakeSlice(v.Type(), 0, len(bs)))
		}
	case reflect.Array:
		// Limited to capacity of existing array.
		isArray = true
	case reflect.Interface:
		set := make([][]byte, len(bs))
		for i, b := range bs {
			if err := d.decodeBinary(b, reflect.ValueOf(&set[i]).Elem()); err != nil {
				return err
			}
		}
		v.Set(reflect.ValueOf(set))
		return nil
	default:
		return &UnmarshalTypeError{Value: "binary set", Type: v.Type()}
	}

	for i := 0; i < v.Cap() && i < len(bs); i++ {
		if !isArray {
			v.SetLen(i + 1)
		}
		u, elem := indirect[Unmarshaler](v.Index(i), indirectOptions{})
		if u != nil {
			if d.options.FixUnmarshalIndividualSetValues {
				err := u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberB{Value: bs[i]})
				if err != nil {
					return err
				}
				continue
			} else {
				return u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberBS{Value: bs})
			}
		}
		if err := d.decodeBinary(bs[i], elem); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) decodeNumber(n string, v reflect.Value, fieldTag tag) error {
	switch v.Kind() {
	case reflect.Interface:
		i, err := d.decodeNumberToInterface(n)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(i))
		return nil
	case reflect.String:
		if isNumberValueType(v) {
			v.SetString(n)
			return nil
		}
		v.SetString(n)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			return err
		}
		if v.OverflowInt(i) {
			return &UnmarshalTypeError{
				Value: fmt.Sprintf("number overflow, %s", n),
				Type:  v.Type(),
			}
		}
		v.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(n, 10, 64)
		if err != nil {
			return err
		}
		if v.OverflowUint(i) {
			return &UnmarshalTypeError{
				Value: fmt.Sprintf("number overflow, %s", n),
				Type:  v.Type(),
			}
		}
		v.SetUint(i)
	case reflect.Float32, reflect.Float64:
		i, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return err
		}
		if v.OverflowFloat(i) {
			return &UnmarshalTypeError{
				Value: fmt.Sprintf("number overflow, %s", n),
				Type:  v.Type(),
			}
		}
		v.SetFloat(i)
	default:
		if v.Type().ConvertibleTo(timeType) && fieldTag.AsUnixTime {
			t, err := decodeUnixTime(n)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(t).Convert(v.Type()))
			return nil
		}
		if v.Type().ConvertibleTo(timeType) {
			t, err := d.options.DecodeTime.N(n)
			if err != nil {
				return err
			}
			v.Set(reflect.ValueOf(t).Convert(v.Type()))
			return nil
		}
		return &UnmarshalTypeError{Value: "number", Type: v.Type()}
	}

	return nil
}

func (d *Decoder) decodeNumberToInterface(n string) (interface{}, error) {
	if d.options.UseNumber {
		return Number(n), nil
	}

	// Default to float64 for all numbers
	return strconv.ParseFloat(n, 64)
}

func (d *Decoder) decodeNumberSet(ns []string, v reflect.Value) error {
	var isArray bool

	switch v.Kind() {
	case reflect.Slice:
		// Make room for the slice elements if needed
		if v.IsNil() || v.Cap() < len(ns) {
			// What about if ignoring nil/empty values?
			v.Set(reflect.MakeSlice(v.Type(), 0, len(ns)))
		}
	case reflect.Array:
		// Limited to capacity of existing array.
		isArray = true
	case reflect.Interface:
		if d.options.UseNumber {
			set := make([]Number, len(ns))
			for i, n := range ns {
				if err := d.decodeNumber(n, reflect.ValueOf(&set[i]).Elem(), tag{}); err != nil {
					return err
				}
			}
			v.Set(reflect.ValueOf(set))
		} else {
			set := make([]float64, len(ns))
			for i, n := range ns {
				if err := d.decodeNumber(n, reflect.ValueOf(&set[i]).Elem(), tag{}); err != nil {
					return err
				}
			}
			v.Set(reflect.ValueOf(set))
		}
		return nil
	default:
		return &UnmarshalTypeError{Value: "number set", Type: v.Type()}
	}

	for i := 0; i < v.Cap() && i < len(ns); i++ {
		if !isArray {
			v.SetLen(i + 1)
		}
		u, elem := indirect[Unmarshaler](v.Index(i), indirectOptions{})
		if u != nil {
			if d.options.FixUnmarshalIndividualSetValues {
				err := u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberN{Value: ns[i]})
				if err != nil {
					return err
				}
				continue
			} else {
				return u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberNS{Value: ns})
			}
		}
		if err := d.decodeNumber(ns[i], elem, tag{}); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) decodeList(avList []types.AttributeValue, v reflect.Value) error {
	var isArray bool

	switch v.Kind() {
	case reflect.Slice:
		// Make room for the slice elements if needed
		if v.IsNil() || v.Cap() < len(avList) {
			// What about if ignoring nil/empty values?
			v.Set(reflect.MakeSlice(v.Type(), 0, len(avList)))
		}
	case reflect.Array:
		// Limited to capacity of existing array.
		isArray = true
	case reflect.Interface:
		s := make([]interface{}, len(avList))
		for i, av := range avList {
			if err := d.decode(av, reflect.ValueOf(&s[i]).Elem(), tag{}); err != nil {
				return err
			}
		}
		v.Set(reflect.ValueOf(s))
		return nil
	default:
		return &UnmarshalTypeError{Value: "list", Type: v.Type()}
	}

	// If v is not a slice, array
	for i := 0; i < v.Cap() && i < len(avList); i++ {
		if !isArray {
			v.SetLen(i + 1)
		}
		if err := d.decode(avList[i], v.Index(i), tag{}); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decoder) decodeMap(avMap map[string]types.AttributeValue, v reflect.Value) (err error) {
	var decodeMapKey func(v string, key reflect.Value, fieldTag tag) error

	switch v.Kind() {
	case reflect.Map:
		decodeMapKey, err = d.getMapKeyDecoder(v.Type().Key())
		if err != nil {
			return err
		}

		if v.IsNil() {
			v.Set(reflect.MakeMap(v.Type()))
		}
	case reflect.Struct:
	case reflect.Interface:
		v.Set(reflect.MakeMap(stringInterfaceMapType))
		decodeMapKey = d.decodeString
		v = v.Elem()
	default:
		return &UnmarshalTypeError{Value: "map", Type: v.Type()}
	}

	if v.Kind() == reflect.Map {
		keyType := v.Type().Key()
		valueType := v.Type().Elem()
		for k, av := range avMap {
			key := reflect.New(keyType).Elem()
			// handle pointer keys
			_, indirectKey := indirect[Unmarshaler](key, indirectOptions{skipUnmarshaler: true})
			if err := decodeMapKey(k, indirectKey, tag{}); err != nil {
				return &UnmarshalTypeError{
					Value: fmt.Sprintf("map key %q", k),
					Type:  keyType,
					Err:   err,
				}
			}

			elem := reflect.New(valueType).Elem()
			if err := d.decode(av, elem, tag{}); err != nil {
				return err
			}

			v.SetMapIndex(key, elem)
		}
	} else if v.Kind() == reflect.Struct {
		fields := unionStructFields(v.Type(), structFieldOptions{
			TagKey: d.options.TagKey,
		})
		for k, av := range avMap {
			if f, ok := fields.FieldByName(k); ok {
				fv := decoderFieldByIndex(v, f.Index)
				if err := d.decode(av, fv, f.tag); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

var numberType = reflect.TypeOf(Number(""))
var textUnmarshalerType = reflect.TypeOf((*encoding.TextUnmarshaler)(nil)).Elem()

func (d *Decoder) getMapKeyDecoder(keyType reflect.Type) (func(string, reflect.Value, tag) error, error) {
	// Test the key type to determine if it implements the TextUnmarshaler interface.
	if reflect.PtrTo(keyType).Implements(textUnmarshalerType) || keyType.Implements(textUnmarshalerType) {
		return func(v string, k reflect.Value, _ tag) error {
			if !k.CanAddr() {
				return fmt.Errorf("cannot take address of map key, %v", k.Type())
			}
			return k.Addr().Interface().(encoding.TextUnmarshaler).UnmarshalText([]byte(v))
		}, nil
	}

	var decodeMapKey func(v string, key reflect.Value, fieldTag tag) error

	switch keyType.Kind() {
	case reflect.Bool:
		decodeMapKey = func(v string, key reflect.Value, fieldTag tag) error {
			b, err := strconv.ParseBool(v)
			if err != nil {
				return err
			}
			return d.decodeBool(b, key)
		}
	case reflect.String:
		// Number type handled as a string
		decodeMapKey = d.decodeString

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		decodeMapKey = d.decodeNumber

	default:
		return nil, &UnmarshalTypeError{
			Value: "map key must be string, number, bool, or TextUnmarshaler",
			Type:  keyType,
		}
	}

	return decodeMapKey, nil
}

func (d *Decoder) decodeNull(v reflect.Value) error {
	if v.IsValid() && v.CanSet() {
		v.Set(reflect.Zero(v.Type()))
	}

	return nil
}

func (d *Decoder) decodeString(s string, v reflect.Value, fieldTag tag) error {
	if fieldTag.AsString {
		return d.decodeNumber(s, v, fieldTag)
	}

	// To maintain backwards compatibility with ConvertFrom family of methods which
	// converted strings to time.Time structs
	if v.Type().ConvertibleTo(timeType) {
		t, err := d.options.DecodeTime.S(s)
		if err != nil {
			return err
		}
		v.Set(reflect.ValueOf(t).Convert(v.Type()))
		return nil
	}

	switch v.Kind() {
	case reflect.String:
		v.SetString(s)
	case reflect.Interface:
		// Ensure type aliasing is handled properly
		v.Set(reflect.ValueOf(s).Convert(v.Type()))
	default:
		return &UnmarshalTypeError{Value: "string", Type: v.Type()}
	}

	return nil
}

func (d *Decoder) decodeStringSet(ss []string, v reflect.Value) error {
	var isArray bool

	switch v.Kind() {
	case reflect.Slice:
		// Make room for the slice elements if needed
		if v.IsNil() || v.Cap() < len(ss) {
			v.Set(reflect.MakeSlice(v.Type(), 0, len(ss)))
		}
	case reflect.Array:
		// Limited to capacity of existing array.
		isArray = true
	case reflect.Interface:
		set := make([]string, len(ss))
		for i, s := range ss {
			if err := d.decodeString(s, reflect.ValueOf(&set[i]).Elem(), tag{}); err != nil {
				return err
			}
		}
		v.Set(reflect.ValueOf(set))
		return nil
	default:
		return &UnmarshalTypeError{Value: "string set", Type: v.Type()}
	}

	for i := 0; i < v.Cap() && i < len(ss); i++ {
		if !isArray {
			v.SetLen(i + 1)
		}
		u, elem := indirect[Unmarshaler](v.Index(i), indirectOptions{})
		if u != nil {
			if d.options.FixUnmarshalIndividualSetValues {
				err := u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberS{Value: ss[i]})
				if err != nil {
					return err
				}
				continue
			} else {
				return u.UnmarshalDynamoDBAttributeValue(&types.AttributeValueMemberSS{Value: ss})
			}
		}
		if err := d.decodeString(ss[i], elem, tag{}); err != nil {
			return err
		}
	}

	return nil
}

func decodeUnixTime(n string) (time.Time, error) {
	v, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return time.Time{}, &UnmarshalError{
			Err: err, Value: n, Type: timeType,
		}
	}

	return time.Unix(v, 0), nil
}

// decoderFieldByIndex finds the field with the provided nested index, allocating
// embedded parent structs if needed
func decoderFieldByIndex(v reflect.Value, index []int) reflect.Value {
	for i, x := range index {
		if i > 0 && v.Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.Struct {
			if v.IsNil() {
				v.Set(reflect.New(v.Type().Elem()))
			}
			v = v.Elem()
		}
		v = v.Field(x)
	}
	return v
}

type indirectOptions struct {
	decodeNull      bool
	skipUnmarshaler bool
}

// indirect will walk a value's interface or pointer value types. Returning
// the final value or the value a unmarshaler is defined on.
//
// Based on the enoding/json type reflect value type indirection in Go Stdlib
// https://golang.org/src/encoding/json/decode.go indirect func.
func indirect[U any](v reflect.Value, opts indirectOptions) (U, reflect.Value) {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}

	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && (!opts.decodeNull || e.Elem().Kind() == reflect.Ptr) {
				haveAddr = false
				v = e
				continue
			}
			if e.Kind() != reflect.Ptr && e.IsValid() {
				var u U
				return u, e
			}
		}
		if v.Kind() != reflect.Ptr {
			break
		}
		if opts.decodeNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if !opts.skipUnmarshaler && v.Type().NumMethod() > 0 && v.CanInterface() {
			if u, ok := v.Interface().(U); ok {
				return u, reflect.Value{}
			}
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}

	var u U
	return u, v
}

// A Number represents a Attributevalue number literal.
type Number string

// Float64 attempts to cast the number to a float64, returning
// the result of the case or error if the case failed.
func (n Number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n), 64)
}

// Int64 attempts to cast the number to a int64, returning
// the result of the case or error if the case failed.
func (n Number) Int64() (int64, error) {
	return strconv.ParseInt(string(n), 10, 64)
}

// Uint64 attempts to cast the number to a uint64, returning
// the result of the case or error if the case failed.
func (n Number) Uint64() (uint64, error) {
	return strconv.ParseUint(string(n), 10, 64)
}

// String returns the raw number represented as a string
func (n Number) String() string {
	return string(n)
}

// An UnmarshalTypeError is an error type representing a error
// unmarshaling the AttributeValue's element to a Go value type.
// Includes details about the AttributeValue type and Go value type.
type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
	Err   error
}

// Unwrap returns the underlying error if any.
func (e *UnmarshalTypeError) Unwrap() error { return e.Err }

// Error returns the string representation of the error.
// satisfying the error interface
func (e *UnmarshalTypeError) Error() string {
	return fmt.Sprintf("unmarshal failed, cannot unmarshal %s into Go value type %s",
		e.Value, e.Type.String())
}

// An InvalidUnmarshalError is an error type representing an invalid type
// encountered while unmarshaling a AttributeValue to a Go value type.
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// Error returns the string representation of the error.
// satisfying the error interface
func (e *InvalidUnmarshalError) Error() string {
	var msg string
	if e.Type == nil {
		msg = "cannot unmarshal to nil value"
	} else if e.Type.Kind() != reflect.Ptr {
		msg = fmt.Sprintf("cannot unmarshal to non-pointer value, got %s", e.Type.String())
	} else {
		msg = fmt.Sprintf("cannot unmarshal to nil value, %s", e.Type.String())
	}

	return fmt.Sprintf("unmarshal failed, %s", msg)
}

// An UnmarshalError wraps an error that occurred while unmarshaling a
// AttributeValue element into a Go type. This is different from
// UnmarshalTypeError in that it wraps the underlying error that occurred.
type UnmarshalError struct {
	Err   error
	Value string
	Type  reflect.Type
}

func (e *UnmarshalError) Unwrap() error {
	return e.Err
}

// Error returns the string representation of the error satisfying the error
// interface.
func (e *UnmarshalError) Error() string {
	return fmt.Sprintf("unmarshal failed, cannot unmarshal %q into %s, %v",
		e.Value, e.Type.String(), e.Err)
}

func defaultDecodeTimeS(v string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return time.Time{}, &UnmarshalError{Err: err, Value: v, Type: timeType}
	}
	return t, nil
}

func defaultDecodeTimeN(v string) (time.Time, error) {
	return decodeUnixTime(v)
}
