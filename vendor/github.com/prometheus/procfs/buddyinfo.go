// Copyright 2017 The Prometheus Authors
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

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// A BuddyInfo is the details parsed from /proc/buddyinfo.
// The data is comprised of an array of free fragments of each size.
// The sizes are 2^n*PAGE_SIZE, where n is the array index.
type BuddyInfo struct {
	Node  string
	Zone  string
	Sizes []float64
}

// BuddyInfo reads the buddyinfo statistics from the specified `proc` filesystem.
func (fs FS) BuddyInfo() ([]BuddyInfo, error) {
	file, err := os.Open(fs.proc.Path("buddyinfo"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseBuddyInfo(file)
}

func parseBuddyInfo(r io.Reader) ([]BuddyInfo, error) {
	var (
		buddyInfo   = []BuddyInfo{}
		scanner     = bufio.NewScanner(r)
		bucketCount = -1
	)

	for scanner.Scan() {
		var err error
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) < 4 {
			return nil, fmt.Errorf("%w: Invalid number of fields, found: %v", ErrFileParse, parts)
		}

		node := strings.TrimSuffix(parts[1], ",")
		zone := strings.TrimSuffix(parts[3], ",")
		arraySize := len(parts[4:])

		if bucketCount == -1 {
			bucketCount = arraySize
		} else {
			if bucketCount != arraySize {
				return nil, fmt.Errorf("%w: mismatch in number of buddyinfo buckets, previous count %d, new count %d", ErrFileParse, bucketCount, arraySize)
			}
		}

		sizes := make([]float64, arraySize)
		for i := 0; i < arraySize; i++ {
			sizes[i], err = strconv.ParseFloat(parts[i+4], 64)
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
				return nil, fmt.Errorf("invalid value in buddyinfo: %w", err)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %w", err)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %w", err)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %w", err)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
				return nil, fmt.Errorf("invalid value in buddyinfo: %s", err)
=======
				return nil, fmt.Errorf("%s: Invalid valid in buddyinfo: %f: %w", ErrFileParse, sizes[i], err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
				return nil, fmt.Errorf("%s: Invalid valid in buddyinfo: %f: %w", ErrFileParse, sizes[i], err)
=======
				return nil, fmt.Errorf("%w: Invalid valid in buddyinfo: %f: %w", ErrFileParse, sizes[i], err)
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
			}
		}

		buddyInfo = append(buddyInfo, BuddyInfo{node, zone, sizes})
	}

	return buddyInfo, scanner.Err()
}
