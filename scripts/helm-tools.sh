#!/bin/bash

set -e

# JSON Schema https://json-schema.org/
# JSON Schema spec https://json-schema.org/draft/2020-12/json-schema-validation
# Helm Schema https://helm.sh/docs/topics/charts/#schema-files

# Execute
# scripts/helm-tools.sh
# scripts/helm-tools.sh -h
# scripts/helm-tools.sh --diff
# scripts/helm-tools.sh --install
# scripts/helm-tools.sh --schema
# scripts/helm-tools.sh --lint

show_help() {
cat << EOF
'external-dns' helm linter helper commands

Usage: $(basename "$0") <options>
    -d, --diff          Run schema diff validation
    --docs              Show available documentation
    -h, --help          Display help
    -i, --install       Install required tooling
    -l, --lint          Lint chart
    -s, --schema        Generate schema
EOF
}

install() {
  if [[ -x $(which helm) ]]; then
      helm plugin install https://github.com/losisin/helm-values-schema-json.git | true
      helm plugin list | grep "schema"
    else
      echo "helm is not installed"
      echo "install helm https://helm.sh/docs/intro/install/ and try again"
      exit 1
  fi
}

update_schema() {
  cd charts/external-dns
  helm schema  -indent 2 \
    -draft 7 \
    -input values.yaml \
    -input ci/schema-values.yaml \
    -output values.schema.json
}

diff_schema() {
  cd charts/external-dns
  helm schema  -indent 2 \
    -draft 7 \
    -input values.yaml \
    -input ci/schema-values.yaml \
    -output diff-schema.schema.json
  trap 'rm -rf -- "diff-schema.schema.json"' EXIT
  CURRENT_SCHEMA=$(cat values.schema.json)
  GENERATED_SCHEMA=$(cat diff-schema.schema.json)
  if [ "$CURRENT_SCHEMA" != "$GENERATED_SCHEMA" ]; then
    echo "Schema must be re-generated! Run 'scripts/helm-tools.sh --schema'" 1>&2
    diff diff-schema.schema.json values.schema.json
    exit 1
  fi
}

lint_chart() {
  cd charts/external-dns
  helm lint . --debug --strict
}

function show_docs() {
  open "https://github.com/losisin/helm-values-schema-json?tab=readme-ov-file"
}

function main() {
  case $1 in
    --docs)
      show_docs
      ;;
    -d|--diff)
      diff_schema
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
