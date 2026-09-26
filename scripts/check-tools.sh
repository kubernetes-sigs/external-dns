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

# Compares the tools on PATH against mise.toml. Installs nothing.
# Warns by default, exits non-zero only with --strict.
# Tool names narrow the check to those tools.
#
# Execute
#   scripts/check-tools.sh
#   scripts/check-tools.sh --strict
#   scripts/check-tools.sh --strict golang yq

set -uo pipefail

STRICT=0
WANTED=()
while [[ $# -gt 0 ]]; do
    case "$1" in
        --strict) STRICT=1 ;;
        -*)       echo "ERROR: unknown option '$1', expected --strict" >&2; exit 2 ;;
        *)        WANTED+=("$1") ;;
    esac
    shift
done

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

binary_for() {
    case "$1" in
        golang)                echo "go" ;;
        helm-ct)               echo "ct" ;;
        kube-controller-tools) echo "controller-gen" ;;
        pipx:*)                echo "${1#pipx:}" ;;
        *)                     echo "$1" ;;
    esac
}

# Every tool prints a bare major.minor.patch somewhere in its version output.
installed_version() {
    local binary="$1" out found
    case "${binary}" in
        go)             out=$(go version 2>&1) ;;
        golangci-lint)  out=$(golangci-lint version --short 2>&1) ;;
        helm)           out=$(helm version --template '{{.Version}}' 2>&1) ;;
        helm-docs)      out=$(helm-docs --version 2>&1) ;;
        yamlfmt)        out=$(yamlfmt -version 2>&1) ;;
        yq)             out=$(yq --version 2>&1) ;;
        controller-gen) out=$(controller-gen --version 2>&1) ;;
        *)
            out=$("${binary}" version 2>&1)
            found=$(grep -oE '[0-9]+\.[0-9]+\.[0-9]+' <<< "${out}" | head -1)
            [[ -z "${found}" ]] && out=$("${binary}" --version 2>&1)
            ;;
    esac
    grep -oE '[0-9]+\.[0-9]+\.[0-9]+' <<< "${out}" | head -1
}

wanted() {
    [[ "${#WANTED[@]}" -eq 0 ]] && return 0
    local name
    for name in "${WANTED[@]}"; do
        [[ "${name}" == "$1" || "${name}" == "$2" ]] && return 0
    done
    return 1
}

drift=0
missing=0
seen=()

# fd 3: the version commands inherit stdin and would consume the list.
while read -r tool want <&3; do
    binary="$(binary_for "${tool}")"
    wanted "${tool}" "${binary}" || continue
    seen+=("${tool}" "${binary}")

    if ! command -v "${binary}" &> /dev/null; then
        echo "  ✗ ${binary} not found on PATH (want ${want})"
        missing=1
        continue
    fi

    got="$(installed_version "${binary}")"
    if [[ "${got}" != "${want}" ]]; then
        echo "  ✗ ${binary} is ${got:-unknown}, mise.toml wants ${want}"
        drift=1
    fi
done 3< <("${REPO_ROOT}/scripts/mise-pins.sh")

# A typo would otherwise check nothing and pass.
for name in ${WANTED[@]+"${WANTED[@]}"}; do
    match=0
    for entry in ${seen[@]+"${seen[@]}"}; do
        [[ "${entry}" == "${name}" ]] && match=1 && break
    done
    if [[ "${match}" -eq 0 ]]; then
        echo "ERROR: '${name}' is not pinned in mise.toml" >&2
        exit 2
    fi
done

if [[ "${drift}" -eq 0 && "${missing}" -eq 0 ]]; then
    echo "  ✅ all required tools"
    exit 0
fi

echo ""
echo "Install the pinned versions with 'mise install' (https://mise.jdx.dev)"

[[ "${STRICT}" -eq 1 ]] && exit 1
exit 0
