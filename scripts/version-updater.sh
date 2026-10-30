#!/bin/bash
set -e

PREV_TAG=$1
NEW_TAG=$2

sed -i.bak -e "s/newTag: .*/newTag: ${NEW_TAG}/g" kustomize/kustomization.yaml
rm -f kustomize/kustomization.yaml.bak
git add kustomize/kustomization.yaml

# Image tags only
# mapfile and bare `sed -i` are unavailable in macOS bash 3.2 / BSD sed
IMAGE_FILES=()
while IFS= read -r f; do
    IMAGE_FILES+=("$f")
done < <(git grep -l "external-dns:${PREV_TAG}" -- '*.md' 'docs/snippets/')
if [ ${#IMAGE_FILES[@]} -gt 0 ]; then
    sed -i.bak -e "s|external-dns:${PREV_TAG}|external-dns:${NEW_TAG}|g" "${IMAGE_FILES[@]}"
    for f in "${IMAGE_FILES[@]}"; do rm -f "$f.bak"; done
    git add "${IMAGE_FILES[@]}"
fi

sed -i.bak -e "s/EXT_DNS_VERSION=\"${PREV_TAG}\"/EXT_DNS_VERSION=\"${NEW_TAG}\"/g" docs/release.md
rm -f docs/release.md.bak
git add docs/release.md

git commit -sm "chore(release): updates kustomize & docs with ${NEW_TAG}"

