package core

// (C) Copyright IBM Corp. 2019.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// common HTTP methods
const (
	POST   = http.MethodPost
	GET    = http.MethodGet
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
	HEAD   = http.MethodHead
)

// common headers
const (
	Accept                  = "Accept"
	APPLICATION_JSON        = "application/json"
	CONTENT_DISPOSITION     = "Content-Disposition"
	CONTENT_ENCODING        = "Content-Encoding"
	CONTENT_TYPE            = "Content-Type"
	FORM_URL_ENCODED_HEADER = "application/x-www-form-urlencoded"

	ERRORMSG_SERVICE_URL_MISSING = "service URL is empty"
	ERRORMSG_SERVICE_URL_INVALID = "error parsing service URL: %s"
	ERRORMSG_PATH_PARAM_EMPTY    = "path parameter '%s' is empty"
)

// FormData stores information for form data.
type FormData struct {
	fileName    string
	contentType string
	contents    interface{}
}

// RequestBuilder is used to build an HTTP Request instance.
type RequestBuilder struct {
	Method string
	URL    *url.URL
	Header http.Header
	Body   io.Reader
	Query  map[string][]string
	Form   map[string][]FormData

	// EnableGzipCompression indicates whether or not request bodies
	// should be gzip-compressed.
	// This field has no effect on response bodies.
	// If enabled, the Body field will be gzip-compressed and
	// the "Content-Encoding" header will be added to the request with the
	// value "gzip".
	EnableGzipCompression bool

	// RequestContext is an optional Context instance to be associated with the
	// http.Request that is constructed by the Build() method.
	ctx context.Context
}

// NewRequestBuilder initiates a new request.
func NewRequestBuilder(method string) *RequestBuilder {
	return &RequestBuilder{
		Method: method,
		Header: make(http.Header),
		Query:  make(map[string][]string),
		Form:   make(map[string][]FormData),
	}
}

// WithContext sets "ctx" as the Context to be associated with
// the http.Request instance that will be constructed by the Build() method.
func (requestBuilder *RequestBuilder) WithContext(ctx context.Context) *RequestBuilder {
	requestBuilder.ctx = ctx
	return requestBuilder
}

// ConstructHTTPURL creates a properly-encoded URL with path parameters.
// This function returns an error if the serviceURL is "" or is an
// invalid URL string (e.g. ":<badscheme>").
func (requestBuilder *RequestBuilder) ConstructHTTPURL(serviceURL string, pathSegments []string, pathParameters []string) (*RequestBuilder, error) {
	if serviceURL == "" {
		return requestBuilder, fmt.Errorf(ERRORMSG_SERVICE_URL_MISSING)
	}
	var URL *url.URL

	URL, err := url.Parse(serviceURL)
	if err != nil {
		return requestBuilder, fmt.Errorf(ERRORMSG_SERVICE_URL_INVALID, err.Error())
	}

	for i, pathSegment := range pathSegments {
		if pathSegment != "" {
			URL.Path += "/" + pathSegment
		}

		if pathParameters != nil && i < len(pathParameters) {
			if pathParameters[i] == "" {
				return requestBuilder, fmt.Errorf(ERRORMSG_PATH_PARAM_EMPTY, fmt.Sprintf("[%d]", i))
			}
			URL.Path += "/" + pathParameters[i]
		}
	}
	requestBuilder.URL = URL
	return requestBuilder, nil
}

//
// ResolveRequestURL creates a properly-encoded URL with path params.
// This function returns an error if the serviceURL is "" or is an
// invalid URL string (e.g. ":<badscheme>").
// Parameters:
// serviceURL - the base URL associated with the service endpoint (e.g. "https://myservice.cloud.ibm.com")
// path - the unresolved path string (e.g. "/resource/{resource_id}/type/{type_id}")
// pathParams - a map containing the path params, keyed by the path param base name
// (e.g. {"type_id": "type-1", "resource_id": "res-123-456-789-abc"})
// The resulting request URL: "https://myservice.cloud.ibm.com/resource/res-123-456-789-abc/type/type-1"
//
func (requestBuilder *RequestBuilder) ResolveRequestURL(serviceURL string, path string, pathParams map[string]string) (*RequestBuilder, error) {
	if serviceURL == "" {
		return requestBuilder, fmt.Errorf(ERRORMSG_SERVICE_URL_MISSING)
	}

	urlString := serviceURL

	// If we have a non-empty "path" input parameter, then process it for possible path param references.
	if path != "" {

		// If path parameter values were passed in, then for each one, replace any references to it
		// within "path" with the path parameter's encoded value.
		if len(pathParams) > 0 {
			for k, v := range pathParams {
				if v == "" {
					return requestBuilder, fmt.Errorf(ERRORMSG_PATH_PARAM_EMPTY, k)
				}
				encodedValue := url.PathEscape(v)
				ref := fmt.Sprintf("{%s}", k)
				path = strings.ReplaceAll(path, ref, encodedValue)
			}
		}

		// Next, we need to append "path" to "urlString".
		// We need to pay particular attention to any trailing slash on "urlString" and
		// a leading slash on "path".  Ultimately, we do not want a double slash.
		if strings.HasSuffix(urlString, "/") {
			// If urlString has a trailing slash, then make sure path does not have a leading slash.
			path = strings.TrimPrefix(path, "/")
		} else {
			// If urlString does not have a trailing slash and path does not have a
			// leading slash, then append a slash to urlString.
			if !strings.HasPrefix(path, "/") {
				urlString += "/"
			}
		}

		urlString += path
	}

	var URL *url.URL

	URL, err := url.Parse(urlString)
	if err != nil {
		return requestBuilder, fmt.Errorf(ERRORMSG_SERVICE_URL_INVALID, err.Error())
	}

	requestBuilder.URL = URL
	return requestBuilder, nil
}

// AddQuery adds a query parameter name and value to the request.
func (requestBuilder *RequestBuilder) AddQuery(name string, value string) *RequestBuilder {
	requestBuilder.Query[name] = append(requestBuilder.Query[name], value)
	return requestBuilder
}

// AddHeader adds a header name and value to the request.
func (requestBuilder *RequestBuilder) AddHeader(name string, value string) *RequestBuilder {
	requestBuilder.Header[name] = []string{value}
	return requestBuilder
}

// AddFormData adds a new mime part (constructed from the input parameters)
// to the request's multi-part form.
func (requestBuilder *RequestBuilder) AddFormData(fieldName string, fileName string, contentType string,
	contents interface{}) *RequestBuilder {
	if fileName == "" {
		if file, ok := contents.(*os.File); ok {
			if !((os.File{}) == *file) { // if file is not empty
				name := filepath.Base(file.Name())
				fileName = name
			}
		}
	}
	requestBuilder.Form[fieldName] = append(requestBuilder.Form[fieldName], FormData{
		fileName:    fileName,
		contentType: contentType,
		contents:    contents,
	})
	return requestBuilder
}

// SetBodyContentJSON sets the body content from a JSON structure.
func (requestBuilder *RequestBuilder) SetBodyContentJSON(bodyContent interface{}) (*RequestBuilder, error) {
	requestBuilder.Body = new(bytes.Buffer)
	err := json.NewEncoder(requestBuilder.Body.(io.Writer)).Encode(bodyContent)
	return requestBuilder, err
}

// SetBodyContentString sets the body content from a string.
func (requestBuilder *RequestBuilder) SetBodyContentString(bodyContent string) (*RequestBuilder, error) {
	requestBuilder.Body = strings.NewReader(bodyContent)
	return requestBuilder, nil
}

// SetBodyContentStream sets the body content from an io.Reader instance.
func (requestBuilder *RequestBuilder) SetBodyContentStream(bodyContent io.Reader) (*RequestBuilder, error) {
	requestBuilder.Body = bodyContent
	return requestBuilder, nil
}

// CreateMultipartWriter initializes a new multipart writer.
func (requestBuilder *RequestBuilder) createMultipartWriter() *multipart.Writer {
	buff := new(bytes.Buffer)
	requestBuilder.Body = buff
	return multipart.NewWriter(buff)
}

// CreateFormFile is a convenience wrapper around CreatePart. It creates
// a new form-data header with the provided field name and file name and contentType.
func createFormFile(formWriter *multipart.Writer, fieldname string, filename string, contentType string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	contentDisposition := fmt.Sprintf(`form-data; name="%s"`, fieldname)
	if filename != "" {
		contentDisposition += fmt.Sprintf(`; filename="%s"`, filename)
	}

	h.Set(CONTENT_DISPOSITION, contentDisposition)
	if contentType != "" {
		h.Set(CONTENT_TYPE, contentType)
	}

	return formWriter.CreatePart(h)
}

// SetBodyContentForMultipart sets the body content for a part in a multi-part form.
func (requestBuilder *RequestBuilder) SetBodyContentForMultipart(contentType string, content interface{}, writer io.Writer) error {
	var err error
	if stream, ok := content.(io.Reader); ok {
		_, err = io.Copy(writer, stream)
	} else if stream, ok := content.(*io.ReadCloser); ok {
		_, err = io.Copy(writer, *stream)
	} else if IsJSONMimeType(contentType) || IsJSONPatchMimeType(contentType) {
		err = json.NewEncoder(writer).Encode(content)
	} else if str, ok := content.(string); ok {
		_, err = writer.Write([]byte(str))
	} else if strPtr, ok := content.(*string); ok {
		_, err = writer.Write([]byte(*strPtr))
	} else {
		err = fmt.Errorf("Error: unable to determine the type of 'content' provided")
	}
	return err
}

// Build builds an HTTP Request object from this RequestBuilder instance.
func (requestBuilder *RequestBuilder) Build() (req *http.Request, err error) {
	// Create multipart form data
	if len(requestBuilder.Form) > 0 {
		// handle both application/x-www-form-urlencoded or multipart/form-data
		contentType := requestBuilder.Header.Get(CONTENT_TYPE)
		if contentType == FORM_URL_ENCODED_HEADER {
			data := url.Values{}
			for fieldName, l := range requestBuilder.Form {
				for _, v := range l {
					data.Add(fieldName, v.contents.(string))
				}
			}
			_, err = requestBuilder.SetBodyContentString(data.Encode())
			if err != nil {
				return
			}
		} else {
			formWriter := requestBuilder.createMultipartWriter()
			for fieldName, l := range requestBuilder.Form {
				for _, v := range l {
					var dataPartWriter io.Writer
					dataPartWriter, err = createFormFile(formWriter, fieldName, v.fileName, v.contentType)
					if err != nil {
						return
					}
					if err = requestBuilder.SetBodyContentForMultipart(v.contentType,
						v.contents, dataPartWriter); err != nil {
						return
					}
				}
			}

			requestBuilder.AddHeader("Content-Type", formWriter.FormDataContentType())
			err = formWriter.Close()
			if err != nil {
				return
			}
		}
	}

	// If we have a request body and gzip is enabled, then wrap the body in a Gzip compression reader
	// and add the "Content-Encoding: gzip" request header.
	if !IsNil(requestBuilder.Body) && requestBuilder.EnableGzipCompression &&
		!SliceContains(requestBuilder.Header[CONTENT_ENCODING], "gzip") {
		newBody, err := NewGzipCompressionReader(requestBuilder.Body)
		if err != nil {
			return nil, err
		}
		requestBuilder.Body = newBody
		requestBuilder.Header.Add(CONTENT_ENCODING, "gzip")
	}

	// Create the request
	req, err = http.NewRequest(requestBuilder.Method, requestBuilder.URL.String(), requestBuilder.Body)
	if err != nil {
		return
	}

	// Headers
	req.Header = requestBuilder.Header

	// If "Host" was specified as a header, we need to explicitly copy it
	// to the request's Host field since the "Host" header will be ignored by Request.Write().
	host := req.Header.Get("Host")
	if host != "" {
		req.Host = host
	}

	// Query
	query := req.URL.Query()
	for k, l := range requestBuilder.Query {
		for _, v := range l {
			query.Add(k, v)
		}
	}

	// Encode query
	req.URL.RawQuery = query.Encode()

	// Finally, if a Context should be associated with the new Request instance, then set it.
	if !IsNil(requestBuilder.ctx) {
		req = req.WithContext(requestBuilder.ctx)
	}

	return
}

// SetBodyContent sets the body content from one of three different sources.
func (requestBuilder *RequestBuilder) SetBodyContent(contentType string, jsonContent interface{}, jsonPatchContent interface{},
	nonJSONContent interface{}) (builder *RequestBuilder, err error) {
	if !IsNil(jsonContent) {
		builder, err = requestBuilder.SetBodyContentJSON(jsonContent)
		if err != nil {
			return
		}
	} else if !IsNil(jsonPatchContent) {
		builder, err = requestBuilder.SetBodyContentJSON(jsonPatchContent)
		if err != nil {
			return
		}
	} else if !IsNil(nonJSONContent) {
		// Set the non-JSON body content based on the type of value passed in,
		// which should be a "string", "*string" or an "io.Reader"
		if str, ok := nonJSONContent.(string); ok {
			builder, err = requestBuilder.SetBodyContentString(str)
		} else if strPtr, ok := nonJSONContent.(*string); ok {
			builder, err = requestBuilder.SetBodyContentString(*strPtr)
		} else if stream, ok := nonJSONContent.(io.Reader); ok {
			builder, err = requestBuilder.SetBodyContentStream(stream)
		} else if stream, ok := nonJSONContent.(*io.ReadCloser); ok {
			builder, err = requestBuilder.SetBodyContentStream(*stream)
		} else {
			builder = requestBuilder
			err = fmt.Errorf("Invalid type for non-JSON body content: %s", reflect.TypeOf(nonJSONContent).String())
		}
	} else {
		builder = requestBuilder
		err = fmt.Errorf("No body content provided")
	}
	return
}

// AddQuerySlice converts the passed in slice 'slice' by calling the ConverSlice method,
// and adds the converted slice to the request's query string. An error is returned when
// conversion fails.
func (requestBuilder *RequestBuilder) AddQuerySlice(param string, slice interface{}) (err error) {
	convertedSlice, err := ConvertSlice(slice)
	if err != nil {
		return
	}

	requestBuilder.AddQuery(param, strings.Join(convertedSlice, ","))

	return
}
