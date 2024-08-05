/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package conversion

import (
	"fmt"
	"reflect"
)

type typePair struct {
	source reflect.Type
	dest   reflect.Type
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
type NameFunc func(t reflect.Type) string

var DefaultNameFunc = func(t reflect.Type) string { return t.Name() }

// ConversionFunc converts the object a into the object b, reusing arrays or objects
// or pointers if necessary. It should return an error if the object cannot be converted
// or if some data is invalid. If you do not wish a and b to share fields or nested
// objects, you must copy a before calling this function.
type ConversionFunc func(a, b interface{}, scope Scope) error

// Converter knows how to convert one type to another.
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredConversions        map[typePair]struct{}
	ignoredUntypedConversions map[typePair]struct{}

	// nameFunc is called to retrieve the name of a type; this name is used for the
	// purpose of deciding whether two types match or not (i.e., will we attempt to
	// do a conversion). The default returns the go type name.
	nameFunc func(t reflect.Type) string
}

// NewConverter creates a new Converter object.
func NewConverter(nameFn NameFunc) *Converter {
	c := &Converter{
		conversionFuncs:           NewConversionFuncs(),
		generatedConversionFuncs:  NewConversionFuncs(),
		ignoredConversions:        make(map[typePair]struct{}),
		ignoredUntypedConversions: make(map[typePair]struct{}),
		nameFunc:                  nameFn,
	}
	c.RegisterUntypedConversionFunc(
		(*[]byte)(nil), (*[]byte)(nil),
		func(a, b interface{}, s Scope) error {
			return Convert_Slice_byte_To_Slice_byte(a.(*[]byte), b.(*[]byte), s)
		},
	)
	return c
}

// WithConversions returns a Converter that is a copy of c but with the additional
// fns merged on top.
func (c *Converter) WithConversions(fns ConversionFuncs) *Converter {
	copied := *c
	copied.conversionFuncs = c.conversionFuncs.Merge(fns)
	return &copied
}

// DefaultMeta returns meta for a given type.
func (c *Converter) DefaultMeta(t reflect.Type) *Meta {
	return &Meta{}
}

// Convert_Slice_byte_To_Slice_byte prevents recursing into every byte
func Convert_Slice_byte_To_Slice_byte(in *[]byte, out *[]byte, s Scope) error {
	if *in == nil {
		*out = nil
		return nil
	}
	*out = make([]byte, len(*in))
	copy(*out, *in)
	return nil
}

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
// If multiple converters exist in the system, Scope will allow you to use the correct one
// from a conversion function--that is, the one your conversion function was called by.
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}

func NewConversionFuncs() ConversionFuncs {
	return ConversionFuncs{
		untyped: make(map[typePair]ConversionFunc),
	}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}

// AddUntyped adds the provided conversion function to the lookup table for the types that are
// supplied as a and b. a and b must be pointers or an error is returned. This method overwrites
// previously defined functions.
func (c ConversionFuncs) AddUntyped(a, b interface{}, fn ConversionFunc) error {
	tA, tB := reflect.TypeOf(a), reflect.TypeOf(b)
	if tA.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", a)
	}
	if tB.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", b)
	}
	c.untyped[typePair{tA, tB}] = fn
	return nil
}

// Merge returns a new ConversionFuncs that contains all conversions from
// both other and c, with other conversions taking precedence.
func (c ConversionFuncs) Merge(other ConversionFuncs) ConversionFuncs {
	merged := NewConversionFuncs()
	for k, v := range c.untyped {
		merged.untyped[k] = v
	}
	for k, v := range other.untyped {
		merged.untyped[k] = v
	}
	return merged
}

// Meta is supplied by Scheme, when it calls Convert.
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *Converter
	meta      *Meta
}

// Convert continues a conversion.
func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest, s.meta)
}

// Meta returns the meta object that was originally passed to Convert.
func (s *scope) Meta() *Meta {
	return s.meta
}

// RegisterUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.conversionFuncs.AddUntyped(a, b, fn)
}

// RegisterGeneratedUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterGeneratedUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.generatedConversionFuncs.AddUntyped(a, b, fn)
}

// RegisterIgnoredConversion registers a "no-op" for conversion, where any requested
// conversion between from and to is ignored.
func (c *Converter) RegisterIgnoredConversion(from, to interface{}) error {
	typeFrom := reflect.TypeOf(from)
	typeTo := reflect.TypeOf(to)
	if reflect.TypeOf(from).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'from' param 0, got: %v", typeFrom)
	}
	if typeTo.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'to' param 1, got: %v", typeTo)
	}
	c.ignoredConversions[typePair{typeFrom.Elem(), typeTo.Elem()}] = struct{}{}
	c.ignoredUntypedConversions[typePair{typeFrom, typeTo}] = struct{}{}
	return nil
}

// Convert will translate src to dest if it knows how. Both must be pointers.
// If no conversion func is registered and the default copying mechanism
// doesn't work on this type pair, an error will be returned.
// 'meta' is given to allow you to pass information to conversion functions,
// it is not used by Convert() other than storing it in the scope.
// Not safe for objects with cyclic references!
func (c *Converter) Convert(src, dest interface{}, meta *Meta) error {
	pair := typePair{reflect.TypeOf(src), reflect.TypeOf(dest)}
	scope := &scope{
		converter: c,
		meta:      meta,
	}

	// ignore conversions of this type
	if _, ok := c.ignoredUntypedConversions[pair]; ok {
		return nil
	}
	if fn, ok := c.conversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}
	if fn, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}

	dv, err := EnforcePtr(dest)
	if err != nil {
		return err
	}
	sv, err := EnforcePtr(src)
	if err != nil {
		return err
	}
	return fmt.Errorf("converting (%s) to (%s): unknown conversion", sv.Type(), dv.Type())
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

||||||| parent of 5ce8c7613 (update vendored files)
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

=======
>>>>>>> 5ce8c7613 (update vendored files)
type NameFunc func(t reflect.Type) string

var DefaultNameFunc = func(t reflect.Type) string { return t.Name() }

// ConversionFunc converts the object a into the object b, reusing arrays or objects
// or pointers if necessary. It should return an error if the object cannot be converted
// or if some data is invalid. If you do not wish a and b to share fields or nested
// objects, you must copy a before calling this function.
type ConversionFunc func(a, b interface{}, scope Scope) error

// Converter knows how to convert one type to another.
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredConversions        map[typePair]struct{}
	ignoredUntypedConversions map[typePair]struct{}

	// nameFunc is called to retrieve the name of a type; this name is used for the
	// purpose of deciding whether two types match or not (i.e., will we attempt to
	// do a conversion). The default returns the go type name.
	nameFunc func(t reflect.Type) string
}

// NewConverter creates a new Converter object.
func NewConverter(nameFn NameFunc) *Converter {
	c := &Converter{
		conversionFuncs:           NewConversionFuncs(),
		generatedConversionFuncs:  NewConversionFuncs(),
		ignoredConversions:        make(map[typePair]struct{}),
		ignoredUntypedConversions: make(map[typePair]struct{}),
		nameFunc:                  nameFn,
	}
	c.RegisterUntypedConversionFunc(
		(*[]byte)(nil), (*[]byte)(nil),
		func(a, b interface{}, s Scope) error {
			return Convert_Slice_byte_To_Slice_byte(a.(*[]byte), b.(*[]byte), s)
		},
	)
	return c
}

// WithConversions returns a Converter that is a copy of c but with the additional
// fns merged on top.
func (c *Converter) WithConversions(fns ConversionFuncs) *Converter {
	copied := *c
	copied.conversionFuncs = c.conversionFuncs.Merge(fns)
	return &copied
}

// DefaultMeta returns meta for a given type.
func (c *Converter) DefaultMeta(t reflect.Type) *Meta {
	return &Meta{}
}

// Convert_Slice_byte_To_Slice_byte prevents recursing into every byte
func Convert_Slice_byte_To_Slice_byte(in *[]byte, out *[]byte, s Scope) error {
	if *in == nil {
		*out = nil
		return nil
	}
	*out = make([]byte, len(*in))
	copy(*out, *in)
	return nil
}

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
// If multiple converters exist in the system, Scope will allow you to use the correct one
// from a conversion function--that is, the one your conversion function was called by.
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}

func NewConversionFuncs() ConversionFuncs {
	return ConversionFuncs{
		untyped: make(map[typePair]ConversionFunc),
	}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}

// AddUntyped adds the provided conversion function to the lookup table for the types that are
// supplied as a and b. a and b must be pointers or an error is returned. This method overwrites
// previously defined functions.
func (c ConversionFuncs) AddUntyped(a, b interface{}, fn ConversionFunc) error {
	tA, tB := reflect.TypeOf(a), reflect.TypeOf(b)
	if tA.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", a)
	}
	if tB.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", b)
	}
	c.untyped[typePair{tA, tB}] = fn
	return nil
}

// Merge returns a new ConversionFuncs that contains all conversions from
// both other and c, with other conversions taking precedence.
func (c ConversionFuncs) Merge(other ConversionFuncs) ConversionFuncs {
	merged := NewConversionFuncs()
	for k, v := range c.untyped {
		merged.untyped[k] = v
	}
	for k, v := range other.untyped {
		merged.untyped[k] = v
	}
	return merged
}

// Meta is supplied by Scheme, when it calls Convert.
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *Converter
	meta      *Meta
}

// Convert continues a conversion.
func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest, s.meta)
}

// Meta returns the meta object that was originally passed to Convert.
func (s *scope) Meta() *Meta {
	return s.meta
}

// RegisterUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.conversionFuncs.AddUntyped(a, b, fn)
}

// RegisterGeneratedUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterGeneratedUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.generatedConversionFuncs.AddUntyped(a, b, fn)
}

// RegisterIgnoredConversion registers a "no-op" for conversion, where any requested
// conversion between from and to is ignored.
func (c *Converter) RegisterIgnoredConversion(from, to interface{}) error {
	typeFrom := reflect.TypeOf(from)
	typeTo := reflect.TypeOf(to)
	if reflect.TypeOf(from).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'from' param 0, got: %v", typeFrom)
	}
	if typeTo.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'to' param 1, got: %v", typeTo)
	}
	c.ignoredConversions[typePair{typeFrom.Elem(), typeTo.Elem()}] = struct{}{}
	c.ignoredUntypedConversions[typePair{typeFrom, typeTo}] = struct{}{}
	return nil
}

// Convert will translate src to dest if it knows how. Both must be pointers.
// If no conversion func is registered and the default copying mechanism
// doesn't work on this type pair, an error will be returned.
// 'meta' is given to allow you to pass information to conversion functions,
// it is not used by Convert() other than storing it in the scope.
// Not safe for objects with cyclic references!
func (c *Converter) Convert(src, dest interface{}, meta *Meta) error {
	pair := typePair{reflect.TypeOf(src), reflect.TypeOf(dest)}
	scope := &scope{
		converter: c,
		meta:      meta,
	}

	// ignore conversions of this type
	if _, ok := c.ignoredUntypedConversions[pair]; ok {
		return nil
	}
	if fn, ok := c.conversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}
	if fn, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}

	dv, err := EnforcePtr(dest)
	if err != nil {
		return err
	}
	sv, err := EnforcePtr(src)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
=======
	return fmt.Errorf("converting (%s) to (%s): unknown conversion", sv.Type(), dv.Type())
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

||||||| parent of 6b7ce455e (update vendored files)
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

=======
>>>>>>> 6b7ce455e (update vendored files)
type NameFunc func(t reflect.Type) string

var DefaultNameFunc = func(t reflect.Type) string { return t.Name() }

// ConversionFunc converts the object a into the object b, reusing arrays or objects
// or pointers if necessary. It should return an error if the object cannot be converted
// or if some data is invalid. If you do not wish a and b to share fields or nested
// objects, you must copy a before calling this function.
type ConversionFunc func(a, b interface{}, scope Scope) error

// Converter knows how to convert one type to another.
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredUntypedConversions map[typePair]struct{}
}

// NewConverter creates a new Converter object.
// Arg NameFunc is just for backward compatibility.
func NewConverter(NameFunc) *Converter {
	c := &Converter{
		conversionFuncs:           NewConversionFuncs(),
		generatedConversionFuncs:  NewConversionFuncs(),
		ignoredUntypedConversions: make(map[typePair]struct{}),
	}
	c.RegisterUntypedConversionFunc(
		(*[]byte)(nil), (*[]byte)(nil),
		func(a, b interface{}, s Scope) error {
			return Convert_Slice_byte_To_Slice_byte(a.(*[]byte), b.(*[]byte), s)
		},
	)
	return c
}

// WithConversions returns a Converter that is a copy of c but with the additional
// fns merged on top.
func (c *Converter) WithConversions(fns ConversionFuncs) *Converter {
	copied := *c
	copied.conversionFuncs = c.conversionFuncs.Merge(fns)
	return &copied
}

// DefaultMeta returns meta for a given type.
func (c *Converter) DefaultMeta(t reflect.Type) *Meta {
	return &Meta{}
}

// Convert_Slice_byte_To_Slice_byte prevents recursing into every byte
func Convert_Slice_byte_To_Slice_byte(in *[]byte, out *[]byte, s Scope) error {
	if *in == nil {
		*out = nil
		return nil
	}
	*out = make([]byte, len(*in))
	copy(*out, *in)
	return nil
}

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
// If multiple converters exist in the system, Scope will allow you to use the correct one
// from a conversion function--that is, the one your conversion function was called by.
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}

func NewConversionFuncs() ConversionFuncs {
	return ConversionFuncs{
		untyped: make(map[typePair]ConversionFunc),
	}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}

// AddUntyped adds the provided conversion function to the lookup table for the types that are
// supplied as a and b. a and b must be pointers or an error is returned. This method overwrites
// previously defined functions.
func (c ConversionFuncs) AddUntyped(a, b interface{}, fn ConversionFunc) error {
	tA, tB := reflect.TypeOf(a), reflect.TypeOf(b)
	if tA.Kind() != reflect.Pointer {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", a)
	}
	if tB.Kind() != reflect.Pointer {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", b)
	}
	c.untyped[typePair{tA, tB}] = fn
	return nil
}

// Merge returns a new ConversionFuncs that contains all conversions from
// both other and c, with other conversions taking precedence.
func (c ConversionFuncs) Merge(other ConversionFuncs) ConversionFuncs {
	merged := NewConversionFuncs()
	for k, v := range c.untyped {
		merged.untyped[k] = v
	}
	for k, v := range other.untyped {
		merged.untyped[k] = v
	}
	return merged
}

// Meta is supplied by Scheme, when it calls Convert.
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *Converter
	meta      *Meta
}

// Convert continues a conversion.
func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest, s.meta)
}

// Meta returns the meta object that was originally passed to Convert.
func (s *scope) Meta() *Meta {
	return s.meta
}

// RegisterUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.conversionFuncs.AddUntyped(a, b, fn)
}

// RegisterGeneratedUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterGeneratedUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.generatedConversionFuncs.AddUntyped(a, b, fn)
}

// RegisterIgnoredConversion registers a "no-op" for conversion, where any requested
// conversion between from and to is ignored.
func (c *Converter) RegisterIgnoredConversion(from, to interface{}) error {
	typeFrom := reflect.TypeOf(from)
	typeTo := reflect.TypeOf(to)
	if typeFrom.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer arg for 'from' param 0, got: %v", typeFrom)
	}
	if typeTo.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer arg for 'to' param 1, got: %v", typeTo)
	}
	c.ignoredUntypedConversions[typePair{typeFrom, typeTo}] = struct{}{}
	return nil
}

// Convert will translate src to dest if it knows how. Both must be pointers.
// If no conversion func is registered and the default copying mechanism
// doesn't work on this type pair, an error will be returned.
// 'meta' is given to allow you to pass information to conversion functions,
// it is not used by Convert() other than storing it in the scope.
// Not safe for objects with cyclic references!
func (c *Converter) Convert(src, dest interface{}, meta *Meta) error {
	pair := typePair{reflect.TypeOf(src), reflect.TypeOf(dest)}
	scope := &scope{
		converter: c,
		meta:      meta,
	}

	// ignore conversions of this type
	if _, ok := c.ignoredUntypedConversions[pair]; ok {
		return nil
	}
	if fn, ok := c.conversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}
	if fn, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}

	dv, err := EnforcePtr(dest)
	if err != nil {
		return err
	}
	sv, err := EnforcePtr(src)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
=======
	return fmt.Errorf("converting (%s) to (%s): unknown conversion", sv.Type(), dv.Type())
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

||||||| parent of 4d7e5ad26 (update vendored files)
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

=======
>>>>>>> 4d7e5ad26 (update vendored files)
type NameFunc func(t reflect.Type) string

var DefaultNameFunc = func(t reflect.Type) string { return t.Name() }

// ConversionFunc converts the object a into the object b, reusing arrays or objects
// or pointers if necessary. It should return an error if the object cannot be converted
// or if some data is invalid. If you do not wish a and b to share fields or nested
// objects, you must copy a before calling this function.
type ConversionFunc func(a, b interface{}, scope Scope) error

// Converter knows how to convert one type to another.
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredUntypedConversions map[typePair]struct{}
}

// NewConverter creates a new Converter object.
// Arg NameFunc is just for backward compatibility.
func NewConverter(NameFunc) *Converter {
	c := &Converter{
		conversionFuncs:           NewConversionFuncs(),
		generatedConversionFuncs:  NewConversionFuncs(),
		ignoredUntypedConversions: make(map[typePair]struct{}),
	}
	c.RegisterUntypedConversionFunc(
		(*[]byte)(nil), (*[]byte)(nil),
		func(a, b interface{}, s Scope) error {
			return Convert_Slice_byte_To_Slice_byte(a.(*[]byte), b.(*[]byte), s)
		},
	)
	return c
}

// WithConversions returns a Converter that is a copy of c but with the additional
// fns merged on top.
func (c *Converter) WithConversions(fns ConversionFuncs) *Converter {
	copied := *c
	copied.conversionFuncs = c.conversionFuncs.Merge(fns)
	return &copied
}

// DefaultMeta returns meta for a given type.
func (c *Converter) DefaultMeta(t reflect.Type) *Meta {
	return &Meta{}
}

// Convert_Slice_byte_To_Slice_byte prevents recursing into every byte
func Convert_Slice_byte_To_Slice_byte(in *[]byte, out *[]byte, s Scope) error {
	if *in == nil {
		*out = nil
		return nil
	}
	*out = make([]byte, len(*in))
	copy(*out, *in)
	return nil
}

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
// If multiple converters exist in the system, Scope will allow you to use the correct one
// from a conversion function--that is, the one your conversion function was called by.
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}

func NewConversionFuncs() ConversionFuncs {
	return ConversionFuncs{
		untyped: make(map[typePair]ConversionFunc),
	}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}

// AddUntyped adds the provided conversion function to the lookup table for the types that are
// supplied as a and b. a and b must be pointers or an error is returned. This method overwrites
// previously defined functions.
func (c ConversionFuncs) AddUntyped(a, b interface{}, fn ConversionFunc) error {
	tA, tB := reflect.TypeOf(a), reflect.TypeOf(b)
	if tA.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", a)
	}
	if tB.Kind() != reflect.Ptr {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", b)
	}
	c.untyped[typePair{tA, tB}] = fn
	return nil
}

// Merge returns a new ConversionFuncs that contains all conversions from
// both other and c, with other conversions taking precedence.
func (c ConversionFuncs) Merge(other ConversionFuncs) ConversionFuncs {
	merged := NewConversionFuncs()
	for k, v := range c.untyped {
		merged.untyped[k] = v
	}
	for k, v := range other.untyped {
		merged.untyped[k] = v
	}
	return merged
}

// Meta is supplied by Scheme, when it calls Convert.
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *Converter
	meta      *Meta
}

// Convert continues a conversion.
func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest, s.meta)
}

// Meta returns the meta object that was originally passed to Convert.
func (s *scope) Meta() *Meta {
	return s.meta
}

// RegisterUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.conversionFuncs.AddUntyped(a, b, fn)
}

// RegisterGeneratedUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterGeneratedUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.generatedConversionFuncs.AddUntyped(a, b, fn)
}

// RegisterIgnoredConversion registers a "no-op" for conversion, where any requested
// conversion between from and to is ignored.
func (c *Converter) RegisterIgnoredConversion(from, to interface{}) error {
	typeFrom := reflect.TypeOf(from)
	typeTo := reflect.TypeOf(to)
	if reflect.TypeOf(from).Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'from' param 0, got: %v", typeFrom)
	}
	if typeTo.Kind() != reflect.Ptr {
		return fmt.Errorf("expected pointer arg for 'to' param 1, got: %v", typeTo)
	}
	c.ignoredUntypedConversions[typePair{typeFrom, typeTo}] = struct{}{}
	return nil
}

// Convert will translate src to dest if it knows how. Both must be pointers.
// If no conversion func is registered and the default copying mechanism
// doesn't work on this type pair, an error will be returned.
// 'meta' is given to allow you to pass information to conversion functions,
// it is not used by Convert() other than storing it in the scope.
// Not safe for objects with cyclic references!
func (c *Converter) Convert(src, dest interface{}, meta *Meta) error {
	pair := typePair{reflect.TypeOf(src), reflect.TypeOf(dest)}
	scope := &scope{
		converter: c,
		meta:      meta,
	}

	// ignore conversions of this type
	if _, ok := c.ignoredUntypedConversions[pair]; ok {
		return nil
	}
	if fn, ok := c.conversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}
	if fn, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}

	dv, err := EnforcePtr(dest)
	if err != nil {
		return err
	}
	sv, err := EnforcePtr(src)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
=======
	return fmt.Errorf("converting (%s) to (%s): unknown conversion", sv.Type(), dv.Type())
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
type typeNamePair struct {
	fieldType reflect.Type
	fieldName string
}

// DebugLogger allows you to get debugging messages if necessary.
type DebugLogger interface {
	Logf(format string, args ...interface{})
}

=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
type NameFunc func(t reflect.Type) string

var DefaultNameFunc = func(t reflect.Type) string { return t.Name() }

// ConversionFunc converts the object a into the object b, reusing arrays or objects
// or pointers if necessary. It should return an error if the object cannot be converted
// or if some data is invalid. If you do not wish a and b to share fields or nested
// objects, you must copy a before calling this function.
type ConversionFunc func(a, b interface{}, scope Scope) error

// Converter knows how to convert one type to another.
type Converter struct {
	// Map from the conversion pair to a function which can
	// do the conversion.
	conversionFuncs          ConversionFuncs
	generatedConversionFuncs ConversionFuncs

	// Set of conversions that should be treated as a no-op
	ignoredUntypedConversions map[typePair]struct{}
}

// NewConverter creates a new Converter object.
// Arg NameFunc is just for backward compatibility.
func NewConverter(NameFunc) *Converter {
	c := &Converter{
		conversionFuncs:           NewConversionFuncs(),
		generatedConversionFuncs:  NewConversionFuncs(),
		ignoredUntypedConversions: make(map[typePair]struct{}),
	}
	c.RegisterUntypedConversionFunc(
		(*[]byte)(nil), (*[]byte)(nil),
		func(a, b interface{}, s Scope) error {
			return Convert_Slice_byte_To_Slice_byte(a.(*[]byte), b.(*[]byte), s)
		},
	)
	return c
}

// WithConversions returns a Converter that is a copy of c but with the additional
// fns merged on top.
func (c *Converter) WithConversions(fns ConversionFuncs) *Converter {
	copied := *c
	copied.conversionFuncs = c.conversionFuncs.Merge(fns)
	return &copied
}

// DefaultMeta returns meta for a given type.
func (c *Converter) DefaultMeta(t reflect.Type) *Meta {
	return &Meta{}
}

// Convert_Slice_byte_To_Slice_byte prevents recursing into every byte
func Convert_Slice_byte_To_Slice_byte(in *[]byte, out *[]byte, s Scope) error {
	if *in == nil {
		*out = nil
		return nil
	}
	*out = make([]byte, len(*in))
	copy(*out, *in)
	return nil
}

// Scope is passed to conversion funcs to allow them to continue an ongoing conversion.
// If multiple converters exist in the system, Scope will allow you to use the correct one
// from a conversion function--that is, the one your conversion function was called by.
type Scope interface {
	// Call Convert to convert sub-objects. Note that if you call it with your own exact
	// parameters, you'll run out of stack space before anything useful happens.
	Convert(src, dest interface{}) error

	// Meta returns any information originally passed to Convert.
	Meta() *Meta
}

func NewConversionFuncs() ConversionFuncs {
	return ConversionFuncs{
		untyped: make(map[typePair]ConversionFunc),
	}
}

type ConversionFuncs struct {
	untyped map[typePair]ConversionFunc
}

// AddUntyped adds the provided conversion function to the lookup table for the types that are
// supplied as a and b. a and b must be pointers or an error is returned. This method overwrites
// previously defined functions.
func (c ConversionFuncs) AddUntyped(a, b interface{}, fn ConversionFunc) error {
	tA, tB := reflect.TypeOf(a), reflect.TypeOf(b)
	if tA.Kind() != reflect.Pointer {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", a)
	}
	if tB.Kind() != reflect.Pointer {
		return fmt.Errorf("the type %T must be a pointer to register as an untyped conversion", b)
	}
	c.untyped[typePair{tA, tB}] = fn
	return nil
}

// Merge returns a new ConversionFuncs that contains all conversions from
// both other and c, with other conversions taking precedence.
func (c ConversionFuncs) Merge(other ConversionFuncs) ConversionFuncs {
	merged := NewConversionFuncs()
	for k, v := range c.untyped {
		merged.untyped[k] = v
	}
	for k, v := range other.untyped {
		merged.untyped[k] = v
	}
	return merged
}

// Meta is supplied by Scheme, when it calls Convert.
type Meta struct {
	// Context is an optional field that callers may use to pass info to conversion functions.
	Context interface{}
}

// scope contains information about an ongoing conversion.
type scope struct {
	converter *Converter
	meta      *Meta
}

// Convert continues a conversion.
func (s *scope) Convert(src, dest interface{}) error {
	return s.converter.Convert(src, dest, s.meta)
}

// Meta returns the meta object that was originally passed to Convert.
func (s *scope) Meta() *Meta {
	return s.meta
}

// RegisterUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.conversionFuncs.AddUntyped(a, b, fn)
}

// RegisterGeneratedUntypedConversionFunc registers a function that converts between a and b by passing objects of those
// types to the provided function. The function *must* accept objects of a and b - this machinery will not enforce
// any other guarantee.
func (c *Converter) RegisterGeneratedUntypedConversionFunc(a, b interface{}, fn ConversionFunc) error {
	return c.generatedConversionFuncs.AddUntyped(a, b, fn)
}

// RegisterIgnoredConversion registers a "no-op" for conversion, where any requested
// conversion between from and to is ignored.
func (c *Converter) RegisterIgnoredConversion(from, to interface{}) error {
	typeFrom := reflect.TypeOf(from)
	typeTo := reflect.TypeOf(to)
	if typeFrom.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer arg for 'from' param 0, got: %v", typeFrom)
	}
	if typeTo.Kind() != reflect.Pointer {
		return fmt.Errorf("expected pointer arg for 'to' param 1, got: %v", typeTo)
	}
	c.ignoredUntypedConversions[typePair{typeFrom, typeTo}] = struct{}{}
	return nil
}

// Convert will translate src to dest if it knows how. Both must be pointers.
// If no conversion func is registered and the default copying mechanism
// doesn't work on this type pair, an error will be returned.
// 'meta' is given to allow you to pass information to conversion functions,
// it is not used by Convert() other than storing it in the scope.
// Not safe for objects with cyclic references!
func (c *Converter) Convert(src, dest interface{}, meta *Meta) error {
	pair := typePair{reflect.TypeOf(src), reflect.TypeOf(dest)}
	scope := &scope{
		converter: c,
		meta:      meta,
	}

	// ignore conversions of this type
	if _, ok := c.ignoredUntypedConversions[pair]; ok {
		return nil
	}
	if fn, ok := c.conversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}
	if fn, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return fn(src, dest, scope)
	}

	dv, err := EnforcePtr(dest)
	if err != nil {
		return err
	}
	sv, err := EnforcePtr(src)
	if err != nil {
		return err
	}
<<<<<<< HEAD
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// Leave something on the stack, so that calls to struct tag getters never fail.
	scope.srcStack.push(scopeStackElem{})
	scope.destStack.push(scopeStackElem{})
	return f(sv, dv, scope)
}

// callUntyped calls predefined conversion func.
func (c *Converter) callUntyped(sv, dv reflect.Value, f ConversionFunc, scope *scope) error {
	if !dv.CanAddr() {
		return scope.errorf("cant addr dest")
	}
	var svPointer reflect.Value
	if sv.CanAddr() {
		svPointer = sv.Addr()
	} else {
		svPointer = reflect.New(sv.Type())
		svPointer.Elem().Set(sv)
	}
	dvPointer := dv.Addr()
	return f(svPointer.Interface(), dvPointer.Interface(), scope)
}

// convert recursively copies sv into dv, calling an appropriate conversion function if
// one is registered.
func (c *Converter) convert(sv, dv reflect.Value, scope *scope) error {
	dt, st := dv.Type(), sv.Type()
	pair := typePair{st, dt}

	// ignore conversions of this type
	if _, ok := c.ignoredConversions[pair]; ok {
		if c.Debug != nil {
			c.Debug.Logf("Ignoring conversion of '%v' to '%v'", st, dt)
		}
		return nil
	}

	// Convert sv to dv.
	pair = typePair{reflect.PtrTo(sv.Type()), reflect.PtrTo(dv.Type())}
	if f, ok := c.conversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}
	if f, ok := c.generatedConversionFuncs.untyped[pair]; ok {
		return c.callUntyped(sv, dv, f, scope)
	}

	if !dv.CanSet() {
		return scope.errorf("Cannot set dest. (Tried to deep copy something with unexported fields?)")
	}

	if !scope.flags.IsSet(AllowDifferentFieldTypeNames) && c.nameFunc(dt) != c.nameFunc(st) {
		return scope.errorf(
			"type names don't match (%v, %v), and no conversion 'func (%v, %v) error' registered.",
			c.nameFunc(st), c.nameFunc(dt), st, dt)
	}

	switch st.Kind() {
	case reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Struct:
		// Don't copy these via assignment/conversion!
	default:
		// This should handle all simple types.
		if st.AssignableTo(dt) {
			dv.Set(sv)
			return nil
		}
		if st.ConvertibleTo(dt) {
			dv.Set(sv.Convert(dt))
			return nil
		}
	}

	if c.Debug != nil {
		c.Debug.Logf("Trying to convert '%v' to '%v'", st, dt)
	}

	scope.srcStack.push(scopeStackElem{value: sv})
	scope.destStack.push(scopeStackElem{value: dv})
	defer scope.srcStack.pop()
	defer scope.destStack.pop()

	switch dv.Kind() {
	case reflect.Struct:
		return c.convertKV(toKVValue(sv), toKVValue(dv), scope)
	case reflect.Slice:
		if sv.IsNil() {
			// Don't make a zero-length slice.
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeSlice(dt, sv.Len(), sv.Cap()))
		for i := 0; i < sv.Len(); i++ {
			scope.setIndices(i, i)
			if err := c.convert(sv.Index(i), dv.Index(i), scope); err != nil {
				return err
			}
		}
	case reflect.Ptr:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.New(dt.Elem()))
		switch st.Kind() {
		case reflect.Ptr, reflect.Interface:
			return c.convert(sv.Elem(), dv.Elem(), scope)
		default:
			return c.convert(sv, dv.Elem(), scope)
		}
	case reflect.Map:
		if sv.IsNil() {
			// Don't copy a nil ptr!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		dv.Set(reflect.MakeMap(dt))
		for _, sk := range sv.MapKeys() {
			dk := reflect.New(dt.Key()).Elem()
			if err := c.convert(sk, dk, scope); err != nil {
				return err
			}
			dkv := reflect.New(dt.Elem()).Elem()
			scope.setKeys(sk.Interface(), dk.Interface())
			// TODO:  sv.MapIndex(sk) may return a value with CanAddr() == false,
			// because a map[string]struct{} does not allow a pointer reference.
			// Calling a custom conversion function defined for the map value
			// will panic. Example is PodInfo map[string]ContainerStatus.
			if err := c.convert(sv.MapIndex(sk), dkv, scope); err != nil {
				return err
			}
			dv.SetMapIndex(dk, dkv)
		}
	case reflect.Interface:
		if sv.IsNil() {
			// Don't copy a nil interface!
			dv.Set(reflect.Zero(dt))
			return nil
		}
		tmpdv := reflect.New(sv.Elem().Type()).Elem()
		if err := c.convert(sv.Elem(), tmpdv, scope); err != nil {
			return err
		}
		dv.Set(reflect.ValueOf(tmpdv.Interface()))
		return nil
	default:
		return scope.errorf("couldn't copy '%v' into '%v'; didn't understand types", st, dt)
	}
	return nil
}

var stringType = reflect.TypeOf("")

func toKVValue(v reflect.Value) kvValue {
	switch v.Kind() {
	case reflect.Struct:
		return structAdaptor(v)
	case reflect.Map:
		if v.Type().Key().AssignableTo(stringType) {
			return stringMapAdaptor(v)
		}
	}

	return nil
}

// kvValue lets us write the same conversion logic to work with both maps
// and structs. Only maps with string keys make sense for this.
type kvValue interface {
	// returns all keys, as a []string.
	keys() []string
	// Will just return "" for maps.
	tagOf(key string) reflect.StructTag
	// Will return the zero Value if the key doesn't exist.
	value(key string) reflect.Value
	// Maps require explicit setting-- will do nothing for structs.
	// Returns false on failure.
	confirmSet(key string, v reflect.Value) bool
}

type stringMapAdaptor reflect.Value

func (a stringMapAdaptor) len() int {
	return reflect.Value(a).Len()
}

func (a stringMapAdaptor) keys() []string {
	v := reflect.Value(a)
	keys := make([]string, v.Len())
	for i, v := range v.MapKeys() {
		if v.IsNil() {
			continue
		}
		switch t := v.Interface().(type) {
		case string:
			keys[i] = t
		}
	}
	return keys
}

func (a stringMapAdaptor) tagOf(key string) reflect.StructTag {
	return ""
}

func (a stringMapAdaptor) value(key string) reflect.Value {
	return reflect.Value(a).MapIndex(reflect.ValueOf(key))
}

func (a stringMapAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

type structAdaptor reflect.Value

func (a structAdaptor) len() int {
	v := reflect.Value(a)
	return v.Type().NumField()
}

func (a structAdaptor) keys() []string {
	v := reflect.Value(a)
	t := v.Type()
	keys := make([]string, t.NumField())
	for i := range keys {
		keys[i] = t.Field(i).Name
	}
	return keys
}

func (a structAdaptor) tagOf(key string) reflect.StructTag {
	v := reflect.Value(a)
	field, ok := v.Type().FieldByName(key)
	if ok {
		return field.Tag
	}
	return ""
}

func (a structAdaptor) value(key string) reflect.Value {
	v := reflect.Value(a)
	return v.FieldByName(key)
}

func (a structAdaptor) confirmSet(key string, v reflect.Value) bool {
	return true
}

// convertKV can convert things that consist of key/value pairs, like structs
// and some maps.
func (c *Converter) convertKV(skv, dkv kvValue, scope *scope) error {
	if skv == nil || dkv == nil {
		// TODO: add keys to stack to support really understandable error messages.
		return fmt.Errorf("Unable to convert %#v to %#v", skv, dkv)
	}

	lister := dkv
	if scope.flags.IsSet(SourceToDest) {
		lister = skv
	}

	var mapping FieldMappingFunc
	if scope.meta != nil && scope.meta.KeyNameMapping != nil {
		mapping = scope.meta.KeyNameMapping
	}

	for _, key := range lister.keys() {
		if found, err := c.checkField(key, skv, dkv, scope); found {
			if err != nil {
				return err
			}
			continue
		}
		stag := skv.tagOf(key)
		dtag := dkv.tagOf(key)
		skey := key
		dkey := key
		if mapping != nil {
			skey, dkey = scope.meta.KeyNameMapping(key, stag, dtag)
		}

		df := dkv.value(dkey)
		sf := skv.value(skey)
		if !df.IsValid() || !sf.IsValid() {
			switch {
			case scope.flags.IsSet(IgnoreMissingFields):
				// No error.
			case scope.flags.IsSet(SourceToDest):
				return scope.errorf("%v not present in dest", dkey)
			default:
				return scope.errorf("%v not present in src", skey)
			}
			continue
		}
		scope.srcStack.top().key = skey
		scope.srcStack.top().tag = stag
		scope.destStack.top().key = dkey
		scope.destStack.top().tag = dtag
		if err := c.convert(sf, df, scope); err != nil {
			return err
		}
	}
	return nil
}

// checkField returns true if the field name matches any of the struct
// field copying rules. The error should be ignored if it returns false.
func (c *Converter) checkField(fieldName string, skv, dkv kvValue, scope *scope) (bool, error) {
	replacementMade := false
	if scope.flags.IsSet(DestFromSource) {
		df := dkv.value(fieldName)
		if !df.IsValid() {
			return false, nil
		}
		destKey := typeNamePair{df.Type(), fieldName}
		// Check each of the potential source (type, name) pairs to see if they're
		// present in sv.
		for _, potentialSourceKey := range c.structFieldSources[destKey] {
			sf := skv.value(potentialSourceKey.fieldName)
			if !sf.IsValid() {
				continue
			}
			if sf.Type() == potentialSourceKey.fieldType {
				// Both the source's name and type matched, so copy.
				scope.srcStack.top().key = potentialSourceKey.fieldName
				scope.destStack.top().key = fieldName
				if err := c.convert(sf, df, scope); err != nil {
					return true, err
				}
				dkv.confirmSet(fieldName, df)
				replacementMade = true
			}
		}
		return replacementMade, nil
	}

	sf := skv.value(fieldName)
	if !sf.IsValid() {
		return false, nil
	}
	srcKey := typeNamePair{sf.Type(), fieldName}
	// Check each of the potential dest (type, name) pairs to see if they're
	// present in dv.
	for _, potentialDestKey := range c.structFieldDests[srcKey] {
		df := dkv.value(potentialDestKey.fieldName)
		if !df.IsValid() {
			continue
		}
		if df.Type() == potentialDestKey.fieldType {
			// Both the dest's name and type matched, so copy.
			scope.srcStack.top().key = fieldName
			scope.destStack.top().key = potentialDestKey.fieldName
			if err := c.convert(sf, df, scope); err != nil {
				return true, err
			}
			dkv.confirmSet(potentialDestKey.fieldName, df)
			replacementMade = true
		}
	}
	return replacementMade, nil
=======
	return fmt.Errorf("converting (%s) to (%s): unknown conversion", sv.Type(), dv.Type())
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
