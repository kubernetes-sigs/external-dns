package attributevalue

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// An UnixTime provides aliasing of time.Time into a type that when marshaled
// and unmarshaled with AttributeValues it will be done so as number
// instead of string in seconds since January 1, 1970 UTC.
//
// This type is useful as an alternative to the struct tag `unixtime` when you
// want to have your time value marshaled as Unix time in seconds into a number
// attribute type instead of the default time.RFC3339Nano.
//
// Important to note that zero value time as unixtime is not 0 seconds
// from January 1, 1970 UTC, but -62135596800. Which is seconds between
// January 1, 0001 UTC, and January 1, 0001 UTC.
//
// Also, important to note: the default UnixTime implementation of the Marshaler
// interface will marshal into an attribute of type of number; therefore,
// it may not be used as a sort key if the attribute value is of type string. Further,
// the time.RFC3339Nano format removes trailing zeros from the seconds field
// and thus may not sort correctly once formatted.
type UnixTime time.Time

// MarshalDynamoDBAttributeValue implements the Marshaler interface so that
// the UnixTime can be marshaled from to a AttributeValue number
// value encoded in the number of seconds since January 1, 1970 UTC.
func (e UnixTime) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberN{
		Value: strconv.FormatInt(time.Time(e).Unix(), 10),
	}, nil
}

// UnmarshalDynamoDBAttributeValue implements the Unmarshaler interface so that
// the UnixTime can be unmarshaled from a AttributeValue number representing
// the number of seconds since January 1, 1970 UTC.
//
// If an error parsing the AttributeValue number occurs UnmarshalError will be
// returned.
func (e *UnixTime) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	tv, ok := av.(*types.AttributeValueMemberN)
	if !ok {
		return &UnmarshalTypeError{
			Value: fmt.Sprintf("%T", av),
			Type:  reflect.TypeOf((*UnixTime)(nil)),
		}
	}

	t, err := decodeUnixTime(tv.Value)
	if err != nil {
		return err
	}

	*e = UnixTime(t)
	return nil
}

// String calls the underlying time.Time.String to return a human readable
// representation.
func (e UnixTime) String() string {
	return time.Time(e).String()
}

// A Marshaler is an interface to provide custom marshaling of Go value types
// to AttributeValues. Use this to provide custom logic determining how a
// Go Value type should be marshaled.
//
//	type CustomIntType struct {
//		Value Int
//	}
//	func (m *CustomIntType) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
//		return &types.AttributeValueMemberN{
//			Value: strconv.Itoa(m.Value),
//		}, nil
//	}
type Marshaler interface {
	MarshalDynamoDBAttributeValue() (types.AttributeValue, error)
}

// Marshal will serialize the passed in Go value type into a AttributeValue
// type. This value can be used in API operations to simplify marshaling
// your Go value types into AttributeValues.
//
// Marshal will recursively traverse the passed in value marshaling its
// contents into a AttributeValue. Marshal supports basic scalars
// (int,uint,float,bool,string), maps, slices, and structs. Anonymous
// nested types are flattened based on Go anonymous type visibility.
//
// Marshaling slices to AttributeValue will default to a List for all
// types except for []byte and [][]byte. []byte will be marshaled as
// Binary data (B), and [][]byte will be marshaled as binary data set
// (BS).
//
// The `time.Time` type is marshaled as `time.RFC3339Nano` format.
//
// `dynamodbav` struct tag can be used to control how the value will be
// marshaled into a AttributeValue.
//
//	// Field is ignored
//	Field int `dynamodbav:"-"`
//
//	// Field AttributeValue map key "myName"
//	Field int `dynamodbav:"myName"`
//
//	// Field AttributeValue map key "myName", and
//	// Field is omitted if the field is a zero value for the type.
//	Field int `dynamodbav:"myName,omitempty"`
//
//	// Field AttributeValue map key "Field", and
//	// Field is omitted if the field is a zero value for the type.
//	Field int `dynamodbav:",omitempty"`
//
//	// Field's elems will be omitted if the elem's value is empty.
//	// only valid for slices, and maps.
//	Field []string `dynamodbav:",omitemptyelem"`
//
//	// Field AttributeValue map key "Field", and
//	// Field is sent as NULL if the field is a zero value for the type.
//	Field int `dynamodbav:",nullempty"`
//
//	// Field's elems will be sent as NULL if the elem's value a zero value
//	// for the type. Only valid for slices, and maps.
//	Field []string `dynamodbav:",nullemptyelem"`
//
//	// Field will be marshaled as a AttributeValue string
//	// only value for number types, (int,uint,float)
//	Field int `dynamodbav:",string"`
//
//	// Field will be marshaled as a binary set
//	Field [][]byte `dynamodbav:",binaryset"`
//
//	// Field will be marshaled as a number set
//	Field []int `dynamodbav:",numberset"`
//
//	// Field will be marshaled as a string set
//	Field []string `dynamodbav:",stringset"`
//
//	// Field will be marshaled as Unix time number in seconds.
//	// This tag is only valid with time.Time typed struct fields.
//	// Important to note that zero value time as unixtime is not 0 seconds
//	// from January 1, 1970 UTC, but -62135596800. Which is seconds between
//	// January 1, 0001 UTC, and January 1, 0001 UTC.
//	Field time.Time `dynamodbav:",unixtime"`
//
// The omitempty tag is only used during Marshaling and is ignored for
// Unmarshal. omitempty will skip any member if the Go value of the member is
// zero. The omitemptyelem tag works the same as omitempty except it applies to
// the elements of maps and slices instead of struct fields, and will not be
// included in the marshaled AttributeValue Map, List, or Set.
//
// The nullempty tag is only used during Marshaling and is ignored for
// Unmarshal. nullempty will serialize a AttributeValueMemberNULL for the
// member if the Go value of the member is zero. nullemptyelem tag works the
// same as nullempty except it applies to the elements of maps and slices
// instead of struct fields, and will not be included in the marshaled
// AttributeValue Map, List, or Set.
//
// All struct fields and with anonymous fields, are marshaled unless the
// any of the following conditions are meet.
//
//   - the field is not exported
//   - json or dynamodbav field tag is "-"
//   - json or dynamodbav field tag specifies "omitempty", and is a zero value.
//
// Pointer and interfaces values are encoded as the value pointed to or
// contained in the interface. A nil value encodes as the AttributeValue NULL
// value unless `omitempty` struct tag is provided.
//
// Channel, complex, and function values are not encoded and will be skipped
// when walking the value to be marshaled.
//
// Error that occurs when marshaling will stop the marshal, and return
// the error.
//
// Marshal cannot represent cyclic data structures and will not handle them.
// Passing cyclic structures to Marshal will result in an infinite recursion.
func Marshal(in interface{}) (types.AttributeValue, error) {
	return NewEncoder().Encode(in)
}

// MarshalWithOptions will serialize the passed in Go value type into a AttributeValue
// type, by using . This value can be used in API operations to simplify marshaling
// your Go value types into AttributeValues.
//
// Use the `optsFns` functional options to override the default configuration.
//
// MarshalWithOptions will recursively traverse the passed in value marshaling its
// contents into a AttributeValue. Marshal supports basic scalars
// (int,uint,float,bool,string), maps, slices, and structs. Anonymous
// nested types are flattened based on Go anonymous type visibility.
//
// Marshaling slices to AttributeValue will default to a List for all
// types except for []byte and [][]byte. []byte will be marshaled as
// Binary data (B), and [][]byte will be marshaled as binary data set
// (BS).
//
// The `time.Time` type is marshaled as `time.RFC3339Nano` format.
//
// `dynamodbav` struct tag can be used to control how the value will be
// marshaled into a AttributeValue.
//
//	// Field is ignored
//	Field int `dynamodbav:"-"`
//
//	// Field AttributeValue map key "myName"
//	Field int `dynamodbav:"myName"`
//
//	// Field AttributeValue map key "myName", and
//	// Field is omitted if the field is a zero value for the type.
//	Field int `dynamodbav:"myName,omitempty"`
//
//	// Field AttributeValue map key "Field", and
//	// Field is omitted if the field is a zero value for the type.
//	Field int `dynamodbav:",omitempty"`
//
//	// Field's elems will be omitted if the elem's value is empty.
//	// only valid for slices, and maps.
//	Field []string `dynamodbav:",omitemptyelem"`
//
//	// Field AttributeValue map key "Field", and
//	// Field is sent as NULL if the field is a zero value for the type.
//	Field int `dynamodbav:",nullempty"`
//
//	// Field's elems will be sent as NULL if the elem's value a zero value
//	// for the type. Only valid for slices, and maps.
//	Field []string `dynamodbav:",nullemptyelem"`
//
//	// Field will be marshaled as a AttributeValue string
//	// only value for number types, (int,uint,float)
//	Field int `dynamodbav:",string"`
//
//	// Field will be marshaled as a binary set
//	Field [][]byte `dynamodbav:",binaryset"`
//
//	// Field will be marshaled as a number set
//	Field []int `dynamodbav:",numberset"`
//
//	// Field will be marshaled as a string set
//	Field []string `dynamodbav:",stringset"`
//
//	// Field will be marshaled as Unix time number in seconds.
//	// This tag is only valid with time.Time typed struct fields.
//	// Important to note that zero value time as unixtime is not 0 seconds
//	// from January 1, 1970 UTC, but -62135596800. Which is seconds between
//	// January 1, 0001 UTC, and January 1, 0001 UTC.
//	Field time.Time `dynamodbav:",unixtime"`
//
// The omitempty tag is only used during Marshaling and is ignored for
// Unmarshal. omitempty will skip any member if the Go value of the member is
// zero. The omitemptyelem tag works the same as omitempty except it applies to
// the elements of maps and slices instead of struct fields, and will not be
// included in the marshaled AttributeValue Map, List, or Set.
//
// The nullempty tag is only used during Marshaling and is ignored for
// Unmarshal. nullempty will serialize a AttributeValueMemberNULL for the
// member if the Go value of the member is zero. nullemptyelem tag works the
// same as nullempty except it applies to the elements of maps and slices
// instead of struct fields, and will not be included in the marshaled
// AttributeValue Map, List, or Set.
//
// All struct fields and with anonymous fields, are marshaled unless the
// any of the following conditions are meet.
//
//   - the field is not exported
//   - json or dynamodbav field tag is "-"
//   - json or dynamodbav field tag specifies "omitempty", and is a zero value.
//
// Pointer and interfaces values are encoded as the value pointed to or
// contained in the interface. A nil value encodes as the AttributeValue NULL
// value unless `omitempty` struct tag is provided.
//
// Channel, complex, and function values are not encoded and will be skipped
// when walking the value to be marshaled.
//
// Error that occurs when marshaling will stop the marshal, and return
// the error.
//
// MarshalWithOptions cannot represent cyclic data structures and will not handle them.
// Passing cyclic structures to Marshal will result in an infinite recursion.
func MarshalWithOptions(in interface{}, optFns ...func(*EncoderOptions)) (types.AttributeValue, error) {
	return NewEncoder(optFns...).Encode(in)
}

// MarshalMap is an alias for Marshal func which marshals Go value type to a
// map of AttributeValues. If the in parameter does not serialize to a map, an
// empty AttributeValue map will be returned.
//
// Use the `optsFns` functional options to override the default configuration.
//
// This is useful for APIs such as PutItem.
func MarshalMap(in interface{}) (map[string]types.AttributeValue, error) {
	av, err := NewEncoder().Encode(in)

	asMap, ok := av.(*types.AttributeValueMemberM)
	if err != nil || av == nil || !ok {
		return map[string]types.AttributeValue{}, err
	}

	return asMap.Value, nil
}

// MarshalMapWithOptions is an alias for MarshalWithOptions func which marshals Go value type to a
// map of AttributeValues. If the in parameter does not serialize to a map, an
// empty AttributeValue map will be returned.
//
// Use the `optsFns` functional options to override the default configuration.
//
// This is useful for APIs such as PutItem.
func MarshalMapWithOptions(in interface{}, optFns ...func(*EncoderOptions)) (map[string]types.AttributeValue, error) {
	av, err := NewEncoder(optFns...).Encode(in)

	asMap, ok := av.(*types.AttributeValueMemberM)
	if err != nil || av == nil || !ok {
		return map[string]types.AttributeValue{}, err
	}

	return asMap.Value, nil
}

// MarshalList is an alias for Marshal func which marshals Go value
// type to a slice of AttributeValues. If the in parameter does not serialize
// to a slice, an empty AttributeValue slice will be returned.
func MarshalList(in interface{}) ([]types.AttributeValue, error) {
	av, err := NewEncoder().Encode(in)

	asList, ok := av.(*types.AttributeValueMemberL)
	if err != nil || av == nil || !ok {
		return []types.AttributeValue{}, err
	}

	return asList.Value, nil
}

// MarshalListWithOptions is an alias for MarshalWithOptions func which marshals Go value
// type to a slice of AttributeValues. If the in parameter does not serialize
// to a slice, an empty AttributeValue slice will be returned.
//
// Use the `optsFns` functional options to override the default configuration.
func MarshalListWithOptions(in interface{}, optFns ...func(*EncoderOptions)) ([]types.AttributeValue, error) {
	av, err := NewEncoder(optFns...).Encode(in)

	asList, ok := av.(*types.AttributeValueMemberL)
	if err != nil || av == nil || !ok {
		return []types.AttributeValue{}, err
	}

	return asList.Value, nil
}

// EncoderOptions is a collection of options used by the marshaler.
type EncoderOptions struct {
	// Support other custom struct tag keys, such as `yaml`, `json`, or `toml`.
	// Note that values provided with a custom TagKey must also be supported
	// by the (un)marshalers in this package.
	//
	// Tag key `dynamodbav` will always be read, but if custom tag key
	// conflicts with `dynamodbav` the custom tag key value will be used.
	TagKey string

	// Will encode any slice being encoded as a set (SS, NS, and BS) as a NULL
	// AttributeValue if the slice is not nil, but is empty but contains no
	// elements.
	//
	// If a type implements the Marshal interface, and returns empty set
	// slices, this option will not modify the returned value.
	//
	// Defaults to enabled, because AttributeValue sets cannot currently be
	// empty lists.
	NullEmptySets bool

	// Will encode time.Time fields
	//
	// Default encoding is time.RFC3339Nano in a DynamoDB String (S) data type.
	EncodeTime func(time.Time) (types.AttributeValue, error)

	// When enabled, the encoder will use implementations of
	// encoding.TextMarshaler and encoding.BinaryMarshaler when present on
	// marshaled values.
	//
	// Implementations are checked in the following order:
	//   - [Marshaler]
	//   - encoding.TextMarshaler
	//   - encoding.BinaryMarshaler
	//
	// The results of a MarshalText call will convert to string (S), results
	// from a MarshalBinary call will convert to binary (B).
	UseEncodingMarshalers bool

	// When enabled, the encoder will omit null (NULL) attribute values
	// returned from custom marshalers tagged with `omitempty`.
	//
	// NULL attribute values returned from the standard marshaling routine will
	// always respect omitempty regardless of this setting.
	OmitNullAttributeValues bool

	// When enabled, the encoder will omit empty time attribute values
	OmitEmptyTime bool
}

// An Encoder provides marshaling Go value types to AttributeValues.
type Encoder struct {
	options EncoderOptions
}

// NewEncoder creates a new Encoder with default configuration. Use
// the `opts` functional options to override the default configuration.
func NewEncoder(optFns ...func(*EncoderOptions)) *Encoder {
	options := EncoderOptions{
		TagKey:        defaultTagKey,
		NullEmptySets: true,
		EncodeTime:    defaultEncodeTime,
	}
	for _, fn := range optFns {
		fn(&options)
	}

	if options.EncodeTime == nil {
		options.EncodeTime = defaultEncodeTime
	}

	return &Encoder{
		options: options,
	}
}

// Encode will marshal a Go value type to an AttributeValue. Returning
// the AttributeValue constructed or error.
func (e *Encoder) Encode(in interface{}) (types.AttributeValue, error) {
	return e.encode(reflect.ValueOf(in), tag{})
}

func (e *Encoder) encode(v reflect.Value, fieldTag tag) (types.AttributeValue, error) {
	// Ignore fields explicitly marked to be skipped.
	if fieldTag.Ignore {
		return nil, nil
	}

	// Zero values are serialized as null, or skipped if omitEmpty.
	if isZeroValue(v) {
		if fieldTag.OmitEmpty && fieldTag.NullEmpty {
			return nil, &InvalidMarshalError{
				msg: "unable to encode AttributeValue for zero value field with incompatible struct tags, omitempty and nullempty"}
		}

		if fieldTag.OmitEmpty {
			return nil, nil
		} else if isNullableZeroValue(v) || fieldTag.NullEmpty {
			return encodeNull(), nil
		}
	}

	// Handle both pointers and interface conversion into types
	v = valueElem(v)

	if v.Kind() != reflect.Invalid {
		if av, err := e.tryMarshaler(v); err != nil {
			return nil, err
		} else if e.options.OmitNullAttributeValues && fieldTag.OmitEmpty && isNullAttributeValue(av) {
			return nil, nil
		} else if av != nil {
			return av, nil
		}
	}

	switch v.Kind() {
	case reflect.Invalid:
		if fieldTag.OmitEmpty {
			return nil, nil
		}
		// Handle case where member type needed to be dereferenced and resulted
		// in a kind that is invalid.
		return encodeNull(), nil

	case reflect.Struct:
		return e.encodeStruct(v, fieldTag)

	case reflect.Map:
		return e.encodeMap(v, fieldTag)

	case reflect.Slice, reflect.Array:
		return e.encodeSlice(v, fieldTag)

	case reflect.Chan, reflect.Func, reflect.UnsafePointer:
		// skip unsupported types
		return nil, nil

	default:
		return e.encodeScalar(v, fieldTag)
	}
}

func (e *Encoder) encodeStruct(v reflect.Value, fieldTag tag) (types.AttributeValue, error) {
	// Time structs have no public members, and instead are converted to
	// RFC3339Nano formatted string, unix time seconds number if struct tag is set.
	if v.Type().ConvertibleTo(timeType) {
		var t time.Time
		t = v.Convert(timeType).Interface().(time.Time)

		if e.options.OmitEmptyTime && fieldTag.OmitEmpty && t.IsZero() {
			return nil, nil
		}

		if fieldTag.AsUnixTime {
			return UnixTime(t).MarshalDynamoDBAttributeValue()
		}
		return e.options.EncodeTime(t)
	}

	m := &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{}}
	fields := unionStructFields(v.Type(), structFieldOptions{
		TagKey: e.options.TagKey,
	})
	for _, f := range fields.All() {
		if f.Name == "" {
			return nil, &InvalidMarshalError{msg: "map key cannot be empty"}
		}

		fv, found := encoderFieldByIndex(v, f.Index)
		if !found {
			continue
		}

		elem, err := e.encode(fv, f.tag)
		if err != nil {
			return nil, err
		} else if elem == nil {
			continue
		}

		m.Value[f.Name] = elem
	}

	return m, nil
}

func (e *Encoder) encodeMap(v reflect.Value, fieldTag tag) (types.AttributeValue, error) {
	m := &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{}}
	for _, key := range v.MapKeys() {
		keyName, err := mapKeyAsString(key, fieldTag)
		if err != nil {
			return nil, err
		}

		elemVal := v.MapIndex(key)
		elem, err := e.encode(elemVal, tag{
			OmitEmpty: fieldTag.OmitEmptyElem,
			NullEmpty: fieldTag.NullEmptyElem,
		})
		if err != nil {
			return nil, err
		} else if elem == nil {
			continue
		}

		m.Value[keyName] = elem
	}

	return m, nil
}

func mapKeyAsString(keyVal reflect.Value, fieldTag tag) (keyStr string, err error) {
	defer func() {
		if err != nil {
			return
		}
		if keyStr == "" {
			err = &InvalidMarshalError{msg: "map key cannot be empty"}
		}
	}()

	if k, ok := keyVal.Interface().(encoding.TextMarshaler); ok {
		b, err := k.MarshalText()
		if err != nil {
			return "", fmt.Errorf("failed to marshal text, %w", err)
		}
		return string(b), err
	}

	switch keyVal.Kind() {
	case reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:

		return fmt.Sprint(keyVal.Interface()), nil

	default:
		return "", &InvalidMarshalError{
			msg: "map key type not supported, must be string, number, bool, or TextMarshaler",
		}
	}
}

func (e *Encoder) encodeSlice(v reflect.Value, fieldTag tag) (types.AttributeValue, error) {
	if v.Type().Elem().Kind() == reflect.Uint8 {
		slice := reflect.MakeSlice(byteSliceType, v.Len(), v.Len())
		reflect.Copy(slice, v)

		return &types.AttributeValueMemberB{
			Value: append([]byte{}, slice.Bytes()...),
		}, nil
	}

	var setElemFn func(types.AttributeValue) error
	var av types.AttributeValue

	if fieldTag.AsBinSet || v.Type() == byteSliceSliceType { // Binary Set
		if v.Len() == 0 && e.options.NullEmptySets {
			return encodeNull(), nil
		}

		bs := &types.AttributeValueMemberBS{Value: make([][]byte, 0, v.Len())}
		av = bs
		setElemFn = func(elem types.AttributeValue) error {
			b, ok := elem.(*types.AttributeValueMemberB)
			if !ok || b == nil || b.Value == nil {
				return &InvalidMarshalError{
					msg: "binary set must only contain non-nil byte slices"}
			}
			bs.Value = append(bs.Value, b.Value)
			return nil
		}

	} else if fieldTag.AsNumSet { // Number Set
		if v.Len() == 0 && e.options.NullEmptySets {
			return encodeNull(), nil
		}

		ns := &types.AttributeValueMemberNS{Value: make([]string, 0, v.Len())}
		av = ns
		setElemFn = func(elem types.AttributeValue) error {
			n, ok := elem.(*types.AttributeValueMemberN)
			if !ok || n == nil {
				return &InvalidMarshalError{
					msg: "number set must only contain non-nil string numbers"}
			}
			ns.Value = append(ns.Value, n.Value)
			return nil
		}

	} else if fieldTag.AsStrSet { // String Set
		if v.Len() == 0 && e.options.NullEmptySets {
			return encodeNull(), nil
		}

		ss := &types.AttributeValueMemberSS{Value: make([]string, 0, v.Len())}
		av = ss
		setElemFn = func(elem types.AttributeValue) error {
			s, ok := elem.(*types.AttributeValueMemberS)
			if !ok || s == nil {
				return &InvalidMarshalError{
					msg: "string set must only contain non-nil strings"}
			}
			ss.Value = append(ss.Value, s.Value)
			return nil
		}

	} else { // List
		l := &types.AttributeValueMemberL{Value: make([]types.AttributeValue, 0, v.Len())}
		av = l
		setElemFn = func(elem types.AttributeValue) error {
			l.Value = append(l.Value, elem)
			return nil
		}
	}

	if err := e.encodeListElems(v, fieldTag, setElemFn); err != nil {
		return nil, err
	}

	return av, nil
}

func (e *Encoder) encodeListElems(v reflect.Value, fieldTag tag, setElem func(types.AttributeValue) error) error {
	for i := 0; i < v.Len(); i++ {
		elem, err := e.encode(v.Index(i), tag{
			OmitEmpty: fieldTag.OmitEmptyElem,
			NullEmpty: fieldTag.NullEmptyElem,
		})
		if err != nil {
			return err
		} else if elem == nil {
			continue
		}

		if err := setElem(elem); err != nil {
			return err
		}
	}

	return nil
}

// Returns if the type of the value satisfies an interface for number like the
// encoding/json#Number and feature/dynamodb/attributevalue#Number
func isNumberValueType(v reflect.Value) bool {
	type numberer interface {
		Float64() (float64, error)
		Int64() (int64, error)
		String() string
	}

	_, ok := v.Interface().(numberer)
	return ok && v.Kind() == reflect.String
}

func (e *Encoder) encodeScalar(v reflect.Value, fieldTag tag) (types.AttributeValue, error) {
	if isNumberValueType(v) {
		if fieldTag.AsString {
			return &types.AttributeValueMemberS{Value: v.String()}, nil
		}
		return &types.AttributeValueMemberN{Value: v.String()}, nil
	}

	switch v.Kind() {
	case reflect.Bool:
		return &types.AttributeValueMemberBOOL{Value: v.Bool()}, nil

	case reflect.String:
		return e.encodeString(v)

	default:
		// Fallback to encoding numbers, will return invalid type if not supported
		av, err := e.encodeNumber(v)
		if err != nil {
			return nil, err
		}

		n, isNumber := av.(*types.AttributeValueMemberN)
		if fieldTag.AsString && isNumber {
			return &types.AttributeValueMemberS{Value: n.Value}, nil
		}
		return av, nil
	}
}

func (e *Encoder) encodeNumber(v reflect.Value) (types.AttributeValue, error) {

	var out string
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		out = encodeInt(v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		out = encodeUint(v.Uint())

	case reflect.Float32:
		out = encodeFloat(v.Float(), 32)

	case reflect.Float64:
		out = encodeFloat(v.Float(), 64)

	default:
		return nil, nil
	}

	return &types.AttributeValueMemberN{Value: out}, nil
}

func (e *Encoder) encodeString(v reflect.Value) (types.AttributeValue, error) {

	switch v.Kind() {
	case reflect.String:
		s := v.String()
		return &types.AttributeValueMemberS{Value: s}, nil

	default:
		return nil, nil
	}
}

func encodeInt(i int64) string {
	return strconv.FormatInt(i, 10)
}
func encodeUint(u uint64) string {
	return strconv.FormatUint(u, 10)
}
func encodeFloat(f float64, bitSize int) string {
	return strconv.FormatFloat(f, 'f', -1, bitSize)
}
func encodeNull() types.AttributeValue {
	return &types.AttributeValueMemberNULL{Value: true}
}

// encoderFieldByIndex finds the field with the provided nested index
func encoderFieldByIndex(v reflect.Value, index []int) (reflect.Value, bool) {
	for i, x := range index {
		if i > 0 && v.Kind() == reflect.Ptr && v.Type().Elem().Kind() == reflect.Struct {
			if v.IsNil() {
				return reflect.Value{}, false
			}
			v = v.Elem()
		}
		v = v.Field(x)
	}
	return v, true
}

func valueElem(v reflect.Value) reflect.Value {
	switch v.Kind() {
	case reflect.Interface, reflect.Ptr:
		for v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
	}

	return v
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func isNullableZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func (e *Encoder) tryMarshaler(v reflect.Value) (types.AttributeValue, error) {
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		v = v.Addr()
	}

	if v.Type().NumMethod() == 0 {
		return nil, nil
	}

	i := v.Interface()
	if m, ok := i.(Marshaler); ok {
		return m.MarshalDynamoDBAttributeValue()
	}
	if e.options.UseEncodingMarshalers {
		return e.tryEncodingMarshaler(i)
	}

	return nil, nil
}

func (e *Encoder) tryEncodingMarshaler(v any) (types.AttributeValue, error) {
	if m, ok := v.(encoding.TextMarshaler); ok {
		s, err := m.MarshalText()
		if err != nil {
			return nil, err
		}

		return &types.AttributeValueMemberS{Value: string(s)}, nil
	}

	if m, ok := v.(encoding.BinaryMarshaler); ok {
		b, err := m.MarshalBinary()
		if err != nil {
			return nil, err
		}

		return &types.AttributeValueMemberB{Value: b}, nil
	}

	return nil, nil
}

// An InvalidMarshalError is an error type representing an error
// occurring when marshaling a Go value type to an AttributeValue.
type InvalidMarshalError struct {
	msg string
}

// Error returns the string representation of the error.
// satisfying the error interface
func (e *InvalidMarshalError) Error() string {
	return fmt.Sprintf("marshal failed, %s", e.msg)
}

func defaultEncodeTime(t time.Time) (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: t.Format(time.RFC3339Nano),
	}, nil
}

func isNullAttributeValue(av types.AttributeValue) bool {
	n, ok := av.(*types.AttributeValueMemberNULL)
	return ok && n.Value
}
