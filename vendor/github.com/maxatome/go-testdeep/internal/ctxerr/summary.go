// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
	"strings"

	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/util"
)

// ErrorSummary is the interface used to render error summaries. See
// Error.Summary.
type ErrorSummary interface {
	AppendSummary(buf *strings.Builder, prefix string, colorized bool)
}

// ErrorSummaryItem implements the [ErrorSummary] interface and allows
// to render a labeled value.
//
// With explanation set:
//
//	Label: value
//	Explanation
//
// With an empty explanation:
//
//	Label: value
type ErrorSummaryItem struct {
	Label       string
	Value       string
	Explanation string
}

var _ ErrorSummary = ErrorSummaryItem{}

// AppendSummary implements the [ErrorSummary] interface.
func (s ErrorSummaryItem) AppendSummary(buf *strings.Builder, prefix string, colorized bool) {
	buf.WriteString(prefix)

	badOn, badOff := "", ""
	if colorized {
		color.Init()
		badOn, badOff = color.BadOn, color.BadOff
		buf.WriteString(color.BadOnBold)
	}
	buf.WriteString(s.Label)
	buf.WriteString(": ")

	util.IndentColorizeStringIn(buf, s.Value, prefix+strings.Repeat(" ", len(s.Label)+2), badOn, badOff)

	if s.Explanation != "" {
		buf.WriteByte('\n')
		buf.WriteString(prefix)
		util.IndentColorizeStringIn(buf, s.Explanation, prefix, badOn, badOff)
	}
}

// ErrorSummaryItems implements the [ErrorSummary] interface and
// allows to render summaries with several labeled values. For example:
//
//	Missing 6 items: the 6 items...
//	  Extra 2 items: the 2 items...
type ErrorSummaryItems []ErrorSummaryItem

var _ ErrorSummary = (ErrorSummaryItems)(nil)

// AppendSummary implements [ErrorSummary] interface.
func (s ErrorSummaryItems) AppendSummary(buf *strings.Builder, prefix string, colorized bool) {
	maxLen := 0
	for _, item := range s {
		if len(item.Label) > maxLen {
			maxLen = len(item.Label)
		}
	}

	for idx, item := range s {
		if idx > 0 {
			buf.WriteByte('\n')
		}
		if len(item.Label) < maxLen {
			item.Label = strings.Repeat(" ", maxLen-len(item.Label)) + item.Label
		}
		item.AppendSummary(buf, prefix, colorized)
	}
}

type errorSummaryString string

var _ ErrorSummary = errorSummaryString("")

func (s errorSummaryString) AppendSummary(buf *strings.Builder, prefix string, colorized bool) {
	badOn, badOff := "", ""
	if colorized {
		color.Init()
		badOn, badOff = color.BadOn, color.BadOff
	}

	buf.WriteString(prefix)
	util.IndentColorizeStringIn(buf, string(s), prefix, badOn, badOff)
}

// NewSummary returns an ErrorSummary composed by the simple string s.
func NewSummary(s string) ErrorSummary {
	return errorSummaryString(s)
}

// NewSummaryReason returns an [ErrorSummary] meaning that the value got
// failed for an (optional) reason.
//
// With a given reason "it is not nil", the generated summary is:
//
//	        value: the_got_value
//	it failed coz: it is not nil
//
// If reason is empty, the generated summary is:
//
//	  value: the_got_value
//	it failed but didn't say why
func NewSummaryReason(got any, reason string) ErrorSummary {
	if reason == "" {
		return ErrorSummaryItem{
			Label:       "  value", // keep 2 indent spaces
			Value:       util.ToString(got),
			Explanation: "it failed but didn't say why",
		}
	}

	return ErrorSummaryItems{
		{
			Label: "value",
			Value: util.ToString(got),
		},
		{
			Label: "it failed coz",
			Value: reason,
		},
	}
}
