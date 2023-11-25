#!/bin/bash

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

for p in ${providers[@]}; do
  for file in provider/$p/*.go; do
    echo "Processing file: $file"
    # ... perform operations on the file ...
    replace1=""
    with1="//go:build all \|\| $p\n// +build all $p"
    sed -i "1i\
    $with1" "$file"
  done
done
