FROM registry.ddbuild.io/images/base/gbi-ubuntu_2204:release

USER root
RUN clean-apt install curl
COPY /go/src/github.com/sigs.k8s.io/external-dns/build/external-dns /usr/local/bin/external-dns

USER nobody

ENTRYPOINT ["/usr/local/bin/external-dns"]
