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
# native FIPS 140-3 module (see https://go.dev/doc/security/fips140).
#
# Usage:
#   ./scripts/build-fips-image.sh [-h|--help]
#   make build.image-fips  # calls this script

show_help() {
cat << EOF
Build a FIPS 140-3 flagged variant of the external-dns image.

Usage: $(basename "$0") [-h|--help]

Env vars (all optional, matching the Makefile's build.push/multiarch vars):
    VERSION       Image tag prefix; final tag is "\${VERSION}-fips" (default: git describe)
    IMAGE         KO_DOCKER_REPO to build into (default: us.gcr.io/k8s-artifacts-prod/external-dns/external-dns,
                  or "external-dns" when IMG_LOCAL=true)
    GIT_REVISION  Recorded in the image's revision label (default: git rev-parse HEAD)
    IMG_PLATFORM  Target platform(s) (default: linux/amd64,linux/arm64,linux/arm/v7,
                  or the host platform when IMG_LOCAL=true)
    IMG_PUSH      Whether to push the built image (default: false; forced false when IMG_LOCAL=true)
    IMG_LOCAL     Load the built image into the local Docker daemon and verify the FIPS
                  module landed in the binary. Set to "false" to just prove the build
                  succeeds, without loading or verifying (default: true)
    GOFIPS140     Go FIPS 140-3 module version to build with (default: v1.0.0, the
                  CMVP-validated module; "latest"/"v1.26.0" is still Pending Review)
EOF
}

case "${1:-}" in
	-h|--help)
		show_help
		exit 0
		;;
esac

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
cd "${PROJECT_ROOT}"

IMG_LOCAL="${IMG_LOCAL:-true}"

VERSION="${VERSION:-$(git describe --tags --always --dirty --match "v*")}"
GIT_REVISION="${GIT_REVISION:-$(git rev-parse HEAD)}"
GOFIPS140="${GOFIPS140:-v1.0.0}"

if [[ "${IMG_LOCAL}" == "true" ]]; then
	IMAGE="${IMAGE:-external-dns}"
	IMG_PLATFORM="${IMG_PLATFORM:-linux/$(go env GOARCH)}"
	IMG_PUSH=false
else
	IMAGE="${IMAGE:-us.gcr.io/k8s-artifacts-prod/external-dns/external-dns}"
	IMG_PLATFORM="${IMG_PLATFORM:-linux/amd64,linux/arm64,linux/arm/v7}"
	IMG_PUSH="${IMG_PUSH:-false}"
fi

tmp_ko_dir="$(mktemp -d)"
verify_container=""
cleanup() {
	rm -rf "${tmp_ko_dir}"
	[[ -n "${verify_container}" ]] && docker rm -f "${verify_container}" >/dev/null 2>&1
	return 0
}
trap cleanup EXIT
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
	--platform="${IMG_PLATFORM}" --push="${IMG_PUSH}" --local="${IMG_LOCAL}" .

if [[ "${IMG_LOCAL}" == "true" ]]; then
	echo ">> verifying FIPS module in ${IMAGE}:${VERSION}-fips"
	verify_container="$(docker create "${IMAGE}:${VERSION}-fips")"
	tmp_binary="${tmp_ko_dir}/external-dns"
	docker cp "${verify_container}:/ko-app/external-dns" "${tmp_binary}" >/dev/null
	docker rm "${verify_container}" >/dev/null
	verify_container=""

	fips_info="$(go version -m "${tmp_binary}" | grep -E '^\s+build\s' | grep -Ei 'fips|godebug' || true)"
	if [[ -z "${fips_info}" ]]; then
		echo "FIPS module not found in built binary" >&2
		exit 1
	fi
	echo "${fips_info}"
fi
