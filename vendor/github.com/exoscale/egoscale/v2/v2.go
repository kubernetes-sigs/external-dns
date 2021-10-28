// Package v2 is the new Exoscale client API binding.
// Reference: https://openapi-v2.exoscale.com/
package v2

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type getter interface {
	get(ctx context.Context, client *Client, zone, id string) (interface{}, error)
}

// validateOperationParams is a helper function that returns an error if
// fields of the struct res tagged with `req-for:"<OPERATIONS|*>"` are set
// to a nil value. Fields tagged with "req-for" MUST be of type pointer.
// The expected format for the `req-for:` tag value is a comma-separated
// list of required operations, or "*" for any operation (i.e. the field
// is always required).
func validateOperationParams(res interface{}, op string) error {
	rv := reflect.ValueOf(res)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("field must be a non-nil pointer value")
	}

	if op == "" {
		return errors.New("no operation specified")
	}

	structValue := reflect.ValueOf(res).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		structField := structValue.Type().Field(i)

		reqOp, required := structField.Tag.Lookup("req-for")
		if required {
			if structField.Type.Kind() != reflect.Ptr {
				return fmt.Errorf(
					"%s.%s field is tagged with req-for but its type is not a pointer",
					structValue.Type().String(),
					structField.Name,
				)
			}

			switch {
			case
				reqOp == op,
				reqOp == "*":
				if structValue.Field(i).IsNil() {
					return fmt.Errorf(
						"%s.%s field is required for this operation",
						structValue.Type().String(),
						structField.Name,
					)
				}

			case strings.Contains(reqOp, ","):
				for _, o := range strings.Split(reqOp, ",") {
					if strings.TrimSpace(o) == op && structValue.Field(i).IsNil() {
						return fmt.Errorf(
							"%s.%s field is required for this operation",
							structValue.Type().String(),
							structField.Name,
						)
					}
				}
			}
		}
	}

	return nil
}
