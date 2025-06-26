#!/bin/bash
set -e

PREV_TAG=$1
NEW_TAG=$2

sed -i -e "s/newTag: .*/newTag: $1/g" kustomize/kustomization.yaml
git add kustomize/kustomization.yaml

sed -i -e "s/${PREV_TAG}/${NEW_TAG}/g" *.md docs/*.md docs/*/*.md
git add *.md docs/*.md docs/*/*.md

git commit -sm "chore(release): updates kustomize & docs with ${NEW_TAG}"
