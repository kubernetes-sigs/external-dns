#!/usr/bin/env bash
# migrate-annotation-prefix.sh — migrate external-dns annotations from the legacy prefix to the GA prefix.
#
# WARNING
#   This script is provided as a possible migration aid and is used at your own risk.
#   It may not suit every operator's environment, workflow, or cluster configuration.
#   Review it carefully, test it against a non-production cluster first, and ensure you
#   have a backup or rollback strategy before running it with --apply in production.
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
#   This script performs the one-time annotation migration before you switch the prefix in your
#   external-dns deployment.
#
# WHAT IT DOES (example)
#   Given a Service with:
#     external-dns.alpha.kubernetes.io/hostname: my-app.example.com
#     external-dns.alpha.kubernetes.io/ttl:      300
#     external-dns.alpha.kubernetes.io/target:   1.2.3.4
#
#   The script adds:
#     external-dns.kubernetes.io/hostname: my-app.example.com
#     external-dns.kubernetes.io/ttl:      300
#     external-dns.kubernetes.io/target:   1.2.3.4
#
#   The old annotations are left in place (harmless — external-dns with the new prefix ignores them).
#   Unrelated annotations are never touched.
#
# USAGE
#   Dry-run (default — no changes made, prints what would happen):
#     ./migrate-annotation-prefix.sh
#
#   Apply changes:
#     ./migrate-annotation-prefix.sh --apply
#
#   Restrict to specific resource types (comma-separated):
#     ./migrate-annotation-prefix.sh --resources=services,ingresses --apply
#
# REQUIREMENTS
#   kubectl, jq

set -euo pipefail

OLD_PREFIX="external-dns.alpha.kubernetes.io/"
NEW_PREFIX="external-dns.kubernetes.io/"
APPLY=false
RESOURCES=("services" "ingresses")

for arg in "$@"; do
  case $arg in
    --apply)
      APPLY=true
      ;;
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

if [ "$APPLY" = false ]; then
  echo "DRY-RUN mode — no changes will be made. Pass --apply to apply."
  echo ""
fi

total_resources=0
total_annotations=0

for resource in "${RESOURCES[@]}"; do
  echo "Scanning $resource..."

  while IFS=$'\t' read -r ns name; do
    total_resources=$((total_resources + 1))
    echo "  $resource $name (namespace: $ns)"

    while IFS= read -r annotation; do
      total_annotations=$((total_annotations + 1))
      echo "    + $annotation"
      if [ "$APPLY" = true ]; then
        kubectl annotate "$resource" -n "$ns" "$name" --overwrite "$annotation"
      fi
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

echo ""
if [ "$total_resources" -eq 0 ]; then
  echo "No resources found with prefix $OLD_PREFIX — nothing to migrate."
  exit 0
fi

echo "Found $total_resources resource(s) with $total_annotations annotation(s) to migrate."

if [ "$APPLY" = false ]; then
  echo "Re-run with --apply to apply the changes."
else
  echo "Migration complete."
fi
