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

package http

import (
	"bytes"
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
	DrainAndClose(rc)
	assert.True(t, rc.closed)

	// Confirm body is fully drained: reader should be at EOF.
	n, err := rc.Read(make([]byte, 1))
	assert.Equal(t, 0, n)
	assert.ErrorIs(t, err, io.EOF)
}

func TestDrainAndClose_EmptyBody(t *testing.T) {
	rc := &trackingReadCloser{Reader: strings.NewReader("")}
	DrainAndClose(rc)
	assert.True(t, rc.closed)
}

func TestDrainAndClose_OversizedBody(t *testing.T) {
	// Body is 1 byte larger than the cap; the excess byte must remain unread so
	// the connection is discarded rather than pooled — but Close must still be called.
	oversized := bytes.Repeat([]byte("x"), drainMaxBytes+1)
	rc := &trackingReadCloser{Reader: bytes.NewReader(oversized)}
	DrainAndClose(rc)
	assert.True(t, rc.closed)

	// Exactly one byte should remain after the capped drain.
	remaining, err := io.ReadAll(rc.Reader)
	assert.NoError(t, err)
	assert.Len(t, remaining, 1, "expected exactly 1 byte past the drain cap to remain unread")
}
