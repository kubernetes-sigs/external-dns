# The TXT registry

The TXT registry is the default registry.
It stores DNS record metadata in TXT records, using the same provider.

> Note:
>
> - If you plan to manage apex domains with external-dns whilst using a txt registry, you should ensure when using `--txt-prefix` that you specify the record type substitution and that it ends in a period (**.**).
>   The record should be created under the same domain as the apex record being managed, i.e. `--txt-prefix=someprefix-%{record_type}.`
> - `--txt-prefix` and `--txt-suffix` contribute to the 63-byte maximum record length. To avoid errors, use them only if absolutely required and keep them as short as possible.

## Record Format Options

### For version `v0.18+`

The TXT registry supports single format for storing DNS record metadata:

- Creates a TXT record with record type information (e.g., 'a-' prefix for A records)

The TXT registry would try to guarantee a consistency in between providers and sources, if provider supports the behaviour.

If configured `--txt-prefix="%{record_type}-abc-."` for apex domain `ex.com` the expected result is

|         Name         |  TYPE   |
| :------------------: | :-----: |
| `cname-abc-.ex.com.` |  `TXT`  |
|      `ex.com.`       | `CNAME` |

For the domain `www.ex.com` the expected result is

|           Name           |  TYPE   |
| :----------------------: | :-----: |
| `cname-abc-.www.ex.com.` |  `TXT`  |
|      `www.ex.com.`       | `CNAME` |

If configured `--txt-suffix="-.%{record_type}"` for apex domain `ex.com`, the expected result would be `ex-.a.com`, which fails to create a TXT record because it does not exist within the managed zone.

For the domain `www.ex.com` the expected result is

|         Name         |  TYPE   |
| :------------------: | :-----: |
| `www-.cname.ex.com.` |  `TXT`  |
|    `www.ex.com.`     | `CNAME` |

### AWS A ALIAS records

[AWS ALIAS records](../tutorials/aws.md#alias) are stored in Route 53 as
[`A`/`AAAA` records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html),
so their ownership TXT uses the matching `a-`/`aaaa-` prefix. For `A` ALIAS records this replaces the
legacy `cname-` prefix: the old `cname-` record is still recognized (ownership is preserved), the `a-`
record is created on the next reconciliation, and the obsolete `cname-` is left in place with a warning.
See [#2903](https://github.com/kubernetes-sigs/external-dns/issues/2903).

### Manually Cleanup Legacy TXT Records

> While deleting registry TXT records won't cause downtime, a well-thought-out migration and cleanup plan is crucial.

Occasionally, it may be necessary to remove outdated TXT records from your registry.

An example script for AWS can be found in [scripts/aws-cleanup-legacy-txt-records.py](../../scripts/aws-cleanup-legacy-txt-records.py) with instructions on how to run it.
The script performs targeted deletion of TXT records that include `ResourceRecords` matching the `heritage=external-dns,external-dns/owner=default` or similar pattern.
In the event of unintended deletion of all TXT records managed by `external-dns`, `external-dns` will initiate a full DNS record regeneration, along with`TXT` and `non-TXT` records. Just be aware, this operation's duration is directly proportional to the DNS estate size."

To remove the obsolete `cname-` records left by the AWS A ALIAS migration, run the script with
`--alias-cname-cleanup`, which deletes a `cname-` record only when its `a-` counterpart exists.

### For version `v0.16.0 & v0.16.1`

The TXT registry supports two formats for storing DNS record metadata:

- Legacy format: Creates a TXT record without record type information
- New format: Creates a TXT record with record type information (e.g., 'a-' prefix for A records)

By default, the TXT registry creates records in both formats for backwards compatibility. You can configure it to use only the new format by using the `--txt-new-format-only` flag. This reduces the number of TXT records created, which can be helpful when working with provider-specific record limits.

Note: The following record types always use only the new format regardless of this setting:

- AAAA records
- Encrypted TXT records (when using `--txt-encrypt-enabled`)

Example:

```sh
# Default behavior - creates both formats
external-dns --provider=aws --source=ingress --managed-record-types=A --managed-record-types=TXT

# Only create new format records (alongside other required flags)
external-dns --provider=aws --source=ingress --managed-record-types=A --managed-record-types=TXT --txt-new-format-only
```

The `--txt-new-format-only` flag should be used in addition to your existing external-dns configuration flags. It does not implicitly configure TXT record handling - you still need to specify `--managed-record-types=TXT` if you want external-dns to manage TXT records.

### Migration to New Format Only

> Note: `external-dns` will not automatically remove legacy format records when switching to new-format-only mode. You'll need to clean up the old records manually if desired.

When transitioning from dual-format to new-format-only records:

- Ensure all your `external-dns` instances support the new format
- Enable the `--txt-new-format-only` flag on your external-dns instances
  Manually clean up any existing legacy format TXT records from your DNS provider

## Prefixes and Suffixes

In order to avoid having the registry TXT records collide with
TXT or CNAME records created from sources, you can configure a fixed prefix or suffix
to be added to the first component of the domain of all registry TXT records.

The prefix or suffix may not be changed after initial deployment,
lest the registry records be orphaned and the metadata be lost.

The prefix or suffix may contain the substring `%{record_type}`, which is replaced with
the record type of the DNS record for which it is storing metadata.

The prefix is specified using the `--txt-prefix` flag and the suffix is specified using
the `--txt-suffix` flag. The two flags are mutually exclusive.

### Zone-aware ownership names

By default, suffixes are added to the first label, preserving existing naming behavior.
With `--registry=txt --txt-zone-aware`, the suffix is instead added immediately before
the authoritative zone boundary. For example, with zone `example.com` and
`--txt-suffix=-txtsuffix`, an A record for `name.192.sub.example.com` uses
`a-name.192.sub-txtsuffix.example.com` rather than `a-name-txtsuffix.192.sub.example.com`.

This opt-in currently requires RFC2136 with explicit, non-root `--rfc2136-zone` values.
Provider caching is supported. Domain filters do not define zone boundaries: a filter
such as `sub.example.com` does not replace the configured zone `example.com`.
Other providers without authoritative zone support are rejected when this flag is enabled.
The DynamoDB registry retains its existing naming behavior.

Existing ownership TXT names are retained on updates and deletes; enabling this flag
does not automatically migrate them. If legacy and zone-aware parsing interpret a TXT
name as different records, or ownership records conflict, reconciliation stops with an
error rather than choosing an owner. Ambiguous new ownership names are rejected before
any changes reach the provider. Resolve such conflicts explicitly before retrying.

In zone-aware mode, recognized ownership TXT records must have exactly one target;
multiple targets are rejected even if they contain identical metadata. Unmatched TXT
names and ordinary TXT values without ExternalDNS ownership metadata are not subject
to this restriction. Typed, untyped, and legacy alias markers applicable to the same
record must agree on all ownership metadata. Distinct typed markers for different
record types remain independent. If one marker applies to multiple existing records
(for example, an untyped marker shared by A and MX records), reconciliation stops
before updates or deletes, including when only one of those types is managed. Resolve
shared ownership explicitly before enabling this mode; no automatic migration or
partial deletion of shared ownership values is performed.

Do not disable the flag after creating zone-aware names without planning an ownership
migration: the default parser may not recognize those records.

## Wildcard Replacement

The `--txt-wildcard-replacement` flag specifies a string to use to replace the "\*" in
registry TXT records for wildcard domains. Without using this, registry TXT records for
wildcard domains will have invalid domain syntax and be rejected by most providers.

## Encryption

Registry TXT records may contain information, such as the internal ingress name or namespace, considered sensitive, , which attackers could exploit to gather information about your infrastructure.
By encrypting TXT records, you can protect this information from unauthorized access.

Encryption is enabled by setting the `--txt-encrypt-enabled`. The 32-byte AES-256-GCM encryption
key must be specified in URL-safe base64 form (recommended) or be a plain text, using the `--txt-encrypt-aes-key=<key>` flag.

Note that the key used for encryption should be a secure key and properly managed to ensure the security of your TXT records.

### Generating the TXT Encryption Key

Python

```python
python -c 'import os,base64; print(base64.standard_b64encode(os.urandom(32)).decode())'
```

Bash

```shell
dd if=/dev/urandom bs=32 count=1 2>/dev/null | base64; echo
```

OpenSSL

```shell
openssl rand -base64 32
```

PowerShell

```powershell
# Add System.Web assembly to session, just in case
Add-Type -AssemblyName System.Web
[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes([System.Web.Security.Membership]::GeneratePassword(32,4)))
```

Terraform

```hcl
resource "random_password" "txt_key" {
  length           = 32
}
```

### Manually Encrypting/Decrypting TXT Records

In some cases you might need to edit registry TXT records. The following example Go code encrypts and decrypts such records.

```go
package main

import (
	b64 "encoding/base64"
	"fmt"

	"sigs.k8s.io/external-dns/endpoint"
)

func main() {
	keys := []string{
		"ZPitL0NGVQBZbTD6DwXJzD8RiStSazzYXQsdUowLURY=", // safe base64 url encoded 44 bytes and 32 when decoded
		"01234567890123456789012345678901",             // plain txt 32 bytes
		"passphrasewhichneedstobe32bytes!",             // plain txt 32 bytes
	}

	for _, k := range keys {
		key := []byte(k)
		if len(key) != 32 {
			// if key is not a plain txt let's decode
			var err error
			if key, err = b64.StdEncoding.DecodeString(string(key)); err != nil || len(key) != 32 {
				fmt.Errorf("the AES Encryption key must have a length of 32 byte")
			}
		}
		encrypted, _ := endpoint.EncryptText(
			"heritage=external-dns,external-dns/owner=example,external-dns/resource=ingress/default/example",
			key,
			nil,
		)
		decrypted, _, err := endpoint.DecryptText(encrypted, key)
		if err != nil {
			fmt.Println("Error decrypting:", err, "for key:", k)
		}
		fmt.Println(decrypted)
	}
}
```

## Caching

The TXT registry can optionally cache DNS records read from the provider. This can mitigate
rate limits imposed by the provider.

Caching is enabled by specifying a cache duration with the `--txt-cache-interval` flag.

## OwnerID migration

> Automating DNS migrations with third-party tools can be risky. DNS is often business-critical, and without deep understanding of the environment, 3rd party automation tools can do more harm than good.

The owner ID of the TXT records managed by external-dns instance can be updated.

When `--migrate-from-txt-owner` is set, it will enable the migration checks
in the run loop using `--txt-owner-id=new-owner-id` and the value you defined for this flag.

If you want to test the outputs of a migration beforehand, you can use the `--dry-run` flag
along with `--migrate-from-txt-owner`.

Example, if you had a standard deployment like so:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.k8s.io/external-dns/external-dns:v0.22.0
        imagePullPolicy: Always
        args:
        - "--txt-prefix=%{record_type}-"
        - "--txt-cache-interval=2m"
        - "--log-level=debug"
        - "--log-format=text"
        - "--txt-owner-id=old-owner"
        - "--policy=sync"
        - "--provider=some-provider"
        - "--registry=txt"
        - "--interval=1m"
        - "--source=ingress"
```

You can update your deployment to migrate like so :

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: external-dns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        imagePullPolicy: Always
        image: registry.k8s.io/external-dns/external-dns:v0.22.0
        args:
        - "--txt-prefix=%{record_type}-"
        - "--txt-cache-interval=2m"
        - "--log-level=debug"
        - "--log-format=text"
        - "--txt-owner-id=new-owner"
        - "--migrate-from-txt-owner=old-owner"
        - "--policy=sync"
        - "--provider=some-provider"
        - "--registry=txt"
        - "--interval=1m"
        - "--source=ingress"
```

If you didn't set the owner ID, the value set by external-dns is `default`. You can set the
`--migrate-from-txt-owner` flag to `default` to migrate the associated records.

### OwnerID migration: multi-cluster considerations

> Warning: The `--migrate-from-txt-owner` flag combined with `policy=sync` can be unsafe in shared hosted zones when multiple clusters previously used the same TXT owner value (for example `default`).

In a shared hosted zone, if one cluster runs ExternalDNS with `policy=sync` and `--migrate-from-txt-owner=default`, it may attempt to delete DNS records that belong to other clusters which still use `owner=default`.
To avoid this, do not share the same TXT owner value across clusters in any zone where `policy=sync` or migration flags will be used.

#### Per-cluster owner IDs

For multi-cluster setups sharing a hosted zone:

- Assign a **unique** `--txt-owner-id` to each cluster (for example `cluster1`, `cluster2`) and document this convention clearly in your platform configuration.
- Avoid using a common owner such as `default` across clusters in a shared zone if any cluster will run with `policy=sync` or use `--migrate-from-txt-owner`.

#### Example migration sequence for shared zones

When migrating from a shared owner (such as `default`) in a shared hosted zone:

1. While still using `policy=upsert-only` (or equivalent), roll out cluster-specific `--txt-owner-id` values and ensure *new* records are created with the cluster’s own owner ID.
2. Avoid `--migrate-from-txt-owner=<old-owner>` unless you can guarantee that only a single cluster has records with `<old-owner>` in that hosted zone, or perform the migration in an isolated zone where only that cluster writes records.

### When to avoid owner migration

The following pattern is **not recommended** and may cause record deletion for other clusters:

- Multiple clusters share a Route53 hosted zone and all existing records use `owner=default`.
- Only one cluster is upgraded to use `policy=sync`, `--txt-owner-id=<cluster-name>`, and `--migrate-from-txt-owner=default`, while other clusters still use `owner=default`.

In this situation, the upgraded cluster can treat other clusters’ records as orphans and schedule them for deletion during synchronization. Prefer per-cluster zones, manual TXT record adjustment, or fully coordinated migration of all clusters if the migration flag must be used.
