// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// license that can be found in the LICENSE file.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// license that can be found in the LICENSE.md file.
=======
// license that can be found in the LICENSE file.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

// +build cmp_debug

package diff

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// The algorithm can be seen running in real-time by enabling debugging:
//	go test -tags=cmp_debug -v
//
// Example output:
//	=== RUN   TestDifference/#34
//	┌───────────────────────────────┐
//	│ \ · · · · · · · · · · · · · · │
//	│ · # · · · · · · · · · · · · · │
//	│ · \ · · · · · · · · · · · · · │
//	│ · · \ · · · · · · · · · · · · │
//	│ · · · X # · · · · · · · · · · │
//	│ · · · # \ · · · · · · · · · · │
//	│ · · · · · # # · · · · · · · · │
//	│ · · · · · # \ · · · · · · · · │
//	│ · · · · · · · \ · · · · · · · │
//	│ · · · · · · · · \ · · · · · · │
//	│ · · · · · · · · · \ · · · · · │
//	│ · · · · · · · · · · \ · · # · │
//	│ · · · · · · · · · · · \ # # · │
//	│ · · · · · · · · · · · # # # · │
//	│ · · · · · · · · · · # # # # · │
//	│ · · · · · · · · · # # # # # · │
//	│ · · · · · · · · · · · · · · \ │
//	└───────────────────────────────┘
//	[.Y..M.XY......YXYXY.|]
//
// The grid represents the edit-graph where the horizontal axis represents
// list X and the vertical axis represents list Y. The start of the two lists
// is the top-left, while the ends are the bottom-right. The '·' represents
// an unexplored node in the graph. The '\' indicates that the two symbols
// from list X and Y are equal. The 'X' indicates that two symbols are similar
// (but not exactly equal) to each other. The '#' indicates that the two symbols
// are different (and not similar). The algorithm traverses this graph trying to
// make the paths starting in the top-left and the bottom-right connect.
//
// The series of '.', 'X', 'Y', and 'M' characters at the bottom represents
// the currently established path from the forward and reverse searches,
// separated by a '|' character.

const (
	updateDelay  = 100 * time.Millisecond
	finishDelay  = 500 * time.Millisecond
	ansiTerminal = true // ANSI escape codes used to move terminal cursor
)

var debug debugger

type debugger struct {
	sync.Mutex
	p1, p2           EditScript
	fwdPath, revPath *EditScript
	grid             []byte
	lines            int
}

func (dbg *debugger) Begin(nx, ny int, f EqualFunc, p1, p2 *EditScript) EqualFunc {
	dbg.Lock()
	dbg.fwdPath, dbg.revPath = p1, p2
	top := "┌─" + strings.Repeat("──", nx) + "┐\n"
	row := "│ " + strings.Repeat("· ", nx) + "│\n"
	btm := "└─" + strings.Repeat("──", nx) + "┘\n"
	dbg.grid = []byte(top + strings.Repeat(row, ny) + btm)
	dbg.lines = strings.Count(dbg.String(), "\n")
	fmt.Print(dbg)

	// Wrap the EqualFunc so that we can intercept each result.
	return func(ix, iy int) (r Result) {
		cell := dbg.grid[len(top)+iy*len(row):][len("│ ")+len("· ")*ix:][:len("·")]
		for i := range cell {
			cell[i] = 0 // Zero out the multiple bytes of UTF-8 middle-dot
		}
		switch r = f(ix, iy); {
		case r.Equal():
			cell[0] = '\\'
		case r.Similar():
			cell[0] = 'X'
		default:
			cell[0] = '#'
		}
		return
	}
}

func (dbg *debugger) Update() {
	dbg.print(updateDelay)
}

func (dbg *debugger) Finish() {
	dbg.print(finishDelay)
	dbg.Unlock()
}

func (dbg *debugger) String() string {
	dbg.p1, dbg.p2 = *dbg.fwdPath, dbg.p2[:0]
	for i := len(*dbg.revPath) - 1; i >= 0; i-- {
		dbg.p2 = append(dbg.p2, (*dbg.revPath)[i])
	}
	return fmt.Sprintf("%s[%v|%v]\n\n", dbg.grid, dbg.p1, dbg.p2)
}

func (dbg *debugger) print(d time.Duration) {
	if ansiTerminal {
		fmt.Printf("\x1b[%dA", dbg.lines) // Reset terminal cursor
	}
	fmt.Print(dbg)
	time.Sleep(d)
}
