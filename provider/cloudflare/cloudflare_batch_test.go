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

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"strings"
	"testing"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/dns"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

func (m *mockCloudFlareClient) BatchDNSRecords(_ context.Context, params dns.RecordBatchParams) (*dns.RecordBatchResponse, error) {
	m.BatchDNSRecordsCalls++
	zoneID := params.ZoneID.Value

	// Snapshot zone state for transactional rollback on error.
	// The real Cloudflare batch API is fully transactional — if any
	// operation fails, the entire batch is rolled back.
	var snapshot map[string]dns.RecordResponse
	if zone, ok := m.Records[zoneID]; ok {
		snapshot = make(map[string]dns.RecordResponse, len(zone))
		maps.Copy(snapshot, zone)
	}
	actionsStart := len(m.Actions)

	var firstErr error

	// Process Deletes first to mirror the real API's ordering.
	for _, del := range params.Deletes.Value {
		recordID := del.ID.Value
		m.Actions = append(m.Actions, MockAction{
			Name:     "Delete",
			ZoneId:   zoneID,
			RecordId: recordID,
		})
		if zone, ok := m.Records[zoneID]; ok {
			if rec, exists := zone[recordID]; exists {
				name := rec.Name
				delete(zone, recordID)
				if strings.HasPrefix(name, "newerror-delete-") && firstErr == nil {
					firstErr = errors.New("failed to delete erroring DNS record")
				}
			}
		}
	}

	// Process Puts (updates) before Posts (creates) to mirror the real API's
	// server-side execution order: Deletes → Patches → Puts → Posts.
	for _, putUnion := range params.Puts.Value {
		id, record := extractBatchPutData(putUnion)
		m.Actions = append(m.Actions, MockAction{
			Name:       "Update",
			ZoneId:     zoneID,
			RecordId:   id,
			RecordData: record,
		})
		if zone, ok := m.Records[zoneID]; ok {
			if _, exists := zone[id]; exists {
				if strings.HasPrefix(record.Name, "newerror-update-") {
					if firstErr == nil {
						firstErr = errors.New("failed to update erroring DNS record")
					}
				} else {
					zone[id] = record
				}
			}
		}
	}

	// Process Posts (creates).
	for _, postUnion := range params.Posts.Value {
		post, ok := postUnion.(dns.RecordBatchParamsPost)
		if !ok {
			continue
		}
		typeStr := string(post.Type.Value)
		record := dns.RecordResponse{
			ID:       generateDNSRecordID(typeStr, post.Name.Value, post.Content.Value),
			Name:     post.Name.Value,
			TTL:      dns.TTL(post.TTL.Value),
			Proxied:  post.Proxied.Value,
			Type:     dns.RecordResponseType(typeStr),
			Content:  post.Content.Value,
			Priority: post.Priority.Value,
		}
		m.Actions = append(m.Actions, MockAction{
			Name:       "Create",
			ZoneId:     zoneID,
			RecordId:   record.ID,
			RecordData: record,
		})
		if zone, ok := m.Records[zoneID]; ok {
			zone[record.ID] = record
		}
		if record.Name == "newerror.bar.com" && firstErr == nil {
			firstErr = fmt.Errorf("failed to create record")
		}
	}

	// Transactional: on error, rollback all state and action changes.
	if firstErr != nil {
		if snapshot != nil {
			m.Records[zoneID] = snapshot
		}
		m.Actions = m.Actions[:actionsStart]
		return nil, firstErr
	}

	return &dns.RecordBatchResponse{}, nil
}

// extractBatchPutData unpacks a BatchPutUnionParam into a record ID and a RecordResponse
// suitable for recording in the mock's Actions list.
func extractBatchPutData(put dns.BatchPutUnionParam) (string, dns.RecordResponse) {
	switch p := put.(type) {
	case dns.BatchPutARecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:      p.ID.Value,
			Name:    p.Name.Value,
			TTL:     p.TTL.Value,
			Proxied: p.Proxied.Value,
			Type:    dns.RecordResponseTypeA,
			Content: p.Content.Value,
		}
	case dns.BatchPutAAAARecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:      p.ID.Value,
			Name:    p.Name.Value,
			TTL:     p.TTL.Value,
			Proxied: p.Proxied.Value,
			Type:    dns.RecordResponseTypeAAAA,
			Content: p.Content.Value,
		}
	case dns.BatchPutCNAMERecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:      p.ID.Value,
			Name:    p.Name.Value,
			TTL:     p.TTL.Value,
			Proxied: p.Proxied.Value,
			Type:    dns.RecordResponseTypeCNAME,
			Content: p.Content.Value,
		}
	case dns.BatchPutTXTRecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:      p.ID.Value,
			Name:    p.Name.Value,
			TTL:     p.TTL.Value,
			Proxied: p.Proxied.Value,
			Type:    dns.RecordResponseTypeTXT,
			Content: p.Content.Value,
		}
	case dns.BatchPutMXRecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:       p.ID.Value,
			Name:     p.Name.Value,
			TTL:      p.TTL.Value,
			Proxied:  p.Proxied.Value,
			Type:     dns.RecordResponseTypeMX,
			Content:  p.Content.Value,
			Priority: p.Priority.Value,
		}
	case dns.BatchPutNSRecordParam:
		return p.ID.Value, dns.RecordResponse{
			ID:      p.ID.Value,
			Name:    p.Name.Value,
			TTL:     p.TTL.Value,
			Proxied: p.Proxied.Value,
			Type:    dns.RecordResponseTypeNS,
			Content: p.Content.Value,
		}
	default:
		panic(fmt.Sprintf("extractBatchPutData: unexpected BatchPutUnionParam type %T", put))
	}
}

// generateDNSRecordID builds the deterministic record ID used by the mock client.
func generateDNSRecordID(rrtype string, name string, content string) string {
	return fmt.Sprintf("%s-%s-%s", name, rrtype, content)
}

func TestBatchFallbackIndividual(t *testing.T) {
	t.Run("batch failure falls back to individual operations", func(t *testing.T) {
		// Create a provider with pre-existing records.
		client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
			"001": {
				{ID: "existing-1", Name: "ok.bar.com", Type: "A", Content: "1.2.3.4", TTL: 120},
			},
		})
		p := &CloudFlareProvider{
			Client: client,
		}

		// Apply changes that include a good create and a bad create.
		// "newerror.bar.com" triggers a batch failure in the mock BatchDNSRecords,
		// then an individual fallback failure in CreateDNSRecord.
		changes := &plan.Changes{
			Create: []*endpoint.Endpoint{
				{DNSName: "good.bar.com", Targets: endpoint.Targets{"5.6.7.8"}, RecordType: "A"},
				{DNSName: "newerror.bar.com", Targets: endpoint.Targets{"9.10.11.12"}, RecordType: "A"},
			},
		}

		err := p.ApplyChanges(t.Context(), changes)
		require.Error(t, err, "should return error when individual fallback has failures")
		assert.Equal(t, 1, client.BatchDNSRecordsCalls, "batch path should be attempted before fallback")

		// The batch should have failed (because of newerror.bar.com), then
		// fallback should have applied "good.bar.com" individually (success)
		// and "newerror.bar.com" individually (failure).

		// Verify the good record was created via individual fallback.
		zone001 := client.Records["001"]
		goodID := generateDNSRecordID("A", "good.bar.com", "5.6.7.8")
		assert.Contains(t, zone001, goodID, "good record should exist after individual fallback")
	})

	t.Run("failed individual delete is reported", func(t *testing.T) {
		// When a batch containing two deletes fails, the fallback replays them
		// individually. The one that ultimately fails should be reported;
		// the one that succeeds should not block the overall zone from converging.
		client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
			"001": {
				{ID: "del-ok", Name: "deleteme.bar.com", Type: "A", Content: "1.2.3.4", TTL: 120},
				{ID: "del-err", Name: "newerror-delete-1.bar.com", Type: "A", Content: "5.6.7.8", TTL: 120},
			},
		})
		p := &CloudFlareProvider{
			Client: client,
		}

		changes := &plan.Changes{
			Delete: []*endpoint.Endpoint{
				{DNSName: "deleteme.bar.com", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: "A"},
				{DNSName: "newerror-delete-1.bar.com", Targets: endpoint.Targets{"5.6.7.8"}, RecordType: "A"},
			},
		}
		err := p.ApplyChanges(t.Context(), changes)
		require.Error(t, err, "should return error for the failing delete")

		// The good delete should have succeeded via individual fallback.
		assert.NotContains(t, client.Records["001"], "del-ok", "successfully deleted record should be gone")
	})

	t.Run("fallback update failure is reported", func(t *testing.T) {
		client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
			"001": {
				{ID: "upd-err", Name: "newerror-update-1.bar.com", Type: "A", Content: "1.2.3.4", TTL: 120},
			},
		})
		p := &CloudFlareProvider{
			Client: client,
		}

		changes := &plan.Changes{
			UpdateNew: []*endpoint.Endpoint{
				{DNSName: "newerror-update-1.bar.com", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: "A", RecordTTL: 300},
			},
			UpdateOld: []*endpoint.Endpoint{
				{DNSName: "newerror-update-1.bar.com", Targets: endpoint.Targets{"1.2.3.4"}, RecordType: "A", RecordTTL: 120},
			},
		}
		err := p.ApplyChanges(t.Context(), changes)
		require.Error(t, err, "should return error for the failing update")
	})
}

func TestChunkBatchChanges(t *testing.T) {
	// Build sample changes and batch params.
	mkDelete := func(id string) dns.RecordBatchParamsDelete {
		return dns.RecordBatchParamsDelete{ID: cloudflare.F(id)}
	}
	mkPost := func(name, content string) dns.RecordBatchParamsPostUnion {
		return dns.RecordBatchParamsPost{
			Name:    cloudflare.F(name),
			Type:    cloudflare.F(dns.RecordBatchParamsPostsTypeA),
			Content: cloudflare.F(content),
		}
	}
	mkPut := func(id, name, content string) dns.BatchPutUnionParam {
		return dns.BatchPutARecordParam{
			ID: cloudflare.F(id),
			ARecordParam: dns.ARecordParam{
				Name:    cloudflare.F(name),
				Type:    cloudflare.F(dns.ARecordTypeA),
				Content: cloudflare.F(content),
			},
		}
	}
	mkChange := func(action changeAction, name, content string) *cloudFlareChange {
		return &cloudFlareChange{
			Action:         action,
			ResourceRecord: dns.RecordResponse{Name: name, Type: "A", Content: content},
		}
	}

	deletes := []dns.RecordBatchParamsDelete{mkDelete("d1"), mkDelete("d2")}
	deleteChanges := []*cloudFlareChange{
		mkChange(cloudFlareDelete, "del1.bar.com", "1.1.1.1"),
		mkChange(cloudFlareDelete, "del2.bar.com", "2.2.2.2"),
	}
	posts := []dns.RecordBatchParamsPostUnion{mkPost("create1.bar.com", "3.3.3.3")}
	createChanges := []*cloudFlareChange{
		mkChange(cloudFlareCreate, "create1.bar.com", "3.3.3.3"),
	}
	puts := []dns.BatchPutUnionParam{mkPut("u1", "update1.bar.com", "4.4.4.4")}
	updateChanges := []*cloudFlareChange{
		mkChange(cloudFlareUpdate, "update1.bar.com", "4.4.4.4"),
	}

	t.Run("single chunk when under limit", func(t *testing.T) {
		bc := batchCollections{
			batchDeletes:  deletes,
			deleteChanges: deleteChanges,
			batchPosts:    posts,
			createChanges: createChanges,
			batchPuts:     puts,
			updateChanges: updateChanges,
		}
		chunks := chunkBatchChanges("zone1", bc, 10)
		require.Len(t, chunks, 1)
		assert.Len(t, chunks[0].deleteChanges, 2)
		assert.Len(t, chunks[0].createChanges, 1)
		assert.Len(t, chunks[0].updateChanges, 1)
	})

	t.Run("splits into multiple chunks at limit", func(t *testing.T) {
		bc := batchCollections{
			batchDeletes:  deletes,
			deleteChanges: deleteChanges,
			batchPosts:    posts,
			createChanges: createChanges,
			batchPuts:     puts,
			updateChanges: updateChanges,
		}
		chunks := chunkBatchChanges("zone1", bc, 2)
		require.Len(t, chunks, 2)
		// First chunk: 2 deletes (fills limit)
		assert.Len(t, chunks[0].deleteChanges, 2)
		assert.Empty(t, chunks[0].updateChanges)
		assert.Empty(t, chunks[0].createChanges)
		// Second chunk: 1 put then 1 post
		assert.Empty(t, chunks[1].deleteChanges)
		assert.Len(t, chunks[1].updateChanges, 1)
		assert.Len(t, chunks[1].createChanges, 1)
	})

	t.Run("preserves operation order across chunk boundaries", func(t *testing.T) {
		bc := batchCollections{
			batchDeletes: []dns.RecordBatchParamsDelete{mkDelete("d1")},
			deleteChanges: []*cloudFlareChange{
				mkChange(cloudFlareDelete, "del1.bar.com", "1.1.1.1"),
			},
			batchPuts: []dns.BatchPutUnionParam{
				mkPut("u1", "update1.bar.com", "2.2.2.2"),
				mkPut("u2", "update2.bar.com", "3.3.3.3"),
			},
			updateChanges: []*cloudFlareChange{
				mkChange(cloudFlareUpdate, "update1.bar.com", "2.2.2.2"),
				mkChange(cloudFlareUpdate, "update2.bar.com", "3.3.3.3"),
			},
			batchPosts: []dns.RecordBatchParamsPostUnion{
				mkPost("create1.bar.com", "4.4.4.4"),
				mkPost("create2.bar.com", "5.5.5.5"),
			},
			createChanges: []*cloudFlareChange{
				mkChange(cloudFlareCreate, "create1.bar.com", "4.4.4.4"),
				mkChange(cloudFlareCreate, "create2.bar.com", "5.5.5.5"),
			},
		}

		chunks := chunkBatchChanges("zone1", bc, 2)
		require.Len(t, chunks, 3)

		assert.Len(t, chunks[0].deleteChanges, 1)
		assert.Len(t, chunks[0].updateChanges, 1)
		assert.Empty(t, chunks[0].createChanges)

		assert.Empty(t, chunks[1].deleteChanges)
		assert.Len(t, chunks[1].updateChanges, 1)
		assert.Len(t, chunks[1].createChanges, 1)

		assert.Empty(t, chunks[2].deleteChanges)
		assert.Empty(t, chunks[2].updateChanges)
		assert.Len(t, chunks[2].createChanges, 1)
	})
}

func TestTagsFromResponse(t *testing.T) {
	t.Run("nil input returns nil", func(t *testing.T) {
		assert.Nil(t, tagsFromResponse(nil))
	})
	t.Run("non-string-slice returns nil", func(t *testing.T) {
		assert.Nil(t, tagsFromResponse(42))
	})
	t.Run("string slice is returned unchanged", func(t *testing.T) {
		tags := []string{"tag1", "tag2"}
		assert.Equal(t, tags, tagsFromResponse(tags))
	})
}

func TestBuildBatchPutParam(t *testing.T) {
	base := dns.RecordResponse{
		Name:    "example.bar.com",
		TTL:     120,
		Proxied: false,
		Comment: "test-comment",
	}

	t.Run("AAAA record", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeAAAA
		r.Content = "2001:db8::1"
		param, ok := buildBatchPutParam("id-aaaa", r)
		require.True(t, ok)
		p, cast := param.(dns.BatchPutAAAARecordParam)
		require.True(t, cast)
		assert.Equal(t, "id-aaaa", p.ID.Value)
		assert.Equal(t, "2001:db8::1", p.Content.Value)
		assert.Equal(t, dns.AAAARecordTypeAAAA, p.Type.Value)
	})

	t.Run("CNAME record", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeCNAME
		r.Content = "target.bar.com"
		param, ok := buildBatchPutParam("id-cname", r)
		require.True(t, ok)
		p, cast := param.(dns.BatchPutCNAMERecordParam)
		require.True(t, cast)
		assert.Equal(t, "id-cname", p.ID.Value)
		assert.Equal(t, "target.bar.com", p.Content.Value)
		assert.Equal(t, dns.CNAMERecordTypeCNAME, p.Type.Value)
	})

	t.Run("TXT record", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeTXT
		r.Content = "v=spf1 include:example.com ~all"
		param, ok := buildBatchPutParam("id-txt", r)
		require.True(t, ok)
		p, cast := param.(dns.BatchPutTXTRecordParam)
		require.True(t, cast)
		assert.Equal(t, "id-txt", p.ID.Value)
		assert.Equal(t, dns.TXTRecordTypeTXT, p.Type.Value)
	})

	t.Run("MX record with priority", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeMX
		r.Content = "mail.example.com"
		r.Priority = 10
		param, ok := buildBatchPutParam("id-mx", r)
		require.True(t, ok)
		p, cast := param.(dns.BatchPutMXRecordParam)
		require.True(t, cast)
		assert.Equal(t, "id-mx", p.ID.Value)
		assert.InDelta(t, float64(10), float64(p.Priority.Value), 0)
		assert.Equal(t, dns.MXRecordTypeMX, p.Type.Value)
	})

	t.Run("NS record", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeNS
		r.Content = "ns1.example.com"
		param, ok := buildBatchPutParam("id-ns", r)
		require.True(t, ok)
		p, cast := param.(dns.BatchPutNSRecordParam)
		require.True(t, cast)
		assert.Equal(t, "id-ns", p.ID.Value)
		assert.Equal(t, dns.NSRecordTypeNS, p.Type.Value)
	})

	t.Run("SRV record falls back (returns nil, false)", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeSRV
		r.Content = "10 20 443 target.bar.com"
		param, ok := buildBatchPutParam("id-srv", r)
		assert.False(t, ok)
		assert.Nil(t, param)
	})

	t.Run("CAA record falls back (returns nil, false)", func(t *testing.T) {
		r := base
		r.Type = dns.RecordResponseTypeCAA
		r.Content = "0 issue letsencrypt.org"
		param, ok := buildBatchPutParam("id-caa", r)
		assert.False(t, ok)
		assert.Nil(t, param)
	})
}

func TestBuildBatchCollections_EdgeCases(t *testing.T) {
	p := &CloudFlareProvider{}

	t.Run("update with missing record ID is skipped", func(t *testing.T) {
		changes := []*cloudFlareChange{
			{
				Action: cloudFlareUpdate,
				ResourceRecord: dns.RecordResponse{
					Name:    "missing.bar.com",
					Type:    dns.RecordResponseTypeA,
					Content: "1.2.3.4",
				},
			},
		}
		// Empty records map — getRecordID will return ""
		bc := p.buildBatchCollections("zone1", changes, make(DNSRecordsMap))
		assert.Empty(t, bc.batchPuts, "missing record should not be added to batch puts")
		assert.Empty(t, bc.updateChanges)
		assert.Empty(t, bc.fallbackUpdates)
	})

	t.Run("SRV update goes to fallbackUpdates", func(t *testing.T) {
		srvRecord := dns.RecordResponse{
			ID:      "srv-1",
			Name:    "srv.bar.com",
			Type:    dns.RecordResponseTypeSRV,
			Content: "10 20 443 target.bar.com",
		}
		records := DNSRecordsMap{
			newDNSRecordIndex(srvRecord): srvRecord,
		}
		changes := []*cloudFlareChange{
			{
				Action:         cloudFlareUpdate,
				ResourceRecord: srvRecord,
			},
		}
		bc := p.buildBatchCollections("zone1", changes, records)
		assert.Empty(t, bc.batchPuts, "SRV should not be in batch puts")
		assert.Empty(t, bc.updateChanges)
		require.Len(t, bc.fallbackUpdates, 1)
		assert.Equal(t, "srv.bar.com", bc.fallbackUpdates[0].ResourceRecord.Name)
	})

	t.Run("delete with missing record ID is skipped", func(t *testing.T) {
		changes := []*cloudFlareChange{
			{
				Action: cloudFlareDelete,
				ResourceRecord: dns.RecordResponse{
					Name:    "gone.bar.com",
					Type:    dns.RecordResponseTypeA,
					Content: "1.2.3.4",
				},
			},
		}
		bc := p.buildBatchCollections("zone1", changes, make(DNSRecordsMap))
		assert.Empty(t, bc.batchDeletes, "missing record should not be added to batch deletes")
		assert.Empty(t, bc.deleteChanges)
	})
}

func TestSubmitDNSRecordChanges_BatchInterval(t *testing.T) {
	// Build 201 creates so they span 2 chunks (defaultBatchChangeSize=200),
	// triggering the time.Sleep(BatchChangeInterval) code path between chunks.
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {},
	})
	p := &CloudFlareProvider{
		Client: client,
		DNSRecordsConfig: DNSRecordsConfig{
			BatchChangeInterval: 1, // 1 nanosecond — non-zero triggers sleep
		},
	}

	const nRecords = defaultBatchChangeSize + 1
	var posts []dns.RecordBatchParamsPostUnion
	var createChanges []*cloudFlareChange
	for i := range nRecords {
		name := fmt.Sprintf("record%d.bar.com", i)
		posts = append(posts, dns.RecordBatchParamsPost{
			Name:    cloudflare.F(name),
			Type:    cloudflare.F(dns.RecordBatchParamsPostsTypeA),
			Content: cloudflare.F("1.2.3.4"),
		})
		createChanges = append(createChanges, &cloudFlareChange{
			Action:         cloudFlareCreate,
			ResourceRecord: dns.RecordResponse{Name: name, Type: "A", Content: "1.2.3.4"},
		})
	}

	bc := batchCollections{
		batchPosts:    posts,
		createChanges: createChanges,
	}

	failed := p.submitDNSRecordChanges(t.Context(), "001", bc, make(DNSRecordsMap))
	assert.False(t, failed, "should not fail")
	assert.Equal(t, 2, client.BatchDNSRecordsCalls, "two chunks should require two batch API calls")
}

func TestSubmitDNSRecordChanges_FallbackUpdates(t *testing.T) {
	t.Run("successful SRV fallback update", func(t *testing.T) {
		srvRecord := dns.RecordResponse{
			ID:      "srv-1",
			Name:    "srv.bar.com",
			Type:    dns.RecordResponseTypeSRV,
			Content: "10 20 443 target.bar.com",
		}
		client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
			"001": {srvRecord},
		})
		p := &CloudFlareProvider{Client: client}

		records := DNSRecordsMap{
			newDNSRecordIndex(srvRecord): srvRecord,
		}
		bc := batchCollections{
			fallbackUpdates: []*cloudFlareChange{
				{Action: cloudFlareUpdate, ResourceRecord: srvRecord},
			},
		}

		failed := p.submitDNSRecordChanges(t.Context(), "001", bc, records)
		assert.False(t, failed, "successful SRV fallback update should not report failure")
		assert.Equal(t, 0, client.BatchDNSRecordsCalls, "batch API not called for fallback-only changes")
	})

	t.Run("failed SRV fallback update is reported", func(t *testing.T) {
		srvRecord := dns.RecordResponse{
			ID:      "newerror-upd-srv",
			Name:    "newerror-update-srv.bar.com",
			Type:    dns.RecordResponseTypeSRV,
			Content: "10 20 443 target.bar.com",
		}
		client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
			"001": {srvRecord},
		})
		p := &CloudFlareProvider{Client: client}

		records := DNSRecordsMap{
			newDNSRecordIndex(srvRecord): srvRecord,
		}
		bc := batchCollections{
			fallbackUpdates: []*cloudFlareChange{
				{Action: cloudFlareUpdate, ResourceRecord: srvRecord},
			},
		}

		failed := p.submitDNSRecordChanges(t.Context(), "001", bc, records)
		assert.True(t, failed, "failed SRV fallback update should be reported")
	})
}

func TestFallbackIndividualChanges_MissingRecord(t *testing.T) {
	client := NewMockCloudFlareClientWithRecords(map[string][]dns.RecordResponse{
		"001": {},
	})
	p := &CloudFlareProvider{Client: client}
	emptyRecords := make(DNSRecordsMap)

	t.Run("delete where record is already gone succeeds silently", func(t *testing.T) {
		chunk := batchChunk{
			deleteChanges: []*cloudFlareChange{
				{
					Action: cloudFlareDelete,
					ResourceRecord: dns.RecordResponse{
						Name:    "gone.bar.com",
						Type:    dns.RecordResponseTypeA,
						Content: "1.2.3.4",
					},
				},
			},
		}
		failed := p.fallbackIndividualChanges(t.Context(), "001", chunk, emptyRecords)
		assert.False(t, failed, "delete of already-absent record should not report failure")
	})

	t.Run("update where record is not found skips gracefully", func(t *testing.T) {
		chunk := batchChunk{
			updateChanges: []*cloudFlareChange{
				{
					Action: cloudFlareUpdate,
					ResourceRecord: dns.RecordResponse{
						Name:    "missing.bar.com",
						Type:    dns.RecordResponseTypeA,
						Content: "1.2.3.4",
					},
				},
			},
		}
		failed := p.fallbackIndividualChanges(t.Context(), "001", chunk, emptyRecords)
		assert.False(t, failed, "update of missing record should not report failure")
	})
}
