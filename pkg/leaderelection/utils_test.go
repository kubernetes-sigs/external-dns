/*
Copyright 2025 The Kubernetes Authors.

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

package leaderelection

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetInClusterNamespace_FileNotExist(t *testing.T) {
	tmp := t.TempDir()
	originalPath := inClusterNamespacePath
	defer func() { inClusterNamespacePath = originalPath }()

	inClusterNamespacePath = fmt.Sprintf("%s/kubernetes.io/serviceaccount/namespace", tmp)
	require.Equal(t, DefaultNamespace, getNamespace())
}

func TestGetInClusterNamespace_InClusterFileExists(t *testing.T) {
	tmp := t.TempDir()
	originalPath := inClusterNamespacePath
	defer func() { inClusterNamespacePath = originalPath }()

	inClusterNamespacePath = fmt.Sprintf("%s/namespace", tmp)
	file, err := os.Create(inClusterNamespacePath)
	require.NoError(t, err)
	defer file.Close()

	ns := "test-namespace"
	_, err = file.WriteString(ns)
	require.Equal(t, ns, getNamespace())
}

func TestGetLeaderId_Success(t *testing.T) {
	id := getLeaderId()
	require.NotEmpty(t, id)
	require.Contains(t, id, "_")
}

func TestGetLeaderId_HostnameFailed(t *testing.T) {
	defer func() { hostname = os.Hostname }()
	hostname = func() (string, error) {
		return "", fmt.Errorf("hostname failed")
	}
	id := getLeaderId()
	require.NotEmpty(t, id)
	require.Contains(t, id, "external-dns_")
}
