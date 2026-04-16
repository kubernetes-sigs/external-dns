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

# Warning: The script deletes all records that match certain values. It could delete both legacy and new records if there is no way to differentiate them.

# This Python script is designed to help migrate DNS management to `external-dns` by cleaning up legacy TXT records in AWS Route 53.
# It identifies and deletes TXT records that match a specified pattern, ensuring that `external-dns` can take over managing these resources.
# The script performs the following steps:
#
# 1. **Setup and Configuration**:
#    - Imports necessary libraries (`boto3`, `argparse`, etc.).
#    - Defines constants and utility functions.
#    - Parses command-line arguments for configuration.
#
# 2. **Record Class**:
#    - Represents a DNS record with methods to check if it should be deleted.
#
# 3. **Main Functionality**:
#    - Connects to AWS Route 53 using `boto3`.
#    - Support single zone cleanup at a time.
#    - Lists and filters TXT records based on the specified pattern.
#    - Deletes the filtered records in batches, with an option for a dry run or actual deletion.
#
# 4. **Execution**:
#    - The script is executed with command-line arguments specifying the hosted zone ID, record pattern, total items to delete, batch size, and whether to perform a dry run or actual deletion.
#    - Check 'To Run script' section for more details

# WARNING: run this script at your own RISK. This will delete all the TXT records that do contain certain string.
# To Run script
# 1. Python, pip and pipenv installed https://pipenv.pypa.io/en/latest/
# 2. AWS Access https://docs.aws.amazon.com/signin/latest/userguide/command-line-sign-in.html
# 3. pipenv shell
# 4. pip install boto3
# 5. python scripts/aws-cleanup-legacy-txt-records.py --help
# 6. DRY RUN python scripts/aws-cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF --record-match text
# 6.1 Before execution consider to stop `external-dns`
# 7. Execute Deletion. First few times with reduced number of items
# - python scripts/aws-cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF --total-items 3 --batch-delete-count 1 --record-match 'external-dns'
# - python scripts/aws-cleanup-legacy-txt-records.py --zone-id ASDFQEQREWRQADF --total-items 10000 --batch-delete-count 50 --run --record-match "external-dns/owner=default"

# python scripts/aws-cleanup-legacy-txt-records.py --help
# python scripts/aws-cleanup-legacy-txt-records.py --zone-id Z06155043AVN8RVC88TYY --total-items 300 --batch-delete-count 20 --record-match "external-dns/owner=default" --run

import boto3
from botocore.config import Config as AwsConfig
import json, argparse, os, uuid, time

MAX_ITEMS=300 # max is 300 https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/route53/client/list_resource_record_sets.html
SLEEP=1 # in seconds, required to make sure Route53 API is not throttled
SESSION_ID=uuid.uuid4()

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

    def is_for_deletion(self, contains):

        if contains in self.resource_record:
            return True
        return False

    def __str__(self):
        return f'record: name: {self.name}, type: {self.type}, records: {self.resource_record}'

class Config:

    def __init__(self, zone_id, contain, total_items, batch, run):
        self.zone_id = zone_id
        self.record_contain = contain
        self.total_items = total_items
        self.batch_size = batch
        self.run = run
        self.contain = contain

def records(config: Config) -> None:
    print(f"calculate TXT records to cleanup for 'zone:{config.zone_id}' and 'max records:{config.total_items}'")
    # https://botocore.amazonaws.com/v1/documentation/api/latest/reference/config.html
    cfg = AwsConfig(
        user_agent=f"ExternalDNS/boto3-{SESSION_ID}",
    )
    r53client = boto3.client('route53', config=cfg)
    dns_records_to_cleanup = []
    items = 0
    try:
        params = {
            'HostedZoneId': config.zone_id,
            'MaxItems': str(MAX_ITEMS),
        }
        dns_in_iteration = r53client.list_resource_record_sets(**params)
        elements = dns_in_iteration['ResourceRecordSets']
        for el in elements:
            if el['Type'] == 'TXT':
                record = Record(el)
                if record.is_for_deletion(config.contain):
                    dns_records_to_cleanup.append(record)
                    print("to cleanup >>", record)
                    items += 1
                    if items >= config.total_items:
                        break

        while len(elements) > 0 and 'NextRecordName' in dns_in_iteration.keys() and items < config.total_items:
            dns_in_iteration = r53client.list_resource_record_sets(
                HostedZoneId= config.zone_id,
                StartRecordName= dns_in_iteration['NextRecordName'],
                MaxItems= str(MAX_ITEMS),
            )
            elements = dns_in_iteration['ResourceRecordSets']
            for el in elements:
                if el['Type'] == 'TXT':
                    record = Record(el)
                    if record.is_for_deletion(config.contain):
                        dns_records_to_cleanup.append(record)
                        print("to cleanup >>", record)
                        items += 1
                        if items >= config.total_items:
                            break

        if len(dns_records_to_cleanup) > 0:
            delete_records(r53client, config, dns_records_to_cleanup)
        else:
            print("No 'TXT' records found to cleanup....")
    except Exception as e:
        print(f"An error occurred: {e}")
        os._exit(os.EX_OSERR)

def delete_records(client: boto3.client, config: Config, records: list[Record]) -> None:
    total=len(records)
    print(f"will cleanup '{total}' records with batch '{config.batch_size}' at a time")
    count = 0

    if config.run:
        print("deletion of records!!")
    else:
        print("dry run execution")

    for i in range(0, total, config.batch_size):
        if config.batch_size <= 0:
            break
        batch = records[i:min(i + config.batch_size, total)]
        count += config.batch_size
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

        if config.run:
            client.change_resource_record_sets(
                HostedZoneId=config.zone_id,
                ChangeBatch={
                    "Comment": "external-dns legacy record cleanup. batch of ",
                    "Changes": changes,
                }
            )
            time.sleep(SLEEP)

        print(f"BATCH deletion(success). {count}/{total}(deleted/total)")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Cleanup legacy TXT records")
    parser.add_argument("--zone-id", type=str, required=True, help="Hosted Zone ID for which to run a cleanup.")
    parser.add_argument("--record-match", type=str, required=True, help="Record to match specific value. Example 'external-dns/owner=default'")
    parser.add_argument("--total-items", type=int, required=False, default=10, help="Number of items to delete. Default to 10")
    parser.add_argument("--batch-delete-count", type=int, required=False, default=2, help="Number of items to delete in single DELETE batch. Default to 2")
    parser.add_argument("--run", action="store_true", help="Execute the cleanup. The tool will do a dry-run if --run is not specified.")

    answer = input("Run this script at your own RISKS!!! Please enter 'yes' or 'no': ")
    if answer != 'yes':
        os._exit(0)

    print(f"Session ID  '{SESSION_ID}'")

    args = parser.parse_args()
    print("arguments:",args)
    cfg = Config(
        zone_id=args.zone_id,
        contain=args.record_match,
        total_items=args.total_items,
        batch=args.batch_delete_count,
        run=args.run,
    )
    records(cfg)
