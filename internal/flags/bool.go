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

const negPrefix = "no-"

// AddNegationToBoolFlags adds the given flag in both `--foo` and `--no-foo` variants.
// --no-foo is hidden from help output.
func AddNegationToBoolFlags(f *pflag.FlagSet, p *bool, name string, value bool, usage string) {
	negativeName := negPrefix + name

	f.BoolVarP(p, name, "", value, usage)
	f.Bool(negativeName, !value, "(negative) "+usage)

	if !value {
		name = negativeName
	}
	if err := f.MarkHidden(name); err != nil {
		panic(err)
	}
}

func ReconcileAndLinkBoolFlags(f *pflag.FlagSet) error {
	var err error
	f.VisitAll(func(flag *pflag.Flag) {
		if err != nil {
			return
		}

		// Walk the "no-" versions of the flags. Make sure we didn't set
		// both, and set the positive value to the opposite of the "no-"
		// value if it exists.
		if strings.HasPrefix(flag.Name, negPrefix) {
			positiveName := flag.Name[len(negPrefix):]
			positive := f.Lookup(positiveName)
			// Non-paired flag, or wrong types
			if positive == nil || positive.Value.Type() != "bool" || flag.Value.Type() != "bool" {
				return
			}
			if flag.Changed {
				if positive.Changed {
					err = fmt.Errorf("only one of --%s and --%s may be specified",
						flag.Name, positiveName)
					return
				}
				err = checkExplicitFalse(flag, positiveName)
				if err != nil {
					return
				}
				err = positive.Value.Set("false")
			} else {
				err = checkExplicitFalse(positive, flag.Name)
			}
		}
	})
	return err
}

func checkExplicitFalse(f *pflag.Flag, betterFlag string) error {
	if !f.Changed {
		return nil
	}
	val, err := strconv.ParseBool(f.Value.String())
	if err != nil {
		return err
	}
	if !val {
		return fmt.Errorf("use --%s instead of providing \"%s\" to --%s",
			betterFlag, f.Value.String(), f.Name)
	}
	return nil
}
