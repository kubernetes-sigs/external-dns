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

package registry

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
	foo                  map[string]string
	fooAsText            string
	fooAsTextWithQuotes  string
	barText              string
	barTextAsMap         map[string]string
	noHeritageText       string
	noHeritageAsMap      map[string]string
	wrongHeritageText    string
	multipleHeritageText string //considered invalid
}

func (suite *ParserSuite) SetupTest() {
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
	suite.barText = "heritage=external-dns,,external-dns/owner=bar-owner,external-dns/resource=bar-resource,external-dns/new-key=bar-new-key,random=stuff,no-equal-sign,," //also has some random gibberish

	suite.noHeritageText = "external-dns/owner=random-owner"
	suite.wrongHeritageText = "heritage=mate,external-dns/owner=random-owner"
	suite.multipleHeritageText = "heritage=mate,heritage=external-dns,external-dns/owner=random-owner"
}

func (suite *ParserSuite) TestSerialize() {
	suite.Equal(suite.fooAsText, serialize(suite.foo, false), "should serialize")
	suite.Equal(suite.fooAsTextWithQuotes, serialize(suite.foo, true), "should serialize")
}

func (suite *ParserSuite) TestDeserialize() {
	foo, err := deserialize(suite.fooAsText)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.foo, foo, "should reconstruct original label map")

	foo, err = deserialize(suite.fooAsTextWithQuotes)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.foo, foo, "should reconstruct original label map")

	bar, err := deserialize(suite.barText)
	suite.NoError(err, "should succeed for valid label text")
	suite.Equal(suite.barTextAsMap, bar, "should reconstruct original label map")

	noHeritage, err := deserialize(suite.noHeritageText)
	suite.Equal(invalidHeritage, err, "should fail if no heritage is found")
	suite.Nil(noHeritage, "should return nil")

	wrongHeritage, err := deserialize(suite.wrongHeritageText)
	suite.Equal(invalidHeritage, err, "should fail if wrong heritage is found")
	suite.Nil(wrongHeritage, "if error should return nil")

	multipleHeritage, err := deserialize(suite.multipleHeritageText)
	suite.Equal(invalidHeritage, err, "should fail if multiple heritage is found")
	suite.Nil(multipleHeritage, "if error should return nil")
}

func TestParser(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}
