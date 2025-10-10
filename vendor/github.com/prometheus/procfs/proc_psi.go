// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package procfs

// The PSI / pressure interface is described at
//   https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/Documentation/accounting/psi.txt
// Each resource (cpu, io, memory, ...) is exposed as a single file.
// Each file may contain up to two lines, one for "some" pressure and one for "full" pressure.
// Each line contains several averages (over n seconds) and a total in µs.
//
// Example io pressure file:
// > some avg10=0.06 avg60=0.21 avg300=0.99 total=8537362
// > full avg10=0.00 avg60=0.13 avg300=0.96 total=8183134

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/prometheus/procfs/internal/util"
)

const lineFormat = "avg10=%f avg60=%f avg300=%f total=%d"

// PSILine is a single line of values as returned by `/proc/pressure/*`.
//
// The Avg entries are averages over n seconds, as a percentage.
// The Total line is in microseconds.
type PSILine struct {
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  uint64
}

// PSIStats represent pressure stall information from /proc/pressure/*
//
// "Some" indicates the share of time in which at least some tasks are stalled.
// "Full" indicates the share of time in which all non-idle tasks are stalled simultaneously.
type PSIStats struct {
	Some *PSILine
	Full *PSILine
}

// PSIStatsForResource reads pressure stall information for the specified
// resource from /proc/pressure/<resource>. At time of writing this can be
// either "cpu", "memory" or "io".
func (fs FS) PSIStatsForResource(resource string) (PSIStats, error) {
	data, err := util.ReadFileNoStat(fs.proc.Path(fmt.Sprintf("%s/%s", "pressure", resource)))
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %q: %w", resource, err)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %q: %w", resource, err)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %q: %w", resource, err)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %q: %w", resource, err)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return PSIStats{}, fmt.Errorf("psi_stats: unavailable for %s", resource)
=======
		return PSIStats{}, fmt.Errorf("%s: psi_stats: unavailable for %q: %w", ErrFileRead, resource, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
		return PSIStats{}, fmt.Errorf("%s: psi_stats: unavailable for %q: %w", ErrFileRead, resource, err)
=======
		return PSIStats{}, fmt.Errorf("%w: psi_stats: unavailable for %q: %w", ErrFileRead, resource, err)
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	}

	return parsePSIStats(bytes.NewReader(data))
}

// parsePSIStats parses the specified file for pressure stall information.
func parsePSIStats(r io.Reader) (PSIStats, error) {
	psiStats := PSIStats{}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l := scanner.Text()
		prefix := strings.Split(l, " ")[0]
		switch prefix {
		case "some":
			psi := PSILine{}
			_, err := fmt.Sscanf(l, fmt.Sprintf("some %s", lineFormat), &psi.Avg10, &psi.Avg60, &psi.Avg300, &psi.Total)
			if err != nil {
				return PSIStats{}, err
			}
			psiStats.Some = &psi
		case "full":
			psi := PSILine{}
			_, err := fmt.Sscanf(l, fmt.Sprintf("full %s", lineFormat), &psi.Avg10, &psi.Avg60, &psi.Avg300, &psi.Total)
			if err != nil {
				return PSIStats{}, err
			}
			psiStats.Full = &psi
		default:
			// If we encounter a line with an unknown prefix, ignore it and move on
			// Should new measurement types be added in the future we'll simply ignore them instead
			// of erroring on retrieval
			continue
		}
	}

	return psiStats, nil
}
