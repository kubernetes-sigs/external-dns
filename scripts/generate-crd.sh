#!/usr/bin/env bash
# Copyright 2026 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# generate-crd.sh
#
# This script generates Kubernetes Custom Resource Definitions (CRDs) and related
# deepcopy code for external-dns using controller-gen from controller-tools.
#
## What this script does:
# 1. Generates DeepCopy methods for types in the endpoint package
# 2. Generates CRD manifests for API types in the apis package
# 3. Copies CRDs to the Helm chart directory
#
# Usage:
#   ./scripts/generate-crd.sh
#   make crd  # calls this script

set -euo pipefail

# Get the script directory
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# Get the project root (parent of scripts directory)
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

cd "${PROJECT_ROOT}"

# Define tool commands (using tools from go.tool.mod)
CONTROLLER_GEN="go tool -modfile=go.tool.mod controller-gen"
YQ="go tool -modfile=go.tool.mod yq"
YAMLFMT="go tool -modfile=go.tool.mod yamlfmt"

echo " Generating CRDs using controller-gen..."

# Step 1: Generate deepcopy methods for endpoint types
# This creates zz_generated.deepcopy.go with DeepCopy/DeepCopyInto/DeepCopyObject methods
# The 'object' generator adds these methods for types marked with +kubebuilder:object markers
echo "  → Generating deepcopy for endpoint package..."
${CONTROLLER_GEN} object crd:crdVersions=v1 paths="./endpoint/..."

# Clean up empty import statements from generated files
# controller-gen sometimes adds empty import() blocks which create noise in diffs
find ./endpoint -name "zz_generated.deepcopy.go" -exec gofmt -s -w {} \;

# Step 2: Generate CRD manifests for API types
# - Generates CRDs from Go types with kubebuilder markers
# - Outputs to stdout, formats with yamlfmt, then splits into individual files
# - Each CRD is saved to config/crd/standard/<crd-name>.yaml
echo "  → Generating CRDs for apis package..."
${CONTROLLER_GEN} object crd:crdVersions=v1 paths="./apis/..." output:crd:stdout | \
    ${YAMLFMT} - | \
    ${YQ} eval '.' --no-doc --split-exp '"./config/crd/standard/" + .metadata.name + ".yaml"'

# Clean up empty import statements from generated files
find ./apis -name "zz_generated.deepcopy.go" -exec gofmt -s -w {} \;

# Step 3: Copy CRDs to Helm chart with filtered annotations
# - Reads CRDs from config/crd/standard/
# - Filters annotations to only keep kubernetes.io/* (removes controller-gen annotations)
# - Splits and saves to charts/external-dns/crds/ for Helm chart packaging
echo "  → Copying CRDs to chart directory..."
${YQ} eval '.metadata.annotations |= with_entries(select(.key | test("kubernetes\.io")))' \
    --no-doc --split-exp '"./charts/external-dns/crds/" + .metadata.name + ".yaml"' \
    ./config/crd/standard/*.yaml

echo -e "  ✅ CRD generation complete"
