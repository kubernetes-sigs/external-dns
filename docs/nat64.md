# Configure NAT64 DNS Records

Some NAT64 configurations are entirely handled outside the Kubernetes cluster, therefore Kubernetes does not know anything about the associated IPv4 addresses. ExternalDNS should also be able to create A records for those cases.
Therefore, we can configure `nat64-networks`, which **must** be a /96 network. You can also specify multiple `nat64-networks` for more complex setups.
This creates an additional A record with a NAT64-translated IPv4 address for each AAAA record pointing to an IPv6 address within the given `nat64-networks`.

This can be configured with the following flag passed to the operator binary. You can also pass multiple `nat64-networks` by using a comma as seperator.

```sh
--nat64-networks="2001:db8:96::/96"
```

## Setup Example

We use an external NAT64 resolver and SIIT (Stateless IP/ICMP Translation). Therefore, our nodes only have IPv6 IP adresses but can reach IPv4 addresses *and* can be reached via IPv4.
Outgoing connections are a classic NAT64 setup, where all IPv6 addresses gets translated to a small pool of IPv4 addresses.
Incoming connnections are mapped on a different IPv4 pool, e.g. `198.51.100.0/24`, which can get translated one-to-one to IPv6 addresses.
We dedicate a `/96` network for this, for example `2001:db8:96::/96`, so `198.51.100.0/24` can translated to `2001:db8:96::c633:6400/120`. Note: `/120` IPv6 network has exactly as many IP addresses as `/24` IPv4 network.

Therefore, the `/96` network can be configured as `nat64-networks`. This means, that `2001:0DB8:96::198.51.100.10` or `2001:db8:96::c633:640a` can be translated to `198.51.100.10`.
Any source can point a record to an IPv6 address within the given `nat64-networks`, for example `2001:db8:96::c633:640a`.
This creates by default an AAAA record and - if `nat64-networks` is configured - also an A record with `198.51.100.10` as target.
