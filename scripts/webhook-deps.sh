#!/usr/bin/env bash
#
# webhook-deps.sh fails when the packages imported by out-of-tree webhook
# providers transitively pull in the Kubernetes client stack. depguard only
# sees direct imports, so a leak through any intermediate package slips past it.

set -eu -o pipefail

cd "$(git rev-parse --show-toplevel)"

PACKAGES="./endpoint ./plan ./provider ./provider/webhook/api"
DENY='^(k8s\.io/client-go|k8s\.io/api/|k8s\.io/apimachinery/pkg/apis|sigs\.k8s\.io/controller-runtime|sigs\.k8s\.io/external-dns/(source|pkg/apis|pkg/events))(/|$)'

# shellcheck disable=SC2086
leaks=$(go list -deps $PACKAGES | grep -E "$DENY" || true)

if [ -n "$leaks" ]; then
  echo "::error::webhook-facing packages pull the Kubernetes client stack:" >&2
  echo "$leaks" >&2
  echo "Find the importer with: go list -f '{{.ImportPath}}: {{join .Imports \" \"}}' -deps $PACKAGES | grep <pkg>" >&2
  exit 1
fi

echo "webhook-facing packages are free of the Kubernetes client stack"
