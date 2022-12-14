#!/bin/bash
set -e

sed -i -e "s/newTag: .*/newTag: $1/g" kustomize/kustomization.yaml
git add kustomize/kustomization.yaml
git commit -sm "updates kustomize with newly released version"
