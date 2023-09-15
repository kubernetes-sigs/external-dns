# Ingress source

The ingress source creates DNS entries based on `Ingress.networking.k8s.io` resources.

## Filtering the Ingresses considered

The `--ingress-class` flag filters Ingress resources by a set of ingress classes.
The flag may be specified multiple times in order to
allow multiple ingress classes.

This source supports the `--label-filter` flag, which filters Ingress resources
by a set of labels.

## Domain names

The domain names of the DNS entries created from an Ingress are sourced from the following places:

* Iterates over the Ingress's `spec.rules`, adding any non-empty `host`.

  This behavior is suppressed if the `--ignore-ingress-rules-spec` flag was specified
or the Ingress had an
`external-dns.alpha.kubernetes.io/ingress-hostname-source: annotation-only` annotation. 

* Iterates over the Ingress's `spec.tls`, adding each member of `hosts`.

  This behavior is suppressed if the `--ignore-ingress-tls-spec` flag was specified
or the Ingress had an
`external-dns.alpha.kubernetes.io/ingress-hostname-source: annotation-only` annotation, 

* Adds the hostnames from any `external-dns.alpha.kubernetes.io/hostname` annotation.

  This behavior is suppressed if the `--ignore-hostname-annotation` flag was specified
or the Ingress had an
`external-dns.alpha.kubernetes.io/ingress-hostname-source: defined-hosts-only` annotation.

* If no endpoints were produced for an Ingress by the previous steps
or the `--combine-fqdn-annotation` flag was specified, then adds hostnames
generated from any`--fqdn-template` flag.

## Targets

The targets of the DNS entries created from an Ingress are sourced from the following places:

* If the Ingress has an `external-dns.alpha.kubernetes.io/target` annotation, uses 
the values from that. 

* Otherwise, iterates over the Ingress's `status.loadBalancer.ingress`, 
adding each non-empty `ip` and `hostname`. 
