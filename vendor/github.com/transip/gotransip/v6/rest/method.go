package rest

// Method is a struct which is used by the client to present a HTTP method of choice
// and the expected return status codes on which you can check to see if the response is correct,
// thus not an error.
type Method struct {
	// Method is where a HTTP method like "GET" would go
	Method string
	// ExpectedStatusCodes are the expected status codes with which
	// you can check if the response status code is correct
	ExpectedStatusCodes []int
}

var (
	// GetMethod is a wrapper with expected status codes around the HTTP "GET" method
	GetMethod = Method{Method: "GET", ExpectedStatusCodes: []int{200}}
	// PostMethod is a wrapper with expected status codes around the HTTP "POST" method
	PostMethod = Method{Method: "POST", ExpectedStatusCodes: []int{200, 201}}
	// PutMethod is a wrapper with expected status codes around the HTTP "PUT" method
	PutMethod = Method{Method: "PUT", ExpectedStatusCodes: []int{204}}
	// PatchMethod is a wrapper with expected status codes around the HTTP "PATCH" method
	PatchMethod = Method{Method: "PATCH", ExpectedStatusCodes: []int{204}}
	// DeleteMethod is a wrapper with expected status codes around the HTTP "DELETE" method
	DeleteMethod = Method{Method: "DELETE", ExpectedStatusCodes: []int{204}}
)

// StatusCodeOK returns true when the status code is correct
// This method used by the rest client to check if the given status code is correct.
func (r *Method) StatusCodeOK(statusCode int) bool {
	return contains(r.ExpectedStatusCodes, statusCode)
}

// contains is used to see if a certain value is part of an array
func contains(haystack []int, needle int) bool {
	for _, a := range haystack {
		if a == needle {
			return true
		}
	}
	return false
}
