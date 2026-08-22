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

# build-fips-image.sh
#
# Builds a FIPS 140-3 flagged variant of the external-dns image, using Go's
# native FIPS 140-3 module (see https://go.dev/doc/security/fips140). No
# BoringCrypto/cgo/OpenSSL provider is involved - GOFIPS140 is a pure Go
# build setting.
#
# Usage:
#   ./scripts/build-fips-image.sh
#   make build.image-fips  # calls this script
#
# Env vars (all optional, matching the Makefile's build.push/multiarch vars):
#   VERSION       Image tag prefix; final tag is "${VERSION}-fips" (default: git describe)
#   IMAGE         KO_DOCKER_REPO to build into (default: us.gcr.io/k8s-artifacts-prod/external-dns/external-dns)
#   GIT_REVISION  Recorded in the image's revision label (default: git rev-parse HEAD)
#   IMG_PLATFORM  Target platform(s) (default: linux/amd64,linux/arm64,linux/arm/v7)
#   IMG_PUSH      Whether to push the built image (default: false)
#   GOFIPS140     Go FIPS 140-3 module version to build with (default: v1.0.0, the
#                 CMVP-validated module; "latest"/"v1.26.0" is still Pending Review)

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
cd "${PROJECT_ROOT}"

VERSION="${VERSION:-$(git describe --tags --always --dirty --match "v*")}"
IMAGE="${IMAGE:-us.gcr.io/k8s-artifacts-prod/external-dns/external-dns}"
GIT_REVISION="${GIT_REVISION:-$(git rev-parse HEAD)}"
IMG_PLATFORM="${IMG_PLATFORM:-linux/amd64,linux/arm64,linux/arm/v7}"
IMG_PUSH="${IMG_PUSH:-false}"
GOFIPS140="${GOFIPS140:-v1.0.0}"

tmp_ko_dir="$(mktemp -d)"
trap 'rm -rf "${tmp_ko_dir}"' EXIT
tmp_ko_config="${tmp_ko_dir}/ko.yaml"

echo ">> generating temporary ko config with GOFIPS140=${GOFIPS140}"
cat > "${tmp_ko_config}" <<EOF
defaultBaseImage: gcr.io/distroless/static-debian12:latest
builds:
- env:
  - CGO_ENABLED=0
  - GOFIPS140=${GOFIPS140}
  flags:
  - -v
  ldflags:
  - -s
  - -w
  - -X sigs.k8s.io/external-dns/pkg/apis/externaldns.Version={{.Env.VERSION}}
EOF

echo ">> building FIPS-flagged image: ${IMAGE}:${VERSION}-fips"
KO_CONFIG_PATH="${tmp_ko_config}" \
KO_DOCKER_REPO="${IMAGE}" \
VERSION="${VERSION}" \
ko build --tags "${VERSION}-fips" --bare --sbom spdx \
	--image-label org.opencontainers.image.source="https://github.com/kubernetes-sigs/external-dns" \
	--image-label org.opencontainers.image.revision="${GIT_REVISION}" \
	--platform="${IMG_PLATFORM}" --push="${IMG_PUSH}" .
