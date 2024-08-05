/*
Copyright 2015 The Kubernetes Authors.

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

package json

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"io"

	kjson "sigs.k8s.io/json"
)

// NewEncoder delegates to json.NewEncoder
// It is only here so this package can be a drop-in for common encoding/json uses
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// Marshal delegates to json.Marshal
// It is only here so this package can be a drop-in for common encoding/json uses
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// limit recursive depth to prevent stack overflow errors
const maxDepth = 10000

<<<<<<< HEAD
// Unmarshal unmarshals the given data
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// If v is a *map[string]interface{}, *[]interface{}, or *interface{} numbers
// are converted to int64 or float64
func Unmarshal(data []byte, v interface{}) error {
	switch v := v.(type) {
	case *map[string]interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertMapNumbers(*v, 0)

	case *[]interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertSliceNumbers(*v, 0)

	case *interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertInterfaceNumbers(v, 0)

	default:
		return json.Unmarshal(data, v)
	}
}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// If v is a *map[string]interface{}, numbers are converted to int64 or float64
||||||| parent of 5ce8c7613 (update vendored files)
// If v is a *map[string]interface{}, numbers are converted to int64 or float64
=======
// If v is a *map[string]interface{}, *[]interface{}, or *interface{} numbers
// are converted to int64 or float64
>>>>>>> 5ce8c7613 (update vendored files)
func Unmarshal(data []byte, v interface{}) error {
	switch v := v.(type) {
	case *map[string]interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertMapNumbers(*v, 0)

	case *[]interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertSliceNumbers(*v, 0)

	case *interface{}:
		// Build a decoder from the given data
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		// Preserve numbers, rather than casting to float64 automatically
		decoder.UseNumber()
		// Run the decode
		if err := decoder.Decode(v); err != nil {
			return err
		}
		// If the decode succeeds, post-process the map to convert json.Number objects to int64 or float64
		return ConvertInterfaceNumbers(v, 0)

	default:
		return json.Unmarshal(data, v)
	}
}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
<<<<<<< HEAD
			err = convertSliceNumbers(v, depth+1)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
			err = convertSliceNumbers(v, depth+1)
=======
			err = ConvertSliceNumbers(v, depth+1)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// If v is a *map[string]interface{}, numbers are converted to int64 or float64
||||||| parent of 6b7ce455e (update vendored files)
// Unmarshal unmarshals the given data
// If v is a *map[string]interface{}, numbers are converted to int64 or float64
=======
// Unmarshal unmarshals the given data.
// Object keys are case-sensitive.
// Numbers decoded into interface{} fields are converted to int64 or float64.
>>>>>>> 6b7ce455e (update vendored files)
func Unmarshal(data []byte, v interface{}) error {
	return kjson.UnmarshalCaseSensitivePreserveInts(data, v)
}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
<<<<<<< HEAD
			err = convertSliceNumbers(v, depth+1)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
			err = convertSliceNumbers(v, depth+1)
=======
			err = ConvertSliceNumbers(v, depth+1)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"bytes"
||||||| parent of 4d7e5ad26 (update vendored files)
	"bytes"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"io"

	kjson "sigs.k8s.io/json"
)

// NewEncoder delegates to json.NewEncoder
// It is only here so this package can be a drop-in for common encoding/json uses
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// Marshal delegates to json.Marshal
// It is only here so this package can be a drop-in for common encoding/json uses
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// limit recursive depth to prevent stack overflow errors
const maxDepth = 10000

// Unmarshal unmarshals the given data.
// Object keys are case-sensitive.
// Numbers decoded into interface{} fields are converted to int64 or float64.
func Unmarshal(data []byte, v interface{}) error {
	return kjson.UnmarshalCaseSensitivePreserveInts(data, v)
}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
<<<<<<< HEAD
			err = convertSliceNumbers(v, depth+1)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
			err = convertSliceNumbers(v, depth+1)
=======
			err = ConvertSliceNumbers(v, depth+1)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"bytes"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"bytes"
=======
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
	"fmt"
	"io"

	kjson "sigs.k8s.io/json"
)

// NewEncoder delegates to json.NewEncoder
// It is only here so this package can be a drop-in for common encoding/json uses
func NewEncoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

// Marshal delegates to json.Marshal
// It is only here so this package can be a drop-in for common encoding/json uses
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// limit recursive depth to prevent stack overflow errors
const maxDepth = 10000

// Unmarshal unmarshals the given data.
// Object keys are case-sensitive.
// Numbers decoded into interface{} fields are converted to int64 or float64.
func Unmarshal(data []byte, v interface{}) error {
	return kjson.UnmarshalCaseSensitivePreserveInts(data, v)
}

// ConvertInterfaceNumbers converts any json.Number values to int64 or float64.
// Values which are map[string]interface{} or []interface{} are recursively visited
func ConvertInterfaceNumbers(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = convertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumbers(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

// ConvertMapNumbers traverses the map, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertMapNumbers(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ConvertSliceNumbers traverses the slice, converting any json.Number values to int64 or float64.
// values which are map[string]interface{} or []interface{} are recursively visited
func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("exceeded max depth of %d", maxDepth)
	}

	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = convertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumbers(v, depth+1)
		case []interface{}:
<<<<<<< HEAD
			err = convertSliceNumbers(v, depth+1)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
			err = convertSliceNumbers(v, depth+1)
=======
			err = ConvertSliceNumbers(v, depth+1)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// convertNumber converts a json.Number to an int64 or float64, or returns an error
func convertNumber(n json.Number) (interface{}, error) {
	// Attempt to convert to an int64 first
	if i, err := n.Int64(); err == nil {
		return i, nil
	}
	// Return a float64 (default json.Decode() behavior)
	// An overflow will return an error
	return n.Float64()
}
