#!/usr/bin/env bash
# Collect external-dns version, startup args, and logs.
#
# Usage:
#   [NAMESPACE=external-dns] [SINCE=5m] ./collect-extdns-info.sh
#
# Examples:
#   ./collect-extdns-info.sh
#   NAMESPACE=my-ns ./collect-extdns-info.sh
#   SINCE=30m ./collect-extdns-info.sh
#   NAMESPACE=my-ns SINCE=1h ./collect-extdns-info.sh
set -euo pipefail

NS="${NAMESPACE:-external-dns}"
SINCE="${SINCE:-5m}"
OUT="extdns-info-$(date +%Y%m%d-%H%M%S).txt"

{
  echo "=== external-dns version ==="
  out=$(kubectl get pod -n "${NS}" \
    -l app.kubernetes.io/name=external-dns \
    -o jsonpath='{range .items[*]}{.metadata.name}{"\t"}{range .spec.containers[*]}{.image}{"\n"}{end}{end}' \
    2>/dev/null)
  echo "${out:-(not found)}"

  echo ""
  echo "=== external-dns startup args ==="
  out=$(kubectl get pod -n "${NS}" \
    -l app.kubernetes.io/name=external-dns \
    -o jsonpath='{range .items[*]}{.metadata.name}{"\n"}{range .spec.containers[*]}{range .args[*]}{@}{"\n"}{end}{end}{end}' \
    2>/dev/null)
  echo "${out:-(not found)}"

  echo ""
  echo "=== external-dns logs (last ${SINCE}) ==="
  out=$(kubectl logs -n "${NS}" \
    -l app.kubernetes.io/name=external-dns \
    --since="${SINCE}" --prefix=true 2>/dev/null)
  echo "${out:-(not found)}"

} | tee "${OUT}"

echo ""
echo "Saved to: ${OUT}"
echo "Review for sensitive data before sharing."
