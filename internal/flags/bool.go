/*
Copyright 2025 The Kubernetes Authors.

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

package flags

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var negPrefix = "no-"

// AddNegationToBoolFlags adds the given flag in both `--foo` and `--no-foo` variants.
// If you do this, make sure you call ReconcileBoolFlags later to catch errors and
// set the relationship between the flag values.
func AddNegationToBoolFlags(f *pflag.FlagSet, p *bool, name string, value bool, usage string) {
	negativeName := negPrefix + name

	f.BoolVarP(p, name, "", value, usage)
	f.Bool(negativeName, !value, "negative flag for --"+name)

	err := f.MarkHidden(negativeName)
	if err != nil {
		panic(err)
	}
}

// ReconcileBoolFlags sets the value of the all the "--foo" flags based on
// "--no-foo" if provided, and returns an error if both were provided or an
// explicit value of false was provided to either (as that's confusing).
func ReconcileBoolFlags(f *pflag.FlagSet) error {
	var err error
	f.VisitAll(func(flag *pflag.Flag) {
		// Return early from our comprehension
		if err != nil {
			return
		}
		// Walk the "no-" versions of the flags. Make sure we didn't set
		// both, and set the positive value to the opposite of the "no-"
		// value if it exists.
		if strings.HasPrefix(flag.Name, "no-") {
			positiveName := flag.Name[len(negPrefix):]
			positive := f.Lookup(positiveName)
			if flag.Changed {
				if positive.Changed {
					err = fmt.Errorf("only one of --%s and --%s may be specified",
						flag.Name, positiveName)
					return
				}
				var noValue bool
				noValue, err = strconv.ParseBool(flag.Value.String())
				if err != nil {
					return
				}
				if !noValue {
					err = fmt.Errorf("use --%s instead of providing false to --%s",
						positiveName, flag.Name)
					if err != nil {
						return
					}
				}
				err = positive.Value.Set(strconv.FormatBool(!noValue))
			} else if positive.Changed {
				// For the positive version, just check it wasn't set to the
				// confusing "false" value.
				var yesValue bool
				yesValue, err = strconv.ParseBool(positive.Value.String())
				if err != nil {
					return
				}
				if !yesValue {
					err = fmt.Errorf("use --%s instead of providing false to --%s",
						flag.Name, positiveName)
					if err != nil {
						return
					}
				}
			}
		}
	})
	return err
}
