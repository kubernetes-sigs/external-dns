#!/usr/bin/env bash
set -e

function cleanup {
  # clean up kubernetes objects
  kubectl --namespace "$NAMESPACE" delete ingress nginx --ignore-not-found
  kubectl --namespace "$NAMESPACE" delete service nginx --ignore-not-found
  kubectl --namespace "$NAMESPACE" delete deployment nginx --ignore-not-found

  kubectl delete namespace "$NAMESPACE" --ignore-not-found

  # remove all records in the target zone
  gcloud --project "$PROJECT" dns record-sets import /dev/null --zone "$ZONE" --delete-all-existing
}
# trap cleanup EXIT

# target google project and hosted zone
PROJECT="zalando-external-dns-test"
ZONE="external-dns-test-gcp-zalan-do"

# set scope for this test
NAMESPACE="external-dns-test"

# remove objects from previous run
# cleanup

# create testing namespace
kubectl create namespace "$NAMESPACE" || true

cat <<EOF | kubectl --namespace "$NAMESPACE" create -f - || true
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: nginx
spec:
  rules:
  - host: nginx-ingress.external-dns-test.gcp.zalan.do
    http:
      paths:
      - backend:
          serviceName: nginx
          servicePort: 80
EOF

# create simple deployment and service
kubectl --namespace "$NAMESPACE" run nginx --image=nginx --replicas=1 --port=80 || true
kubectl --namespace "$NAMESPACE" expose deployment nginx --port=80 --target-port=80 --type=ClusterIP || true

# wait until service gets external IP
LOAD_BALANCER="<no value>"
while [[ $LOAD_BALANCER == "<no value>" ]]
do
  sleep 5

  LOAD_BALANCER=$(kubectl --namespace "$NAMESPACE" get ingress nginx -o go-template --template='{{.status.loadBalancer.ingress}}')
  echo "waiting for load balancer"
done

# get the external IP
INGRESS_IP=$(kubectl --namespace "$NAMESPACE" get ingress nginx -o go-template --template='{{(index .status.loadBalancer.ingress 0).ip}}')
echo "ingress ip:" $INGRESS_IP

# run single external-dns sync loop
go run main.go --zone "$ZONE" --source "ingress" --provider "google" --google-project "$PROJECT" --namespace "$NAMESPACE" --once --dry-run=true --health-port 9092

# wait until DNS propagated
DNS_TARGET=""
while [[ $DNS_TARGET == "" ]]
do
  sleep 5

  DNS_TARGET=$(dig +short nginx-ingress.external-dns-test.gcp.zalan.do.)
  echo "dns target:" $DNS_TARGET
done

# check if resolved IP matches service IP
if [[ "$INGRESS_IP" == "$DNS_TARGET" ]]; then
  echo "success"
else
  echo "failure"
fi
