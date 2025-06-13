FROM registry.ci.openshift.org/ocp/builder:rhel-9-golang-1.22-openshift-4.17 AS builder
WORKDIR /sigs.k8s.io/external-dns
COPY . .
RUN make build

FROM registry.ci.openshift.org/ocp/4.17:base-rhel9
COPY --from=builder /sigs.k8s.io/external-dns/build/external-dns /usr/bin/
ENTRYPOINT ["/usr/bin/external-dns"]
LABEL io.openshift.release.operator="true"
LABEL io.k8s.display-name="Kubernetes ExternalDNS" \
      io.k8s.description="Synchronizes exposed Kubernetes Services and Ingresses with DNS providers." \
      maintainer="<aos-network-edge-staff@redhat.com>"
