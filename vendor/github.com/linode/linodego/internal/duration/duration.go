package duration

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func UnmarshalTimeRemaining(m json.RawMessage) *int {
	jsonBytes, err := m.MarshalJSON()
	if err != nil {
		panic(jsonBytes)
	}

	if len(jsonBytes) == 4 && string(jsonBytes) == "null" {
		return nil
	}

	var timeStr string
	if err := json.Unmarshal(jsonBytes, &timeStr); err == nil && len(timeStr) > 0 {
		if dur, err := durationToSeconds(timeStr); err != nil {
			panic(err)
		} else {
			return &dur
		}
	} else {
		var intPtr int
		if err := json.Unmarshal(jsonBytes, &intPtr); err == nil {
			return &intPtr
		}
	}

	log.Println("[WARN] Unexpected unmarshalTimeRemaining value: ", jsonBytes)

	return nil
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds.
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds
=======
// durationToSeconds takes a hh:mm:ss string and returns the number of seconds.
>>>>>>> 4d7e5ad26 (update vendored files)
func durationToSeconds(s string) (int, error) {
	multipliers := [3]int{60 * 60, 60, 1}
	segs := strings.Split(s, ":")

	if len(segs) > len(multipliers) {
		return 0, fmt.Errorf("too many ':' separators in time duration: %s", s)
	}

	var d int

	l := len(segs)

	for i := 0; i < l; i++ {
		m, err := strconv.Atoi(segs[i])
		if err != nil {
			return 0, err
		}

		d += m * multipliers[i+len(multipliers)-l]
	}

	return d, nil
}
