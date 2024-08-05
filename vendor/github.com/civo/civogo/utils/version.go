package utils

import (
	"runtime/debug"
)

const packagePath = "github.com/civo/civogo"

var (
	// Version is the default version of the package
	Version = "dev"
)

// GetVersion init attempts to source the version from the build info injected
// at runtime and sets the DefaultUserAgent.
func GetVersion() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if ok {
		for _, dep := range buildInfo.Deps {
			if dep.Path == packagePath {
				if dep.Replace != nil {
					Version = dep.Replace.Version
				}
				Version = dep.Version
				break
			}
		}
	}

	return Version
}
