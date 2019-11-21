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

# builder image
FROM golang:1.13 as builder

#ARG  http_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"
#ARG https_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"

WORKDIR /github.com/kubernetes-sigs/external-dns

COPY . .
RUN go mod vendor && \
    make test && \
    make build

# final image
FROM alpine:3.10
LABEL maintainer="Team Teapot @ Zalando SE <team-teapot@zalando.de>"

#RUN http_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128" apk add ca-certificates && update-ca-certificates
RUN apk add ca-certificates && update-ca-certificates

RUN apk add --no-cache ca-certificates && \
    update-ca-certificates

COPY --from=builder /github.com/kubernetes-sigs/external-dns/build/external-dns /bin/external-dns

# Run as UID for nobody since k8s pod securityContext runAsNonRoot can't resolve the user ID:
# https://github.com/kubernetes/kubernetes/issues/40958
USER 65534

ENTRYPOINT ["/bin/external-dns"]
