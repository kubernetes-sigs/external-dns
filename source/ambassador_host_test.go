/*
Copyright 2019 The Kubernetes Authors.

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

package source

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type AmbassadorSuite struct {
	suite.Suite
}

func TestAmbassadorSource(t *testing.T) {
	suite.Run(t, new(AmbassadorSuite))
	t.Run("Interface", testAmbassadorSourceImplementsSource)
}

// testAmbassadorSourceImplementsSource tests that ambassadorHostSource is a valid Source.
func testAmbassadorSourceImplementsSource(t *testing.T) {
	require.Implements(t, (*Source)(nil), new(ambassadorHostSource))
}

// TestParseAmbLoadBalancerService tests our parsing of Ambassador service info.
func TestParseAmbLoadBalancerService(t *testing.T) {
	vectors := []struct {
		input  string
		ns     string
		svc    string
		errstr string
	}{
		{"svc", "default", "svc", ""},
		{"ns/svc", "ns", "svc", ""},
		{"svc.ns", "ns", "svc", ""},
		{"svc.ns.foo.bar", "ns.foo.bar", "svc", ""},
		{"ns/svc/foo/bar", "", "", "invalid external-dns service: ns/svc/foo/bar"},
		{"ns/svc/foo.bar", "", "", "invalid external-dns service: ns/svc/foo.bar"},
		{"ns.foo/svc/bar", "", "", "invalid external-dns service: ns.foo/svc/bar"},
	}

	for _, v := range vectors {
		ns, svc, err := parseAmbLoadBalancerService(v.input)

		errstr := ""

		if err != nil {
			errstr = err.Error()
		}

		if v.ns != ns {
			t.Errorf("%s: got ns \"%s\", wanted \"%s\"", v.input, ns, v.ns)
		}

		if v.svc != svc {
			t.Errorf("%s: got svc \"%s\", wanted \"%s\"", v.input, svc, v.svc)
		}

		if v.errstr != errstr {
			t.Errorf("%s: got err \"%s\", wanted \"%s\"", v.input, errstr, v.errstr)
		}
	}
}
