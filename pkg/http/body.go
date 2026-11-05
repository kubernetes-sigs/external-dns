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
	"errors"
	"io"
)

// ErrResponseTooLarge distinguishes an over-limit body from a malformed or
// truncated one, which io.LimitReader would report as a plain EOF.
var ErrResponseTooLarge = errors.New("webhook response body exceeds the configured max size")

// LimitBody caps response body reads at maxBodySize, failing with
// ErrResponseTooLarge once exceeded; a non-positive limit disables the cap.
func LimitBody(r io.Reader, maxBodySize int64) io.Reader {
	if maxBodySize <= 0 {
		return r
	}
	return &maxBytesReader{r: r, remaining: maxBodySize}
}

// maxBytesReader is a client-side http.MaxBytesReader: it errors on the byte
// past the limit instead of reporting EOF like io.LimitReader.
type maxBytesReader struct {
	r         io.Reader
	remaining int64 // bytes still allowed; negative once the limit is passed
}

func (l *maxBytesReader) Read(p []byte) (int, error) {
	if l.remaining < 0 {
		return 0, ErrResponseTooLarge
	}
	// Read one byte past the limit to detect an over-limit body; compare
	// against remaining so the +1 can't overflow near math.MaxInt64.
	if l.remaining < int64(len(p)) {
		p = p[:l.remaining+1]
	}
	n, err := l.r.Read(p)
	l.remaining -= int64(n)
	if l.remaining < 0 {
		return n, ErrResponseTooLarge
	}
	return n, err
}
