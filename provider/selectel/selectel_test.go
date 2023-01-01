/*
Copyright 2023 The Kubernetes Authors.

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

package selectel

import (
	"context"
	"os"
	"sigs.k8s.io/external-dns/endpoint"
	"testing"
)

func TestNewSelectelProvider(t *testing.T) {
	_ = os.Setenv("SELECTEL_API_KEY", "testkey")
	_, err := NewSelectelProvider(context.Background(), endpoint.NewDomainFilter([]string{"test.selectel.ru"}), true)
	if err != nil {
		t.Errorf("failed : %s", err)
	}

	_ = os.Unsetenv("SELECTEL_API_KEY")
	_, err = NewSelectelProvider(context.Background(), endpoint.NewDomainFilter([]string{"test.selectel.ru"}), true)
	if err == nil {
		t.Errorf("expected to fail")
	}
}
