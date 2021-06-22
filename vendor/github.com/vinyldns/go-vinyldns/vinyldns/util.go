/*
Copyright 2018 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	awsauth "github.com/smartystreets/go-aws-auth"
)

func concat(arr []string) string {
	var buff bytes.Buffer

	for _, elem := range arr {
		buff.WriteString(elem)
	}

	return buff.String()
}

func concatStrs(delim string, str ...string) string {
	return strings.Join(str, delim)
}

func resourceRequest(c *Client, url, method string, body []byte, responseStruct interface{}) error {
	if logRequests() {
		fmt.Printf("Request url: \n\t%s\nrequest body: \n\t%s \n\n", url, string(body))
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")

	awsauth.Sign4(req, awsauth.Credentials{
		AccessKeyID:     c.AccessKey,
		SecretAccessKey: c.SecretKey,
	})

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if logRequests() {
		fmt.Printf("Response status: \n\t%d\nresponse body: \n\t%s \n\n", resp.StatusCode, bodyContents)
	}
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted {
		dError := &Error{}
		dError.RequestURL = url
		dError.RequestMethod = method
		dError.RequestBody = string(body)
		dError.ResponseCode = resp.StatusCode
		dError.ResponseBody = string(bodyContents)
		return dError
	}
	err = json.Unmarshal(bodyContents, responseStruct)
	if err != nil {
		return err
	}
	return nil
}
