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
	"io"
)

const (
	// drainMaxBytes caps how much of a response body we drain to return the
	// connection to the pool. A buggy or adversarial server could stream an
	// unbounded body; reading it all would block indefinitely and waste memory.
	// On success paths the JSON decoder has already consumed the payload before
	// the deferred DrainAndClose runs, so only trailing bytes remain. On error paths
	// the body is typically a short error message. 1 MiB is generous for either case.
	drainMaxBytes = 1 << 20 // 1 MiB
)

// DrainAndClose drains up to drainMaxBytes of the response body before
// closing it so the underlying TCP connection can be reused by the HTTP
// client's connection pool. Bytes beyond the cap are left unread; the
// connection will be discarded rather than pooled in that case, which is
// acceptable for oversized or malformed responses.
func DrainAndClose(body io.ReadCloser) {
	_, _ = io.Copy(io.Discard, io.LimitReader(body, drainMaxBytes))
	_ = body.Close()
}
