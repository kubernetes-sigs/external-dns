#!/usr/bin/env python

# Copyright 2018 The Kubernetes Authors.
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
#
# pip install kubernetes boto3

import boto3
from kubernetes import client, config

# replace with your hosted zone id
hosted_zone_id = ''
# replace with your txt-owner-id you are using
# inside of your external-dns controller
txt_owner_id = ''

# change to false if you have external-dns not looking at services
external_dns_manages_services = True

# change to false if you have external-dns not looking at ingresses
external_dns_manages_ingresses = True

config.load_kube_config()

# grab all the domains that k8s thinks it is going to 
# manage (services with domainName specified and 
# ingress hosts)
k8s_domains = []

if external_dns_manages_services:
    v1 = client.CoreV1Api()
    svcs = v1.list_service_for_all_namespaces()
    for i in svcs.items:
        annotations = i.metadata.annotations
        if annotations is not None and 'domainName' in annotations:
            k8s_domains.extend(annotations['domainName'].split(','))

if external_dns_manages_ingresses:
    ev1 = client.ExtensionsV1beta1Api()
    ings = ev1.list_ingress_for_all_namespaces()
    for i in ings.items:
        for r in i.spec.rules:
            if r.host not in k8s_domains:
                k8s_domains.append(r.host)


r53client = boto3.client('route53')

# grab the existing route53 domains and identify gaps where a domain may be 
# missing a txt record pair
existing_r53_txt_domains=[]
existing_r53_domains=[]
has_next = True
next_record_name, next_record_type='',''

while has_next:
    if next_record_name is not '' and next_record_type is not '':
        resource_records = r53client.list_resource_record_sets(HostedZoneId=hosted_zone_id, 
                                                               StartRecordName=next_record_name, 
                                                               StartRecordType=next_record_type)
    else:
        resource_records = r53client.list_resource_record_sets(HostedZoneId=hosted_zone_id)

    for r in resource_records['ResourceRecordSets']:
        if r['Type'] == 'TXT':
            existing_r53_txt_domains.append(r['Name'][:-1])
        elif r['Type'] == 'A':
            existing_r53_domains.append(r['Name'][:-1])
    has_next = resource_records['IsTruncated']
    if has_next:
        next_record_name, next_record_type = resource_records['NextRecordName'], resource_records['NextRecordType']

# grab only the domains in route53 that kubernetes is managing
r53_k8s_domains = [r for r in k8s_domains if r in existing_r53_domains]
# from those find the ones that do not have matching txt entries
missing_k8s_txt = [r for r in r53_k8s_domains if r not in existing_r53_txt_domains]

# make the change batch for the route53 call, modify this as needed
change_batch=[]
for r in missing_k8s_txt:
    change_batch.append(
        {
            'Action': 'CREATE',
            'ResourceRecordSet': {
                'Name': r,
                'Type': 'TXT',
                'TTL': 300,
                'ResourceRecords': [
                    {
                        'Value': '\heritage=external-dns,owner="' + txt_owner_id + '\"'
                    },
                ]
            }
        })

print('This will create the following resources')
print(change_batch)
response = input("Good to go? ")  

if response.lower() in ['y', 'yes', 'yup', 'ok', 'sure', 'why not', 'why not?']:
    print('Updating route53')
    change_response = r53client.change_resource_record_sets(
                            HostedZoneId=hosted_zone_id,
                            ChangeBatch={
                              'Changes': change_batch
                        })
    print('Submitted change request to route53. Details below.')
    print(change_response)
else:
    print('No changes were made')
