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
	"context"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

// RequestWithRetry sends req, retrying connection errors and 5xx responses
// with the given backoff. Any other non-2xx status is returned as an error at
// once, since repeating the request would not change the answer. On success
// the caller owns the response body.
func RequestWithRetry(client *http.Client, req *http.Request, backoff wait.Backoff) (*http.Response, error) {
	var resp *http.Response
	var lastErr error
	err := wait.ExponentialBackoffWithContext(req.Context(), backoff, func(context.Context) (bool, error) {
		// Reset the body before each attempt so retries send the full payload.
		// GetBody is set automatically by http.NewRequest for in-memory body types;
		// it is nil for GET requests (no body), so the block is safely skipped.
		if req.GetBody != nil {
			body, err := req.GetBody()
			if err != nil {
				return false, fmt.Errorf("failed to reset request body: %w", err)
			}
			req.Body = body
		}

		r, err := client.Do(req)
		if err != nil {
			log.Debugf("Request to %s failed: %v", req.URL, err)
			lastErr = err
			return false, nil
		}
		// 5xx: retryable server error
		if r.StatusCode >= http.StatusInternalServerError {
			DrainAndClose(r.Body)
			lastErr = fmt.Errorf("server error: status code %d", r.StatusCode)
			return false, nil
		}
		// 3xx/4xx: permanent error, not worth retrying. A 3xx shows up here only
		// when the client was told not to follow redirects or for a non-redirect
		// status such as 304.
		if r.StatusCode >= http.StatusMultipleChoices {
			DrainAndClose(r.Body)
			return false, fmt.Errorf("unexpected status code %d", r.StatusCode)
		}
		resp = r
		return true, nil
	})
	if wait.Interrupted(err) {
		if cause := context.Cause(req.Context()); cause != nil {
			return nil, cause
		}
		if lastErr != nil {
			return nil, lastErr
		}
	}
	return resp, err
}
