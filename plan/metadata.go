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
// and colliding with them will corrupt Plan behaviour.
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
