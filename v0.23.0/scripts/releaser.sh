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

# section <title> <grep args...>
function section {
  local title="$1"
  shift
  local body
  body=$(grep -E "$@" "${MERGED_PRS}" | label_by_scope | sort) || true
  [ -n "${body}" ] || return 0
  printf '\n%s\n\n%s\n' "${title}" "${body}"
}

# Everything left, folded per type: the type is the summary, so entries keep only their scope.
function others {
  # "chore(deps)" is a dependency bump, not a chore: fold it with the "deps" type.
  local rest
  rest=$(grep -vE -e "${BREAKING_RE}" -e "${FEAT_RE}" -e "${FIX_RE}" -e "${DOCS_RE}" "${MERGED_PRS}" \
    | sed -E 's/^- chore\(deps\)/- deps/') || true
  [ -n "${rest}" ] || return 0

  printf '\n## :package: Others\n'
  local type
  for type in $(printf '%s\n' "${rest}" | sed -nE 's/^- ([a-z]+)[(:].*/\1/p' | sort -u); do
    printf '\n<details><summary>%s</summary>\n\n%s\n\n</details>\n' \
      "${type}" "$(printf '%s\n' "${rest}" | grep -E "^- ${type}[(:]" | label_by_scope | sort)"
  done

  # Titles that are not conventional commits at all, left as they are.
  local misc
  misc=$(printf '%s\n' "${rest}" | grep -vE '^- [a-z]+[(:]' | sort) || true
  if [ -n "${misc}" ]; then
    printf '\n%s\n' "${misc}"
  fi
}

function generate_changelog {
  section "## :warning: Breaking Changes" "${BREAKING_RE}"
  section "## :rocket: Features" "${FEAT_RE}"
  section "## :bug: Bug fixes" "${FIX_RE}"
  section "## :memo: Documentation" "${DOCS_RE}"

  printf '\n## :package: Docker Image\n\n'
  echo '```sh'
  echo "# This pull command only works when it's released"
  echo "docker pull registry.k8s.io/external-dns/external-dns:${VERSION}"
  echo '```'

  others
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
