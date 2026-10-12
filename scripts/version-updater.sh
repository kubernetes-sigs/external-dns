#!/bin/bash
set -e

PREV_TAG=$1
NEW_TAG=$2

sed -i -e "s/newTag: .*/newTag: ${NEW_TAG}/g" kustomize/kustomization.yaml
git add kustomize/kustomization.yaml

# Image tags only
mapfile -t IMAGE_FILES < <(git grep -l "external-dns:${PREV_TAG}" -- '*.md' 'docs/snippets/')
if [ ${#IMAGE_FILES[@]} -gt 0 ]; then
    sed -i -e "s|external-dns:${PREV_TAG}|external-dns:${NEW_TAG}|g" "${IMAGE_FILES[@]}"
    git add "${IMAGE_FILES[@]}"
fi

sed -i -e "s/EXT_DNS_VERSION=\"${PREV_TAG}\"/EXT_DNS_VERSION=\"${NEW_TAG}\"/g" docs/release.md
git add docs/release.md

git commit -sm "chore(release): updates kustomize & docs with ${NEW_TAG}"
