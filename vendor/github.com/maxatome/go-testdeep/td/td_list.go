// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"reflect"

	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdList struct {
	baseOKNil
	items []reflect.Value
}

func newList(items ...interface{}) tdList {
	return tdList{
		baseOKNil: newBaseOKNil(4),
		items:     flat.Values(items),
	}
}

func (l *tdList) String() string {
	return util.SliceToBuffer(bytes.NewBufferString(l.GetLocation().Func), l.items).
		String()
}
