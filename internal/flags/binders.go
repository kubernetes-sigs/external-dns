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
	"regexp"
	"strconv"
	"time"

	"github.com/alecthomas/kingpin/v2"
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
}

// KingpinBinder implements FlagBinder using github.com/alecthomas/kingpin/v2.
type KingpinBinder struct {
	App *kingpin.Application
}

// NewKingpinBinder creates a FlagBinder backed by a kingpin Application.
func NewKingpinBinder(app *kingpin.Application) *KingpinBinder {
	return &KingpinBinder{App: app}
}

func (b *KingpinBinder) StringVar(name, help, def string, target *string) {
	b.App.Flag(name, help).Default(def).StringVar(target)
}

func (b *KingpinBinder) BoolVar(name, help string, def bool, target *bool) {
	if def {
		b.App.Flag(name, help).Default("true").BoolVar(target)
	} else {
		b.App.Flag(name, help).Default("false").BoolVar(target)
	}
}

func (b *KingpinBinder) DurationVar(name, help string, def time.Duration, target *time.Duration) {
	b.App.Flag(name, help).Default(def.String()).DurationVar(target)
}

func (b *KingpinBinder) IntVar(name, help string, def int, target *int) {
	b.App.Flag(name, help).Default(strconv.Itoa(def)).IntVar(target)
}

func (b *KingpinBinder) Int64Var(name, help string, def int64, target *int64) {
	b.App.Flag(name, help).Default(strconv.FormatInt(def, 10)).Int64Var(target)
}

func (b *KingpinBinder) StringsVar(name, help string, def []string, target *[]string) {
	if len(def) > 0 {
		b.App.Flag(name, help).Default(def...).StringsVar(target)
		return
	}
	b.App.Flag(name, help).StringsVar(target)
}

func (b *KingpinBinder) EnumVar(name, help, def string, target *string, allowed ...string) {
	b.App.Flag(name, help).Default(def).EnumVar(target, allowed...)
}

func (b *KingpinBinder) StringsEnumVar(name, help string, def []string, target *[]string, allowed ...string) {
	if len(def) > 0 {
		b.App.Flag(name, help).Default(def...).EnumsVar(target, allowed...)
		return
	}
	b.App.Flag(name, help).EnumsVar(target, allowed...)
}

func (b *KingpinBinder) StringMapVar(name, help string, target *map[string]string) {
	b.App.Flag(name, help).StringMapVar(target)
}

func (b *KingpinBinder) RegexpVar(name, help string, def *regexp.Regexp, target **regexp.Regexp) {
	defStr := ""
	if def != nil {
		defStr = def.String()
	}
	b.App.Flag(name, help).Default(defStr).RegexpVar(target)
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
