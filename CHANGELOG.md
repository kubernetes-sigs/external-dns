Features:

  - Support creation of CNAME records when endpoint target is a hostname.
  - Allow omitting the trailing dot in Service annotations.
  - Expose basic Go metrics via Prometheus.

Bug fixes:

  - AWS Route 53: Do not submit request when there are no changes.

Documentation:

  - Add documentation on how to setup ExternalDNS for Services on AWS.

## v0.1.0 - 2017-03-30 (KubeCon)

Features:

  - Manage DNS records for Services with `Type=LoadBalancer` on Google CloudDNS.
