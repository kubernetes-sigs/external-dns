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
		Preference: new(int32(preference)),
		Exchange:   new(exchange),
	}, nil
}

// Helper function (shared with test code)
func parseSrvTarget[T dns.SrvRecord | privatedns.SrvRecord](srvTarget string) (T, error) {
	targetParts := strings.SplitN(srvTarget, " ", 4)
	if len(targetParts) != 4 {
		return T{}, fmt.Errorf("srv target needs to be of form '10 5 5060 example.com.'")
	}

	priority, err := strconv.ParseInt(targetParts[0], 10, 32)
	if err != nil {
		return T{}, fmt.Errorf("invalid srv priority specified")
	}
	weight, err := strconv.ParseInt(targetParts[1], 10, 32)
	if err != nil {
		return T{}, fmt.Errorf("invalid srv weight specified")
	}
	port, err := strconv.ParseInt(targetParts[2], 10, 32)
	if err != nil {
		return T{}, fmt.Errorf("invalid srv port specified")
	}
	target := targetParts[3]

	return T{
		Priority: new(int32(priority)),
		Weight:   new(int32(weight)),
		Port:     new(int32(port)),
		Target:   new(target),
	}, nil
}
