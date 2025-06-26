#!/bin/bash

set -e

# JSON Schema https://json-schema.org/
# JSON Schema spec https://json-schema.org/draft/2020-12/json-schema-validation
# Helm Schema https://helm.sh/docs/topics/charts/#schema-files

# Execute
# scripts/helm-tools.sh
# scripts/helm-tools.sh -h
# scripts/helm-tools.sh --install
# scripts/helm-tools.sh --diff
# scripts/helm-tools.sh --schema
# scripts/helm-tools.sh --lint
# scripts/helm-tools.sh --docs
# scripts/helm-tools.sh --helm-template
# scripts/helm-tools.sh --helm-unittest

show_help() {
cat << EOF
'external-dns' helm linter helper commands

Usage: $(basename "$0") <options>
    -d, --diff          Schema diff validation
    --docs              Re-generate helm documentation
    -h, --help          Display help
    -i, --install       Install required tooling
    -l, --lint          Lint chart
    -s, --schema        Generate schema
    --helm-unittest     Run helm unittest(s)
    --helm-template     Run helm template
    --show-docs         Show available documentation
EOF
}

install() {
  if [[ -x $(which helm) ]]; then
      echo "installing https://github.com/losisin/helm-values-schema-json.git plugin"
      helm plugin install https://github.com/losisin/helm-values-schema-json.git | true
      helm plugin update schema
      helm plugin list | grep "schema"

      helm plugin install https://github.com/helm-unittest/helm-unittest.git | true
      helm plugin update unittest
      helm plugin list | grep "unittest"

      echo "installing helm-docs"
      go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest | true

      if [[ -x $(which brew) ]]; then
        echo "installing chart-testing https://github.com/helm/chart-testing"
        brew install chart-testing
      fi
    else
      echo "helm is not installed"
      echo "install helm https://helm.sh/docs/intro/install/ and try again"
      exit 1
  fi
}

update_schema() {
  cd charts/external-dns
  # uses .schema.yamle
  helm schema
}

diff_schema() {
  cd charts/external-dns
  helm schema  \
    -output diff-schema.schema.json
  trap 'rm -rf -- "diff-schema.schema.json"' EXIT
  CURRENT_SCHEMA=$(cat values.schema.json)
  GENERATED_SCHEMA=$(cat diff-schema.schema.json)
  if [ "$CURRENT_SCHEMA" != "$GENERATED_SCHEMA" ]; then
    echo "Schema must be re-generated! Run 'scripts/helm-tools.sh --schema'" 1>&2
    diff -Nau diff-schema.schema.json values.schema.json
    exit 1
  fi
}

lint_chart() {
  cd charts/external-dns
  helm lint . --debug --strict \
  --values values.yaml \
  --values ci/ci-values.yaml
  # lint with chart testing tool
  ct lint --target-branch=master --check-version-increment=false
}

helm_docs() {
  cd charts/external-dns
  helm-docs
}

helm_unittest() {
  helm unittest -f 'tests/*_test.yaml' --color charts/external-dns
}

helm_template() {
  helm template external-dns charts/external-dns \
		--output-dir _scratch \
		-n kube-system
}

show_docs() {
  open "https://github.com/losisin/helm-values-schema-json?tab=readme-ov-file"
}

function main() {
  case $1 in
    --show-docs)
      show_docs
      ;;
    --helm-unittest)
      helm_unittest
      ;;
    --helm-template)
      helm_template
      ;;
    -d|--diff)
      diff_schema
      ;;
    --docs)
      helm_docs
      ;;
    -i|--install)
      install
      ;;
    -l|--lint)
      lint_chart
      ;;
    -s|--schema)
      update_schema
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
