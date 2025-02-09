#!/usr/bin/env bash

# Copyright 2025 The Kubernetes Authors.
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

# Execute
# scripts/api-docs.sh test

GOPATH=${GOPATH:-$(go env GOPATH)}

# "go env" doesn't print anything if GOBIN is the default, so we
# have to manually default it.
GOBIN=${GOBIN:-$(go env GOBIN)}
GOBIN=${GOBIN:-${GOPATH}/bin}

readonly HERE=$(cd $(dirname $0) && pwd)
readonly REPO=$(cd ${HERE}/../.. && pwd)

gendoc::build() {
    go install github.com/ahmetb/gen-crd-api-reference-docs
}

# Exec the doc generator.
gendoc::exec() {
    local readonly confdir="scripts/"

    ${GOBIN}/gen-crd-api-reference-docs \
        -template-dir ${confdir} \
        -config ${confdir}/config.json \
        "$@"
}

gendoc::build
gendoc::exec \
    -api-dir "sigs.k8s.io/external-dns/endpoint" \
    -out-file "${1}"