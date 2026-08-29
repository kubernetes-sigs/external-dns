#!/usr/bin/env bash
# Compares Go major.minor version in cloudbuild.yaml against go.mod and mise.toml

set -euo pipefail

if ! command -v yq &>/dev/null; then
    echo "ERROR: yq is not installed" >&2
    exit 1
fi

REPO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

extract_major_minor() {
    [[ $1 =~ ([0-9]+\.[0-9]+) ]] && echo "${BASH_REMATCH[1]}"
}

image_name=$(yq '.steps[].name | select(. == "docker.io/library/golang*")' "$REPO_ROOT/cloudbuild.yaml" | head -1)
if [[ -z "$image_name" ]]; then
    echo "ERROR: Go version not found in cloudbuild.yaml" >&2
    exit 1
fi

cloudbuild_ver=$(extract_major_minor "${image_name##*:}")
gomod_ver=$(extract_major_minor "$(awk '/^go [0-9]/{print $2; exit}' "$REPO_ROOT/go.mod")")
misetoml_ver=$(extract_major_minor "$(awk -F' *= *' '$1 == "golang" { gsub(/"/, "", $2); print $2; exit }' "$REPO_ROOT/mise.toml")")

echo "cloudbuild.yaml : $cloudbuild_ver"
echo "go.mod          : $gomod_ver"
echo "mise.toml       : $misetoml_ver"

ok=true

if [[ "$cloudbuild_ver" != "$gomod_ver" ]]; then
    echo "Go version in cloudbuild.yaml ($cloudbuild_ver) doesn't match go.mod ($gomod_ver)"
    ok=false
fi

if [[ "$cloudbuild_ver" != "$misetoml_ver" ]]; then
    echo "Go version in cloudbuild.yaml ($cloudbuild_ver) doesn't match mise.toml ($misetoml_ver)"
    ok=false
fi

if [[ "$gomod_ver" != "$misetoml_ver" ]]; then
    echo "Go version in go.mod ($gomod_ver) doesn't match mise.toml ($misetoml_ver)"
    ok=false
fi

if $ok; then
    echo "OK: all versions match ($cloudbuild_ver)"
else
    exit 1
fi
