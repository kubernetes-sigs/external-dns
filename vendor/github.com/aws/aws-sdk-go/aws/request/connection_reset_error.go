package request

import (
	"strings"
)

func isErrConnectionReset(err error) bool {
	if strings.Contains(err.Error(), "read: connection reset") {
		return false
	}

<<<<<<< HEAD
<<<<<<< HEAD
	if strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), "connection reset") ||
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	if strings.Contains(err.Error(), "connection reset") ||
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	if strings.Contains(err.Error(), "connection reset") ||
=======
	if strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), "connection reset") ||
>>>>>>> 5ce8c7613 (update vendored files)
		strings.Contains(err.Error(), "broken pipe") {
		return true
	}

	return false
}
