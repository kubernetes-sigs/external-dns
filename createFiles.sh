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
  file="main_${p}.go"
  # tail -n +16 "$file" > temp_file && mv temp_file "$file"
  multiline_text=$(cat <<EOF
//go:build all || $p
// +build all $p
EOF
)
echo "$multiline_text" | cat - "$file" > temp_file && mv temp_file "$file"
  
done
