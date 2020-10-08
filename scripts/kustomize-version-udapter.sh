#!/bin/bash
set -e

sed -i -e "s/newTag: */newTag: $1/g" kustomize/kustomization.yaml
git commit -sam "updates kustomize with newly released version"
