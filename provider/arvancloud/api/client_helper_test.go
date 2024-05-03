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

package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RoundTripFunc func(req *http.Request) (*http.Response, error)

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func newHttpClientMock(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

var _ json.Marshaler = (*inputDataNewRequestInvalid)(nil)

type inputDataNewRequestInvalid struct {
	ID int
}

func (i inputDataNewRequestInvalid) MarshalJSON() ([]byte, error) {
	return nil, errors.New("invalid data")
}

type inputDataNewRequestValid struct {
	ID int
}

type outputDataUnmarshalResponseValid struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type outputReaderUnmarshalResponseInvalid struct{}

func (r *outputReaderUnmarshalResponseInvalid) Read(p []byte) (n int, err error) {
	return 0, errors.New("response body read error")
}

type mockReadCloserErr struct{}

func (m *mockReadCloserErr) Read(p []byte) (n int, err error) {
	return 0, errors.New("fail to read data")
}

func (m *mockReadCloserErr) Close() error {
	return errors.New("fail to close reader")
}
