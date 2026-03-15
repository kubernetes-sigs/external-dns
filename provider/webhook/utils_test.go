/*
Copyright 2026 The Kubernetes Authors.

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

package webhook

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type trackingReadCloser struct {
	io.Reader
	closed bool
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

func TestDrainAndClose_DrainsThenCloses(t *testing.T) {
	rc := &trackingReadCloser{Reader: strings.NewReader("remaining body data")}
	drainAndClose(rc)
	assert.True(t, rc.closed)

	// Confirm body is fully drained: reader should be at EOF.
	n, err := rc.Read(make([]byte, 1))
	assert.Equal(t, 0, n)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDrainAndClose_EmptyBody(t *testing.T) {
	rc := &trackingReadCloser{Reader: strings.NewReader("")}
	drainAndClose(rc)
	assert.True(t, rc.closed)
}
