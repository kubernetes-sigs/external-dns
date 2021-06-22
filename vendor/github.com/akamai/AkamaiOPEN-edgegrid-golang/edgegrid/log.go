// Copyright 2018. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package edgegrid

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	logstd "log"

	log "github.com/sirupsen/logrus"
)

var logBuffer *bufio.Writer
var LogFile *os.File
var EdgegridLog *log.Logger

func SetupLogging() {

	if EdgegridLog != nil {
		return // already configured
	}

	EdgegridLog = log.New()
	EdgegridLog.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation:    true,
		EnvironmentOverrideColors: true,
	})
	// Log file destination specified? If not, use default stdout
	if logFileName := os.Getenv("AKAMAI_LOG_FILE"); logFileName != "" {
		// If the file doesn't exist, create it, or append to the file
		LogFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		EdgegridLog.SetOutput(LogFile)
	}

	EdgegridLog.SetLevel(log.PanicLevel)
	if logLevel := os.Getenv("AKAMAI_LOG"); logLevel != "" {
		level, err := log.ParseLevel(logLevel)
		if err == nil {
			EdgegridLog.SetLevel(level)
		} else {
			log.Warningln("[WARN] Unknown AKAMAI_LOG value. Allowed values: panic, fatal, error, warn, info, debug, trace")

		}
	}

	defer LogFile.Close()
}

func LogMultiline(f func(args ...interface{}), args ...string) {
	for _, str := range args {
		for _, str := range strings.Split(strings.Trim(str, "\n"), "\n") {
			f(str)
		}
	}
}

func LogMultilineln(f func(args ...interface{}), args ...string) {
	LogMultiline(f, args...)
}

func LogMultilinef(f func(formatter string, args ...interface{}), formatter string, args ...interface{}) {
	str := fmt.Sprintf(formatter, args...)
	for _, str := range strings.Split(strings.Trim(str, "\n"), "\n") {
		f(str)
	}
}

// Utility func to print http req
func PrintHttpRequest(req *http.Request, body bool) {

	if req == nil {
		return
	}
	b, err := httputil.DumpRequestOut(req, body)
	if err == nil {
		LogMultiline(EdgegridLog.Traceln, string(b))
	}
}

func PrintHttpRequestCorrelation(req *http.Request, body bool, correlationid string) {

	if req == nil {
		return
	}
	b, err := httputil.DumpRequestOut(req, body)
	if err == nil {
		LogMultiline(EdgegridLog.Traceln, string(b))
		PrintfCorrelation("[DEBUG] REQUEST", correlationid, prettyPrintJsonLines(b))
	}
}

// Utility func to print http response
func PrintHttpResponse(res *http.Response, body bool) {

	if res == nil {
		return
	}
	b, err := httputil.DumpResponse(res, body)
	if err == nil {
		LogMultiline(EdgegridLog.Traceln, string(b))
	}
}

func PrintHttpResponseCorrelation(res *http.Response, body bool, correlationid string) {

	if res == nil {
		return
	}
	b, err := httputil.DumpResponse(res, body)
	if err == nil {
		LogMultiline(EdgegridLog.Traceln, string(b))
		PrintfCorrelation("[DEBUG] RESPONSE ", correlationid, prettyPrintJsonLines(b))
	}
}

func PrintfCorrelation(level string, correlationid string, msg string) {

	if correlationid == "" {
		logstd.Printf("%s  %s\n", level, msg)
	} else {
		logstd.SetFlags(0)
		logstd.Printf("%v %s\n", correlationid, msg)
	}

}

// prettyPrintJsonLines iterates through a []byte line-by-line,
// transforming any lines that are complete json into pretty-printed json.
func prettyPrintJsonLines(b []byte) string {
	parts := strings.Split(string(b), "\n")
	for i, p := range parts {
		if b := []byte(p); json.Valid(b) {
			var out bytes.Buffer
			json.Indent(&out, b, "", " ")
			parts[i] = out.String()
		}
	}
	return strings.Join(parts, "\n")
}
