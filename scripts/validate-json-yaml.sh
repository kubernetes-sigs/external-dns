#!/usr/bin/env bash
set -euo pipefail

# Validates all JSON and YAML files in the repository.
# Excludes Helm templates and mkdocs.yml which use non-standard syntax.
#
# Requirements (pre-installed on GitHub Actions ubuntu-latest):
#   - python3 (for JSON validation via json.tool)
#   - yq (for YAML validation, supports multi-document files)

EXCLUDE_PATTERN='(charts/external-dns/templates)'

errors=0

while IFS= read -r -d '' file; do
  if [[ "$file" =~ $EXCLUDE_PATTERN ]]; then
    continue
  fi
  if ! yq '.' "$file" > /dev/null 2>&1; then
    echo "FAIL: $file"
    errors=$((errors + 1))
  fi
done < <(find . \( -name '*.json' -o -name '*.yaml' -o -name '*.yml' \) -print0)

if [ "$errors" -gt 0 ]; then
  echo ""
  echo "$errors file(s) failed validation."
  exit 1
fi

echo "All JSON and YAML files are valid."
