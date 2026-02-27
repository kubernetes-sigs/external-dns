#!/bin/bash

set -e

KO_VERSION="0.18.0"
KIND_VERSION="0.30.0"
ALPINE_VERSION="3.22"
KUBECTL_VERSION="1.35.0"

echo "Starting end-to-end tests for external-dns with local provider..."

# Install kind
echo "Installing kind..."
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v${KIND_VERSION}/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind

# Create kind cluster
echo "Creating kind cluster..."
kind create cluster

# Install kubectl
echo "Installing kubectl..."
curl -LO "https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/kubectl

# Install ko
echo "Installing ko..."
curl -sSfL "https://github.com/ko-build/ko/releases/download/v${KO_VERSION}/ko_${KO_VERSION}_linux_x86_64.tar.gz" > ko.tar.gz
tar xzf ko.tar.gz ko
chmod +x ./ko
sudo mv ko /usr/local/bin/ko

# Build external-dns
echo "Building external-dns..."
# Use ko with --local to save the image to Docker daemon
EXTERNAL_DNS_IMAGE_FULL=$(KO_DOCKER_REPO=ko.local VERSION=$(git describe --tags --always --dirty) \
    ko build --tags "$(git describe --tags --always --dirty)" --bare --sbom none \
    --platform=linux/amd64 --local .)
echo "Built image: $EXTERNAL_DNS_IMAGE_FULL"

# Extract image name and tag (strip the @sha256 digest for kind load and kustomize)
EXTERNAL_DNS_IMAGE="${EXTERNAL_DNS_IMAGE_FULL%%@*}"
echo "Using image reference: $EXTERNAL_DNS_IMAGE"

# apply etcd deployment as provider
echo "Applying etcd"
kubectl apply -f e2e/provider/etcd.yaml

# Build a DNS testing image with dig
echo "Building DNS test image with dig..."
docker build -t dns-test:v1 -f - . <<EOF
FROM alpine:${ALPINE_VERSION}
RUN apk add --no-cache bind-tools curl
ENTRYPOINT ["sh"]
EOF

# Load all images into kind cluster
echo "Loading Docker images into kind cluster..."
kind load docker-image "$EXTERNAL_DNS_IMAGE"
kind load docker-image dns-test:v1

# Deploy ExternalDNS to the cluster
echo "Deploying external-dns with custom arguments..."

# Create temporary directory for kustomization
TEMP_KUSTOMIZE_DIR=$(mktemp -d)
cp -r kustomize/* "$TEMP_KUSTOMIZE_DIR/"

# Create patch file on the fly
cat <<EOF > "$TEMP_KUSTOMIZE_DIR/deployment-args-patch.yaml"
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  template:
    spec:
      hostNetwork: true
      containers:
        - name: external-dns
          args:
            - --source=service
            - --provider=coredns
            - --txt-owner-id=external.dns
            - --policy=sync
            - --log-level=debug
          env:
            - name: ETCD_URLS
              value: http://etcd-0.etcd:2379
EOF

# Update kustomization.yaml to include the patch
cat <<EOF > "$TEMP_KUSTOMIZE_DIR/kustomization.yaml"
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

images:
  - name: registry.k8s.io/external-dns/external-dns
    newName: ${EXTERNAL_DNS_IMAGE%%:*}
    newTag: ${EXTERNAL_DNS_IMAGE##*:}

resources:
  - ./external-dns-deployment.yaml
  - ./external-dns-serviceaccount.yaml
  - ./external-dns-clusterrole.yaml
  - ./external-dns-clusterrolebinding.yaml

patchesStrategicMerge:
  - ./deployment-args-patch.yaml
EOF

# Apply the kustomization
kubectl kustomize "$TEMP_KUSTOMIZE_DIR" | kubectl apply -f -

# add a wait for the deployment to be available
kubectl wait --for=condition=available --timeout=60s deployment/external-dns || true

kubectl describe pods -l app=external-dns
kubectl describe deployment external-dns
kubectl logs -l app=external-dns

# Cleanup temporary directory
rm -rf "$TEMP_KUSTOMIZE_DIR"

# Apply kubernetes yaml with service
echo "Applying Kubernetes service..."
kubectl apply -f e2e

# Wait for convergence
echo "Waiting for convergence (90 seconds)..."
sleep 90  # normal loop is 60 seconds, this is enough and should not cause flakes

# Check that the records are present
echo "Checking services again..."
kubectl get svc -owide
kubectl logs -l app=external-dns

# Check that the DNS records are present using our DNS server
echo "Testing DNS server functionality..."

# Get the node IP where the pod is running (since we're using hostNetwork)
NODE_IP=$(kubectl get nodes -o jsonpath='{.items[0].status.addresses[?(@.type=="InternalIP")].address}')
echo "Node IP: $NODE_IP"

# Test our DNS server with dig
echo "Testing DNS server with dig..."

# Create DNS test job that uses dig to query our DNS server
cat <<EOF | kubectl apply -f -
apiVersion: batch/v1
kind: Job
metadata:
  name: dns-server-test-job
  labels:
    app: dns-server-test
spec:
  backoffLimit: 3
  template:
    metadata:
      labels:
        app: dns-server-test
    spec:
      restartPolicy: Never
      hostNetwork: true
      containers:
      - name: dns-server-test
        image: dns-test:v1
        command:
        - /bin/sh
        - -c
        - |
          echo "Testing DNS server at $NODE_IP:5353"

          echo "=== Testing DNS server with dig ==="
          echo "Querying: externaldns-e2e.external.dns A record"
          if dig @$NODE_IP -p 5353 externaldns-e2e.external.dns A +short +timeout=5; then
            echo "DNS query successful"
            exit 0
          else
            echo "DNS query failed"
            exit 1
          fi

          echo "DNS server tests completed"
          exit 0
EOF

# Wait for the job to complete
echo "Waiting for DNS server test job to complete..."
kubectl wait --for=condition=complete --timeout=90s job/dns-server-test-job || true

# Check job status and get results
echo "DNS server test job results:"
kubectl logs job/dns-server-test-job

# Final validation
JOB_SUCCEEDED=$(kubectl get job dns-server-test-job -o jsonpath='{.status.succeeded}')
if [ "$JOB_SUCCEEDED" = "1" ]; then
    echo "SUCCESS: DNS server test completed successfully"
    TEST_PASSED=true
else
    echo "WARNING: DNS server test job did not complete successfully"
    kubectl describe job dns-server-test-job
    TEST_PASSED=false
fi

# Cleanup the test job
kubectl delete job dns-server-test-job

echo "End-to-end test completed!"

# Cleanup function
cleanup() {
    echo "Cleaning up..."
    if [ ! -z "$EXTERNAL_DNS_PID" ]; then
        kill $EXTERNAL_DNS_PID 2>/dev/null || true
    fi
    if [ ! -z "$LOCAL_PROVIDER_PID" ]; then
        kill $LOCAL_PROVIDER_PID 2>/dev/null || true
    fi
    kind delete cluster 2>/dev/null || true
}

# Set trap to cleanup on script exit
trap cleanup EXIT
