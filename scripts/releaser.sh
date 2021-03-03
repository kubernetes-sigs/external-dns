#!/bin/bash
set -e

current_tag="${GITHUB_REF#refs/tags/}"
start_ref="HEAD"

function generate_changelog {
  # Find the previous release on the same branch, skipping prereleases if the
  # current tag is a full release
  previous_tag=""
  while [[ -z $previous_tag || ( $previous_tag == *-* && $current_tag != *-* ) ]]; do
    previous_tag="$(git describe --tags "$start_ref"^ --abbrev=0)"
    start_ref="$previous_tag"
  done

  git log "$previous_tag".. --reverse --merges --oneline --grep='Merge pull request #' | \
    while read -r sha title; do
      pr_num="$(grep -o '#[[:digit:]]\+' <<<"$title")"
      pr_desc="$(git show -s --format=%b "$sha" | sed -n '1,/^$/p' | tr $'\n' ' ')"
      pr_author="$(gh pr view "$pr_num" | grep author | awk '{ print $2 }' | tr $'\n' ' ')"
      printf "* %s (%s) @%s\n\n" "$pr_desc" "$pr_num" "$pr_author"
    done

  git log "$previous_tag".. --reverse --oneline --grep='(#' | \
    while read -r sha title; do
      pr_num="$(grep -o '#[[:digit:]]\+' <<<"$title")"
      pr_desc="$(git show -s --format=%s "$sha")"
      pr_author="$(gh pr view "$pr_num" | grep author | awk '{ print $2 }' | tr $'\n' ' ')"
      printf "* %s (%s) @%s\n\n" "$pr_desc" "$pr_num" "$pr_author"
    done
}

function create_release {
  generate_changelog | gh release create "$1" -t "$1" -F -
}

if [ $# -ne 1 ]; then
    echo "$0: usage: releaser [release number]"
    echo "example: ./releaser.sh v0.7.5"
    exit 1
fi

create_release "$1"
