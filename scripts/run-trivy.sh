#! /bin/bash

trivy image --exit-code 1 us.gcr.io/k8s-artifacts-prod/external-dns/external-dns:$(git describe --tags --always --dirty)
