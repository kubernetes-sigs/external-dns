/*
Copyright 2026 The Kubernetes Authors.

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

package plan

import (
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

type MetaKey string

// Metadata holds arbitrary per-endpoint binary metadata used by the Plan.
//
// It is stored in the Plan.Metadata field and consulted during planning
// (for example to remember registry type or TTL expiry of an endpoint).
//
// Keys should be generated with the helpers in this package
// (e.g. genMetakeyEndpointTTL). Do not use the predefined MetaKey*
// constants as raw map keys — they are only the suffix part of a full key
// and colliding with them will corrupt Plan behavior.
type Metadata map[MetaKey][]byte

const (
	MetaKeyRegistryType = "registry"
)

func isRegistryNoOp(meta Metadata) bool {
	v, ok := meta[MetaKeyRegistryType]
	if !ok {
		return false
	}

	return string(v) == externaldns.RegistryNoop
}

func isRegistryAWSSD(meta Metadata) bool {
	v, ok := meta[MetaKeyRegistryType]
	if !ok {
		return false
	}

	return string(v) == externaldns.RegistryAWSSD
}
