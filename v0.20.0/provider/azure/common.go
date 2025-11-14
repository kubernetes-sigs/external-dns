/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//nolint:staticcheck // Required due to the current dependency on a deprecated version of azure-sdk-for-go
package azure

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	dns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"
	privatedns "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/privatedns/armprivatedns"
)

// Helper function (shared with test code)
func parseMxTarget[T dns.MxRecord | privatedns.MxRecord](mxTarget string) (T, error) {
	targetParts := strings.SplitN(mxTarget, " ", 2)
	if len(targetParts) != 2 {
		return T{}, fmt.Errorf("mx target needs to be of form '10 example.com'")
	}

	preferenceRaw, exchange := targetParts[0], targetParts[1]
	preference, err := strconv.ParseInt(preferenceRaw, 10, 32)
	if err != nil {
		return T{}, fmt.Errorf("invalid preference specified")
	}

	return T{
		Preference: to.Ptr(int32(preference)),
		Exchange:   to.Ptr(exchange),
	}, nil
}
