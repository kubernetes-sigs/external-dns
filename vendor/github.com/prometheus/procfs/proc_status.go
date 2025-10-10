// Copyright 2018 The Prometheus Authors
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
	"bytes"
	"math/bits"
	"sort"
	"strconv"
	"strings"

	"github.com/prometheus/procfs/internal/util"
)

// ProcStatus provides status information about the process,
// read from /proc/[pid]/status.
type ProcStatus struct {
	// The process ID.
	PID int
	// The process name.
	Name string

	// Thread group ID.
	TGID int
	// List of Pid namespace.
	NSpids []uint64

	// Peak virtual memory size.
	VmPeak uint64 // nolint:revive
	// Virtual memory size.
	VmSize uint64 // nolint:revive
	// Locked memory size.
	VmLck uint64 // nolint:revive
	// Pinned memory size.
	VmPin uint64 // nolint:revive
	// Peak resident set size.
	VmHWM uint64 // nolint:revive
	// Resident set size (sum of RssAnnon RssFile and RssShmem).
	VmRSS uint64 // nolint:revive
	// Size of resident anonymous memory.
	RssAnon uint64 // nolint:revive
	// Size of resident file mappings.
	RssFile uint64 // nolint:revive
	// Size of resident shared memory.
	RssShmem uint64 // nolint:revive
	// Size of data segments.
	VmData uint64 // nolint:revive
	// Size of stack segments.
	VmStk uint64 // nolint:revive
	// Size of text segments.
	VmExe uint64 // nolint:revive
	// Shared library code size.
	VmLib uint64 // nolint:revive
	// Page table entries size.
	VmPTE uint64 // nolint:revive
	// Size of second-level page tables.
	VmPMD uint64 // nolint:revive
	// Swapped-out virtual memory size by anonymous private.
	VmSwap uint64 // nolint:revive
	// Size of hugetlb memory portions
	HugetlbPages uint64

	// Number of voluntary context switches.
	VoluntaryCtxtSwitches uint64
	// Number of involuntary context switches.
	NonVoluntaryCtxtSwitches uint64

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs)
	UIDs [4]uint64
	// GIDs of the process (Real, effective, saved set, and filesystem GIDs)
<<<<<<< HEAD
	GIDs [4]string
}

// NewStatus returns the current status information of the process.
func (p Proc) NewStatus() (ProcStatus, error) {
	data, err := util.ReadFileNoStat(p.path("status"))
	if err != nil {
		return ProcStatus{}, err
	}

	s := ProcStatus{PID: p.PID}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if !bytes.Contains([]byte(line), []byte(":")) {
			continue
		}

		kv := strings.SplitN(line, ":", 2)

		// removes spaces
		k := string(strings.TrimSpace(kv[0]))
		v := string(strings.TrimSpace(kv[1]))
		// removes "kB"
		v = string(bytes.Trim([]byte(v), " kB"))

		// value to int when possible
		// we can skip error check here, 'cause vKBytes is not used when value is a string
		vKBytes, _ := strconv.ParseUint(v, 10, 64)
		// convert kB to B
		vBytes := vKBytes * 1024

		s.fillStatus(k, v, vKBytes, vBytes)
	}

	return s, nil
}

func (s *ProcStatus) fillStatus(k string, vString string, vUint uint64, vUintBytes uint64) {
	switch k {
	case "Tgid":
		s.TGID = int(vUint)
	case "Name":
		s.Name = vString
	case "Uid":
		copy(s.UIDs[:], strings.Split(vString, "\t"))
	case "Gid":
		copy(s.GIDs[:], strings.Split(vString, "\t"))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
||||||| parent of 5ce8c7613 (update vendored files)
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs)
>>>>>>> 5ce8c7613 (update vendored files)
	UIDs [4]string
	// GIDs of the process (Real, effective, saved set, and filesystem GIDs)
	GIDs [4]string
}

// NewStatus returns the current status information of the process.
func (p Proc) NewStatus() (ProcStatus, error) {
	data, err := util.ReadFileNoStat(p.path("status"))
	if err != nil {
		return ProcStatus{}, err
	}

	s := ProcStatus{PID: p.PID}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if !bytes.Contains([]byte(line), []byte(":")) {
			continue
		}

		kv := strings.SplitN(line, ":", 2)

		// removes spaces
		k := string(strings.TrimSpace(kv[0]))
		v := string(strings.TrimSpace(kv[1]))
		// removes "kB"
		v = string(bytes.Trim([]byte(v), " kB"))

		// value to int when possible
		// we can skip error check here, 'cause vKBytes is not used when value is a string
		vKBytes, _ := strconv.ParseUint(v, 10, 64)
		// convert kB to B
		vBytes := vKBytes * 1024

		s.fillStatus(k, v, vKBytes, vBytes)
	}

	return s, nil
}

func (s *ProcStatus) fillStatus(k string, vString string, vUint uint64, vUintBytes uint64) {
	switch k {
	case "Tgid":
		s.TGID = int(vUint)
	case "Name":
		s.Name = vString
	case "Uid":
		copy(s.UIDs[:], strings.Split(vString, "\t"))
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	case "Gid":
		copy(s.GIDs[:], strings.Split(vString, "\t"))
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
||||||| parent of 6b7ce455e (update vendored files)
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs)
>>>>>>> 6b7ce455e (update vendored files)
	UIDs [4]string
	// GIDs of the process (Real, effective, saved set, and filesystem GIDs)
	GIDs [4]string
}

// NewStatus returns the current status information of the process.
func (p Proc) NewStatus() (ProcStatus, error) {
	data, err := util.ReadFileNoStat(p.path("status"))
	if err != nil {
		return ProcStatus{}, err
	}

	s := ProcStatus{PID: p.PID}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if !bytes.Contains([]byte(line), []byte(":")) {
			continue
		}

		kv := strings.SplitN(line, ":", 2)

		// removes spaces
		k := string(strings.TrimSpace(kv[0]))
		v := string(strings.TrimSpace(kv[1]))
		// removes "kB"
		v = string(bytes.Trim([]byte(v), " kB"))

		// value to int when possible
		// we can skip error check here, 'cause vKBytes is not used when value is a string
		vKBytes, _ := strconv.ParseUint(v, 10, 64)
		// convert kB to B
		vBytes := vKBytes * 1024

		s.fillStatus(k, v, vKBytes, vBytes)
	}

	return s, nil
}

func (s *ProcStatus) fillStatus(k string, vString string, vUint uint64, vUintBytes uint64) {
	switch k {
	case "Tgid":
		s.TGID = int(vUint)
	case "Name":
		s.Name = vString
	case "Uid":
		copy(s.UIDs[:], strings.Split(vString, "\t"))
<<<<<<< HEAD
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	case "Gid":
		copy(s.GIDs[:], strings.Split(vString, "\t"))
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
||||||| parent of 4d7e5ad26 (update vendored files)
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs)
>>>>>>> 4d7e5ad26 (update vendored files)
	UIDs [4]string
	// GIDs of the process (Real, effective, saved set, and filesystem GIDs)
	GIDs [4]string
}

// NewStatus returns the current status information of the process.
func (p Proc) NewStatus() (ProcStatus, error) {
	data, err := util.ReadFileNoStat(p.path("status"))
	if err != nil {
		return ProcStatus{}, err
	}

	s := ProcStatus{PID: p.PID}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if !bytes.Contains([]byte(line), []byte(":")) {
			continue
		}

		kv := strings.SplitN(line, ":", 2)

		// removes spaces
		k := string(strings.TrimSpace(kv[0]))
		v := string(strings.TrimSpace(kv[1]))
		// removes "kB"
		v = string(bytes.Trim([]byte(v), " kB"))

		// value to int when possible
		// we can skip error check here, 'cause vKBytes is not used when value is a string
		vKBytes, _ := strconv.ParseUint(v, 10, 64)
		// convert kB to B
		vBytes := vKBytes * 1024

		s.fillStatus(k, v, vKBytes, vBytes)
	}

	return s, nil
}

func (s *ProcStatus) fillStatus(k string, vString string, vUint uint64, vUintBytes uint64) {
	switch k {
	case "Tgid":
		s.TGID = int(vUint)
	case "Name":
		s.Name = vString
	case "Uid":
		copy(s.UIDs[:], strings.Split(vString, "\t"))
<<<<<<< HEAD
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	case "Gid":
		copy(s.GIDs[:], strings.Split(vString, "\t"))
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs (GIDs))
=======
	// UIDs of the process (Real, effective, saved set, and filesystem UIDs)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	UIDs [4]string
	// GIDs of the process (Real, effective, saved set, and filesystem GIDs)
	GIDs [4]string
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	GIDs [4]string
=======
	GIDs [4]uint64
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)

	// CpusAllowedList: List of cpu cores processes are allowed to run on.
	CpusAllowedList []uint64
}

// NewStatus returns the current status information of the process.
func (p Proc) NewStatus() (ProcStatus, error) {
	data, err := util.ReadFileNoStat(p.path("status"))
	if err != nil {
		return ProcStatus{}, err
	}

	s := ProcStatus{PID: p.PID}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if !bytes.Contains([]byte(line), []byte(":")) {
			continue
		}

		kv := strings.SplitN(line, ":", 2)

		// removes spaces
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		// removes "kB"
		v = strings.TrimSuffix(v, " kB")

		// value to int when possible
		// we can skip error check here, 'cause vKBytes is not used when value is a string
		vKBytes, _ := strconv.ParseUint(v, 10, 64)
		// convert kB to B
		vBytes := vKBytes * 1024

		err = s.fillStatus(k, v, vKBytes, vBytes)
		if err != nil {
			return ProcStatus{}, err
		}
	}

	return s, nil
}

func (s *ProcStatus) fillStatus(k string, vString string, vUint uint64, vUintBytes uint64) error {
	switch k {
	case "Tgid":
		s.TGID = int(vUint)
	case "Name":
		s.Name = vString
	case "Uid":
<<<<<<< HEAD
		copy(s.UIDs[:], strings.Split(vString, "\t"))
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
		copy(s.UIDs[:], strings.Split(vString, "\t"))
=======
		var err error
		for i, v := range strings.Split(vString, "\t") {
			s.UIDs[i], err = strconv.ParseUint(v, 10, bits.UintSize)
			if err != nil {
				return err
			}
		}
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	case "Gid":
		var err error
		for i, v := range strings.Split(vString, "\t") {
			s.GIDs[i], err = strconv.ParseUint(v, 10, bits.UintSize)
			if err != nil {
				return err
			}
		}
	case "NSpid":
		s.NSpids = calcNSPidsList(vString)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	case "VmPeak":
		s.VmPeak = vUintBytes
	case "VmSize":
		s.VmSize = vUintBytes
	case "VmLck":
		s.VmLck = vUintBytes
	case "VmPin":
		s.VmPin = vUintBytes
	case "VmHWM":
		s.VmHWM = vUintBytes
	case "VmRSS":
		s.VmRSS = vUintBytes
	case "RssAnon":
		s.RssAnon = vUintBytes
	case "RssFile":
		s.RssFile = vUintBytes
	case "RssShmem":
		s.RssShmem = vUintBytes
	case "VmData":
		s.VmData = vUintBytes
	case "VmStk":
		s.VmStk = vUintBytes
	case "VmExe":
		s.VmExe = vUintBytes
	case "VmLib":
		s.VmLib = vUintBytes
	case "VmPTE":
		s.VmPTE = vUintBytes
	case "VmPMD":
		s.VmPMD = vUintBytes
	case "VmSwap":
		s.VmSwap = vUintBytes
	case "HugetlbPages":
		s.HugetlbPages = vUintBytes
	case "voluntary_ctxt_switches":
		s.VoluntaryCtxtSwitches = vUint
	case "nonvoluntary_ctxt_switches":
		s.NonVoluntaryCtxtSwitches = vUint
	case "Cpus_allowed_list":
		s.CpusAllowedList = calcCpusAllowedList(vString)
	}

	return nil
}

// TotalCtxtSwitches returns the total context switch.
func (s ProcStatus) TotalCtxtSwitches() uint64 {
	return s.VoluntaryCtxtSwitches + s.NonVoluntaryCtxtSwitches
}

func calcCpusAllowedList(cpuString string) []uint64 {
	s := strings.Split(cpuString, ",")

	var g []uint64

	for _, cpu := range s {
		// parse cpu ranges, example: 1-3=[1,2,3]
		if l := strings.Split(strings.TrimSpace(cpu), "-"); len(l) > 1 {
			startCPU, _ := strconv.ParseUint(l[0], 10, 64)
			endCPU, _ := strconv.ParseUint(l[1], 10, 64)

			for i := startCPU; i <= endCPU; i++ {
				g = append(g, i)
			}
		} else if len(l) == 1 {
			cpu, _ := strconv.ParseUint(l[0], 10, 64)
			g = append(g, cpu)
		}

	}

	sort.Slice(g, func(i, j int) bool { return g[i] < g[j] })
	return g
}

func calcNSPidsList(nspidsString string) []uint64 {
	s := strings.Split(nspidsString, " ")
	var nspids []uint64

	for _, nspid := range s {
		nspid, _ := strconv.ParseUint(nspid, 10, 64)
		if nspid == 0 {
			continue
		}
		nspids = append(nspids, nspid)
	}

	return nspids
}
