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
	"os"
	"path"
	"runtime"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/stretchr/testify/assert"
)

func TestGetCloudConfiguration(t *testing.T) {
	tests := map[string]struct {
		cloudName string
		expected  cloud.Configuration
		setEnv    map[string]string
	}{
		"AzureChinaCloud":   {"AzureChinaCloud", cloud.AzureChina, nil},
		"AzurePublicCloud":  {"", cloud.AzurePublic, nil},
		"AzureUSGovernment": {"AzureUSGovernmentCloud", cloud.AzureGovernment, nil},
		"AzureCustomCloud":  {"AzureCustomCloud", cloud.Configuration{ActiveDirectoryAuthorityHost: "https://custom.microsoftonline.com/"}, map[string]string{"AZURE_AD_ENDPOINT": "https://custom.microsoftonline.com/"}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if test.setEnv != nil {
				for key, value := range test.setEnv {
					os.Setenv(key, value)
					defer os.Unsetenv(key)
				}
			}

			cloudCfg, err := getCloudConfiguration(test.cloudName)
			if err != nil {
				t.Errorf("got unexpected err %v", err)
			}
			if cloudCfg.ActiveDirectoryAuthorityHost != test.expected.ActiveDirectoryAuthorityHost {
				t.Errorf("got %v, want %v", cloudCfg, test.expected)
			}
		})
	}
}

func TestOverrideConfiguration(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	configFile := path.Join(path.Dir(filename), "config_test.json")
	cfg, err := getConfig(configFile, "subscription-override", "rg-override", "")
	if err != nil {
		t.Errorf("got unexpected err %v", err)
	}
	assert.Equal(t, cfg.SubscriptionID, "subscription-override")
	assert.Equal(t, cfg.ResourceGroup, "rg-override")
}
