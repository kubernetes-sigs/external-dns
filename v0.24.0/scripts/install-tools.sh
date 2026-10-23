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
GOLANG_CI_LINTER_VERSION=v2.14.0
GOLANG_CI_LINTER_INSTALL_SCRIPT_COMMIT=114493f9b3e7257d29e4130f2b4a4aadefbb6845
GOLANG_CI_LINTER_INSTALL_SCRIPT_SHA256=1022ddb4d87ed252350ed03fc9677e250a4ae95cc6bcd4658c2a20a8a23d390f

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
      local script
      script=$(mktemp)
      curl -sSfL "https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANG_CI_LINTER_INSTALL_SCRIPT_COMMIT}/install.sh" -o "$script"
      echo "${GOLANG_CI_LINTER_INSTALL_SCRIPT_SHA256}  ${script}" | sha256sum --check --strict
      sh "$script" -b "$(go env GOPATH)/bin" "${GOLANG_CI_LINTER_VERSION}"
      rm -f "$script"
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
