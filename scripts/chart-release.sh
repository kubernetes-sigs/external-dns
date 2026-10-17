#!/bin/bash
# Publish the external-dns Helm chart as a GitHub release, then update the chart
# index. Works around chart-releaser not supporting GitHub immutable releases:
# the asset must be attached while the release is still a draft, then the draft
# is published. Drop this script and restore chart-releaser-action once
# https://github.com/helm/chart-releaser/issues/591 lands.
#
# Usage: scripts/chart-release.sh <version> <release-notes-file>
#
# Requires cr, gh and jq on PATH. CI provides GITHUB_REPOSITORY, GITHUB_SHA and
# GH_TOKEN / CR_TOKEN.

set -euo pipefail

VERSION="${1:?version required}"
RELEASE_NOTES_PATH="${2:?release notes file required}"

REPO="${GITHUB_REPOSITORY:-kubernetes-sigs/external-dns}"
OWNER="${REPO%%/*}"
NAME="${REPO##*/}"

RELEASE_TAG="external-dns-helm-chart-${VERSION}"
CHART_ASSET=".cr-release-packages/external-dns-${VERSION}.tgz"
ASSET_NAME="$(basename "${CHART_ASSET}")"

# retry <attempts> <command...> with exponential backoff.
retry() {
  local attempts="$1"; shift
  local delay=5 attempt
  for attempt in $(seq 1 "${attempts}"); do
    "$@" && return 0
    [[ "${attempt}" == "${attempts}" ]] && return 1
    sleep "${delay}"; delay=$((delay * 2))
  done
}

package_chart() {
  rm -rf .cr-release-packages .cr-index
  mkdir -p .cr-release-packages .cr-index
  cr package charts/external-dns --package-path .cr-release-packages
}

# Echo {id, draft} for the release, creating a draft release if it is missing.
ensure_release() {
  if gh release view "${RELEASE_TAG}" --repo "${REPO}" \
       --json databaseId,isDraft --jq '{id: .databaseId, draft: .isDraft}' 2>/dev/null; then
    return 0
  fi
  gh api -X POST "repos/${REPO}/releases" \
    -f tag_name="${RELEASE_TAG}" \
    -f target_commitish="${GITHUB_SHA}" \
    -f name="${RELEASE_TAG}" \
    -f body="$(<"${RELEASE_NOTES_PATH}")" \
    -F draft=true -f make_latest=false \
    --jq '{id: .id, draft: .draft}'
}

# True if the chart asset is fully uploaded to the given release.
asset_ready() {
  local asset
  asset="$(gh api "repos/${REPO}/releases/${1}/assets" \
    --jq ".[] | select(.name == \"${ASSET_NAME}\")" | jq -s '.[0] // empty')"
  [[ -n "${asset}" ]] && jq -e '.state == "uploaded" and .size > 0' <<<"${asset}" >/dev/null
}

upload_asset() {
  asset_ready "${1}" && { echo "Asset ${ASSET_NAME} already attached."; return 0; }
  # --clobber replaces any partial upload left by a previous failed attempt.
  gh release upload "${RELEASE_TAG}" "${CHART_ASSET}" --repo "${REPO}" --clobber
}

publish_release() {
  if [[ "$(gh api "repos/${REPO}/releases/${1}" --jq '.draft')" == "false" ]]; then
    echo "Release ${RELEASE_TAG} already published."
    return 0
  fi
  gh api -X PATCH "repos/${REPO}/releases/${1}" -F draft=false -f make_latest=false >/dev/null
}

update_index() {
  retry 5 cr index \
    --owner "${OWNER}" --git-repo "${NAME}" \
    --package-path .cr-release-packages \
    --release-name-template "external-dns-helm-chart-{{ .Version }}" \
    --push
}

main() {
  package_chart

  local release release_id is_draft
  release="$(retry 5 ensure_release)"
  release_id="$(jq -r '.id' <<<"${release}")"
  is_draft="$(jq -r '.draft' <<<"${release}")"

  if [[ "${is_draft}" == "true" ]]; then
    retry 5 upload_asset "${release_id}"
    retry 5 publish_release "${release_id}"
  elif retry 5 asset_ready "${release_id}"; then
    echo "Release ${RELEASE_TAG} already published with ${ASSET_NAME}; updating index."
  else
    echo "Release ${RELEASE_TAG} published without ${ASSET_NAME}; refusing to modify immutable release." >&2
    exit 1
  fi

  update_index
}

main "$@"
