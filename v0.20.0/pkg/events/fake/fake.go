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

package fake

import (
	"github.com/stretchr/testify/mock"

	"sigs.k8s.io/external-dns/pkg/events"
)

type EventEmitter struct {
	mock.Mock
}

func (m *EventEmitter) Add(events ...events.Event) {
	m.Called(events[0])
}

func NewFakeEventEmitter() *EventEmitter {
	m := &EventEmitter{}
	m.On("Add", mock.AnythingOfType("events.Event"))
	return m
}
