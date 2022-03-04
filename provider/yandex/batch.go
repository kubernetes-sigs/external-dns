/*
Copyright 2022 The Kubernetes Authors.

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
