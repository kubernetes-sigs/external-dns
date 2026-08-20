#!/bin/bash
set -euo pipefail

# Conventional-commit matchers, anchored on the "- " list marker.
BREAKING_RE='^- [a-z]+(\([^)]*\))?!: '
FEAT_RE='^- feat(\([^)]*\))?: '
FIX_RE='^- fix(\([^)]*\))?: '
DOCS_RE='^- docs(\([^)]*\))?: '

# The Helm chart has its own release cycle.
CHART_RE='charts?|helm'

# "feat(aws): add X" -> "[aws] add X" ; "feat: add X" -> "add X"
function label_by_scope {
  sed -E -e 's/^- [a-z]+\(([^)]*)\)!?: */- [\1] /' -e 's/^- [a-z]+!?: */- /'
}

# Same, but "Others" keeps the type: "ci(e2e): use X" -> "[ci] [e2e] use X".
# Except "chore", which carries no meaning.
function label_by_type {
  sed -E \
    -e 's/^- chore\(([^)]*)\)!?: */- [\1] /' \
    -e 's/^- chore!?: */- /' \
    -e 's/^- ([a-z]+)\(([^)]*)\)!?: */- [\1] [\2] /' \
    -e 's/^- ([a-z]+)!?: */- [\1] /'
}

# section <title> <formatter> <grep args...>
function section {
  local title="$1" formatter="$2"
  shift 2
  local body
  body=$(grep -E "$@" "${MERGED_PRS}" | "${formatter}" | sort) || true
  [ -n "${body}" ] || return 0
  printf '\n%s\n\n%s\n' "${title}" "${body}"
}

function generate_changelog {
  section "## :warning: Breaking Changes" label_by_scope "${BREAKING_RE}"
  section "## :rocket: Features" label_by_scope "${FEAT_RE}"
  section "## :bug: Bug fixes" label_by_scope "${FIX_RE}"
  section "## :memo: Documentation" label_by_scope "${DOCS_RE}"

  printf '\n## :package: Docker Image\n\n'
  echo '```sh'
  echo "# This pull command only works when it's released"
  echo "docker pull registry.k8s.io/external-dns/external-dns:${VERSION}"
  echo '```'

  section "## :package: Others" label_by_type -v -e "${BREAKING_RE}" -e "${FEAT_RE}" -e "${FIX_RE}" -e "${DOCS_RE}"
}

LATEST=$(gh release list -L 10 --json name,isLatest,publishedAt --jq '.[] | select(.isLatest) | "\(.name)\t\(.publishedAt)\t\(.publishedAt | fromdateiso8601)"')
IFS=$'\t' read -r RELEASE_NAME RELEASE_DATE TIMESTAMP <<< "${LATEST}"

if [ $# -ne 1 ]; then
    echo "** DRY RUN **"
fi

printf "Latest release: %s (%s)\n" "${RELEASE_NAME}" "${RELEASE_DATE}"

ALL_PRS=$(mktemp)
MERGED_PRS=$(mktemp)
trap 'rm -f "${ALL_PRS}" "${MERGED_PRS}"' EXIT

gh pr list \
  --state merged \
  --json author,number,mergeCommit,mergedAt,url,title \
  --limit 999 \
  --jq ".[] |
    select (.mergedAt | fromdateiso8601 > ${TIMESTAMP}) | \
    \"- \(.title | sub(\"^\\\\s+\";\"\") | sub(\"\\\\s+$\";\"\")) by @\(.author.login) in #\(.number)\"
  " | sort > "${ALL_PRS}"

grep -viwE "${CHART_RE}" "${ALL_PRS}" > "${MERGED_PRS}" || true

# On stderr, so it never lands in the release body.
if grep -iwE "${CHART_RE}" "${ALL_PRS}" >&2; then
  printf "^^ excluded from the changelog: chart-only changes, released separately\n\n" >&2
fi

if [ $# -ne 1 ]; then
  export VERSION="v0.x.0"
  generate_changelog
  echo
  echo "** DRY RUN **"
  echo
  echo "To create a release: ./releaser.sh v0.x.0"
else
  export VERSION="$1"
  generate_changelog | gh release create "${VERSION}" -t "${VERSION}" -p -F -
fi
