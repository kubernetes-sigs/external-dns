// Copyright (c) 2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of 6b7ce455e (update vendored files)
=======
//go:build go1.15
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
//go:build go1.15
>>>>>>> 4d7e5ad26 (update vendored files)
// +build go1.15

package td

import "strconv"

// strconv.ParseComplex is only available from go 1.15.
func init() {
	parseComplex = strconv.ParseComplex
}
