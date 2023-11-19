#!/usr/bin/env bash

# Copyright 2022 The Kubernetes Authors.
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

providers=(
 akamai
 alibabacloud
 aws-sd
 aws
 azure-private-dns
 azure
 bluecat
 vinyldns
 vultr
 ultradns
 civo
 cloudflare
 rcodezero
 google
 digitalocean
 ovh
 linode
 dnsimple
 infoblox
 dyn
 coredns
 rdns
 exoscale
 inmemory
 designate
 pdns
 oci
 rfc2136
 ns1
 transip
 scaleway
 godaddy
 gandi
 pihole
 ibmcloud
 safedns
 plural
 tencentcloud
 webhook
)

TARGET=$1

# if target is "binary or image" then check for one arg passed. if target="push" then check for two args passed.
if [ "$TARGET" == "binary" ] || [ "$TARGET" == "image" ]; then
  if [ $# -ne 1 ]; then
    echo "Usage: $0 <binary|image>"
    exit 1
  fi
fi
if [ "$TARGET" == "push" ]; then
  if [ $# -ne 2 ]; then
    echo "Usage: $0 <push> <registry>"
    exit 1
  fi
  REGISTRY=$2
fi

if [ "$TARGET" == "binary" ]; then
  makeArgs="BUILD_FLAGS="
  makeTarget="build"
elif [ "$TARGET" == "image" ]; then
  makeArgs=""
  makeTarget="build.image/multiarch"
elif [ "$TARGET" == "push" ]; then
  makeTarget="build.push/multiarch"
  makeArgs="REGISTRY=$REGISTRY"
else
  echo "Usage: $0 <binary|image>"
  exit 1
fi

make clean
# all providers in one
make $makeArgs $makeTarget

# each provider individually
for p in ${providers[@]}; do
  make $makeArgs BUILD_TAGS=$p $makeTarget
done
