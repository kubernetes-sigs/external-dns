FROM golang:1.24 as builder
WORKDIR /sigs.k8s.io/external-dns
COPY . .
RUN make build

FROM registry.access.redhat.com/ubi9/ubi-minimal:latest
COPY --from=builder /sigs.k8s.io/external-dns/build/external-dns /usr/bin/
ENTRYPOINT ["/usr/bin/external-dns"]
LABEL io.openshift.release.operator="true"
LABEL io.k8s.display-name="Kubernetes ExternalDNS" \
      io.k8s.description="Synchronizes exposed Kubernetes Services and Ingresses with DNS providers." \
      maintainer="<aos-network-edge-staff@redhat.com>"
