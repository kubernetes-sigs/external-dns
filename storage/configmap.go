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

package storage

import (
	"github.com/kubernetes-incubator/external-dns/dnsprovider"
)

/**
ConfigMap - implements storage interface via Kubernetes ConfigMap resource
1. ConfigMap name is passed to the external-dns
2. If such configmap does not exist then it is created in the same namespace
3. ConfigMap fetches data from the specified dns provider
4. And if record is created/updated/deleted by external-dns, it should be accordingly updated in the ConfigMap after dnsprovider API request is successful
5. Periodically configmap resyncs with the dns provider
*/

// ConfigMap implementation of storage via Kubernetes ConfigMap resource
type ConfigMap struct {
	DNSProvider dnsprovider.DNSProvider
}

// Records reads the records from the configmap and converts them to the Entry
func (cm *ConfigMap) Records() []*Entry {
	return nil
}

// Update updates the ConfigMap by overwriting information in the configmap
func (cm *ConfigMap) Update(entries []*Entry) error {
	return nil
}

// WaitForSync fetches the data from dns provider and updates the config map accordingly
// also makes sure that the config map is created if does not exist
func (cm *ConfigMap) WaitForSync() error {
	return nil
}
