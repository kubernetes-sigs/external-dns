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

package endpoint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LabelsSuite struct {
	suite.Suite
	foo                  Labels
	fooAsText            string
	fooAsTextWithQuotes  string
	barText              string
	barTextAsMap         Labels
	noHeritageText       string
	wrongHeritageText    string
	multipleHeritageText string // considered invalid
}

func (suite *LabelsSuite) SetupTest() {
	suite.foo = map[string]string{
		"owner":    "foo-owner",
		"resource": "foo-resource",
	}
	suite.fooAsText = "heritage=external-dns,external-dns/owner=foo-owner,external-dns/resource=foo-resource"
	suite.fooAsTextWithQuotes = fmt.Sprintf(`"%s"`, suite.fooAsText)

	suite.barTextAsMap = map[string]string{
		"owner":    "bar-owner",
		"resource": "bar-resource",
		"new-key":  "bar-new-key",
	}
	suite.barText = "heritage=external-dns,,external-dns/owner=bar-owner,external-dns/resource=bar-resource,external-dns/new-key=bar-new-key,random=stuff,no-equal-sign,," // also has some random gibberish

	suite.noHeritageText = "external-dns/owner=random-owner"
	suite.wrongHeritageText = "heritage=mate,external-dns/owner=random-owner"
	suite.multipleHeritageText = "heritage=mate,heritage=external-dns,external-dns/owner=random-owner"
}

func (suite *LabelsSuite) TestSerialize() {
	suite.Equal(suite.fooAsText, suite.foo.Serialize(false), "should serializeLabel")
	suite.Equal(suite.fooAsTextWithQuotes, suite.foo.Serialize(true), "should serializeLabel")
}

func (suite *LabelsSuite) TestDeserialize() {
	foo, err := NewLabelsFromString(suite.fooAsText)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.foo, foo, "should reconstruct original label map")

	foo, err = NewLabelsFromString(suite.fooAsTextWithQuotes)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.foo, foo, "should reconstruct original label map")

	bar, err := NewLabelsFromString(suite.barText)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.barTextAsMap, bar, "should reconstruct original label map")

	noHeritage, err := NewLabelsFromString(suite.noHeritageText)
	suite.Equal(ErrInvalidHeritage, err, "should fail if no heritage is found")
	suite.Nil(noHeritage, "should return nil")

	wrongHeritage, err := NewLabelsFromString(suite.wrongHeritageText)
	suite.Equal(ErrInvalidHeritage, err, "should fail if wrong heritage is found")
	suite.Nil(wrongHeritage, "if error should return nil")

	multipleHeritage, err := NewLabelsFromString(suite.multipleHeritageText)
	suite.Equal(ErrInvalidHeritage, err, "should fail if multiple heritage is found")
	suite.Nil(multipleHeritage, "if error should return nil")
}

func TestLabels(t *testing.T) {
	suite.Run(t, new(LabelsSuite))
}
