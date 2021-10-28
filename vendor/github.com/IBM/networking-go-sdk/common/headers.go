/**
 * (C) Copyright IBM Corp. 2019.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"fmt"
	"runtime"
)

const (
	sdkName = "networking-go-sdk"
	headerNameUserAgent = "User-Agent"
)

//
// GetSdkHeaders - returns the set of SDK-specific headers to be included in an outgoing request.
//
// This function is invoked by generated service methods (i.e. methods which implement the REST API operations
// defined within the API definition). The purpose of this function is to give the SDK implementor the opportunity
// to provide SDK-specific HTTP headers that will be sent with an outgoing REST API request.
// This function is invoked for each invocation of a generated service method,
// so the set of HTTP headers could be request-specific.
// As an optimization, if your SDK will be returning the same set of HTTP headers for each invocation of this
// function, it is recommended that you initialize the returned map just once (perhaps by using
// lazy initialization) and simply return it each time the function is invoked, instead of building it each time
// as in the example below.
//
// If you plan to gather metrics for your SDK, the User-Agent header value must
// be a string similar to the following:
// my-go-sdk/0.0.1 (lang=go; arch=x86_64; os=Linux; go.version=1.12.9)
//
// In the example above, the analytics tool will parse the user-agent header and
// use the following properties:
// "my-go-sdk" - the name of your sdk
// "0.0.1"- the version of your sdk
// "lang=go" - the language of the current sdk
// "arch=x86_64; os=Linux; go.version=1.12.9" - system information
//
// Note: It is very important that the sdk name ends with the string `-sdk`,
// as the analytics data collector uses this to gather usage data.
//
// Parameters:
//   serviceName - the name of the service as defined in the API definition (e.g. "MyService1")
//   serviceVersion - the version of the service as defined in the API definition (e.g. "V1")
//   operationId - the operationId as defined in the API definition (e.g. getContext)
//
// Returns:
//   a Map which contains the set of headers to be included in the REST API request
//
func GetSdkHeaders(serviceName string, serviceVersion string, operationId string) map[string]string {
	sdkHeaders := make(map[string]string)

	sdkHeaders[headerNameUserAgent] = GetUserAgentInfo()

	return sdkHeaders
}

var userAgent string = fmt.Sprintf("%s/%s %s", sdkName, Version, GetSystemInfo())

func GetUserAgentInfo() string {
	return userAgent
}

var systemInfo = fmt.Sprintf("(lang=go; arch=%s; os=%s; go.version=%s)", runtime.GOARCH, runtime.GOOS, runtime.Version())

func GetSystemInfo() string {
	return systemInfo
}
