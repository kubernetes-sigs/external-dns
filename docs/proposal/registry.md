# Registry
#### [Old name: storage]

Initial discussion - https://github.com/kubernetes-sigs/external-dns/issues/44

## Purpose

One should not be afraid to use external-dns, because it can delete or overwrite the records preexisting in the DNS provider.

**Why we need it?**

DNS provider (AWS Route53, Google DNS, etc.) stores dns records which are created via various means. Integration of External-DNS should be safe and should not delete or overwrite the records which it is not responsible for. Moreover, it should certainly be possible for multiple kubernetes clusters to share the same hosted zone within the dns provider, additionally multiple external-dns instances inside the same cluster should be able to co-exist without messing with the same set of records.

Registry provides a mechanism to ensure the safe management of DNS records

This proposal introduces multiple possible implementation with the details depending on the setup.

## Requirements and assumptions

1. Pre-existing records should not be modified by external-dns
2. External-dns instance only creates/modifies/deletes records which are created by this instance
3. It should be possible to transfer the ownership to another external-dns instance
4. ~~Any integrated DNS provider should provide at least a single way to implement the registry~~ Noop registry can be used to disable ownership
5. Lifetime of the records should not be limited to lifetime of external-dns
6. External-dns should have its identifier for marking the managed records - **`owner-id`**

## Types of registry

The following presents two ways to implement the registry, and we are planning to implement both for compatibility purposes.

### TXT records

This implementation idea is borrowed from [Mate](https://github.com/linki/mate/)

Each record created by external-dns is accompanied by the TXT record, which internally stores the external-dns identifier. For example, if external dns with `owner-id="external-dns-1"` record to be created with dns name `foo.zone.org`, external-dns will create a TXT record with the same dns name `foo.zone.org` and injected value of `"external-dns-1"`. The transfer of ownership can be done by modifying the value of the TXT record.  If no TXT record exists for the record or the value does not match its own `owner-id`, then external-dns will simply ignore it.


#### Goods
1. Easy to guarantee cross-cluster ownership safety
2. Data lifetime is not limited to cluster or external-dns lifetime
3. Supported by major DNS providers
4. TXT record are created alongside other records in a batch request. Hence **eliminating possibility of having inconsistent ownership information and dns provider state**

#### Bads
1. TXT record cannot co-exist with CNAME records (this can be solved by creating a TXT record with another domain name, e.g. `foo.org->foo.txt.org`)
2. Introduces complexity to the logic
3. Difficult to do the transfer of ownership
4. Too easy to mess up with manual modifications

### ConfigMap

**This implementation cannot be considered 100% error free**, hence use with caution [see **Possible failure scenario** below] 

Store the state in the configmap. ConfigMap is created and managed by each external-dns individually, i.e. external-dns with **`owner-id=external-dns-1`** will create and operate on `extern-dns-1-registry` ConfigMap. ConfigMap will store **all** the records present in the DNS provider as serialized JSON. For example:

```
kind: ConfigMap
apiVersion: v1
metadata:
  creationTimestamp: 2016-03-09T19:14:38Z
  name: external-dns-1-storage
  namespace: same-as-external-dns-1
data:
  records: "[{
	\"dnsname\": \"foo.org\",
	\"owner\": \"external-dns-1\",
	\"target\": \"loadbalancer1.com\",
	\"type\": \"A\"
}, {
	\"dnsname\": \"bar.org\",
	\"owner\": \"external-dns-2\",
	\"target\": \"loadbalancer2.com\",
	\"type\": \"A\"
}, {
	\"dnsname\": \"unmanaged.org\",
	\"owner\": \"\",
	\"target\": \"loadbalancer2.com\",
	\"type\": \"CNAME\"
}]"

```

ConfigMap will be periodically resynced with the dns provider by fetching the dns records and comparing it with the data currently stored and hence rebuilding the ownership information.

#### Goods
1. Not difficult to implement and easy to do the ownership transfer
2. ConfigMap is a first class citizen in kubernetes world
3. Does not create dependency/restriction on DNS provider
4. Cannot be easily messed with by other parties

#### Bads
1. ConfigMap might be out of sync with dns provider state
2. LifeTime is obviously limited to the cluster lifetime
3. Not supported in older kubernetes clusters
4. Bloated ConfigMap - cannot be paginated and will grow very big on DNS provider with thousands of records

#### Failure scenario

It is possible that the configmap will go out of sync with the dns provider state. In the implementation flow external-dns will first modify records on dns provider side to subsequently update configmap. And if ExternalDNS will crash in-between two operation created records will be left unmanaged and not viable for update/deletion by External DNS.

## Component integration

Components:
* Source - all endpoints ( collection of ingress, service[type=LoadBalancer] etc.)
* [Plan](https://github.com/kubernetes-sigs/external-dns/issues/13) - object responsible for the create of change lists in external-dns
* Provider - interface to access the DNS provider API

Registry will serve as wrapper around `Provider` providing additional information regarding endpoint ownership. Ownership will further taken into account by `Plan` to filter out records to include only records managed by current ExternalDNS instance (having same `owner-id` value)

A single loop iteration of external-dns operation:

1. Get all endpoints ( collection ingress, service[type=LoadBalancer] etc.) into collection of `endpoints`
2. Get registry `Records()` (makes the call to DNSProvider and also build ownership information)
3. Pass `Records` (including ownership information) and list of endpoints to `Plan` to do the calculation
4. Call registry `ApplyChanges()` method (which subsequently calls DNS Provider Apply method to update records)
5. If ConfigMap implementation of Registry is used, then ConfigMap needs to be updated separately

~~In case of configmap, Registry gets updated all the time via `Poll`. `Plan` does not call DNS provider directly. Good value of the `Poll` is to have simple rate limiting mechanism on DNS provider.~~

#### Notes:

1. DNS Provider should use batch operations
2. DNS Provider should be called with CREATE operation (not UPSERT!) when the record does not yet exist!

### Resources:

1. Registry implementation draft: 

[Flow](https://lh3.googleusercontent.com/BNUZZQ8XivYkXyYVPDgPCoZpwYv0pOyoyfBKbOnYJGsqueeB-EUXfzBZLk7xP-E_GDo7YHiTlA4XgPEs6ao_Ex0TY2SN66-yg5iRmn5Tc2EXVqs_yS9CtumhE1T4krZc4Z8_1gHOirDxCegU-Fk0K3fvg-J3UpzdKmGDG-JZwdzRyP4WyORWUQilJO9jErh-HP8AtM8p2ZjiqN9B3-VXdYuHbsiR6EHNFw43aOQAk52muDf2AgjqX2YUSbN9eO0Akt39ien3euT2HsZJlPvm5s8v2a_ZqTSW0DVcGaRhLQbZXcogSEP-ebbuGunuVbz45Ws8X6zJhZpASNQ-jknhGZEhZkSAQdwvihZpTsDdUuJx9RFDXNwA0lEaE_xediW119uJGywSNc6w8hnJZ6Xo49YQStuGbJKRAieQMvEhZXofiqCKyOUXSlsO7j9iE-rzis0JRSHWB8acA3AlcXqBj9D70AHfRHC_HfBLw9lcusy4dInmK2OCzGqXV11PoqibiZPqh-oNED31pToZQk4NB1xbOuUC_Tjf8UR_xAyhJ3yKzS09K898uCf-87Ra4iqRDCz3N35b=w2560-h1260)
