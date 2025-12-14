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

package externaldns

import (
	"fmt"
	"regexp"
	"time"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns/flags"
)

// FlagBinder abstracts flag registration for different CLI backends.
type FlagBinder interface {
	StringVar(name, help, def string, target *string)
	BoolVar(name, help string, def bool, target *bool)
	DurationVar(name, help string, def time.Duration, target *time.Duration)
	IntVar(name, help string, def int, target *int)
	Int64Var(name, help string, def int64, target *int64)
	StringsVar(name, help string, def []string, target *[]string)
	EnumVar(name, help, def string, target *string, allowed ...string)
	// StringsEnumVar binds a repeatable string flag with an allowed set.
	// Implementations may not enforce allowed values.
	StringsEnumVar(name, help string, def []string, target *[]string, allowed ...string)
	// StringMapVar binds key=value repeatable flags into a map.
	StringMapVar(name, help string, target *map[string]string)
	// RegexpVar binds a regular expression value.
	RegexpVar(name, help string, def *regexp.Regexp, target **regexp.Regexp)
	Flags() *flag.FlagSet
}

type regexpValue struct {
	target **regexp.Regexp
}

func (rv *regexpValue) String() string {
	if rv == nil || rv.target == nil || *rv.target == nil {
		return ""
	}
	return (*rv.target).String()
}

func (rv *regexpValue) Set(s string) error {
	re, err := regexp.Compile(s)
	if err != nil {
		return err
	}
	*rv.target = re
	return nil
}

func (rv *regexpValue) Type() string { return "regexp" }

type regexpSetter interface {
	Set(string) error
}

func setRegexpDefault(rs regexpSetter, def *regexp.Regexp, name string) {
	if def != nil {
		if err := rs.Set(def.String()); err != nil {
			panic(fmt.Errorf("invalid default regexp for flag %s: %w", name, err))
		}
	}
}

// CobraBinder implements FlagBinder using github.com/spf13/cobra.
type CobraBinder struct {
	Cmd *cobra.Command
}

// NewCobraBinder creates a FlagBinder backed by a Cobra command.
func NewCobraBinder(cmd *cobra.Command) *CobraBinder {
	return &CobraBinder{Cmd: cmd}
}

func (b *CobraBinder) StringVar(name, help, def string, target *string) {
	b.Cmd.Flags().StringVar(target, name, def, help)
}

func (b *CobraBinder) BoolVar(name, help string, def bool, target *bool) {
	flags.AddNegationToBoolFlags(b.Cmd.Flags(), target, name, "", def, help)
}

func (b *CobraBinder) DurationVar(name, help string, def time.Duration, target *time.Duration) {
	b.Cmd.Flags().DurationVar(target, name, def, help)
}

func (b *CobraBinder) IntVar(name, help string, def int, target *int) {
	b.Cmd.Flags().IntVar(target, name, def, help)
}

func (b *CobraBinder) Int64Var(name, help string, def int64, target *int64) {
	b.Cmd.Flags().Int64Var(target, name, def, help)
}

func (b *CobraBinder) StringsVar(name, help string, def []string, target *[]string) {
	// Preserve repeatable flag semantics.
	b.Cmd.Flags().StringArrayVar(target, name, def, help)
}

func (b *CobraBinder) EnumVar(name, help, def string, target *string, allowed ...string) {
	b.Cmd.Flags().StringVar(target, name, def, help)
}

func (b *CobraBinder) StringsEnumVar(name, help string, def []string, target *[]string, allowed ...string) {
	// pflag does not enforce enums.
	b.Cmd.Flags().StringArrayVar(target, name, def, help)
}

func (b *CobraBinder) StringMapVar(name, help string, target *map[string]string) {
	// Use StringToStringVar for key=value pairs.
	b.Cmd.Flags().StringToStringVar(target, name, map[string]string{}, help)
}

func (b *CobraBinder) RegexpVar(name, help string, def *regexp.Regexp, target **regexp.Regexp) {
	rv := &regexpValue{target: target}
	// set default value to mirror kingpin's Default(def.String()) behavior
	setRegexpDefault(rv, def, name)
	b.Cmd.Flags().Var(rv, name, help)
}

func (b *CobraBinder) Flags() *flag.FlagSet {
	return b.Cmd.Flags()
}
