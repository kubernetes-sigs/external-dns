FROM registry.ddbuild.io/images/base/gbi-ubuntu_2204:release

USER root
RUN clean-apt install curl
COPY build/external-dns /usr/local/bin/external-dns

USER nobody

ENTRYPOINT ["/usr/local/bin/external-dns"]
