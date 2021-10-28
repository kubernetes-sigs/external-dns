// Copyright (c) 2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

// +build go1.15

package td

import "strconv"

// strconv.ParseComplex is only available from go 1.15.
func init() {
	parseComplex = strconv.ParseComplex
}
