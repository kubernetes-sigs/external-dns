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

# renovate: datasource=github-releases depName=golangci/golangci-lint
GOLANG_CI_LINTER_VERSION=v2.7.2

# Execute
# scripts/install-tools.sh
# scripts/install-tools.sh -h
# scripts/install-tools.sh --generator
# scripts/install-tools.sh --golangci

show_help() {
cat << EOF
'external-dns' helm linter helper commands

Usage: $(basename "$0") <options>
    -h, --help          Display help
    --generator         Install generator
    --golangci          Install golangci linter
EOF
}

install_golangci() {
  local install=false
  if [[ -x $(which golangci-lint) ]]; then
      local version=$(golangci-lint version --short)
      if [[ "${version}" == "${GOLANG_CI_LINTER_VERSION#v}" ]]; then
          install=false
        else
          install=true
      fi
    else
      install=true
  fi
  if [[ "$install" == true ]]; then
      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/9f61b0f53f80672872fced07b6874397c3ed197b/install.sh \
        | sh -s -- -b $(go env GOPATH)/bin "${GOLANG_CI_LINTER_VERSION}"
  fi
}

function main() {
  case $1 in
    --golangci)
      install_golangci
      ;;
    -h|--help)
      show_help
      ;;
    *)
      echo "unknown sub-command" >&2
      show_help
      exit 1
      ;;
  esac
}

main "$@"
