#!/bin/bash

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

set -o errexit
set -o nounset
set -o pipefail

DEST_CHART_DIR=${DEST_CHART_DIR:-bin/}
HELM=${HELM:-helm}
HELM_CHART_REPO=${HELM_CHART_REPO:-gcr.io/k8s-staging-external-dns/charts}

readonly chart_dir="charts/external-dns"
readonly chart_file="${chart_dir}/Chart.yaml"

chart_version=$(grep -E '^version: ' "${chart_file}" | cut -d' ' -f2)

mkdir -p "${DEST_CHART_DIR}"
${HELM} package "${chart_dir}" -d "${DEST_CHART_DIR}"

if [ "${HELM_CHART_PUSH:-false}" = "true" ]; then
  ${HELM} push "${DEST_CHART_DIR}/external-dns-${chart_version}.tgz" "oci://${HELM_CHART_REPO}"
fi
