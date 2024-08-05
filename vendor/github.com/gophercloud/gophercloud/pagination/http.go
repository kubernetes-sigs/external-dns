package pagination

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/gophercloud/gophercloud"
)

// PageResult stores the HTTP response that returned the current page of results.
type PageResult struct {
	gophercloud.Result
	url.URL
}

// PageResultFrom parses an HTTP response as JSON and returns a PageResult containing the
// results, interpreting it as JSON if the content type indicates.
func PageResultFrom(resp *http.Response) (PageResult, error) {
	var parsedBody interface{}

	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return PageResult{}, err
	}

	if strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		err = json.Unmarshal(rawBody, &parsedBody)
		if err != nil {
			return PageResult{}, err
		}
	} else {
		parsedBody = rawBody
	}

	return PageResultFromParsed(resp, parsedBody), err
}

// PageResultFromParsed constructs a PageResult from an HTTP response that has already had its
// body parsed as JSON (and closed).
func PageResultFromParsed(resp *http.Response, body interface{}) PageResult {
	return PageResult{
		Result: gophercloud.Result{
			Body:       body,
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
		},
		URL: *resp.Request.URL,
	}
}

// Request performs an HTTP request and extracts the http.Response from the result.
func Request(client *gophercloud.ServiceClient, headers map[string]string, url string) (*http.Response, error) {
	return client.Get(url, nil, &gophercloud.RequestOpts{
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		MoreHeaders:      headers,
		OkCodes:          []int{200, 204, 300},
		KeepResponseBody: true,
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
=======
		MoreHeaders:      headers,
		OkCodes:          []int{200, 204, 300},
		KeepResponseBody: true,
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
=======
		MoreHeaders:      headers,
		OkCodes:          []int{200, 204, 300},
		KeepResponseBody: true,
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
=======
		MoreHeaders:      headers,
		OkCodes:          []int{200, 204, 300},
		KeepResponseBody: true,
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		MoreHeaders: headers,
		OkCodes:     []int{200, 204, 300},
=======
		MoreHeaders:      headers,
		OkCodes:          []int{200, 204, 300},
		KeepResponseBody: true,
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	})
}
