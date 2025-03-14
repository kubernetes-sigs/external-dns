#!/usr/bin/env python

# Copyright 2025 The Kubernetes Authors.
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

# This is a script that we wrote to try to help the migration over to using external-dns.
# This script looks at kubernetes ingresses and services (which are the two things we have
# external-dns looking at) and compares them to existing TXT and A records in route53 to
# find out where there are gaps.  It then assigns the heritage and owner TXT records where
# needed so external-dns can take over managing those resources. You can modify the script
# to only look at one or the other if needed.

# To run script
# 1. Python, pip and pipenv installed https://pipenv.pypa.io/en/latest/
# 2. AWS Access https://docs.aws.amazon.com/signin/latest/userguide/command-line-sign-in.html
# 3. pipenv shell
# 4. pip install boto3
# 5. python scripts/cleanup-legacy-txt-records.py --help
# WARNING: run this script at your own RISK
# 6. DRY RUN python scripts/cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF
# 7. Execute Deletion. First few times with reduced number of items
# - python scripts/cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF --total-items 3 --batch-delete-count 1 --run
# - python scripts/cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF --total-items 10000 --batch-delete-count 50 --run

# python scripts/cleanup-legacy-txt-records.py --help
# python scripts/cleanup-legacy-txt-records.py --zone-id Z06155043AVN8RVC88TYY --total-items 3 --batch-delete-count 1 --record-contain "external-dns/owner=default"

import boto3
from botocore.config import Config
import json, argparse, os, uuid, time

MAX_ITEMS=300 # max is 300 https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/route53/client/list_resource_record_sets.html
SLEEP=1 # in seconds, required to make sure Route53 API is not throttled
OPERATION_UUID=uuid.uuid4()

def json_prettify(data):
    return json.dumps(data, indent=4, default=str)

class Record:

    def __init__(self, record):
        # static
        self.type = 'TXT'
        self.record = record
        self.name = record['Name']
        self.resource_records = record['ResourceRecords']
        resource_record = ''
        for r in self.resource_records:
            resource_record += r['Value']
        self.resource_record = resource_record

    def __str__(self):
        return f'record: name: {self.name}, type: {self.type}, records: {self.resource_record} '

def records(zone_id, record_contain, max_items, batch, run):
    print(f"calculate TXT records to cleanup for 'zone:{zone_id}' and 'max records:{max_items}'")
    # https://botocore.amazonaws.com/v1/documentation/api/latest/reference/config.html
    cfg = Config(
       user_agent=f"ExternalDNS/boto3-{OPERATION_UUID}",

    )
    r53client = boto3.client('route53', config=cfg)
    dns_records_to_cleanup = []
    items = 0
    try:
      params = {
        'HostedZoneId': zone_id,
        'MaxItems': str(MAX_ITEMS),
      }
      dns_in_iteration = r53client.list_resource_record_sets(**params)
      elements = dns_in_iteration['ResourceRecordSets']
      for el in elements:
        if el['Type'] == 'TXT':
            record = Record(el)
            dns_records_to_cleanup.append(record)
            print("to cleanup >>", record)
            items += 1
            if items >= max_items:
                break

      while len(elements) > 0 and 'NextRecordName' in dns_in_iteration.keys() and items < max_items:
        dns_in_iteration = r53client.list_resource_record_sets(
            HostedZoneId= zone_id,
            StartRecordName= dns_in_iteration['NextRecordName'],
            MaxItems= str(MAX_ITEMS),
        )
        elements = dns_in_iteration['ResourceRecordSets']
        for el in elements:
            if el['Type'] == 'TXT':
                record = Record(el)
                dns_records_to_cleanup.append(record)
                print("to cleanup >>", record)
                items += 1
                if items >= max_items:
                    break

      deleteRecords(r53client, zone_id, dns_records_to_cleanup, batch, run)
    except Exception as e:
      print(f"An error occurred: {e}")
      os._exit(os.EX_OSERR)

def deleteRecords(client, zone_id, records, batch_size, run):
    total=len(records)
    print(f"will cleanup '{total}' records with batch '{batch_size}' at a time")
    count = 0
    for i in range(0, total, batch_size):
        if batch_size <= 0:
            break
        batch = records[i:min(i + batch_size, total)]
        print("NEW BATCH")
        count += batch_size
        if count >= total:
            count = total

        changes = []

        for el in batch:
            # https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/route53/client/change_resource_record_sets.html
            changes.append({
                            'Action': 'DELETE',
                            'ResourceRecordSet': el.record
                        })

        print(f"BATCH deletion(start). {len(changes)} records > {changes}")

        if run:
            client.change_resource_record_sets(
                HostedZoneId=zone_id,
                ChangeBatch={
                    "Comment": "external-dns legacy record cleanup. batch of ",
                    "Changes": changes,
                }
            )
            time.sleep(SLEEP)
        else:
            print("dry run execution")

        print(f"BATCH deletion(success). {count}/{total}(deleted/total)")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Cleanup legacy TXT records")
    parser.add_argument("--zone-id", type=str, required=True, help="Hosted Zone ID for which to run a cleanup.")
    parser.add_argument("--record-contain", type=str, required=True, help="Record to contain value. Example 'external-dns/owner=default'")
    parser.add_argument("--total-items", type=int, required=False, default=100, help="Number of items to delete. Default to 10")
    parser.add_argument("--batch-delete-count", type=int, required=False, default=2, help="Number of items to delete in single DELETE batch. Default to 2")
    parser.add_argument("--run", action="store_true", help="Execute the cleanup. The tool will do a dry-run if --run is not specified.")

    print("Run this script at your own RISKS!!!!")
    print(f"Session ID  '{OPERATION_UUID}'")

    args = parser.parse_args()
    print("arguments:",args)
    records(args.zone_id, args.record_contain, args.total_items, args.batch_delete_count, args.run)
