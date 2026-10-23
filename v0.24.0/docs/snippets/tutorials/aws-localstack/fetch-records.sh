#!/bin/bash

set -e

# run docs/snippets/tutorials/aws-localstack/check-records.sh

export AWS_REGION=eu-west-1
export AWS_ACCESS_KEY_ID=foo
export AWS_SECRET_ACCESS_KEY=bar
export AWS_ENDPOINT_URL=http://127.0.0.1:32379

MATCH="${1:-}"   # optional positional argument to filter records by name

zones=$(aws route53 list-hosted-zones-by-name --query "HostedZones[].Id"  --output json)

echo "$zones" | jq -r '.[]' | while IFS= read -r hosted_zone_id; do
  zone=${hosted_zone_id#"/hostedzone/"}
  echo "Checking records for zone: $zone"

  if [ -z "$MATCH" ]; then
    # default behaviour (unchanged)
    aws route53 list-resource-record-sets \
      --hosted-zone-id "$zone" \
      --query "ResourceRecordSets[].{Name:Name, Type:Type, Value:ResourceRecords[*].Value, TTL:TTL}" \
      --output json
  else
    # filtered behaviour
    aws route53 list-resource-record-sets \
      --hosted-zone-id "$zone" \
      --query "ResourceRecordSets[?contains(Name, \`${MATCH}\`)].{Name:Name, Type:Type, Value:ResourceRecords[*].Value, TTL:TTL}" \
      --output json
  fi
done
