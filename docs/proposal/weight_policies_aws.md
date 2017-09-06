# Weight Policies For AWS

## Overview of Weight on AWS Route53 Records

The purpose of adding weight to AWS Route53 records is to distribute requests to more than a single DNS name and control
the request load on each of the DNS records for the same name.

## How It Works

Amazon Route 53 calculates the sum of the weights for the resource record sets that have the same combination of DNS
name and type. Amazon Route 53 then responds to queries based on the ratio of a resource's weight to the total.

Source: http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values-weighted.html#rrsets-values-weighted-weight

For example to have weighted records for the domain name `nginx.example.com` where one record distrubutes twice as many
requests to one backend than to another, one will have two DNS records both with the name of `nginx.example.com`.

The first record will look like:

```
Name: nginx.example.com
Type: A
Weight: 2
Target: 8.8.8.8
```

The second record will look like:
```
Name: nginx.example.com
Type: A
Weight: 1
Target: 9.9.9.9
```

Note that weight cannot be applied to different targets within the same DNS name.

## Purpose

The reasoning for this is outlined well in https://github.com/kubernetes-incubator/external-dns/issues/196, but for a
brief overview: The option to apply weight to DNS records is needed when either running Federated services and wanting
to distribute the load between clusters, or when running multiple services in the same cluster and wanting to distribute
the load between those services (e.g. canary deployments).

## Requirements

* No change should be made to records that do not support or do not include the option to add weight (opt-in only)
* We are capable of adding weight amount via annotations on the service, which are then reflected in the record that is
created (route 53)
* We are capable of running multiple records with the same owner (TXT record)
* We are able to delete records and the TXT record will only be removed when there are no longer any remaining records
under its ownership

## Types of Ownship and Scope of Feature

For now this will focus on multiple records per the same owner, referred to as "cluster" ownership. In the future
the plan will be to support multiple records for multiple owners, referred to as "global" ownership.

## Code Changes

### Extend the Endpoint Type to Include Policies

The endpoint.Endpoint type will include a Policy like so:

```go
type Endpoint struct {
  // The hostname of the DNS record
  DNSName string
  // The target the DNS record points to
  Target string
  // RecordType type of record, e.g. CNAME, A, TXT etc
  RecordType string
  // Policy stores policies for the endpoint
  Policy EndpointPolicy
  // Labels stores labels defined for the Endpoint
  Labels map[string]string
}
```

The EndpointPolicy may include a number of policies depending on the source. In the case the source has the appropriate
annotations, it may include a WeightPolicy like so:

```go
type EndpointPolicy struct {
  // Weight is the weight policy for the endpoint
  Weight WeightPolicy
}
```

### Manage Weight Policies as an Interface

We'll assume there will be other types other than AWS ELBs that will need weights as well. So WeightPolicy is an
interface like:

```go
type Weight interface {
  Amount() int64
  ID() string
}
```

The route53WeightPolicy type implments the WeightPolicy interface. It is specific to Route53 and can be created via the
following:

```go
type route53WeightPolicy struct {
  // scope is can be cluster or none
  // TODO: Implement global weight scope in addition to cluster-only
  scope string
  // weight is the weight of the record
  weight int64
  // setidentifier for the weight policy
  setIdentifier string
}

func NewRoute53WeightPolicy(weight int64, setIdentifier string) Weight {
  return route53WeightPolicy{
    scope: parseScope(setIdentifier),
    weight: weight,
    setIdentifier: setIdentifier,
  }
}
```

### How to Set Weight Policies from the Provider

Policies can be added to an endpoint through methods on the Endpoint type. For example, inside the AWS provider:

```go
newEndpoint := endpoint.NewEndpoint(wildcardUnescape(aws.StringValue(r.Name)), aws.StringValue(rr.Value), aws.StringValue(r.Type))
newEndpoint.AttachWeightPolicy(endpoint.policies.NewRoute53WeightPolicy(aws.Int64Value(r.Weight), aws.aws.StringValue(r.SetIdentifier)))
endpoints = append(endpoints, newEndpoint)
```

This assumes `r.SetIdentifier` is something like `cluster/service-a` or `global/us-east-1`. These are set in the next
section.

### How to Set Weight Polcies from the Source

Policies can be added to an endpoint from the Source in the same fashion as the provider, but using values specified via
the following annotations:

```go
  // The annotation used for defining the desired weight for the record
  weightAnnotationKey = "external-dns.alpha.kubernetes.io/weight"
  // The annotation used for defining the desired weight scope for the record
  weightScopeAnnotationKey = "external-dns.alpha.kubernetes.io/weight-scope"
  // The annotation used for setting the value of suffix of the ID for the weight set
  weightSuffixAnnotationKey = "external-dns.alpha.kubernetes.io/weight-suffix"
```

A raw example would be:

```go
if svc.Annotations[weightScopeAnnotationKey] != "" && svc.Annotations[weightSuffixAnnotationKey] != "" && svc.Annotations[weightAnnotationKey] != "" {
  weight, _ := strconv.ParseInt(svc.Annotations[weightAnnotationKey], 10, 64)
  newEndpoint.AttachWeightPolicy(endpoint.policies.NewRoute53WeightPolicy(weight, fmt.Sprintf("%s/%s", svc.Annotations[weightScopeAnnotationKey], svc.Annotations[weightSuffixAnnotationKey])))
}
```

### Usage Inside the Provider

The provider creates the records using the policy if it exists. For example:

```go
if endpoint.Policy.HasWeightPolicy() {
  change.ResourceRecordSet.Weight = aws.Int64(endpoint.Policy.Weight.Amount())
  change.ResourceRecordSet.SetIdentifier = aws.String(endpoint.Policy.Weight.ID())
}
```

### How to Handle Multiple Records Per Owner

#### Plan

In order for this to work, we must allow a single TXT record (owner) to manage multiple records. This is done by adding
logic to the plan to do the following:

* Create the record if it doesn't exist (existing behavior)
* If the record already exists, check to see if the record supports weighting
* If the record supports weighting, check to see if the ID of the weight policy matches the expected value
* If the record does not match the expected value, flag the record as having an owner (via `Endpoint.Labels`) and create
a new record with the desired ID

To delete records that have been removed:

* If the record exists but should not, check to see if the record supports weighting
* Delete the record if it does not support weighting, continue if it does
* Check to see if the ID of the weight policy matches the expected value of the record that should be deleted
* If the record still exists but should not, keep a reference count for all records that use the same TXT owner
* Add the value to the reference counter
* After all records have been processed, loop over the reference counter and:
* Delete the record if the reference counter is equal to one, or:
* Mark the record as having an owner which should not be deleted (via `Endpoint.Labels`) and:
* Delete the record only

#### TXT Records

The txt `ApplyChanges` method will be adjusted to not inject the TXT record into the plan's changes if either of the
following is true:

* The record is set to be created but the endpoint has been marked by the plan as having an owner
* The record is set to be delete but the endpoint has been marked by the plan as having an owner which is managing valid
 records

