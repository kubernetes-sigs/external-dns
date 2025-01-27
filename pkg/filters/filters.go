package filters

import (
	"sigs.k8s.io/external-dns/pkg/filters/zonetagfilter"
)

var (
	NewZoneTagFilter = zonetagfilter.NewZoneTagFilter
)

type ZoneTagFilter = zonetagfilter.ZoneTagFilter
