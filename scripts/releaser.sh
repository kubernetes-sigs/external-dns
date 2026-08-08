#!/bin/bash
set -e



function generate_changelog {
  MERGED_PRS="$1"

  echo
  echo "## :warning: Breaking Changes"
  echo
  cat "${MERGED_PRS}" | grep "\!" || true # no breaking change, section should be removed.

  echo
  echo "## :rocket: Features"
  echo
  cat "${MERGED_PRS}" | grep feat[:\(]

  echo
  echo "## :bug: Bug fixes"
  echo
  cat "${MERGED_PRS}" | grep fix[:\(]

  echo
  echo "## :memo: Documentation"
  echo
  cat "${MERGED_PRS}" | grep docs[:\(]

  echo
  echo "## :package: Others"
  echo
  cat "${MERGED_PRS}" | grep -v "\!" | grep -v feat[:\(] | grep -v fix[:\(] | grep -v docs[:\(]

  echo
  echo "## :package: Docker Image"
  echo
  echo "\`\`\`sh"
  echo "# This pull command only works when it's released"
  echo "docker pull registry.k8s.io/external-dns/external-dns:${VERSION}"
  echo "\`\`\`"

}

function create_release {
  generate_changelog | sort # | gh release create "$1" -t "$1" -F -
}

function latest_release {
  gh release list -L 10 --json name,isLatest --jq '.[] | select(.isLatest)|.name'
}

function latest_release_date {
  gh release list -L 10 --json name,isLatest,publishedAt --jq '.[] | select(.isLatest)|.publishedAt'
}

function latest_release_ts {
  gh release list -L 10 --json name,isLatest,publishedAt --jq '.[] | select(.isLatest)|.publishedAt | fromdateiso8601'
}

if [ $# -ne 1 ]; then
    echo "** DRY RUN **"
fi

printf "Latest release: %s (%s)\n" $(latest_release) $(latest_release_date)

TIMESTAMP=$(latest_release_ts)
MERGED_PRS=$(mktemp)
gh pr list \
  --state merged \
  --json author,number,mergeCommit,mergedAt,url,title \
  --limit 999 \
  --jq ".[] |
    select (.mergedAt | fromdateiso8601 > ${TIMESTAMP}) | \
    \"- \(.title) by @\(.author.login) in #\(.number)\"
  " | sort > "${MERGED_PRS}"

if [ $# -ne 1 ]; then
  export VERSION="v0.x.0"
  generate_changelog "${MERGED_PRS}"
  echo "** DRY RUN **"
  echo
  echo "To create a release: ./releaser.sh v0.x.0"
else
  export VERSION="$1"
  generate_changelog "${MERGED_PRS}" | gh release create "${VERSION}" -t "${VERSION}" -p -F -
fi

rm -f "${MERGED_PRS}"
