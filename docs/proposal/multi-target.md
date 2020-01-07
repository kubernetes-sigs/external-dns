# Multiple Targets per hostname 
*(November 2017)*

## Purpose 

One should be able to define multiple targets (IPs/Hostnames) in the **same** Kubernetes resource object and expect 
ExternalDNS create DNS record(s) with a specified hostname and all targets. So far the connection between k8s resources (ingress/services) and DNS records 
were not streamlined. This proposal aims to make the connection explicit, making k8s resources acquire or release certain DNS names. As long as the resource
ingress/service owns the record it can have multiple targets enable iff they are specified in the same resource.    

## Use cases  

See https://github.com/kubernetes-sigs/external-dns/issues/239

## Current behaviour 
*(as of the moment of writing)*

Central piece of enabling multi-target is having consistent and correct behaviour in `plan` component in regards to how endpoints generated 
from kubernetes resources are mapped to dns records. Current implementation of the `plan` has inconsistent behaviour in the following scenarios, all
of which must be resolved before multi-target support can be enabled in the provider implementations: 

1.  No records registered so far. Two **different** ingresses request same hostname but different targets, e.g. Ingress A: example.com -> 1.1.1.1 and Ingress B: example.com -> 2.2.2.2  
    * *Current Behaviour*: both are added to the "Create" (records to be created) list and passed to Provider
    * *Expected Behaviour*: only one (random/ or according to predefined strategy) should be chosen and passed to Provider  
    
    **NOTE**: while this seems to go against multi-target support, this is done so no other resource can "hijack" already created DNS record. Multi targets are supported only 
on per single resource basis  

2. Now let's say Ingress A was chosen and successfully created, but both ingress A and B are still there. So on next iteration ExternalDNS would see both again in the Desired list. 
    * *Current Behaviour*: DNS record target will change to that of Ingress B.
    * *Expected Behaviour*: Ingress A should stay unchanged. Ingress B record is not created

3. DNS record for Ingress A was created but its target has changed. Ingress B is still there
    * *Current Behaviour*: Undetermined behaviour based on which ingress will be parsed last. 
    * *Expected Behaviour*: DNS record should point to the new target specified in A. Ingress B should still be ignored.
    
    **NOTE**: both 2. and 3. can be resolved if External DNS is aware which resource has already acquired DNS record 
    
4. Ingress C has multiple targets: 1.1.1.1 and 2.2.2.2
    * *Current Behaviour*: Both targets are split into different endpoints and we end up in one of the cases above 
    * *Expected Behaviour*: Endpoint should contain list of targets and treated as one ingress object. 

## Requirements and assumptions

For this feature to work we have to make sure that: 

1. DNS records are now owned by certain ingress/service resources. For External DNS it would mean that TXT records now 
should store back-reference for the resource this record was created for, i.e. `"heritage=external-dns,external-dns/resource=ingress/default/my-ingress-object-name"` 
2. DNS records are updated only: 

    - If owning resource target list has changed 

    - If owning resource record is not found in the desired list (meaning it was deleted), therefore it will now be owned by another record. So its target list will be updated

    - Changes related to other record properties (e.g. TTL)  

4. All of the issues described in `Current Behaviour` sections are resolved 

Once Create/Update/Delete lists are calculated correctly (this is where conflicts based on requested DNS names are resolved) they are passed to `provider`, where
`provider` specific implementation will decide how to convert the structures into required formats. If DNS provider does not (or partially) support multi targets
then it is up to the provider to make sure that the change list of records passed to the DNS provider API is valid. **TODO**: explain best strategy.    

Additionally see https://github.com/kubernetes-sigs/external-dns/issues/258

## Implementation plan

Brief summary of open PRs and what they are trying to address:

### PRs 

1. https://github.com/kubernetes-sigs/external-dns/pull/243 - first attempt to add support for multiple targets. It is lagging far behind from master tip
    
    *what it does*: unfinished attempt to extend `Endpoint` struct, for it to allow multiple targets (essentially `target string -> targets []string`)
    
    *action*: evaluate if rebasing makes sense, or we can just close it. 
    
2. https://github.com/kubernetes-sigs/external-dns/pull/261 - attempt to rework `plan` to make it work correctly with multiple targets. 
    
    *what it does* : attempts to fix issues with `plan` described in `Current Behaviour` section above. Included tests reveal the current problem with `plan`
    
    *action*: rebase on master and make necessary changes to satisfy requirements listed in this document including back-reference to owning record
    
3. https://github.com/kubernetes-sigs/external-dns/pull/326 - attempt to add multiple target support. 
    
    *what it does*: for each pair `DNS Name` + `Record Type` it aggregates **all** targets from the cluster and passes them to Provider. It adds basic support
    for DO, Azura, Cloudflare, AWS, GCP, however those are not tested (?). (DNSSimple and Infoblox providers were not updated)
    
    *action*: the `plan` logic will probably needs to be reworked, however the rest concerning support in Providers and extending `Endpoint` struct can be reused. 
    Rebase on master and add missing pieces. Depends on `2`. 
    
 Related PRs: https://github.com/kubernetes-sigs/external-dns/pull/331/files,  https://github.com/kubernetes-sigs/external-dns/pull/347/files - aiming at AWS Route53 weighted records.
These PRs should be considered after common agreement about the way to address multi-target support is achieved. Related discussion:  https://github.com/kubernetes-sigs/external-dns/issues/196

### How to proceed from here

The following steps are needed: 
1. Make sure consensus regarding the approach is achieved via collaboration on the current document 
2. Notify all PR (see above) authors about the agreed approach
3. Implementation: 
    
    a. `Plan` is working as expected - either based on #261 above or from scratch. `Plan` should be working correctly regardless of multi-target support
    
    b. Extensive testing making sure new `plan` does not introduce any breaking changes
    
    c. Change Endpoint struct to support multiple targets - based on #326 - integrate it with new `plan` @sethpollack
    
    d. Make sure new endpoint format can still be used in providers which have only partial support for multi targets ~~**TODO**: how ?~~ . This is to be done by simply using first target in the targets list. 
    
    e. Add support for multi target which are already addressed in #326. It goes alongside c. and can be based on the same PR @sethpollack. New providers 
    added since then should maintain same functionality.  

5. Extensive testing on **all** providers before making new release
6. Update all related documentation and explain how multi targets are supported on per provider basis 
7. Think of introducing weighted records (see PRs section above) and making them configurable. 
  
## Open questions 

- Handling cases when ingress/service targets include both hostnames and IPs - postpone this until use cases occurs
- "Weighted records scope": https://github.com/kubernetes-sigs/external-dns/issues/196 - this should be considered once multi-target support is implemented
