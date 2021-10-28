package dns

import "fmt"

// Version is current version of this library.
<<<<<<< HEAD
<<<<<<< HEAD
var Version = v{1, 1, 48}
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
var Version = v{1, 1, 35}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
var Version = v{1, 1, 35}
=======
var Version = v{1, 1, 48}
>>>>>>> 4d7e5ad26 (update vendored files)

// v holds the version of this library.
type v struct {
	Major, Minor, Patch int
}

func (v v) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
