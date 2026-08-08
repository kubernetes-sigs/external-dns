#!/usr/bin/env bash
# Collect Kubernetes resources relevant to your external-dns source type.
#
# Usage:
#   RESOURCE=<kubectl-resource> ./collect-resources.sh
#
# Examples:
#   RESOURCE=ingress ./collect-resources.sh
#   RESOURCE="ingress,service" ./collect-resources.sh
#   RESOURCE="gateway,httproute" ./collect-resources.sh
#   RESOURCE=dnsendpoint ./collect-resources.sh
set -euo pipefail

RESOURCE="${RESOURCE:?Usage: RESOURCE=<kubectl-resource> $0}"
OUT="extdns-resources-$(date +%Y%m%d-%H%M%S).txt"

{
  echo "=== ${RESOURCE} (all namespaces) ==="
  kubectl get "${RESOURCE}" -A -o yaml 2>/dev/null || echo "(not found)"

} | tee "${OUT}"

echo ""
echo "Saved to: ${OUT}"
echo "Review for sensitive data before sharing."
