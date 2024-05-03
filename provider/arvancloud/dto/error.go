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
	"strings"
)

const (
	ApiTokenFromFileActErr = "api-token-from-file"
	ApiTokenFromFileMsgErr = "failed to read api token from file."
	ApiTokenRequireActErr  = "api-token-require"
	ApiTokenRequireMsgErr  = "the Arvan cloud api token is required"
	ApiVersionActErr       = "api-version"
	ApiVersionMsgErr       = "the Arvan cloud api version is not match with semantic version, Please fill with X.Y (like 4.0)"
	UnauthorizedActErr     = "unauthorized"
	UnauthorizedMsgErr     = "you must authenticate yourself to get the requested response"
	UnknownActErr          = "unknown"
	UnknownMsgErr          = "unknown error happened"
	ParseRecordActErr      = "parse-record"
	ParseRecordMsgErr      = "fail to parse data record to object"
	NonPointerActErr       = "non-pointer-type"
	NonPointerMsgErr       = "attempt to decode into a non-pointer"
	ParseMXRecordActErr    = "parse-mx-record"
	ParseMXRecordMsgErr    = "fail to parse mx record"
	ParseSRVRecordActErr   = "parse-rsv-record"
	ParseSRVRecordMsgErr   = "fail to parse srv record"
	ParseTLSARecordActErr  = "parse-tlsa-record"
	ParseTLSARecordMsgErr  = "fail to parse tlsa record"
	ParseCAARecordActErr   = "parse-caa-record"
	ParseCAARecordMsgErr   = "fail to parse caa record"
)

type ProviderError struct {
	err         error
	action      string
	msg         string
	operational bool
}

var _ error = &ProviderError{}

func (b *ProviderError) Error() string {
	var sb strings.Builder
	sb.WriteString(b.msg)
	sb.WriteString(" (action: ")
	sb.WriteString(b.action)
	sb.WriteString(")")

	if b.err != nil {
		sb.WriteString(": ")
		sb.WriteString(b.err.Error())
	}

	return sb.String()
}

func (b *ProviderError) GetAction() string {
	return b.action
}

func (b *ProviderError) GetMessage() string {
	return b.msg
}

func (b *ProviderError) IsOperational() bool {
	return b.operational
}

func NewBusinessError(msg string, action string) *ProviderError {
	return &ProviderError{msg: msg, action: action, operational: true}
}

func NewInfraError(err error, msg string, action string) *ProviderError {
	return &ProviderError{err: err, msg: msg, action: action, operational: false}
}

func NewApiTokenFromFileError(err error) *ProviderError {
	return NewInfraError(err, ApiTokenFromFileMsgErr, ApiTokenFromFileActErr)
}

func NewApiVersionError() *ProviderError {
	return NewBusinessError(ApiVersionMsgErr, ApiVersionActErr)
}

func NewApiTokenRequireError() *ProviderError {
	return NewBusinessError(ApiTokenRequireMsgErr, ApiTokenRequireActErr)
}

func NewUnknownError(err error) *ProviderError {
	return NewInfraError(err, UnknownMsgErr, UnknownActErr)
}

func NewParseRecordError() *ProviderError {
	return NewInfraError(nil, ParseRecordMsgErr, ParseRecordActErr)
}

func NewUnauthorizedError() *ProviderError {
	return NewInfraError(nil, UnauthorizedMsgErr, UnauthorizedActErr)
}

func NewNonPointerError() *ProviderError {
	return NewInfraError(nil, NonPointerMsgErr, NonPointerActErr)
}

func NewParseMXRecordError() *ProviderError {
	return NewBusinessError(ParseMXRecordMsgErr, ParseMXRecordActErr)
}

func NewParseSRVRecordError() *ProviderError {
	return NewBusinessError(ParseSRVRecordMsgErr, ParseSRVRecordActErr)
}

func NewParseTLSARecordError() *ProviderError {
	return NewBusinessError(ParseTLSARecordMsgErr, ParseTLSARecordActErr)
}

func NewCAARecordError() *ProviderError {
	return NewBusinessError(ParseCAARecordMsgErr, ParseCAARecordActErr)
}
