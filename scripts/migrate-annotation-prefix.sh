#!/usr/bin/env bash
# migrate-annotation-prefix.sh — scan for resources that need annotation migration.
#
# WARNING
#   This script is provided as a possible migration aid and is used at your own risk.
#   It may not suit every operator's environment, workflow, or cluster configuration.
#   Review the output carefully and test against a non-production cluster first.
#
# BACKGROUND
#   The default annotation prefix changed from:
#     external-dns.alpha.kubernetes.io/
#   to:
#     external-dns.kubernetes.io/
#
#   Starting from the release that introduced this change, --annotation-prefix is a required flag.
#   Running external-dns with the new prefix against a cluster where resources still carry the old
#   prefix annotations causes silent DNS record deletion: the controller sees records it owns in the
#   registry but no source claiming them, and removes them.
#
#   This script performs a read-only scan and prints the kubectl annotate commands needed to migrate
#   resources before you switch the prefix in your external-dns deployment. No changes are made.
#
# WHAT IT DOES (example)
#   Given a Service with:
#     external-dns.alpha.kubernetes.io/hostname: my-app.example.com
#     external-dns.alpha.kubernetes.io/ttl:      300
#     external-dns.alpha.kubernetes.io/target:   1.2.3.4
#
#   The script prints the command to add:
#     external-dns.kubernetes.io/hostname: my-app.example.com
#     external-dns.kubernetes.io/ttl:      300
#     external-dns.kubernetes.io/target:   1.2.3.4
#
#   The old annotations are left in place (harmless — external-dns with the new prefix ignores them).
#   Unrelated annotations are never touched.
#
# USAGE
#   Scan default resource types (services):
#     ./migrate-annotation-prefix.sh
#
#   Scan additional resource types:
#     ./migrate-annotation-prefix.sh --resources=services,ingresses,nodes,pods
#
#   Flags:
#     --resources=<list>       comma-separated list of Kubernetes resource types to scan
#                              (default: services)
#                              common values: services, ingresses, nodes, pods
#
# REQUIREMENTS
#   kubectl, jq

set -euo pipefail

OLD_PREFIX="external-dns.alpha.kubernetes.io/"
NEW_PREFIX="external-dns.kubernetes.io/"
RESOURCES=("services")

for arg in "$@"; do
  case $arg in
    --resources=*)
      IFS=',' read -ra RESOURCES <<< "${arg#*=}"
      ;;
    *)
      echo "Unknown argument: $arg" >&2
      exit 1
      ;;
  esac
done

for cmd in kubectl jq; do
  if ! command -v "$cmd" &>/dev/null; then
    echo "Error: $cmd is required but not installed." >&2
    exit 1
  fi
done

echo "WARNING: This script is provided as a possible migration aid and is used at your own risk."
echo "         It may not suit every operator's environment. Test on a non-production cluster first."
echo ""
echo "Scanning for resources with prefix: $OLD_PREFIX"
echo ""

total_resources=0
total_annotations=0

for resource in "${RESOURCES[@]}"; do
  while IFS=$'\t' read -r ns name; do
    annotations=()
    while IFS= read -r annotation; do
      annotations+=("$annotation")
    done < <(
      kubectl get "$resource" -n "$ns" "$name" -o json | jq -r \
        --arg old "$OLD_PREFIX" \
        --arg new "$NEW_PREFIX" '
          .metadata.annotations
          | to_entries
          | map(select(.key | startswith($old)))
          | map("\(.key | gsub($old; $new))=\(.value)")
          | .[]
        '
    )

    if [ "${#annotations[@]}" -eq 0 ]; then
      continue
    fi

    total_resources=$((total_resources + 1))
    total_annotations=$((total_annotations + ${#annotations[@]}))

    echo "$resource/$name (namespace: $ns)"
    for annotation in "${annotations[@]}"; do
      echo "  $annotation"
    done
    echo ""
  done < <(
    kubectl get "$resource" --all-namespaces -o json | jq -r \
      --arg old "$OLD_PREFIX" '
        .items[]
        | select(any(.metadata.annotations // {} | keys[]; startswith($old)))
        | [.metadata.namespace, .metadata.name]
        | @tsv
      '
  )
done

if [ "$total_resources" -eq 0 ]; then
  echo "No resources found with prefix $OLD_PREFIX — nothing to migrate."
  exit 0
fi

echo "Found $total_resources resource(s) with $total_annotations annotation(s) to migrate."
