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

# Prints "<tool> <version>" per line from the [tools] table of mise.toml.
# Other tables are skipped, so their keys are not read as tools.
#
# Execute
#   scripts/mise-pins.sh

set -euo pipefail

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PINS="${REPO_ROOT}/mise.toml"

if [[ ! -f "${PINS}" ]]; then
    echo "ERROR: ${PINS} not found" >&2
    exit 1
fi

awk '
    # Any table header ends [tools].
    /^[[:space:]]*\[/ { in_tools = ($0 ~ /^[[:space:]]*\[tools\][[:space:]]*(#.*)?$/); next }
    !in_tools { next }

    {
        line = $0
        sub(/[[:space:]]*#.*$/, "", line)
        eq = index(line, "=")
        if (eq == 0) next

        key = substr(line, 1, eq - 1)
        val = substr(line, eq + 1)
        gsub(/^[[:space:]]+|[[:space:]]+$/, "", key)
        gsub(/^[[:space:]]+|[[:space:]]+$/, "", val)
        gsub(/"/, "", key)

        # tool = "1.2.3" and tool = { version = "1.2.3", ... } are both valid.
        if (val ~ /^\{/) {
            if (!match(val, /version[[:space:]]*=[[:space:]]*"[^"]+"/)) next
            val = substr(val, RSTART, RLENGTH)
            sub(/^version[[:space:]]*=[[:space:]]*"/, "", val)
            sub(/"$/, "", val)
        } else {
            gsub(/"/, "", val)
        }

        if (key != "" && val != "") print key, val
    }
' "${PINS}"
