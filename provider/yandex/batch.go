package yandex

import (
	dnsInt "github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

type upsertBatch struct {
	ZoneID   string
	ZoneName string
	Deletes  []*dnsInt.RecordSet
	Updates  []*dnsInt.RecordSet
	Creates  []*dnsInt.RecordSet
}

type upsertBatchMap map[string]*upsertBatch

func (m upsertBatchMap) GetOrCreate(zoneID, zoneName string) *upsertBatch {
	batch, ok := m[zoneID]

	if !ok {
		batch = &upsertBatch{
			ZoneID:   zoneID,
			ZoneName: zoneName,
			Creates:  make([]*dnsInt.RecordSet, 0),
			Deletes:  make([]*dnsInt.RecordSet, 0),
			Updates:  make([]*dnsInt.RecordSet, 0),
		}
		m[zoneID] = batch
	}

	return batch
}

func (m upsertBatchMap) ApplyChanges(mapper provider.ZoneIDName, changes []*endpoint.Endpoint, handler func(*upsertBatch, *dnsInt.RecordSet)) {
	for _, change := range changes {
		zoneID, zoneName := mapper.FindZone(change.DNSName)
		if zoneID == "" || zoneName == "" {
			continue
		}

		batch := m.GetOrCreate(zoneID, zoneName)
		handler(batch, toRecordSet(change))
	}
}
