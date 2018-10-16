package egoscale

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

const (
	jsonContentType = "application/json"
)

func TestJobStatusTypeString(t *testing.T) {
	if Failure != JobStatusType(2) {
		t.Error("bad enum value", (int)(Failure), 2)
	}

	if Failure.String() != "Failure" {
		t.Error("mismatch", Failure, "Failure")
	}
	s := JobStatusType(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad state", s.String())
	}
}

func TestErrorCodeString(t *testing.T) {
	if ParamError != ErrorCode(431) {
		t.Error("bad enum value", (int)(ParamError), 431)
	}

	if ParamError.String() != "ParamError" {
		t.Error("mismatch", ParamError, "ParamError")
	}
	s := ErrorCode(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad state", s.String())
	}
}

func TestCSErrorCodeString(t *testing.T) {
	if CloudException != CSErrorCode(4275) {
		t.Error("bad enum value", (int)(CloudException), 4275)
	}

	if CloudException.String() != "CloudException" {
		t.Error("mismatch", CloudException, "CloudException")
	}
	s := CSErrorCode(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad state", s.String())
	}
}

func TestRequest(t *testing.T) {
	params := url.Values{}
	params.Set("command", "listApis")
	params.Set("apikey", "KEY")
	params.Set("name", "dummy")
	params.Set("response", "json")
	ts := newGetServer(params, jsonContentType, `
{
	"listapisresponse": {
		"api": [{
			"name": "dummy",
			"description": "this is a test",
			"isasync": false,
			"since": "4.4",
			"type": "list",
			"name": "listDummies",
			"params": [],
			"related": "",
			"response": []
		}],
		"count": 1
	}
}
	`)
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	req := &ListAPIs{
		Name: "dummy",
	}
	resp, err := cs.Request(req)
	if err != nil {
		t.Fatalf(err.Error())
	}
	apis := resp.(*ListAPIsResponse)
	if apis.Count != 1 {
		t.Errorf("Expected exactly one API, got %d", apis.Count)
	}
}

func TestRequestSignatureFailure(t *testing.T) {
	ts := newServer(response{401, jsonContentType, `
{"createsshkeypairresponse" : {
	"uuidList":[],
	"errorcode":401,
	"errortext":"unable to verify usercredentials and/or request signature"
}}
	`})
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &CreateSSHKeyPair{
		Name: "123",
	}

	if _, err := cs.Request(req); err == nil {
		t.Errorf("This should have failed?")
		r, ok := err.(*ErrorResponse)
		if !ok {
			t.Errorf("A CloudStack error was expected, got %v", err)
		}
		if r.ErrorCode != Unauthorized {
			t.Errorf("Unauthorized error was expected")
		}
	}
}

func TestBooleanAsyncRequest(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{
	"expungevirtualmachineresponse": {
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobresult": {},
		"jobstatus": 0
	}
}
	`}, response{200, jsonContentType, `
{
	"queryasyncjobresultresponse": {
		"accountid": "7bd023bf-d55c-4b27-9918-680f03efc26c",
		"cmd": "expunge",
		"created": "2018-04-03T22:40:04+0200",
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobprocstatus": 0,
		"jobresult": {
			"success": true
		},
		"jobresultcode": 0,
		"jobresulttype": "object",
		"jobstatus": 1,
		"userid": "87cf2ef9-0d1b-4763-96d8-c33676371f29"
	}
}
	`})
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &ExpungeVirtualMachine{
		ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
	}
	if err := cs.BooleanRequest(req); err != nil {
		t.Error(err)
	}
}

func TestBooleanAsyncRequestFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{
	"expungevirtualmachineresponse": {
		"jobid": "7bd023bf-d55c-4b27-9918-680f03efc26c",
		"jobresult": {},
		"jobstatus": 0
	}
}
	`}, response{200, jsonContentType, `
{
	"queryasyncjobresultresponse": {
		"accountid": "7bd023bf-d55c-4b27-9918-680f03efc26c",
		"cmd": "expunge",
		"created": "2018-04-03T22:40:04+0200",
		"jobid": "87cf2ef9-0d1b-4763-96d8-c33676371f29",
		"jobprocstatus": 0,
		"jobresult": {
			"errorcode": 531,
			"errortext": "fail",
			"cserrorcode": 9999
		},
		"jobresultcode": 0,
		"jobresulttype": "object",
		"jobstatus": 2,
		"userid": "a0d421e9-8d99-4896-b22c-ea77abaebf06"
	}
}
	`})
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &ExpungeVirtualMachine{
		ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
	}
	if err := cs.BooleanRequest(req); err == nil {
		t.Error("An error was expected")
	}
}

func TestBooleanAsyncRequestWithContext(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{
	"expungevirtualmachineresponse": {
		"jobid": "a0d421e9-8d99-4896-b22c-ea77abaebf06",
		"jobresult": {},
		"jobstatus": 0
	}
}
	`}, response{200, jsonContentType, `
{
	"queryasyncjobresultresponse": {
		"accountid": "a0d421e9-8d99-4896-b22c-ea77abaebf06",
		"cmd": "expunge",
		"created": "2018-04-03T22:40:04+0200",
		"jobid": "7bd023bf-d55c-4b27-9918-680f03efc26c",
		"jobprocstatus": 0,
		"jobresult": {
			"success": true
		},
		"jobresultcode": 0,
		"jobresulttype": "object",
		"jobstatus": 1,
		"userid": "87cf2ef9-0d1b-4763-96d8-c33676371f29"
	}
}
	`})
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &ExpungeVirtualMachine{
		ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
	}

	// WithContext
	if err := cs.BooleanRequestWithContext(context.Background(), req); err != nil {
		t.Error(err)
	}
}

func TestBooleanRequestTimeout(t *testing.T) {
	ts := newSleepyServer(time.Second, 200, jsonContentType, `
{
	"expungevirtualmachine": {
		"jobid": "87cf2ef9-0d1b-4763-96d8-c33676371f29",
		"jobresult": {
			"success": false
		},
		"jobstatus": 0
	}
}
	`)
	defer ts.Close()
	done := make(chan bool)

	go func() {
		cs := NewClient(ts.URL, "TOKEN", "SECRET")
		cs.HTTPClient.Timeout = time.Millisecond

		req := &ExpungeVirtualMachine{
			ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
		}
		err := cs.BooleanRequest(req)

		if err == nil {
			t.Error("An error was expected")
		}

		// We expect the HTTP Client to timeout
		msg := err.Error()
		if !strings.HasPrefix(msg, "Get") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}

		done <- true
	}()

	<-done
}

func TestSyncRequestWithoutContext(t *testing.T) {

	ts := newServer(
		response{200, jsonContentType, `{
	"deployvirtualmachineresponse": {
		"jobid": "6c4077e3-4ec2-4e6d-9806-4ab1a30138ba",
		"jobresult": {},
		"jobstatus": 0
	}
}`},
	)

	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &DeployVirtualMachine{
		Name:              "test",
		ServiceOfferingID: MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470"),
		ZoneID:            MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11"),
		TemplateID:        MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108"),
	}

	// WithContext
	resp, err := cs.SyncRequest(req)
	if err != nil {
		t.Error(err)
	}
	result, ok := resp.(*AsyncJobResult)
	if !ok {
		t.Error("wrong type")
	}

	id := MustParseUUID("6c4077e3-4ec2-4e6d-9806-4ab1a30138ba")
	if !result.JobID.Equal(*id) {
		t.Errorf("wrong job id, expected %s, got %s", id, result.JobID)
	}
}

func TestAsyncRequestWithoutContext(t *testing.T) {

	ts := newServer(
		response{200, jsonContentType, `{
	"deployvirtualmachineresponse": {
		"jobid": "56bb40d1-4c65-4608-803b-2fdfcb21fa3b",
		"jobresult": {},
		"jobstatus": 0
	}
}`},
		response{200, jsonContentType, `{
	"queryasyncjobresultresponse": {
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobresult": {
			"virtualmachine": {
				"id": "f344b886-2a8b-4d2c-9662-1f18e5cdde6f",
				"serviceofferingid": "71004023-bb72-4a97-b1e9-bc66dfce9470",
				"templateid": "78c2cbe6-8e11-4722-b01f-bf06f4e28108",
				"zoneid": "1128bd56-b4d9-4ac6-a7b9-c715b187ce11",
				"jobid": "220504ac-b9e7-4fee-b402-47b3c4155fdb"
			}
		},
		"jobstatus": 1
	}
}`},
	)

	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &DeployVirtualMachine{
		Name:              "test",
		ServiceOfferingID: MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470"),
		ZoneID:            MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11"),
		TemplateID:        MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108"),
	}

	resp := &VirtualMachine{}

	// WithContext
	cs.AsyncRequest(req, func(j *AsyncJobResult, err error) bool {
		if err != nil {
			t.Error(err)
			return false
		}

		if j.JobStatus == Success {
			if r := j.Result(resp); r != nil {
				t.Error(r)
			}
			return false
		}
		return true
	})

	id := MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470")
	if !resp.ServiceOfferingID.Equal(*id) {
		t.Errorf("Expected ServiceOfferingID %q, got %q", id, resp.ServiceOfferingID)
	}
}

func TestAsyncRequestWithoutContextFailure(t *testing.T) {
	ts := newServer(
		response{200, jsonContentType, `{
	"deployvirtualmachineresponse": {
		"jobid": "0a1be26b-415b-4c17-87b6-ffb06c507f8b",
		"jobresult": {},
		"jobstatus": 0
	}
}`},
		response{200, jsonContentType, `{
	"queryasyncjobresultresponse": {
		"jobid": "be805478-3c23-460f-a712-11cc8df6da48",
		"jobresult": {
			"virtualmachine": []
		},
		"jobstatus": 1
	}
}`},
	)

	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &DeployVirtualMachine{
		Name:              "test",
		ServiceOfferingID: MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470"),
		ZoneID:            MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11"),
		TemplateID:        MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108"),
	}

	resp := &VirtualMachine{}

	// WithContext
	cs.AsyncRequest(req, func(j *AsyncJobResult, err error) bool {
		if err != nil {
			t.Fatal(err)
		}

		if j.JobStatus == Success {

			if r := j.Result(resp); r != nil {
				return false
			}
			t.Errorf("Expected an error, got <nil>")
		}
		return true
	})
}

func TestAsyncRequestWithoutContextFailureNext(t *testing.T) {

	ts := newServer(
		response{200, jsonContentType, `{
	"deployvirtualmachineresponse: {
		"jobid": "c3f8457b-a10b-4935-a837-68fb53b29008",
		"jobresult": {},
		"jobstatus": 0
	}
}`},
	)

	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &DeployVirtualMachine{
		Name:              "test",
		ServiceOfferingID: MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470"),
		ZoneID:            MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11"),
		TemplateID:        MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108"),
	}

	cs.AsyncRequest(req, func(j *AsyncJobResult, err error) bool {
		return err == nil
	})
}

func TestAsyncRequestWithoutContextFailureNextNext(t *testing.T) {

	ts := newServer(
		response{200, jsonContentType, `{
	"deployvirtualmachineresponse": {
		"jobid": "8933fb8c-ff24-4ca4-b7d1-ba2154e9f2c4",
		"jobresult": {
			"virtualmachine": {}
		},
		"jobstatus": 2
	}
}`},
		response{200, jsonContentType, `{
	"queryasyncjobresultresponse": {
		"jobid": "d05893e0-815b-4396-9189-5b5d8380b380",
		"jobresult": {},
		"jobstatus": 0
	}
}`},
		response{200, jsonContentType, `{
	"queryasyncjobresultresponse": {
		"jobid": "92c82b1d-b57e-4338-94bb-1ffe2130993b",
		"jobresult": [],
		"jobstatus": 1
	}
}`},
	)
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")
	req := &DeployVirtualMachine{
		Name:              "test",
		ServiceOfferingID: MustParseUUID("71004023-bb72-4a97-b1e9-bc66dfce9470"),
		ZoneID:            MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11"),
		TemplateID:        MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108"),
	}

	resp := &VirtualMachine{}

	cs.AsyncRequest(req, func(j *AsyncJobResult, err error) bool {
		if err != nil {
			t.Fatal(err)
		}

		if j.JobStatus == Success {

			j.JobStatus = Failure
			if r := j.Result(resp); r != nil {
				return false
			}
			t.Errorf("Expected an error, got <nil>")
		}
		return true
	})
}

func TestBooleanRequestWithContext(t *testing.T) {
	ts := newSleepyServer(time.Second, 200, jsonContentType, `
{
	"expungevirtualmachine": {
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobresult": {
			"success": false
		},
		"jobstatus": 0
	}
}
	`)
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	done := make(chan bool)

	go func() {
		cs := NewClient(ts.URL, "TOKEN", "SECRET")
		req := &ExpungeVirtualMachine{
			ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
		}

		err := cs.BooleanRequestWithContext(ctx, req)

		if err == nil {
			t.Error("An error was expected")
		}

		// We expect the context to timeout
		msg := err.Error()
		if !strings.HasPrefix(msg, "Get") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}

		done <- true
	}()

	<-done
}

func TestRequestWithContextTimeoutPost(t *testing.T) {
	ts := newSleepyServer(time.Second, 200, jsonContentType, `
{
	"deployvirtualmachineresponse": {
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobresult": {
			"success": false
		},
		"jobstatus": 0
	}
}
	`)
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	done := make(chan bool)

	userData := make([]byte, 1<<11)
	_, err := rand.Read(userData)
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		cs := NewClient(ts.URL, "TOKEN", "SECRET")
		req := &DeployVirtualMachine{
			ServiceOfferingID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
			TemplateID:        MustParseUUID("d2fcd819-7f6e-462d-b8c0-bfae83e4d273"),
			UserData:          base64.StdEncoding.EncodeToString(userData),
			ZoneID:            MustParseUUID("68f0e13a-2ba8-4f8f-81c0-bd78491d81ea"),
		}

		_, err := cs.RequestWithContext(ctx, req)

		if err == nil {
			t.Error("An error was expected")
		}

		// We expect the context to timeout
		msg := err.Error()
		if !strings.HasPrefix(msg, "Post") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}

		done <- true
	}()

	<-done
}

func TestBooleanRequestWithContextAndTimeout(t *testing.T) {
	ts := newSleepyServer(time.Second, 200, jsonContentType, `
{
	"expungevirtualmachine": {
		"jobid": "01ed7adc-8b81-4e33-a0f2-4f55a3b880cd",
		"jobresult": {
			"success": false
		},
		"jobstatus": 0
	}
}
	`)
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		cs := NewClient(ts.URL, "TOKEN", "SECRET")
		cs.HTTPClient.Timeout = time.Millisecond

		req := &ExpungeVirtualMachine{
			ID: MustParseUUID("925207d3-beea-4c56-8594-ef351b526dd3"),
		}
		err := cs.BooleanRequestWithContext(ctx, req)

		if err == nil {
			t.Error("An error was expected")
		}

		// We expect the client to timeout
		msg := err.Error()
		if !strings.HasPrefix(msg, "Get") || !strings.Contains(msg, "net/http: request canceled") {
			t.Errorf("Unexpected error message: %s", err.Error())
		}

		done <- true
	}()

	<-done
}

func TestWrongBodyResponse(t *testing.T) {
	ts := newServer(response{200, "text/html", `
		<html>
		<header><title>This is title</title></header>
		<body>
		Hello world
		</body>
		</html>
	`})
	defer ts.Close()

	cs := NewClient(ts.URL, "TOKEN", "SECRET")

	_, err := cs.Request(&ListZones{})
	if err == nil {
		t.Error("an error was expected but got nil error")
	}

	if err.Error() != fmt.Sprintf("body content-type response expected %q, got %q", jsonContentType, "text/html") {
		t.Error("body content-type error response expected")
	}
}

func TestRequestNilCommand(t *testing.T) {
	cs := NewClient("URL", "TOKEN", "SECRET")

	_, err := cs.Request((*ListZones)(nil))
	if err == nil {
		t.Error("an error was expected bot got nil error")
	}
}

// helpers

type response struct {
	code        int
	contentType string
	body        string
}

func newServer(responses ...response) *httptest.Server {
	i := 0
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if i >= len(responses) {
			w.Header().Set("Content-Type", jsonContentType)
			w.WriteHeader(500)
			w.Write([]byte("{}"))
			return
		}
		w.Header().Set("Content-Type", responses[i].contentType)
		w.WriteHeader(responses[i].code)
		w.Write([]byte(responses[i].body))
		i++
	})
	return httptest.NewServer(mux)
}

func newSleepyServer(sleep time.Duration, code int, contentType, response string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleep)
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(code)
		w.Write([]byte(response))
	})
	return httptest.NewServer(mux)
}

func newGetServer(params url.Values, contentType, response string) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		errors := make(map[string][]string)
		query := r.URL.Query()
		for k, expected := range params {
			if values, ok := query[k]; ok {
				for i, value := range values {
					e := expected[i]
					if e != value {
						if _, ok := errors[k]; !ok {
							errors[k] = make([]string, len(values))
						}
						errors[k][i] = fmt.Sprintf("%s expected %v, got %v", k, e, value)
					}
				}
			} else {
				errors[k] = make([]string, 1)
				errors[k][0] = fmt.Sprintf("%s was expected", k)
			}
		}

		if len(errors) == 0 {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(200)
			w.Write([]byte(response))
		} else {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(400)
			body, _ := json.Marshal(errors)
			w.Write(body)
		}
	})
	return httptest.NewServer(mux)
}
