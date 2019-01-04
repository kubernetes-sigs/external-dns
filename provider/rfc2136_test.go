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

package provider

import (
	"strings"
	"testing"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type rfc2136Stub struct {
	output     []*dns.Envelope
	updateMsgs []*dns.Msg
	createMsgs []*dns.Msg
}

func newStub() *rfc2136Stub {
	return &rfc2136Stub{
		output:     make([]*dns.Envelope, 0),
		updateMsgs: make([]*dns.Msg, 0),
		createMsgs: make([]*dns.Msg, 0),
	}
}

const (
	searchPattern = "AUTHORITY SECTION:"
)

func (r *rfc2136Stub) SendMessage(msg *dns.Msg) error {

	a := msg.String()

	p0 := strings.Index(a, searchPattern)
	a = strings.Trim(a[p0+len(searchPattern):], "\n")
	a = strings.Replace(a, "\t", " ", -1)
	log.Info(a)

	if strings.Contains(a, " CLASS255 ") {
		r.updateMsgs = append(r.updateMsgs, msg)
	} else if strings.Contains(a, " IN ") {
		r.createMsgs = append(r.createMsgs, msg)
	}

	return nil
}

func (r *rfc2136Stub) setOutput(output []string) error {
	r.output = make([]*dns.Envelope, len(output))
	for i, e := range output {
		rr, err := dns.NewRR(e)
		if err != nil {
			return err
		}
		r.output[i] = &dns.Envelope{
			RR: []dns.RR{rr},
		}
	}
	return nil
}

func (r *rfc2136Stub) IncomeTransfer(m *dns.Msg, a string) (env chan *dns.Envelope, err error) {
	outChan := make(chan *dns.Envelope)
	go func() {
		for _, e := range r.output {
			outChan <- e
		}
		close(outChan)
	}()

	return outChan, nil
}

func createRfc2136StubProvider(stub *rfc2136Stub) (Provider, error) {
	return NewRfc2136Provider("", 0, "", false, "key", "secret", "hmac-sha512", true, DomainFilter{}, false, stub)
}

func TestRfc2136GetRecords(t *testing.T) {
	stub := newStub()
	err := stub.setOutput([]string{
		"v4.barfoo.com 3600 TXT test1",
		"v1.foo.com 3600 TXT test2",
		"v2.bar.com 3600 A 8.8.8.8",
		"v3.bar.com 3600 TXT bbbb",
		"v2.foo.com 3600 CNAME cccc",
		"v1.foobar.com 3600 TXT dddd",
	})
	assert.NoError(t, err)

	provider, err := createRfc2136StubProvider(stub)
	assert.NoError(t, err)

	recs, err := provider.Records()
	assert.NoError(t, err)

	assert.Equal(t, 6, len(recs))
	assert.True(t, contains(recs, "v1.foo.com"))
	assert.True(t, contains(recs, "v2.bar.com"))
	assert.True(t, contains(recs, "v2.foo.com"))
}

func TestRfc2136ApplyChanges(t *testing.T) {
	stub := newStub()
	provider, err := createRfc2136StubProvider(stub)
	assert.NoError(t, err)

	p := &plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "v1.foo.com",
				RecordType: "A",
				Targets:    []string{"1.2.3.4"},
			},
			{
				DNSName:    "v1.foobar.com",
				RecordType: "TXT",
				Targets:    []string{"boom"},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "v2.foo.com",
				RecordType: "A",
				Targets:    []string{""},
			},
			{
				DNSName:    "v2.foobar.com",
				RecordType: "TXT",
				Targets:    []string{""},
			},
		},
	}

	err = provider.ApplyChanges(p)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(stub.createMsgs))
	assert.True(t, strings.Contains(stub.createMsgs[0].String(), "v1.foo.com"))
	assert.True(t, strings.Contains(stub.createMsgs[0].String(), "1.2.3.4"))

	assert.True(t, strings.Contains(stub.createMsgs[1].String(), "v1.foobar.com"))
	assert.True(t, strings.Contains(stub.createMsgs[1].String(), "boom"))

	assert.Equal(t, 2, len(stub.updateMsgs))
	assert.True(t, strings.Contains(stub.updateMsgs[0].String(), "v2.foo.com"))
	assert.True(t, strings.Contains(stub.updateMsgs[1].String(), "v2.foobar.com"))

}
