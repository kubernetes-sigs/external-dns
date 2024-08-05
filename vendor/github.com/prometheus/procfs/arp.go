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

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Learned from include/uapi/linux/if_arp.h.
const (
	// completed entry (ha valid).
	ATFComplete = 0x02
	// permanent entry.
	ATFPermanent = 0x04
	// Publish entry.
	ATFPublish = 0x08
	// Has requested trailers.
	ATFUseTrailers = 0x10
	// Obsoleted: Want to use a netmask (only for proxy entries).
	ATFNetmask = 0x20
	// Don't answer this addresses.
	ATFDontPublish = 0x40
)

// ARPEntry contains a single row of the columnar data represented in
// /proc/net/arp.
type ARPEntry struct {
	// IP address
	IPAddr net.IP
	// MAC address
	HWAddr net.HardwareAddr
	// Name of the device
	Device string
	// Flags
	Flags byte
}

// GatherARPEntries retrieves all the ARP entries, parse the relevant columns,
// and then return a slice of ARPEntry's.
func (fs FS) GatherARPEntries() ([]ARPEntry, error) {
	data, err := os.ReadFile(fs.proc.Path("net/arp"))
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		return nil, fmt.Errorf("error reading arp %q: %w", fs.proc.Path("net/arp"), err)
	}

	return parseARPEntries(data)
}

func parseARPEntries(data []byte) ([]ARPEntry, error) {
	lines := strings.Split(string(data), "\n")
	entries := make([]ARPEntry, 0)
	var err error
	const (
		expectedDataWidth   = 6
		expectedHeaderWidth = 9
	)
	for _, line := range lines {
		columns := strings.Fields(line)
		width := len(columns)

		if width == expectedHeaderWidth || width == 0 {
			continue
		} else if width == expectedDataWidth {
			entry, err := parseARPEntry(columns)
			if err != nil {
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %w", err)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
||||||| parent of 5ce8c7613 (update vendored files)
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
=======
		return nil, fmt.Errorf("error reading arp %q: %w", fs.proc.Path("net/arp"), err)
>>>>>>> 5ce8c7613 (update vendored files)
	}

	return parseARPEntries(data)
}

func parseARPEntries(data []byte) ([]ARPEntry, error) {
	lines := strings.Split(string(data), "\n")
	entries := make([]ARPEntry, 0)
	var err error
	const (
		expectedDataWidth   = 6
		expectedHeaderWidth = 9
	)
	for _, line := range lines {
		columns := strings.Fields(line)
		width := len(columns)

		if width == expectedHeaderWidth || width == 0 {
			continue
		} else if width == expectedDataWidth {
			entry, err := parseARPEntry(columns)
			if err != nil {
<<<<<<< HEAD
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
=======
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %w", err)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
||||||| parent of 6b7ce455e (update vendored files)
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
=======
		return nil, fmt.Errorf("error reading arp %q: %w", fs.proc.Path("net/arp"), err)
>>>>>>> 6b7ce455e (update vendored files)
	}

	return parseARPEntries(data)
}

func parseARPEntries(data []byte) ([]ARPEntry, error) {
	lines := strings.Split(string(data), "\n")
	entries := make([]ARPEntry, 0)
	var err error
	const (
		expectedDataWidth   = 6
		expectedHeaderWidth = 9
	)
	for _, line := range lines {
		columns := strings.Fields(line)
		width := len(columns)

		if width == expectedHeaderWidth || width == 0 {
			continue
		} else if width == expectedDataWidth {
			entry, err := parseARPEntry(columns)
			if err != nil {
<<<<<<< HEAD
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
=======
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %w", err)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
||||||| parent of 4d7e5ad26 (update vendored files)
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
=======
		return nil, fmt.Errorf("error reading arp %q: %w", fs.proc.Path("net/arp"), err)
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	return parseARPEntries(data)
}

func parseARPEntries(data []byte) ([]ARPEntry, error) {
	lines := strings.Split(string(data), "\n")
	entries := make([]ARPEntry, 0)
	var err error
	const (
		expectedDataWidth   = 6
		expectedHeaderWidth = 9
	)
	for _, line := range lines {
		columns := strings.Fields(line)
		width := len(columns)

		if width == expectedHeaderWidth || width == 0 {
			continue
		} else if width == expectedDataWidth {
			entry, err := parseARPEntry(columns)
			if err != nil {
<<<<<<< HEAD
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
=======
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %w", err)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return nil, fmt.Errorf("error reading arp %s: %s", fs.proc.Path("net/arp"), err)
=======
		return nil, fmt.Errorf("%s: error reading arp %s: %w", ErrFileRead, fs.proc.Path("net/arp"), err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	return parseARPEntries(data)
}

func parseARPEntries(data []byte) ([]ARPEntry, error) {
	lines := strings.Split(string(data), "\n")
	entries := make([]ARPEntry, 0)
	var err error
	const (
		expectedDataWidth   = 6
		expectedHeaderWidth = 9
	)
	for _, line := range lines {
		columns := strings.Fields(line)
		width := len(columns)

		if width == expectedHeaderWidth || width == 0 {
			continue
		} else if width == expectedDataWidth {
			entry, err := parseARPEntry(columns)
			if err != nil {
<<<<<<< HEAD
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
				return []ARPEntry{}, fmt.Errorf("failed to parse ARP entry: %s", err)
=======
				return []ARPEntry{}, fmt.Errorf("%s: Failed to parse ARP entry: %v: %w", ErrFileParse, entry, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
			}
			entries = append(entries, entry)
		} else {
			return []ARPEntry{}, fmt.Errorf("%s: %d columns found, but expected %d: %w", ErrFileParse, width, expectedDataWidth, err)
		}

	}

	return entries, err
}

func parseARPEntry(columns []string) (ARPEntry, error) {
	entry := ARPEntry{Device: columns[5]}
	ip := net.ParseIP(columns[0])
	entry.IPAddr = ip

	if mac, err := net.ParseMAC(columns[3]); err == nil {
		entry.HWAddr = mac
	} else {
		return ARPEntry{}, err
	}

	if flags, err := strconv.ParseUint(columns[2], 0, 8); err == nil {
		entry.Flags = byte(flags)
	} else {
		return ARPEntry{}, err
	}

	return entry, nil
}

// IsComplete returns true if ARP entry is marked with complete flag.
func (entry *ARPEntry) IsComplete() bool {
	return entry.Flags&ATFComplete != 0
}
