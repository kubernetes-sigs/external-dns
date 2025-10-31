/*
Copyright 2025 The Kubernetes Authors.

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

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/miekg/dns"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider/webhook/api"
)

// DNSRecord represents a single DNS record with its type and targets
type DNSRecord struct {
	Type    string   `json:"type"`
	Targets []string `json:"targets"`
}

// DomainRecords holds all record types for a specific domain
type DomainRecords struct {
	Records map[string]DNSRecord `json:"records"` // recordType -> DNSRecord
}

// DNSRecordStore holds all DNS records in memory
type DNSRecordStore struct {
	mu      sync.RWMutex
	domains map[string]*DomainRecords // domain -> DomainRecords
}

// NewDNSRecordStore creates a new in-memory record store
func NewDNSRecordStore() *DNSRecordStore {
	return &DNSRecordStore{
		domains: make(map[string]*DomainRecords),
	}
}

// AddRecord adds or updates a DNS record
func (s *DNSRecordStore) AddRecord(domain, recordType string, targets []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.domains[domain] == nil {
		s.domains[domain] = &DomainRecords{
			Records: make(map[string]DNSRecord),
		}
	}
	s.domains[domain].Records[recordType] = DNSRecord{
		Type:    recordType,
		Targets: targets,
	}
}

// RemoveRecord removes a DNS record
func (s *DNSRecordStore) RemoveRecord(domain, recordType string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.domains[domain] != nil {
		delete(s.domains[domain].Records, recordType)
		if len(s.domains[domain].Records) == 0 {
			delete(s.domains, domain)
		}
	}
}

// GetRecord retrieves targets for a specific domain and record type
func (s *DNSRecordStore) GetRecord(domain, recordType string) []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.domains[domain] != nil {
		if record, exists := s.domains[domain].Records[recordType]; exists {
			return record.Targets
		}
	}
	return nil
}

// GetAllRecords returns all records as endpoints for webhook responses
func (s *DNSRecordStore) GetAllRecords() []endpoint.Endpoint {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var endpoints []endpoint.Endpoint
	for domain, domainRecords := range s.domains {
		for _, record := range domainRecords.Records {
			endpoints = append(endpoints, endpoint.Endpoint{
				DNSName:    domain,
				RecordType: record.Type,
				Targets:    record.Targets,
			})
		}
	}
	return endpoints
}

// startDNSServer starts the DNS server in a goroutine
func startDNSServer(store *DNSRecordStore, address string, port int, defaultTTL uint32) error {
	mux := dns.NewServeMux()
	mux.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		handleDNSQuery(store, w, r, defaultTTL)
	})

	server := &dns.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Net:     "udp",
		Handler: mux,
	}

	log.Printf("Starting DNS server on %s (UDP)\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("DNS server error: %v", err)
		}
	}()

	// Also start TCP server
	tcpServer := &dns.Server{
		Addr:    fmt.Sprintf("%s:%d", address, port),
		Net:     "tcp",
		Handler: mux,
	}

	go func() {
		if err := tcpServer.ListenAndServe(); err != nil {
			log.Printf("DNS TCP server error: %v", err)
		}
	}()

	return nil
}

// handleDNSQuery handles incoming DNS queries
func handleDNSQuery(store *DNSRecordStore, w dns.ResponseWriter, r *dns.Msg, defaultTTL uint32) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Authoritative = true

	for _, q := range r.Question {
		domain := strings.TrimSuffix(q.Name, ".")
		recordType := dns.TypeToString[q.Qtype]

		targets := store.GetRecord(domain, recordType)

		if len(targets) > 0 {
			for _, target := range targets {
				switch q.Qtype {
				case dns.TypeA:
					if ip := net.ParseIP(target); ip != nil && ip.To4() != nil {
						rr := &dns.A{
							Hdr: dns.RR_Header{
								Name:   q.Name,
								Rrtype: dns.TypeA,
								Class:  dns.ClassINET,
								Ttl:    defaultTTL,
							},
							A: ip.To4(),
						}
						m.Answer = append(m.Answer, rr)
					}
				case dns.TypeAAAA:
					if ip := net.ParseIP(target); ip != nil && ip.To16() != nil && ip.To4() == nil {
						rr := &dns.AAAA{
							Hdr: dns.RR_Header{
								Name:   q.Name,
								Rrtype: dns.TypeAAAA,
								Class:  dns.ClassINET,
								Ttl:    defaultTTL,
							},
							AAAA: ip.To16(),
						}
						m.Answer = append(m.Answer, rr)
					}
				case dns.TypeCNAME:
					rr := &dns.CNAME{
						Hdr: dns.RR_Header{
							Name:   q.Name,
							Rrtype: dns.TypeCNAME,
							Class:  dns.ClassINET,
							Ttl:    defaultTTL,
						},
						Target: dns.Fqdn(target),
					}
					m.Answer = append(m.Answer, rr)
				}
			}
		} else {
			// No records found, set NXDOMAIN
			m.Rcode = dns.RcodeNameError
		}
	}

	w.WriteMsg(m)
}

func main() {
	listenAddress := flag.String("listen-address", "127.0.0.1", "Address to listen on")
	port := flag.Int("port", 8888, "Port to listen on")
	dnsAddress := flag.String("dns-address", "127.0.0.1", "DNS server address")
	dnsPort := flag.Int("dns-port", 5353, "DNS server port")
	dnsTTL := flag.Int("dns-ttl", 300, "Default TTL for DNS responses")
	flag.Parse()

	// Create shared record store
	recordStore := NewDNSRecordStore()

	// Start DNS server
	if err := startDNSServer(recordStore, *dnsAddress, *dnsPort, uint32(*dnsTTL)); err != nil {
		log.Fatalf("Failed to start DNS server: %v", err)
	}

	// Setup HTTP handlers
	http.HandleFunc("/", negotiateHandler)
	http.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {
		recordsHandler(w, r, recordStore)
	})
	http.HandleFunc("/adjustendpoints", adjustEndpointsHandler)
	http.HandleFunc("/healthz", healthzHandler)

	addr := fmt.Sprintf("%s:%d", *listenAddress, *port)
	log.Printf("Starting webhook provider on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func negotiateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", api.MediaTypeFormatAndVersion)
	// Return your supported DomainFilter here
	json.NewEncoder(w).Encode(endpoint.DomainFilter{})
}

func recordsHandler(w http.ResponseWriter, r *http.Request, store *DNSRecordStore) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", api.MediaTypeFormatAndVersion)
		endpoints := store.GetAllRecords()
		json.NewEncoder(w).Encode(endpoints)
		return
	}

	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", api.MediaTypeFormatAndVersion)
		var changes plan.Changes
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &changes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Process deletions
		for _, del := range changes.Delete {
			store.RemoveRecord(del.DNSName, del.RecordType)
		}

		// Process updates (treat as delete + create)
		for _, update := range changes.UpdateOld {
			store.RemoveRecord(update.DNSName, update.RecordType)
		}
		for _, update := range changes.UpdateNew {
			if len(update.Targets) > 0 {
				store.AddRecord(update.DNSName, update.RecordType, update.Targets)
			}
		}

		// Process creations
		for _, create := range changes.Create {
			if len(create.Targets) > 0 {
				store.AddRecord(create.DNSName, create.RecordType, create.Targets)
			}
		}

		w.WriteHeader(http.StatusNoContent)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func adjustEndpointsHandler(w http.ResponseWriter, r *http.Request) {
	// read the endpoints from the input, return them straight back
	var endpoints []endpoint.Endpoint
	if err := json.NewDecoder(r.Body).Decode(&endpoints); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", api.MediaTypeFormatAndVersion)
	json.NewEncoder(w).Encode(endpoints)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
