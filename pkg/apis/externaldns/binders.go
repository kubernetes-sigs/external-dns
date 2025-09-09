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
	"strconv"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/spf13/cobra"
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
	b.Cmd.Flags().BoolVar(target, name, def, help)
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
