// Copyright 2020 Datawire.  All rights reserved
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

///////////////////////////////////////////////////////////////////////////
// Important: Run "make update-yaml" to regenerate code after modifying
// this file.
///////////////////////////////////////////////////////////////////////////

// I'm not sure where a better place to put this is, so I'm putting it here:
//
// # API design guidelines
//
// Ambassador's API has inconsistencies because it has historical
// baggage.  Not all of Ambassador's existing API (or even most of
// it!?) follow these guidelines, but new additions to the API should.
// If/when we advance to getambassador.io/v3 and we can break
// compatibility, these are things that we should apply everywhere.
//
// - Prefer `camelCase` to `snake_case`
//   * Exception: Except for consistency with existing fields in the
//     same resource, or symmetry with identical fields in another
//     resource.
//   * Justification: Kubernetes style is to use camelCase. But
//     historically Ambassador used snake_case for everything.
//
// - Give _every_ field a `json:""` struct tag.
//   * Justification: Marshaling and unmarshaling are key to what we
//     do, and it's critical to carefully define how it happens.
//   * Notes: This is not optional. Do it for _every field_. (It's OK
//     if the tag is literally `json:""` for fields that must never be
//     exposed during marshaling.)
//
// - Prefer `*int`, `*bool`, and `*BoolOrString`, rather than just
//   `int`, `bool`, and `BoolOrString`.
//   * Justification: The Ambassador API is rooted in Python, where
//     it is always possible to tell if a given element was present in
//     in a CRD, or left unset. This is at odds with Go's `omitempty`
//     specifier, which really means "omit if empty _or if set to the
//     default value". For int in particular, this results in a value
//     of 0 being omitted, and for many Ambassador fields, 0 is not
//     the correct default value.
//
//     This resulted in a lot of bugs in the 1.10 timeframe, so be
//     careful going forward.
//
// - Prefer for object references to not support namespacing
//   * Exception: If there's a real use-case for it.
//   * Justification: Most native Kubernetes resources don't support
//     referencing things in a different namespace.  We should be
//     opinionated and not support it either, unless there's a good
//     reason to in a specific case.
//
// - Prefer to use `corev1.LocalObjectReference` or
//   `corev1.SecretReference` references instead of
//   `{name}.{namespace}` strings.
//   * Justification: The `{name}.{namespace}` thing evolved "an
//     opaque DNS name" in the `service` field of Mappings, and that
//     was generalized to other things.  Outside of the context of
//     "this is usable as a DNS name to make a request to", it's just
//     confusing and introduces needless ambiguity.  Nothing other
//     than Ambassador uses that notation.
//   * Notes: For things that don't support cross-namespace references
//     (see above), use LocalObjectReference; if you really must
//     support cross-namespace references, then use SecretReference.
//
// - Prefer to use `metav1.Duration` fields instead of "_s" or "_ms"
//   numeric fields.
//
// - Don't have Ambassador populate anything in the `.spec` or
//   `.metadata` of something a user might edit, only let Ambassador
//   set things in the `.status`.
//   * Exception: If Ambassador 100% owns the resource and a user will
//     never edit it.
//   * Notes: I didn't write "Prefer" on this one.  Don't violate it.
//     Just don't do it.  Ever.  Designing the Host resource in
//     violation of this was a HUGE mistake and one that I regret very
//     much.  Learn from my mistakes.
//   * Justification: Having Ambassador-set things in a subresource
//     from user-set things:
//     1. avoids races between the user updating the spec and us
//        updating the status
//     2. allows watt/whatever to only pay attention to
//        .metadata.generation instead of .metadata.resourceVersion;
//        avoiding pointless reconfigures.
//     3. allows the RBAC to be simpler
//     4. avoids the whole class of bugs where we need to make sure
//        that everything round-trips correctly
//     5. provides clarity on which things a user is expected to know
//        how to fill in

package v2

import (
	"encoding/json"
)

// The old `k8s.io/kube-openapi/cmd/openapi-gen` command had ways to
// specify custom schemas for your types (1: define a "OpenAPIDefinition"
// method, or 2: define a "OpenAPIV3Definition" method, or 3: define
// "OpenAPISchemaType" and "OpenAPISchemaFormat" methods).  But the new
// `sigs.k8s.io/controller-tools/controller-gen` command doesn't; it just
// has a small number of "+kubebuilder:" magic comments ("markers") that we
// can use to influence the schema it generates.
//
// So, for example, we'd like to define the AmbassadorID schema as:
//
//    oneOf:
//    - type: "string"
//    - type: "array"
//    items:             # only matters if type=array
//      type: "string"
//
// but if we're going to use just vanilla controller-gen, we're forced to
// be dumb and say `+kubebuilder:validation:Type=""`, to define its schema
// as
//
//    # no `type:` setting because of the +kubebuilder marker
//    items:
//      type: "string"  # because of the raw type
//
// and then kubectl and/or the apiserver won't be able to validate
// AmbassadorID, because it won't be validated until we actually go to
// UnmarshalJSON it when it makes it to Ambassador.  That's pretty much
// what Kubernetes itself[1] does for the JSON Schema types that are unions
// like that.
//
//  > Aside: Some recent work in controller-gen[2] *strongly* suggests that
//  > setting `+kubebuilder:validation:Type=Any` instead of `:Type=""` is
//  > the proper thing to do.  But, um, it doesn't work... kubectl would
//  > say things like:
//  >
//  >    Invalid value: "array": spec.ambassador_id in body must be of type Any: "array"
//
// But honestly that's dumb, and we can do better than that.
//
// So, option one choice would be to send the controller-tools folks a PR
// to support the openapi-gen methods to allow that customization.  That's
// probably the Right Thing, but that seemed like more work than option
// two.  FIXME(lukeshu): Send the controller-tools folks a PR.
//
// Option two: Say something nonsensical like
// `+kubebuilder:validation:Type="d6e-union"`, and teach the `fix-crds`
// script to notice that and delete that nonsensical `type`, replacing it
// with the appropriate `oneOf: [type: A, type: B]` (note that the version
// of JSONSchema that OpenAPI/Kubernetes uses doesn't support type being an
// array).  And so that's what I did.
//
// FIXME(lukeshu): But all of that is still terrible.  Because the very
// structure of our data inherently means that we must have a
// non-structural[3] schema.  With "apiextensions.k8s.io/v1beta1" CRDs,
// non-structural schemas disable several features; and in v1 CRDs,
// non-structural schemas are entirely forbidden.  I mean it doesn't
// _really_ matter right now, because we give out v1beta1 CRDs anyway
// because v1 only became available in Kubernetes 1.16 and we still support
// down to Kubernetes 1.11; but I don't think that we want to lock
// ourselves out from v1 forever.  So I guess that means when it comes time
// for `getambassador.io/v3` (`ambassadorlabs.com/v1`?), we need to
// strictly avoid union types, in order to avoid violating rule 3 of
// structural schemas.  Or hope that the Kubernetes folks decide to relax
// some of the structural-schema rules.
//
// [1]: https://github.com/kubernetes/apiextensions-apiserver/blob/kubernetes-1.18.4/pkg/apis/apiextensions/v1beta1/types_jsonschema.go#L195-L206
// [2]: https://github.com/kubernetes-sigs/controller-tools/pull/427
// [3]: https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definitions/#specifying-a-structural-schema

type CircuitBreaker struct {
	// +kubebuilder:validation:Enum={"default", "high"}
	Priority           string `json:"priority,omitempty"`
	MaxConnections     *int   `json:"max_connections,omitempty"`
	MaxPendingRequests *int   `json:"max_pending_requests,omitempty"`
	MaxRequests        *int   `json:"max_requests,omitempty"`
	MaxRetries         *int   `json:"max_retries,omitempty"`
}

// ErrorResponseTextFormatSource specifies a source for an error response body
type ErrorResponseTextFormatSource struct {
	// The name of a file on the Ambassador pod that contains a format text string.
	Filename string `json:"filename"`
}

// ErorrResponseOverrideBody specifies the body of an error response
type ErrorResponseOverrideBody struct {
	// A format string representing a text response body.
	// Content-Type can be set using the `content_type` field below.
	ErrorResponseTextFormat string `json:"text_format,omitempty"`

	// A JSON response with content-type: application/json. The values can
	// contain format text like in text_format.
	ErrorResponseJsonFormat map[string]string `json:"json_format,omitempty"`

	// A format string sourced from a file on the Ambassador container.
	// Useful for larger response bodies that should not be placed inline
	// in configuration.
	ErrorResponseTextFormatSource *ErrorResponseTextFormatSource `json:"text_format_source,omitempty"`

	// The content type to set on the error response body when
	// using text_format or text_format_source. Defaults to 'text/plain'.
	ContentType string `json:"content_type,omitempty"`
}

// A response rewrite for an HTTP error response
type ErrorResponseOverride struct {
	// The status code to match on -- not a pointer because it's required.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=400
	// +kubebuilder:validation:Maximum=599
	OnStatusCode int `json:"on_status_code,omitempty"`

	// The new response body
	// +kubebuilder:validation:Required
	Body ErrorResponseOverrideBody `json:"body,omitempty"`
}

// AmbassadorID declares which Ambassador instances should pay
// attention to this resource.  May either be a string or a list of
// strings.  If no value is provided, the default is:
//
//    ambassador_id:
//    - "default"
//
// +kubebuilder:validation:Type="d6e-union:string,array"
type AmbassadorID []string

func (aid *AmbassadorID) UnmarshalJSON(data []byte) error {
	return (*StringOrStringList)(aid).UnmarshalJSON(data)
}

// StringOrStringList is just what it says on the tin, but note that it will always
// marshal as a list of strings right now.
// +kubebuilder:validation:Type="d6e-union:string,array"
type StringOrStringList []string

func (sl *StringOrStringList) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*sl = nil
		return nil
	}

	var err error
	var list []string
	var single string

	if err = json.Unmarshal(data, &single); err == nil {
		*sl = StringOrStringList([]string{single})
		return nil
	}

	if err = json.Unmarshal(data, &list); err == nil {
		*sl = StringOrStringList(list)
		return nil
	}

	return err
}

// BoolOrString is a type that can hold a Boolean or a string.
//
// +kubebuilder:validation:Type="d6e-union:string,boolean"
type BoolOrString struct {
	String *string
	Bool   *bool
}

// MarshalJSON is important both so that we generate the proper
// output, and to trigger controller-gen to not try to generate
// jsonschema for our sub-fields:
// https://github.com/kubernetes-sigs/controller-tools/pull/427
func (o BoolOrString) MarshalJSON() ([]byte, error) {
	switch {
	case o.String == nil && o.Bool == nil:
		return json.Marshal(nil)
	case o.String == nil && o.Bool != nil:
		return json.Marshal(o.Bool)
	case o.String != nil && o.Bool == nil:
		return json.Marshal(o.String)
	case o.String != nil && o.Bool != nil:
		panic("invalid BoolOrString")
	}
	panic("not reached")
}

func (o *BoolOrString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = BoolOrString{}
		return nil
	}

	var err error

	var b bool
	if err = json.Unmarshal(data, &b); err == nil {
		*o = BoolOrString{Bool: &b}
		return nil
	}

	var str string
	if err = json.Unmarshal(data, &str); err == nil {
		*o = BoolOrString{String: &str}
		return nil
	}

	return err
}

// UntypedDict is relatively opaque as a Go type, but it preserves its contents in a roundtrippable
// way.
// +kubebuilder:validation:Type="object"
type UntypedDict struct {
	Values map[string]UntypedValue
}

func (u UntypedDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Values)
}

func (u *UntypedDict) UnmarshalJSON(data []byte) error {
	var values map[string]UntypedValue
	err := json.Unmarshal(data, &values)
	if err != nil {
		return err
	}
	*u = UntypedDict{Values: values}
	return nil
}

type UntypedValue struct {
	raw json.RawMessage
}

func (u UntypedValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.raw)
}

func (u *UntypedValue) UnmarshalJSON(data []byte) error {
	*u = UntypedValue{raw: json.RawMessage(data)}
	return nil
}
