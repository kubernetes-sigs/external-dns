# Running ExternalDNS with limited privileges

You can run ExternalDNS with reduced privileges since `v0.5.6` using the following `SecurityContext`.

```yaml
[[% include 'security-context/extdns-limited-privilege.yaml' %]]
```
