package gqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	client "github.com/Yamashou/gqlgenc/clientv2"
	"github.com/Yamashou/gqlgenc/graphqljson"
	"github.com/schollz/progressbar/v3"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Upload File represents a file to upload.
type Upload struct {
	Field string
	Name  string
	R     io.Reader
}

func WithFiles(files []Upload, httpClient *http.Client) func(ctx context.Context, req *http.Request, gqlInfo *client.GQLRequestInfo, res interface{}) error {
	return func(ctx context.Context, req *http.Request, gqlInfo *client.GQLRequestInfo, res interface{}) error {
		bd := client.Request{}
		err := json.NewDecoder(req.Body).Decode(&bd)
		if err != nil {
			fmt.Printf("decode body %v", err)
		}
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		if err := bodyWriter.WriteField("query", bd.Query); err != nil {
			fmt.Printf("write query field %v", err)
		}

		var variablesBuf bytes.Buffer
		if len(bd.Variables) > 0 {
			variablesField, err := bodyWriter.CreateFormField("variables")
			if err != nil {
				fmt.Printf("create variables field %v", err)
			}
			if err := json.NewEncoder(io.MultiWriter(variablesField, &variablesBuf)).Encode(bd.Variables); err != nil {
				fmt.Printf("encode variables %v", err)
			}
		}

		for _, f := range files {
			part, err := bodyWriter.CreateFormFile(f.Field, f.Name)
			if err != nil {
				fmt.Printf("create form file %v", err)
			}

			if _, err := io.Copy(part, f.R); err != nil {
				fmt.Printf("preparing file %v", err)
			}
		}
		bodyWriter.Close()

		bar := progressbar.DefaultBytes(
			int64(bodyBuf.Len()),
			"upload progress",
		)
		reader := progressbar.NewReader(bodyBuf, bar)
		defer reader.Close()

		req.ContentLength = int64(bodyBuf.Len())
		req.Body = &reader
		req.Header.Set("Accept", "application/json; charset=utf-8")
		req.Header.Set("Content-Type", bodyWriter.FormDataContentType())

		resp, err := httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("request failed: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		return parseResponse(body, resp.StatusCode, res)
	}
}

func parseResponse(body []byte, httpCode int, result interface{}) error {
	errResponse := &ErrorResponse{}
	isKOCode := httpCode < 200 || 299 < httpCode
	if isKOCode {
		errResponse.NetworkError = &HTTPError{
			Code:    httpCode,
			Message: fmt.Sprintf("Response body %s", string(body)),
		}
	}

	// some servers return a graphql error with a non OK http code, try anyway to parse the body
	if err := unmarshal(body, result); err != nil {
		if gqlErr, ok := err.(*GqlErrorList); ok {
			errResponse.GqlErrors = &gqlErr.Errors
		} else if !isKOCode { // if is KO code there is already the http error, this error should not be returned
			return err
		}
	}

	if errResponse.HasErrors() {
		return errResponse
	}

	return nil
}

// response is a GraphQL layer response from a handler.
type response struct {
	Data   json.RawMessage `json:"data"`
	Errors json.RawMessage `json:"errors"`
}

func unmarshal(data []byte, res interface{}) error {
	resp := response{}
	if err := json.Unmarshal(data, &resp); err != nil {
		return fmt.Errorf("failed to decode data %s: %w", string(data), err)
	}

	if resp.Errors != nil && len(resp.Errors) > 0 {
		// try to parse standard graphql error
		errors := &GqlErrorList{}
		if e := json.Unmarshal(data, errors); e != nil {
			return fmt.Errorf("faild to parse graphql errors. Response content %s - %w ", string(data), e)
		}

		return errors
	}

	if err := graphqljson.UnmarshalData(resp.Data, res); err != nil {
		return fmt.Errorf("failed to decode data into response %s: %w", string(data), err)
	}

	return nil
}

// GqlErrorList is the struct of a standard graphql error response
type GqlErrorList struct {
	Errors gqlerror.List `json:"errors"`
}

func (e *GqlErrorList) Error() string {
	return e.Errors.Error()
}

// HTTPError is the error when a GqlErrorList cannot be parsed
type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse represent an handled error
type ErrorResponse struct {
	// populated when http status code is not OK
	NetworkError *HTTPError `json:"networkErrors"`
	// populated when http status code is OK but the server returned at least one graphql error
	GqlErrors *gqlerror.List `json:"graphqlErrors"`
}

// HasErrors returns true when at least one error is declared
func (er *ErrorResponse) HasErrors() bool {
	return er.NetworkError != nil || er.GqlErrors != nil
}

func (er *ErrorResponse) Error() string {
	content, err := json.Marshal(er)
	if err != nil {
		return err.Error()
	}

	return string(content)
}
