// Copyright 2020 The Prometheus Authors
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
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

<<<<<<< HEAD
// NetStat contains statistics for all the counters from one file
type NetStat struct {
	Filename string
	Stats    map[string][]uint64
}

// NetStat retrieves stats from /proc/net/stat/
func (fs FS) NetStat() ([]NetStat, error) {
	statFiles, err := filepath.Glob(fs.proc.Path("net/stat/*"))
	if err != nil {
		return nil, err
	}

	var netStatsTotal []NetStat

	for _, filePath := range statFiles {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		netStatFile := NetStat{
			Filename: filepath.Base(filePath),
			Stats:    make(map[string][]uint64),
		}
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		// First string is always a header for stats
		var headers []string
		headers = append(headers, strings.Fields(scanner.Text())...)

		// Other strings represent per-CPU counters
		for scanner.Scan() {
			for num, counter := range strings.Fields(scanner.Text()) {
				value, err := strconv.ParseUint(counter, 16, 32)
				if err != nil {
					return nil, err
				}
				netStatFile.Stats[headers[num]] = append(netStatFile.Stats[headers[num]], value)
			}
		}
		netStatsTotal = append(netStatsTotal, netStatFile)
	}
	return netStatsTotal, nil
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
// NetStat contains statistics for all the counters from one file.
type NetStat struct {
	Stats    map[string][]uint64
	Filename string
}

// NetStat retrieves stats from `/proc/net/stat/`.
func (fs FS) NetStat() ([]NetStat, error) {
	statFiles, err := filepath.Glob(fs.proc.Path("net/stat/*"))
	if err != nil {
		return nil, err
	}

	var netStatsTotal []NetStat

	for _, filePath := range statFiles {
		procNetstat, err := parseNetstat(filePath)
		if err != nil {
			return nil, err
		}
		procNetstat.Filename = filepath.Base(filePath)

		netStatsTotal = append(netStatsTotal, procNetstat)
	}
	return netStatsTotal, nil
}

// parseNetstat parses the metrics from `/proc/net/stat/` file
// and returns a NetStat structure.
func parseNetstat(filePath string) (NetStat, error) {
	netStat := NetStat{
		Stats: make(map[string][]uint64),
	}
	file, err := os.Open(filePath)
	if err != nil {
		return netStat, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// First string is always a header for stats
	var headers []string
	headers = append(headers, strings.Fields(scanner.Text())...)

	// Other strings represent per-CPU counters
	for scanner.Scan() {
		for num, counter := range strings.Fields(scanner.Text()) {
			value, err := strconv.ParseUint(counter, 16, 64)
			if err != nil {
				return NetStat{}, err
			}
			netStat.Stats[headers[num]] = append(netStat.Stats[headers[num]], value)
		}
	}

	return netStat, nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
