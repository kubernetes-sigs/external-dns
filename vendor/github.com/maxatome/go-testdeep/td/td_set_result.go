// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdSetResultKind uint8

const (
	itemsSetResult tdSetResultKind = iota
	keysSetResult
)

// Implements fmt.Stringer.
func (k tdSetResultKind) String() string {
	switch k {
	case itemsSetResult:
		return "item"
	case keysSetResult:
		return "key"
	default:
		return "?"
	}
}

type tdSetResult struct {
	types.TestDeepStamp
	Missing []reflect.Value
	Extra   []reflect.Value
	Kind    tdSetResultKind
	Sort    bool
}

func (r tdSetResult) IsEmpty() bool {
	return len(r.Missing) == 0 && len(r.Extra) == 0
}

func (r tdSetResult) Summary() ctxerr.ErrorSummary {
	var summary ctxerr.ErrorSummaryItems

	if len(r.Missing) > 0 {
		var missing string

		if len(r.Missing) > 1 {
			if r.Sort {
				sort.Stable(tdutil.SortableValues(r.Missing))
			}
			missing = fmt.Sprintf("Missing %d %ss", len(r.Missing), r.Kind)
		} else {
			missing = fmt.Sprintf("Missing %s", r.Kind)
		}

		summary = append(summary, ctxerr.ErrorSummaryItem{
			Label: missing,
			Value: util.ToString(r.Missing),
		})
	}

	if len(r.Extra) > 0 {
		var extra string

		if len(r.Extra) > 1 {
			if r.Sort {
				sort.Stable(tdutil.SortableValues(r.Extra))
			}
			extra = fmt.Sprintf("Extra %d %ss", len(r.Extra), r.Kind)
		} else {
			extra = fmt.Sprintf("Extra %s", r.Kind)
		}

		summary = append(summary, ctxerr.ErrorSummaryItem{
			Label: extra,
			Value: util.ToString(r.Extra),
		})
	}

	return summary
}
