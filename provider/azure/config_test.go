/*
Copyright 2017 The Kubernetes Authors.

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

package azure

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
)

func TestGetAzureEnvironmentConfig(t *testing.T) {
	tmp, err := ioutil.TempFile("", "azureconf")
	if err != nil {
		t.Errorf("couldn't write temp file %v", err)
	}
	defer os.Remove(tmp.Name())

	tests := map[string]struct {
		cloud string
		err   error
	}{
		"AzureChinaCloud":   {"AzureChinaCloud", nil},
		"AzureGermanCloud":  {"AzureGermanCloud", nil},
		"AzurePublicCloud":  {"", nil},
		"AzureUSGovernment": {"AzureUSGovernmentCloud", nil},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, _ = tmp.Seek(0, 0)
			_, _ = tmp.Write([]byte(fmt.Sprintf(`{"cloud": "%s"}`, test.cloud)))
			got, err := getConfig(tmp.Name(), "", "")
			if err != nil {
				t.Errorf("got unexpected err %v", err)
			}

			if test.cloud == "" {
				test.cloud = "AzurePublicCloud"
			}
			want, err := azure.EnvironmentFromName(test.cloud)
			if err != nil {
				t.Errorf("couldn't get azure environment from provided name %v", err)
			}

			if !reflect.DeepEqual(want, got.Environment) {
				t.Errorf("got %v, want %v", got.Environment, want)
			}
		})
	}
}
