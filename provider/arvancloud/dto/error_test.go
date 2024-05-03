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

package dto

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArvanCloud_NewApiTokenFromFileError(t *testing.T) {
	t.Run("Should make error when can't open file for read api token", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewApiTokenFromFileError(errors.New("fail to open"))

		assert.Error(t, err)
		assert.EqualError(t, err, "failed to read api token from file. (action: api-token-from-file): fail to open")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ApiTokenFromFileActErr, providerError.GetAction())
		assert.Equal(t, ApiTokenFromFileMsgErr, providerError.GetMessage())
		assert.Equal(t, false, providerError.IsOperational())
	})
}

func TestArvanCloud_NewNewApiVersionError(t *testing.T) {
	t.Run("Should make error when api version not valid", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewApiVersionError()

		assert.Error(t, err)
		assert.EqualError(t, err, "the Arvan cloud api version is not match with semantic version, Please fill with X.Y (like 4.0) (action: api-version)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ApiVersionActErr, providerError.GetAction())
		assert.Equal(t, ApiVersionMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}

func TestArvanCloud_NewApiTokenRequireError(t *testing.T) {
	t.Run("Should make error when api token is empty", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewApiTokenRequireError()

		assert.Error(t, err)
		assert.EqualError(t, err, "the Arvan cloud api token is required (action: api-token-require)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ApiTokenRequireActErr, providerError.GetAction())
		assert.Equal(t, ApiTokenRequireMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}

func TestArvanCloud_NewUnknownError(t *testing.T) {
	t.Run("Should make error when unknown error is happened", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewUnknownError(errors.New("error is happened"))

		assert.Error(t, err)
		assert.EqualError(t, err, "unknown error happened (action: unknown): error is happened")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, UnknownActErr, providerError.GetAction())
		assert.Equal(t, UnknownMsgErr, providerError.GetMessage())
		assert.Equal(t, false, providerError.IsOperational())
	})
}

func TestArvanCloud_NewParseRecordError(t *testing.T) {
	t.Run("Should make error when parse record", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewParseRecordError()

		assert.Error(t, err)
		assert.EqualError(t, err, "fail to parse data record to object (action: parse-record)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ParseRecordActErr, providerError.GetAction())
		assert.Equal(t, ParseRecordMsgErr, providerError.GetMessage())
		assert.Equal(t, false, providerError.IsOperational())
	})
}

func TestArvanCloud_NewUnauthorizedError(t *testing.T) {
	t.Run("Should make error when unauthorized error is happened", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewUnauthorizedError()

		assert.Error(t, err)
		assert.EqualError(t, err, "you must authenticate yourself to get the requested response (action: unauthorized)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, UnauthorizedActErr, providerError.GetAction())
		assert.Equal(t, UnauthorizedMsgErr, providerError.GetMessage())
		assert.Equal(t, false, providerError.IsOperational())
	})
}

func TestArvanCloud_NewNonPointerError(t *testing.T) {
	t.Run("Should make error when use non pointer", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewNonPointerError()

		assert.Error(t, err)
		assert.EqualError(t, err, "attempt to decode into a non-pointer (action: non-pointer-type)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, NonPointerActErr, providerError.GetAction())
		assert.Equal(t, NonPointerMsgErr, providerError.GetMessage())
		assert.Equal(t, false, providerError.IsOperational())
	})
}

func TestArvanCloud_NewParseMXRecordError(t *testing.T) {
	t.Run("Should make error when parse MX record", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewParseMXRecordError()

		assert.Error(t, err)
		assert.EqualError(t, err, "fail to parse mx record (action: parse-mx-record)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ParseMXRecordActErr, providerError.GetAction())
		assert.Equal(t, ParseMXRecordMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}

func TestArvanCloud_NewParseSRVRecordError(t *testing.T) {
	t.Run("Should make error when parse SRV record", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewParseSRVRecordError()

		assert.Error(t, err)
		assert.EqualError(t, err, "fail to parse srv record (action: parse-rsv-record)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ParseSRVRecordActErr, providerError.GetAction())
		assert.Equal(t, ParseSRVRecordMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}

func TestArvanCloud_NewParseTLSARecordError(t *testing.T) {
	t.Run("Should make error when parse TLSA record", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewParseTLSARecordError()

		assert.Error(t, err)
		assert.EqualError(t, err, "fail to parse tlsa record (action: parse-tlsa-record)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ParseTLSARecordActErr, providerError.GetAction())
		assert.Equal(t, ParseTLSARecordMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}

func TestArvanCloud_NewNewCAARecordError(t *testing.T) {
	t.Run("Should make error when parse CAA record", func(t *testing.T) {
		providerError := &ProviderError{}

		err := NewCAARecordError()

		assert.Error(t, err)
		assert.EqualError(t, err, "fail to parse caa record (action: parse-caa-record)")
		assert.ErrorAs(t, err, &providerError)
		assert.Equal(t, ParseCAARecordActErr, providerError.GetAction())
		assert.Equal(t, ParseCAARecordMsgErr, providerError.GetMessage())
		assert.Equal(t, true, providerError.IsOperational())
	})
}
