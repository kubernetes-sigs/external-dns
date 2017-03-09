# Storage

## Purpose

Provides a persistent storage with additional information to track the records created by the external-dns. Initial discussion - https://github.com/kubernetes-incubator/external-dns/issues/44

**Why we need it?**

DNS provider (AWS Route53, Google DNS, etc.) stores dns records which are created via various means. Integration of External-DNS should be safe and should not delete or overwrite the records which it is not responsible for. Moreover, it should certainly be possible for multiple kubernetes clusters to share the same hosted zone within the dns provider, additionally multiple external-dns instances inside the same cluster should be able to co-exist without messing with the same set of records. 

This proposal introduces multiple possible implementation with the details depending on the setup. 

## Requirements and assumptions

1. Pre-existing records should not be modified by external-dns
2. External-dns instance only creates/modifies/deletes records which are created by this instance
3. It should be possible to transfer the ownership to another external-dns instance
4. Any integrated dns-provider should provide at least a single way to implement the storage
5. Lifetime of the records should not be limited to lifetime of external-dns
6. External-dns should have its identifier for marking the managed records - **`storage-id`**

## Types of storage

The following presents two ways to implement the storage, and we are planning to implement both for compatibility purposes.

### TXT records

This implementation idea is borrowed from [Mate](https://github.com/zalando-incubator/mate/)

Each record created by external-dns is accompanied by the TXT record, which internally stores the external-dns identifier. For example, if external dns with `storage-id="external-dns-1"` record to be created with dns name `foo.zone.org`, external-dns will create a TXT record with the same dns name `foo.zone.org` and injected value of `"external-dns-1"`. The transfer of ownership can be done by modifying value in the TXT record.  If no TXT record exists for the record, then external-dns will simply ignore it. 


#### Goods
1. Easy to guarantee cross-cluster ownership safety
2. Data lifetime is not limited to cluster or external-dns lifetime
3. Supported by major DNS providers

#### Bads
1. TXT record cannot co-exist with CNAME records (this can be solved by creating a TXT record with another domain name, e.g. `foo.org->foo.txt.org`)
2. Introduces complexity to the logic
3. Difficult to do the transfer of ownership
4. Too easy to mess up with manual modifications

### ConfigMap

Store the state in the configmap. ConfigMap is created and managed by each external-dns individually, i.e. external-dns with **`storage-id=external-dns-1`** will create and operate on `extern-dns-1-storage` ConfigMap. ConfigMap will store **all** the records present in the DNS provider as serialized JSON. For example: 

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
2. ConfigMap is a first class citized in kubernetes world
3. Does not create dependency/restriction on DNS provider
4. Cannot be easily messed with by other parties

#### Bads
1. ConfigMap might be out of sync with dns provider state
2. LifeTime is obviously limited to the cluster lifetime
3. Not supported in older kubernetes clusters


## Component integration

Components: 
* Source - all endpoints ( collection of ingress, service[type=LoadBalancer] etc.)
* [Plan](https://github.com/kubernetes-incubator/external-dns/issues/13) - object responsible for the create of change lists in external-dns 
* DNSProvider - interface to access the DNS provider API 

A single loop iteration of external-dns operation: 

1. Get all endpoints ( collection ingress, service[type=LoadBalancer] etc.) into collection of `endpoints` 
2. Get storage `Records()` 
3. Pass `Records` (including ownership information) and list of endpoints to `Plan` to do the calculation
4. Make a call to DNS provider with `Plan` provided change list
5. If call succeeded pass the change list pass to storage `Assign()` to mark the records that are created 

Storage gets updated all the time via `Poll`.  

#### Notes:

1. DNS Provider should use batch operations
2. DNS Provider should be called with CREATE operation (not UPSERT!) when the record does not yet exist! 
3. Storage does not need to be in complete sync with DNS provider due to #2. Hence resolving the potential caveats of ConfigMap implementation 


## Implementation

Basic implementation of the storage interface: 

1. Storage has the dnsprovider object to retrieve the list of records, but it never makes the call to modify the records (think layer to help out with endpoint filtering)
2. Record() - returns whatever is stored in the storage
3. Assign(endpoints) - called when the records are registered with dns provider - hence storage need to mark its ownership. Therefore DNSProvider serves as a safe-guard from race conditions
4. WaitForSync() - called in the beginning to populate the storage, in case of configmap would be the configmap creation and fetching the dns provider records
5. Poll() - resync loop to stay-up-to-date with dns provider state


### Example:

We will provide `InMemoryStorage` non-persistent storage which should help us with testing.

```
type InMemoryStorage struct {
	registry dnsprovider.DNSProvider
	zone     string
	owner    string 
	cache    map[string]*SharedEndpoint
	sync.Mutex
}

func (im *InMemoryStorage) Poll(stopChan <-chan struct{}) {
	for {
		select {
		case <-time.After(resyncPeriod):
			err := im.refreshCache()
			if err != nil {
				log.Errorf("failed to refresh cache: %v", err)
			}
		case <-stopChan:
			log.Infoln("terminating storage polling")
			return
		}
	}
}

func (im *InMemoryStorage) refreshCache() error {
	im.Lock()
	defer im.Unlock()

	records, err := im.registry.Records(im.zone)
  ...

	im.cache = map[string]*SharedEndpoint{} //drop the current cache
	for _, newCacheRecord := range newCache {
		im.cache[newCacheRecord.DNSName] = newCacheRecord
	}

	return nil
}
```

Initial PR: https://github.com/kubernetes-incubator/external-dns/pull/57
