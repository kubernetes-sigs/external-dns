package common

import (
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseFromHttpResponse(hr *http.Response, response Response) (err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewTencentCloudSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewTencentCloudSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	//log.Printf("[DEBUG] Response Body=%s", body)
	err = response.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	return
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"bufio"
	"bytes"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/json"
	"io/ioutil"
	"log"
	"strconv"

	//"log"
	"net/http"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
)

type Response interface {
	ParseErrorFromHTTPResponse(body []byte) error
}

type BaseResponse struct {
}

type SSEResponse interface {
	Response
	setEventsChannel(ch chan SSEvent)
}

type SSEvent struct {
	Event string
	Data  []byte
	Id    string
	Retry int64
	Err   error
}

type BaseSSEResponse struct {
	BaseResponse
	Events chan SSEvent
}

func (r *BaseSSEResponse) setEventsChannel(ch chan SSEvent) {
	r.Events = ch
}

type ErrorResponse struct {
	Response struct {
		Error struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
		RequestId string `json:"RequestId"`
	} `json:"Response"`
}

type DeprecatedAPIErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	CodeDesc string `json:"codeDesc"`
}

func (r *BaseResponse) ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func ParseErrorFromHTTPResponse(body []byte) (err error) {
	resp := &ErrorResponse{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if resp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(resp.Response.Error.Code, resp.Response.Error.Message, resp.Response.RequestId)
	}

	deprecated := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, deprecated)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if deprecated.Code != 0 {
		return errors.NewTencentCloudSDKError(deprecated.CodeDesc, deprecated.Message, "")
	}
	return nil
}

func TryReadErr(resp *http.Response) (err error) {
	switch resp.Header.Get("Content-Type") {
	case "text/event-stream", "application/octet-stream":
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read http body, because: %s", err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	resp.Body.Close()

	errResp := &ErrorResponse{}
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", string(body), err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if errResp.Response.Error.Code != "" {
		return errors.NewTencentCloudSDKError(errResp.Response.Error.Code, errResp.Response.Error.Message, errResp.Response.RequestId)
	}

	depResp := &DeprecatedAPIErrorResponse{}
	err = json.Unmarshal(body, &depResp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", string(body), err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	if depResp.Code != 0 {
		return errors.NewTencentCloudSDKError(depResp.CodeDesc, depResp.Message, "")
	}
	resp.Body = ioutil.NopCloser(bytes.NewReader(body))
	return nil
}

func ParseFromHttpResponse(hr *http.Response, resp Response) error {
	switch hr.Header.Get("Content-Type") {
	case "text/event-stream":
		return parseFromSSE(hr, resp)
	default:
		return parseFromJson(hr, resp)
	}
}

func parseFromJson(hr *http.Response, resp Response) error {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		msg := fmt.Sprintf("Fail to read response body because %s", err)
		return errors.NewTencentCloudSDKError("ClientError.IOError", msg, "")
	}
	if hr.StatusCode != 200 {
		msg := fmt.Sprintf("Request fail with http status code: %s, with body: %s", hr.Status, body)
		return errors.NewTencentCloudSDKError("ClientError.HttpStatusCodeError", msg, "")
	}
	err = resp.ParseErrorFromHTTPResponse(body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		msg := fmt.Sprintf("Fail to parse json content: %s, because: %s", body, err)
		return errors.NewTencentCloudSDKError("ClientError.ParseJsonError", msg, "")
	}
	return nil
}

func parseFromSSE(hr *http.Response, resp Response) error {
	reqId := hr.Header.Get("X-TC-RequestId")
	r, ok := resp.(SSEResponse)
	if !ok {
		return errors.NewTencentCloudSDKError("ClientError.TypeError",
			"Response type does not implement SSEResponse", reqId)
	}

	ch := make(chan SSEvent)
	r.setEventsChannel(ch)

	// parser
	go func() {
		defer hr.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(hr.Body)
		scanner.Split(bufio.ScanLines)

		event := SSEvent{}
		for scanner.Scan() {
			line := scanner.Bytes()

			// SSE use empty line to indicate message end
			if len(line) == 0 {
				select {
				case ch <- event:
					event = SSEvent{}
					continue
				case <-hr.Request.Context().Done():
					select {
					case ch <- SSEvent{Err: hr.Request.Context().Err()}:
					default:
						log.Println(hr.Request.Context().Err())
					}
					return
				}
			}

			// comment
			if line[0] == ':' {
				continue
			}

			idx := bytes.IndexByte(line, ':')
			if idx == -1 {
				select {
				case ch <- SSEvent{Err: fmt.Errorf("SSE.InvalidLine:%s", line)}:
				default:
				}
				return
			}

			key := string(line[:idx])
			val := line[idx+1:]
			switch key {
			case "event":
				event.Event = string(val)
			case "data":
				// The spec allows for multiple data fields per event, concat them with "\n".
				if len(event.Data) > 0 {
					event.Data = append(event.Data, '\n')
				}
				event.Data = append(event.Data, val...)
			case "id":
				event.Id = string(val)
			case "retry":
				retry, err := strconv.ParseInt(string(val), 10, 64)
				if err != nil {
					select {
					case ch <- SSEvent{Err: fmt.Errorf("SSE.InvalidRetry:%s", line)}:
					default:
					}
					return
				}
				event.Retry = retry
			}
		}

		if err := scanner.Err(); err != nil {
			select {
			case ch <- SSEvent{Err: err}:
			default:
				log.Println(err)
			}
		}
	}()

	return nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
