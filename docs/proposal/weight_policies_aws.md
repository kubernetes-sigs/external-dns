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
// Endpoint is a high-level way of a connection between a service and an IP
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

The EndpointPolicy may include a number of policies depending on the source or provider. In the case the source has the
appropriate annotations, it may include a policy like so:

```go
// AWSRoute53Policy stores the policy attributes for the Route53 record
type AWSRoute53Policy struct {
  // Weight is the weight of the RecordSet
  Weight int64
  // SetIdentifier for the RecordSet
  SetIdentifier string
}
```


The AWSRoute53Policy is specific to Route53 and can be created via the following:

```go
// NewAWSRoute53Policy does basic validation according to AWS requirements and returns the Route53 policy
func NewAWSRoute53Policy(weight int64, setIdentifier string) (*AWSRoute53Policy, error) {
  if weight < 0 || weight > 255 {
    return &AWSRoute53Policy{}, fmt.Errorf("Weight must be between 0-255. Actual: %d", weight)
  }

  if len(setIdentifier) < 1 || len(setIdentifier) > 128 {
    return &AWSRoute53Policy{}, fmt.Errorf("Set Identifier must be between 1-128 characters. Actual: %s",
      setIdentifier)
  }

  return &AWSRoute53Policy{
    Weight:        weight,
    SetIdentifier: setIdentifier,
  }, nil
}
```

### How to Set Policies from the Provider

Policies can be added to an endpoint through methods on the Endpoint type. For example, inside the AWS provider:

```go
if route53WeightPolicy, err := endpoint.NewAWSRoute53Policy(aws.Int64Value(r.Weight), aws.StringValue(r.SetIdentifier)); err == nil {
  newEndpoint.Policy.AttachAWSRoute53Policy(route53WeightPolicy)
}
```

This assumes `r.SetIdentifier` is something like `service-a` or `us-east-1`. These are set in the next section.

### How to Set Policies from the Source

Policies can be added to an endpoint from the Source in the same fashion as the provider, but using values specified via
the following annotations:

```go
// The annotation used for defining the desired weight scope for the record
weightScopeAnnotationKey = "external-dns.alpha.kubernetes.io/weight-scope"
// The annotation used for defining the desired weight for the record
awsRoute53WeightAnnotationKey = "external-dns.alpha.kubernetes.io/aws-route53-weight"
// The value of the suffix for the weighted policy id
awsRoute53SetIdentifierAnnotationKey = "external-dns.alpha.kubernetes.io/aws-route53-set-identifier"
```

An example would be:

```go
func attachAWSRoute53Policy(svc *v1.Service, ep *endpoint.Endpoint) (*endpoint.Endpoint, error) {
  weight, err := strconv.ParseInt(svc.Annotations[awsRoute53WeightAnnotationKey], 10, 64)
  if err != nil {
    return nil, err
  }

  awsRoute53Policy, err := endpoint.NewAWSRoute53Policy(weight,
    svc.Annotations[awsRoute53SetIdentifierAnnotationKey])

  if err != nil {
    return nil, err
  }

  ep.Policy.AttachAWSRoute53Policy(awsRoute53Policy)
  return ep, nil
}
```

### Usage Inside the Provider

The provider creates the records using the policy if it exists. For example:

```go
if endpoint.Policy.HasAWSRoute53Policy() {
  change.ResourceRecordSet.Weight = aws.Int64(endpoint.Policy.AWSRoute53.Weight)
  change.ResourceRecordSet.SetIdentifier = aws.String(endpoint.Policy.AWSRoute53.SetIdentifier)
}
```

### How to Handle Multiple Records Per Owner

#### Plan

In order for this to work, we must allow a single TXT record (owner) to manage multiple records. This is done by adding
logic to the plan to do the following:

* Create the record if it doesn't exist (existing behavior)
* If the record already exists, check to see if the record has a policy
* If the record has a policy, check to see if the policy matches the expected policy
* If the record does not match the expected policy, flag the record as having an owner (via `Endpoint.Labels`) and
create a new record with the desired ID

To delete records that have been removed:

* If the record exists but should not, check to see if the record has a policy
* Delete the record if it does not have a policy, continue if it does
* Check to see if the ID of the policy matches the expected value of the record that should be deleted
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
* The record is set to be deleted but the endpoint has been marked by the plan as having an owner which is managing
valid records
