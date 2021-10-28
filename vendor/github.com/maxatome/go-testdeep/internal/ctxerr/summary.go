// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
	"bytes"
	"strings"

	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/util"
)

// ErrorSummary is the interface used to render error summaries. See
// Error.Summary.
type ErrorSummary interface {
	AppendSummary(buf *bytes.Buffer, prefix string)
}

// ErrorSummaryItem implements the ErrorSummary interface and allows
// to render a labeled value.
//
// With explanation set:
//
//   Label: value
//   Explanation
//
// With an empty explantion:
//
//   Label: value
type ErrorSummaryItem struct {
	Label       string
	Value       string
	Explanation string
}

var _ ErrorSummary = ErrorSummaryItem{}

// AppendSummary implements the ErrorSummary interface.
func (s ErrorSummaryItem) AppendSummary(buf *bytes.Buffer, prefix string) {
	color.Init()

	buf.WriteString(prefix)
	buf.WriteString(color.BadOnBold)
	buf.WriteString(s.Label)
	buf.WriteString(": ")

	buf.WriteString(color.BadOn)
	util.IndentStringIn(buf, s.Value, prefix+strings.Repeat(" ", len(s.Label)+2), color.BadOn, color.BadOff)

	if s.Explanation != "" {
		buf.WriteString(color.BadOff)
		buf.WriteByte('\n')
		buf.WriteString(prefix)
		buf.WriteString(color.BadOn)
		util.IndentStringIn(buf, s.Explanation, prefix, color.BadOn, color.BadOff)
	}

	buf.WriteString(color.BadOff)
}

// ErrorSummaryItems implements the ErrorSummary interface and allows
// to render summaries with several labeled values. For example:
//
//   Missing 6 items: the 6 items...
//     Extra 2 items: the 2 items...
type ErrorSummaryItems []ErrorSummaryItem

var _ ErrorSummary = (ErrorSummaryItems)(nil)

// AppendSummary implements ErrorSummary interface.
func (s ErrorSummaryItems) AppendSummary(buf *bytes.Buffer, prefix string) {
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
		item.AppendSummary(buf, prefix)
	}
}

type errorSummaryString string

var _ ErrorSummary = errorSummaryString("")

func (s errorSummaryString) AppendSummary(buf *bytes.Buffer, prefix string) {
	color.Init()

	buf.WriteString(prefix)
	buf.WriteString(color.BadOn)
	util.IndentStringIn(buf, string(s), prefix, color.BadOn, color.BadOff)
	buf.WriteString(color.BadOff)
}

// NewSummary returns an ErrorSummary composed by the simple string s.
func NewSummary(s string) ErrorSummary {
	return errorSummaryString(s)
}

// NewSummaryReason returns an ErrorSummary meaning that the value got
// failed for an (optional) reason.
//
// With a given reason "it is not nil", the generated summary will be:
//
//           value: the_got_value
//   it failed coz: it is not nil
//
// If reason is empty, the generated summary will be:
//
//     value: the_got_value
//   it failed but didn't say why
func NewSummaryReason(got interface{}, reason string) ErrorSummary {
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
