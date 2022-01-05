#! /bin/bash
set -e

# install trivy
cd /tmp
curl -LO https://github.com/aquasecurity/trivy/releases/download/v0.20.2/trivy_0.20.2_Linux-64bit.tar.gz 
echo "38a6de48e21a34e0fa0d2cf63439c0afcbbae0e78fb3feada7a84a9cf6e7f60c trivy_0.20.2_Linux-64bit.tar.gz" | sha256sum -c 
tar -xvf trivy_0.20.2_Linux-64bit.tar.gz
chmod +x trivy

# run trivy
cd - 
/tmp/trivy image --exit-code 1 us.gcr.io/k8s-artifacts-prod/external-dns/external-dns:$(git describe --tags --always --dirty)
