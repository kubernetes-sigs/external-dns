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
	"time"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
	log "github.com/sirupsen/logrus"
)

const (
	// defaultBatchChangeSize is the default maximum number of DNS record
	// operations included in each Cloudflare batch request.
	defaultBatchChangeSize = 200
)

// batchCollections groups the parallel slices that are assembled while
// classifying per-zone changes. It is passed as a unit to
// submitDNSRecordChanges and chunkBatchChanges, replacing the previous
// eight-parameter signatures and making it clear which slices travel
// together.
type batchCollections struct {
	// Batch API parameters in server-execution order: deletes → puts → posts.
	batchDeletes []dns.RecordBatchParamsDelete
	batchPosts   []dns.RecordBatchParamsPostUnion
	batchPuts    []dns.BatchPutUnionParam

	// Parallel change slices — one entry per batch param, in the same order,
	// so that a failed batch chunk can be replayed with per-record fallback.
	deleteChanges []*cloudFlareChange
	createChanges []*cloudFlareChange
	updateChanges []*cloudFlareChange

	// fallbackUpdates holds changes for record types whose batch-put param
	// requires structured Data fields (e.g. SRV, CAA). These are submitted
	// via individual UpdateDNSRecord calls instead of the batch API.
	fallbackUpdates []*cloudFlareChange
}

// batchChunk holds a DNS record batch request alongside the source changes
// that produced it, enabling per-record fallback when a batch fails.
type batchChunk struct {
	params        dns.RecordBatchParams
	deleteChanges []*cloudFlareChange
	createChanges []*cloudFlareChange
	updateChanges []*cloudFlareChange
}

// BatchDNSRecords submits a batch of DNS record changes to the Cloudflare API.
func (z zoneService) BatchDNSRecords(ctx context.Context, params dns.RecordBatchParams) (*dns.RecordBatchResponse, error) {
	return z.service.DNS.Records.Batch(ctx, params)
}

// getUpdateDNSRecordParam returns the RecordUpdateParams for an individual update.
func getUpdateDNSRecordParam(zoneID string, cfc cloudFlareChange) dns.RecordUpdateParams {
	return dns.RecordUpdateParams{
		ZoneID: cloudflare.F(zoneID),
		Body: dns.RecordUpdateParamsBody{
			Name:     cloudflare.F(cfc.ResourceRecord.Name),
			TTL:      cloudflare.F(cfc.ResourceRecord.TTL),
			Proxied:  cloudflare.F(cfc.ResourceRecord.Proxied),
			Type:     cloudflare.F(dns.RecordUpdateParamsBodyType(cfc.ResourceRecord.Type)),
			Content:  cloudflare.F(cfc.ResourceRecord.Content),
			Priority: cloudflare.F(cfc.ResourceRecord.Priority),
			Comment:  cloudflare.F(cfc.ResourceRecord.Comment),
			Tags:     cloudflare.F(cfc.ResourceRecord.Tags),
		},
	}
}

// getCreateDNSRecordParam returns the RecordNewParams for an individual create.
func getCreateDNSRecordParam(zoneID string, cfc *cloudFlareChange) dns.RecordNewParams {
	return dns.RecordNewParams{
		ZoneID: cloudflare.F(zoneID),
		Body: dns.RecordNewParamsBody{
			Name:     cloudflare.F(cfc.ResourceRecord.Name),
			TTL:      cloudflare.F(cfc.ResourceRecord.TTL),
			Proxied:  cloudflare.F(cfc.ResourceRecord.Proxied),
			Type:     cloudflare.F(dns.RecordNewParamsBodyType(cfc.ResourceRecord.Type)),
			Content:  cloudflare.F(cfc.ResourceRecord.Content),
			Priority: cloudflare.F(cfc.ResourceRecord.Priority),
			Comment:  cloudflare.F(cfc.ResourceRecord.Comment),
			Tags:     cloudflare.F(cfc.ResourceRecord.Tags),
		},
	}
}

// chunkBatchChanges splits DNS record batch operations into batchChunks,
// each containing at most <limit> total operations. Operations are distributed
// in server-execution order: deletes first, then puts, then posts.
// The parallel change slices track which cloudFlareChange produced each batch
// param so that individual fallback is possible when a chunk fails.
func chunkBatchChanges(zoneID string, bc batchCollections, limit int) []batchChunk {
	deletes, deleteChanges := bc.batchDeletes, bc.deleteChanges
	posts, createChanges := bc.batchPosts, bc.createChanges
	puts, updateChanges := bc.batchPuts, bc.updateChanges

	var chunks []batchChunk
	di, pi, ui := 0, 0, 0
	for di < len(deletes) || pi < len(posts) || ui < len(puts) {
		remaining := limit
		chunk := batchChunk{
			params: dns.RecordBatchParams{ZoneID: cloudflare.F(zoneID)},
		}

		if di < len(deletes) && remaining > 0 {
			end := min(di+remaining, len(deletes))
			chunk.params.Deletes = cloudflare.F(deletes[di:end])
			chunk.deleteChanges = deleteChanges[di:end]
			remaining -= end - di
			di = end
		}

		if ui < len(puts) && remaining > 0 {
			end := min(ui+remaining, len(puts))
			chunk.params.Puts = cloudflare.F(puts[ui:end])
			chunk.updateChanges = updateChanges[ui:end]
			remaining -= end - ui
			ui = end
		}

		if pi < len(posts) && remaining > 0 {
			end := min(pi+remaining, len(posts))
			chunk.params.Posts = cloudflare.F(posts[pi:end])
			chunk.createChanges = createChanges[pi:end]
			pi = end
		}

		chunks = append(chunks, chunk)
	}
	return chunks
}

// tagsFromResponse converts a RecordResponse Tags field (any) to the typed tag slice.
func tagsFromResponse(tags any) []dns.RecordTagsParam {
	if ts, ok := tags.([]string); ok {
		return ts
	}
	return nil
}

// buildBatchPostParam constructs a RecordBatchParamsPost for creating a DNS record in a batch.
func buildBatchPostParam(r dns.RecordResponse) dns.RecordBatchParamsPost {
	return dns.RecordBatchParamsPost{
		Name:     cloudflare.F(r.Name),
		TTL:      cloudflare.F(r.TTL),
		Type:     cloudflare.F(dns.RecordBatchParamsPostsType(r.Type)),
		Content:  cloudflare.F(r.Content),
		Proxied:  cloudflare.F(r.Proxied),
		Priority: cloudflare.F(r.Priority),
		Comment:  cloudflare.F(r.Comment),
		Tags:     cloudflare.F[any](tagsFromResponse(r.Tags)),
	}
}

// buildBatchPutParam constructs a BatchPutUnionParam for updating a DNS record in a batch.
// Returns (nil, false) for record types that use structured Data fields (e.g. SRV, CAA),
// which fall back to individual UpdateDNSRecord calls.
func buildBatchPutParam(id string, r dns.RecordResponse) (dns.BatchPutUnionParam, bool) {
	tags := tagsFromResponse(r.Tags)
	comment := r.Comment
	switch r.Type {
	case dns.RecordResponseTypeA:
		return dns.BatchPutARecordParam{
			ID: cloudflare.F(id),
			ARecordParam: dns.ARecordParam{
				Name:    cloudflare.F(r.Name),
				TTL:     cloudflare.F(r.TTL),
				Type:    cloudflare.F(dns.ARecordTypeA),
				Content: cloudflare.F(r.Content),
				Proxied: cloudflare.F(r.Proxied),
				Comment: cloudflare.F(comment),
				Tags:    cloudflare.F(tags),
			},
		}, true
	case dns.RecordResponseTypeAAAA:
		return dns.BatchPutAAAARecordParam{
			ID: cloudflare.F(id),
			AAAARecordParam: dns.AAAARecordParam{
				Name:    cloudflare.F(r.Name),
				TTL:     cloudflare.F(r.TTL),
				Type:    cloudflare.F(dns.AAAARecordTypeAAAA),
				Content: cloudflare.F(r.Content),
				Proxied: cloudflare.F(r.Proxied),
				Comment: cloudflare.F(comment),
				Tags:    cloudflare.F(tags),
			},
		}, true
	case dns.RecordResponseTypeCNAME:
		return dns.BatchPutCNAMERecordParam{
			ID: cloudflare.F(id),
			CNAMERecordParam: dns.CNAMERecordParam{
				Name:    cloudflare.F(r.Name),
				TTL:     cloudflare.F(r.TTL),
				Type:    cloudflare.F(dns.CNAMERecordTypeCNAME),
				Content: cloudflare.F(r.Content),
				Proxied: cloudflare.F(r.Proxied),
				Comment: cloudflare.F(comment),
				Tags:    cloudflare.F(tags),
			},
		}, true
	case dns.RecordResponseTypeTXT:
		return dns.BatchPutTXTRecordParam{
			ID: cloudflare.F(id),
			TXTRecordParam: dns.TXTRecordParam{
				Name:    cloudflare.F(r.Name),
				TTL:     cloudflare.F(r.TTL),
				Type:    cloudflare.F(dns.TXTRecordTypeTXT),
				Content: cloudflare.F(r.Content),
				Proxied: cloudflare.F(r.Proxied),
				Comment: cloudflare.F(comment),
				Tags:    cloudflare.F(tags),
			},
		}, true
	case dns.RecordResponseTypeMX:
		return dns.BatchPutMXRecordParam{
			ID: cloudflare.F(id),
			MXRecordParam: dns.MXRecordParam{
				Name:     cloudflare.F(r.Name),
				TTL:      cloudflare.F(r.TTL),
				Type:     cloudflare.F(dns.MXRecordTypeMX),
				Content:  cloudflare.F(r.Content),
				Proxied:  cloudflare.F(r.Proxied),
				Comment:  cloudflare.F(comment),
				Tags:     cloudflare.F(tags),
				Priority: cloudflare.F(r.Priority),
			},
		}, true
	case dns.RecordResponseTypeNS:
		return dns.BatchPutNSRecordParam{
			ID: cloudflare.F(id),
			NSRecordParam: dns.NSRecordParam{
				Name:    cloudflare.F(r.Name),
				TTL:     cloudflare.F(r.TTL),
				Type:    cloudflare.F(dns.NSRecordTypeNS),
				Content: cloudflare.F(r.Content),
				Proxied: cloudflare.F(r.Proxied),
				Comment: cloudflare.F(comment),
				Tags:    cloudflare.F(tags),
			},
		}, true
	default:
		// Record types that use structured Data fields (SRV, CAA, etc.) are not
		// supported in the generic batch put and fall back to individual updates.
		return nil, false
	}
}

// buildBatchCollections classifies per-zone changes into batch collections.
// Custom hostname side-effects are handled separately by
// processCustomHostnameChanges before this is called.
func (p *CloudFlareProvider) buildBatchCollections(
	zoneID string,
	changes []*cloudFlareChange,
	records DNSRecordsMap,
) batchCollections {
	var bc batchCollections

	for _, change := range changes {
		logFields := log.Fields{
			"record": change.ResourceRecord.Name,
			"type":   change.ResourceRecord.Type,
			"ttl":    change.ResourceRecord.TTL,
			"action": change.Action.String(),
			"zone":   zoneID,
		}

		switch change.Action {
		case cloudFlareCreate:
			bc.batchPosts = append(bc.batchPosts, buildBatchPostParam(change.ResourceRecord))
			bc.createChanges = append(bc.createChanges, change)

		case cloudFlareDelete:
			recordID := p.getRecordID(records, change.ResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
				continue
			}
			bc.batchDeletes = append(bc.batchDeletes, dns.RecordBatchParamsDelete{ID: cloudflare.F(recordID)})
			bc.deleteChanges = append(bc.deleteChanges, change)

		case cloudFlareUpdate:
			recordID := p.getRecordID(records, change.ResourceRecord)
			if recordID == "" {
				log.WithFields(logFields).Errorf("failed to find previous record: %v", change.ResourceRecord)
				continue
			}
			if putParam, ok := buildBatchPutParam(recordID, change.ResourceRecord); ok {
				bc.batchPuts = append(bc.batchPuts, putParam)
				bc.updateChanges = append(bc.updateChanges, change)
			} else {
				log.WithFields(logFields).Debugf("batch PUT not supported for type %s, using individual update", change.ResourceRecord.Type)
				bc.fallbackUpdates = append(bc.fallbackUpdates, change)
			}
		}
	}

	return bc
}

// submitDNSRecordChanges submits the pre-built batch collections and any
// fallback individual updates for a single zone. When a batch chunk fails,
// the provider falls back to individual API calls for that chunk's changes
// (since the batch is transactional — failure means full rollback).
// Returns true if any operation fails.
func (p *CloudFlareProvider) submitDNSRecordChanges(
	ctx context.Context,
	zoneID string,
	bc batchCollections,
	records DNSRecordsMap,
) bool {
	failed := false
	if len(bc.batchDeletes) > 0 || len(bc.batchPosts) > 0 || len(bc.batchPuts) > 0 {
		limit := max(p.DNSRecordsConfig.BatchChangeSize, defaultBatchChangeSize)
		chunks := chunkBatchChanges(zoneID, bc, limit)
		for i, chunk := range chunks {
			log.Debugf("Submitting batch DNS records for zone %s (chunk %d/%d): %d deletes, %d creates, %d updates",
				zoneID, i+1, len(chunks),
				len(chunk.params.Deletes.Value),
				len(chunk.params.Posts.Value),
				len(chunk.params.Puts.Value),
			)
			if _, err := p.Client.BatchDNSRecords(ctx, chunk.params); err != nil {
				log.Warnf("Batch DNS operation failed for zone %s (chunk %d/%d): %v — falling back to individual operations",
					zoneID, i+1, len(chunks), convertCloudflareError(err))
				if p.fallbackIndividualChanges(ctx, zoneID, chunk, records) {
					failed = true
				}
			} else {
				log.Debugf("Successfully submitted batch DNS records for zone %s (chunk %d/%d)", zoneID, i+1, len(chunks))
			}
			if i < len(chunks)-1 && p.DNSRecordsConfig.BatchChangeInterval > 0 {
				time.Sleep(p.DNSRecordsConfig.BatchChangeInterval)
			}
		}
	}
	for _, change := range bc.fallbackUpdates {
		logFields := log.Fields{
			"record": change.ResourceRecord.Name,
			"type":   change.ResourceRecord.Type,
			"ttl":    change.ResourceRecord.TTL,
			"action": change.Action.String(),
			"zone":   zoneID,
		}
		recordID := p.getRecordID(records, change.ResourceRecord)
		recordParam := getUpdateDNSRecordParam(zoneID, *change)
		if _, err := p.Client.UpdateDNSRecord(ctx, recordID, recordParam); err != nil {
			failed = true
			log.WithFields(logFields).Errorf("failed to update record: %v", err)
		} else {
			log.WithFields(logFields).Debugf("individual update succeeded")
		}
	}
	return failed
}

// fallbackIndividualChanges replays a failed (rolled-back) batch chunk as
// individual API calls. Because the batch API is transactional, a failure means
// zero state was changed in that chunk, so these individual calls are the first
// real mutations. Individual calls return Cloudflare's own per-record error
// details.
//
// Execution order matches the batch contract: deletes → updates → creates.
// Returns true if any operation failed.
func (p *CloudFlareProvider) fallbackIndividualChanges(
	ctx context.Context,
	zoneID string,
	chunk batchChunk,
	records DNSRecordsMap,
) bool {
	failed := false

	// Process in batch execution order: deletes → updates → creates.
	groups := []struct {
		changes []*cloudFlareChange
	}{
		{chunk.deleteChanges},
		{chunk.updateChanges},
		{chunk.createChanges},
	}

	for _, group := range groups {
		for _, change := range group.changes {
			logFields := log.Fields{
				"record":  change.ResourceRecord.Name,
				"type":    change.ResourceRecord.Type,
				"content": change.ResourceRecord.Content,
				"action":  change.Action.String(),
				"zone":    zoneID,
			}

			var err error
			switch change.Action {
			case cloudFlareCreate:
				params := getCreateDNSRecordParam(zoneID, change)
				_, err = p.Client.CreateDNSRecord(ctx, params)

			case cloudFlareDelete:
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					// Record is already absent — the desired state is achieved.
					log.WithFields(logFields).Info("fallback: record already gone, treating delete as success")
					continue
				}
				err = p.Client.DeleteDNSRecord(ctx, recordID, dns.RecordDeleteParams{
					ZoneID: cloudflare.F(zoneID),
				})

			case cloudFlareUpdate:
				recordID := p.getRecordID(records, change.ResourceRecord)
				if recordID == "" {
					// Record is gone; let the next sync cycle issue a fresh CREATE.
					log.WithFields(logFields).Info("fallback: record unexpectedly not found for update, will re-evaluate on next sync")
					continue
				}
				params := getUpdateDNSRecordParam(zoneID, *change)
				_, err = p.Client.UpdateDNSRecord(ctx, recordID, params)
			}

			if err != nil {
				failed = true
				log.WithFields(logFields).Errorf("fallback: individual %s failed: %v", change.Action, convertCloudflareError(err))
			} else {
				log.WithFields(logFields).Debugf("fallback: individual %s succeeded", change.Action)
			}
		}
	}

	return failed
}
