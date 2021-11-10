# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Copyright (C) 2020, 2021, Oracle Corporation and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

# builder image
ARG ARCH
FROM ghcr.io/oracle/oraclelinux:7-slim as builder
ARG ARCH

# Install golang via Oracle's yum servers
RUN yum update -y \
    && yum-config-manager --save --setopt=ol7_ociyum_config.skip_if_unavailable=true \
    && yum install -y oracle-golang-release-el7 \
    && yum-config-manager --enable ol7_developer_golang113 \
    && yum-config-manager --add-repo http://yum.oracle.com/repo/OracleLinux/OL7/developer/golang113/x86_64 \
    && yum install -y git gcc make golang-1.13.3-1.el7.x86_64 \
    && yum clean all \
    && rm -rf /var/cache/yum \
    && go version

# Compile to /usr/bin
ENV GOBIN=/usr/bin

# Set go path
ENV GOPATH=/go

WORKDIR /sigs.k8s.io/external-dns

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make test build.$ARCH

# final image
FROM ghcr.io/oracle/oraclelinux:7-slim

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
RUN yum update -y && yum-config-manager --enable ol7_u8_security_validation \
    && yum install -y openssl

COPY --from=builder /sigs.k8s.io/external-dns/build/external-dns /bin/external-dns

# COPY LICENSE and other files to the image
COPY LICENSE README.md THIRD_PARTY_LICENSES.txt SECURITY.md /licenses/

# Run as UID for nobody since k8s pod securityContext runAsNonRoot can't resolve the user ID:
# https://github.com/kubernetes/kubernetes/issues/40958
USER 65534

ENTRYPOINT ["/bin/external-dns"]

