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

package plan

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"sigs.k8s.io/external-dns/endpoint"
)

var _ ConflictResolver = PerResource{}

type ResolverSuite struct {
	// resolvers
	perResource PerResource
	suite.Suite
}

func (suite *ResolverSuite) SetupTest() {
	suite.perResource = PerResource{}
}

func (suite *ResolverSuite) TestStrictResolveUpdateEmpty() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{}
	candidates := []*endpoint.Endpoint{}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Empty(changes.Create)
	suite.Empty(changes.UpdateNew)
	suite.Empty(changes.UpdateOld)
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateNothing() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Empty(changes.Create)
	suite.Empty(changes.UpdateNew)
	suite.Empty(changes.UpdateOld)
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSeeMerge() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.1"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
			Labels: map[string]string{
				"server": "blaster",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "aws",
					Value: "xxxxx",
				},
			},
		},
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.2"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
			Labels: map[string]string{
				"2blaster": "blaster",
				"3blaster": "source",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "2blaster",
					Value: "xxxxx",
				},
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.3"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
			Labels: map[string]string{
				"2blaster": "override",
				"3blaster": "winner",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "3blaster",
					Value: "xxxxx",
				},
			},
		},
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.1"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
			Labels: map[string]string{
				"a":        "b",
				"1looser":  "1looser",
				"3blaster": "1looser",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "x1",
					Value: "x1",
				},
				{
					Name:  "lower",
					Value: "1.1.1.1",
				},
			},
		},
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.1"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
			Labels: map[string]string{
				"a":        "x2",
				"1looser":  "1winner",
				"c":        "d",
				"3blaster": "2looser",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "x1",
					Value: "x2",
				},
				{
					Name:  "x2",
					Value: "x2",
				},
			},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Create, 0)
	suite.Len(changes.Delete, 0)
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(changes.UpdateNew, []*endpoint.Endpoint{
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.1", "1.1.1.3"},
			RecordType:    "A",
			SetIdentifier: "server",
			RecordTTL:     300,
			Labels: map[string]string{
				"3blaster": "winner",
				"c":        "d",
				"server":   "blaster",
				"a":        "x2",
				"1looser":  "1winner",
				"2blaster": "override",
			},
			ProviderSpecific: []endpoint.ProviderSpecificProperty{
				{
					Name:  "x1",
					Value: "x1",
				},
				{
					Name:  "lower",
					Value: "1.1.1.1",
				},
				{
					Name:  "x2",
					Value: "x2",
				},
				{
					Name:  "3blaster",
					Value: "xxxxx",
				},
			},
		},
	})
	suite.Len(changes.UpdateOld, 2)
	suite.Equal(changes.UpdateOld, current)
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateUpdateOld() {
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"server": "1.1.1.1",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"server": "1.1.1.2",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"server": "1.1.1.1",
			},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Empty(changes.Create)
	suite.Empty(changes.UpdateNew)
	suite.Equal(changes.UpdateOld, []*endpoint.Endpoint{current[1]})
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateNothingDuplicated() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},

		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},

		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Empty(changes.Create)
	suite.Empty(changes.UpdateNew)
	suite.Empty(changes.UpdateOld)
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleCreate() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"v2"},
			RecordType: "MX",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"v1"},
			RecordType: "MX",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Create, 1)
	suite.Equal(*changes.Create[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"v1", "v2"},
		RecordType: "MX",
		RecordTTL:  300,
	})
	suite.Empty(changes.UpdateNew)
	suite.Empty(changes.UpdateOld)
	suite.Empty(changes.Delete)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleCreateCNAMEError() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"v2"},
			RecordType: "CNAME",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"v1"},
			RecordType: "CNAME",
			RecordTTL:  300,
		},
	}
	_, err := suite.perResource.Resolve(current, candidates)
	suite.ErrorContains(err, "inconsistent targets for")
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleDelete() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Empty(changes.UpdateNew)
	suite.Empty(changes.UpdateOld)
	suite.Equal(changes.Delete, current)
	suite.Empty(changes.Create)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleUpdateTarget() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: endpoint.Labels{
				"a": "b",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(*changes.UpdateOld[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1"},
		RecordType: "A",
		RecordTTL:  300,
		Labels: endpoint.Labels{
			"a": "b",
		},
	})
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(*changes.UpdateNew[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.2"},
		RecordType: "A",
		RecordTTL:  300,
		Labels: endpoint.Labels{
			"a": "b",
		},
	})
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleUpdateTTL() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: endpoint.Labels{
				"a": "b",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  400,
			Labels: endpoint.Labels{
				"1-2": "1-2",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.3", "1.1.1.1", "1.1.1.2"},
			RecordType: "A",
			RecordTTL:  400,
			Labels: endpoint.Labels{
				"2-x": "3-x",
			},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(changes.UpdateOld, current)
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(changes.UpdateNew[0], &endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2", "1.1.1.3"},
		RecordType: "A",
		RecordTTL:  400,
		Labels: endpoint.Labels{
			"a":   "b",
			"1-2": "1-2",
			"2-x": "3-x",
		},
	})
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleUpdateLabels() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"a": "b",
				"u": "v",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels:     map[string]string{"a": "c"},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(*changes.UpdateOld[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1"},
		RecordType: "A",
		RecordTTL:  300,
		Labels: map[string]string{
			"a": "b",
			"u": "v",
		},
	})

	suite.Len(changes.UpdateNew, 1)
	suite.Equal(*changes.UpdateNew[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1"},
		RecordType: "A",
		RecordTTL:  300,
		Labels: map[string]string{
			"a": "c",
			"u": "v",
		},
	})
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
}

func (suite *ResolverSuite) TestStrictResolveUpdateSimpleUpdateProviderSpecific() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "a",
					Value: "b",
				},
				{
					Name:  "puh",
					Value: "luh",
				},
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			ProviderSpecific: endpoint.ProviderSpecific{
				{
					Name:  "a",
					Value: "c",
				},
				{
					Name:  "hac",
					Value: "sma",
				},
			},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(*changes.UpdateOld[0], *current[0])
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(*changes.UpdateNew[0], *candidates[0])
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
}

func (suite *ResolverSuite) TestStrictResolveUpdateErrorRecordTTL() {
	current := []*endpoint.Endpoint{}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.3"},
			RecordType: "A",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  400,
		},
	}
	_, err := suite.perResource.Resolve(current, candidates)
	suite.ErrorContains(err, "inconsistent RecordTTL")
}
func (suite *ResolverSuite) TestStrictResolveUpdateErrorSetIdentifier() {
	current := []*endpoint.Endpoint{}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.3"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "server",
		},
		{
			DNSName:       "wl.example.com",
			Targets:       endpoint.Targets{"1.1.1.1"},
			RecordType:    "A",
			RecordTTL:     300,
			SetIdentifier: "x1",
		},
	}
	_, err := suite.perResource.Resolve(current, candidates)
	suite.ErrorContains(err, "cannot append endpoint")
}

func (suite *ResolverSuite) TestStrictResolveUpdateMixed() {
	// we delete other Types (is this right)
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"fd00::4"},
			RecordType: "AAAA",
			RecordTTL:  300,
		},
	}
	_, err := suite.perResource.Resolve(current, candidates)
	suite.ErrorContains(err, "cannot append endpoint")
	// suite.Len(changes.UpdateNew, 0)
	// suite.Len(changes.UpdateOld, 0)
	// suite.Len(changes.Delete, 1)
	// suite.Equal(*changes.Delete[0], *current[0])
	// suite.Len(changes.Create, 1)
	// suite.Equal(*changes.Create[0], *candidates[0])
}

func (suite *ResolverSuite) TestStrictResolveUpdateMultipleCandidateTargets() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2", "1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(*changes.UpdateOld[0], *current[0])
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(*changes.UpdateNew[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2"},
		RecordType: "A",
		RecordTTL:  300,
	})
}

func (suite *ResolverSuite) TestStrictResolveUpdateSplitCandidateTargets() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
	suite.Len(changes.UpdateOld, 1)
	suite.Equal(*changes.UpdateOld[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1"},
		RecordType: "A",
		RecordTTL:  300,
	})
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(*changes.UpdateNew[0], endpoint.Endpoint{
		DNSName:    "wl.example.com",
		Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2"},
		RecordType: "A",
		RecordTTL:  300,
	})
}

func (suite *ResolverSuite) TestStrictResolveUpdateOldOnly() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"winner": "winner",
				// "foo":    "bar",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"fix":    "box",
				"winner": "looser",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"fix":    "box",
				"stay":   "here",
				"winner": "looser2",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"fix":    "box",
				"stay":   "here",
				"winner": "looser2",
				// "foo":    "bar",
			},
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
	suite.Len(changes.UpdateOld, 2)
	suite.Equal(changes.UpdateOld, []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"winner": "winner",
				// "foo":    "bar",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"fix":    "box",
				"winner": "looser",
			},
		},
	})

	suite.Len(changes.UpdateNew, 0)

}

func (suite *ResolverSuite) TestStrictResolveUpdateOldMerge() {
	// empty current and candidates -> create == empty, create == empty, delete == empty
	current := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"winner": "winner",
				"foo":    "bar",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1", "1.1.1.2"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"fix":    "box",
				"winner": "looser",
			},
		},
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"stay":   "here",
				"winner": "looser2",
			},
		},
	}
	candidates := []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
		},
	}
	changes, err := suite.perResource.Resolve(current, candidates)
	suite.Nil(err)
	suite.Len(changes.Delete, 0)
	suite.Len(changes.Create, 0)
	suite.Len(changes.UpdateNew, 1)
	suite.Equal(changes.UpdateNew, []*endpoint.Endpoint{
		{
			DNSName:    "wl.example.com",
			Targets:    endpoint.Targets{"1.1.1.1"},
			RecordType: "A",
			RecordTTL:  300,
			Labels: map[string]string{
				"winner": "looser", //due to target sort 1.1.1.2
				"stay":   "here",
				"foo":    "bar",
				"fix":    "box",
			},
		},
	})
	suite.Len(changes.UpdateOld, 3)
	suite.Equal(changes.UpdateOld, current)
}

func TestConflictResolver(t *testing.T) {
	suite.Run(t, new(ResolverSuite))
}
