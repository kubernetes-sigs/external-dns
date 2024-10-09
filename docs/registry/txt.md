# The TXT registry

The TXT registry is the default registry.
It stores DNS record metadata in TXT records, using the same provider.


## TXT Format

### Metadata

The metadata format is used to construct the registry record. It can be selected by setting the `--txt-format` flag to `only-metadata`.

Under this condition, ExternalDNS creates one TXT record holding the ownership information using the following format:

```
{record_type}._metadata.{record_name}
```

Or, in case the original record has a wildcard:

```
*.{record_type}._metadata.{record_name}
```

ExternalDNS reads the legacy formats formats as well in order to facilitate the migration, but note that it does not delete old ownership records.

### Transition

With `--txt-format` set to `transition`, ExternalDNS creates tree distinct registry records:

1. `{record_type}._metadata.{record_name}`
2. `{record_type}.{record_name}` (previously refered as the `new` format).
3. `{record_name}` (previously refered as the `old` format).

Formats 2 and 3 are cosidered legacy and will be removed in the future.

#### Prefixes and Suffixes

In order to avoid having the registry TXT records collide with
TXT or CNAME records created from sources, you can configure a fixed prefix or suffix
to be added to the first component of the domain of all registry TXT records.

The prefix or suffix may not be changed after initial deployment,
lest the registry records be orphaned and the metadata be lost.

The prefix or suffix may contain the substring `%{record_type}`, which is replaced with
the record type of the DNS record for which it is storing metadata.

The prefix is specified using the `--txt-prefix` flag and the suffix is specified using
the `--txt-suffix` flag. The two flags are mutually exclusive.

## Wildcard Replacement

The `--txt-wildcard-replacement` flag specifies a string to use to replace the "*" in
registry TXT records for wildcard domains. Without using this, registry TXT records for
wildcard domains will have invalid domain syntax and be rejected by most providers.

## Encryption

Registry TXT records may contain information, such as the internal ingress name or namespace, considered sensitive, , which attackers could exploit to gather information about your infrastructure. 
By encrypting TXT records, you can protect this information from unauthorized access.

Encryption is enabled by using the `--txt-encrypt-enabled` flag. The 32-byte AES-256-GCM encryption
key must be specified in URL-safe base64 form, using the `--txt-encrypt-aes-key` flag.

Note that the key used for encryption should be a secure key and properly managed to ensure the security of your TXT records.

### Generating the TXT Encryption Key
Python
```python
python -c 'import os,base64; print(base64.urlsafe_b64encode(os.urandom(32)).decode())'
```

Bash
```shell
dd if=/dev/urandom bs=32 count=1 2>/dev/null | base64 | tr -d -- '\n' | tr -- '+/' '-_'; echo
```

OpenSSL
```shell
openssl rand -base64 32 | tr -- '+/' '-_'
```

PowerShell
```powershell
# Add System.Web assembly to session, just in case
Add-Type -AssemblyName System.Web
[Convert]::ToBase64String([System.Text.Encoding]::UTF8.GetBytes([System.Web.Security.Membership]::GeneratePassword(32,4))).Replace("+","-").Replace("/","_")
```

Terraform
```hcl
resource "random_password" "txt_key" {
  length           = 32
  override_special = "-_"
}
```

### Manually Encrypting/Decrypting TXT Records

In some cases you might need to edit registry TXT records. The following example Go code encrypts and decrypts such records.

```go
package main

import (
	"fmt"
	"sigs.k8s.io/external-dns/endpoint"
)

func main() {
	key := []byte("testtesttesttesttesttesttesttest")
	encrypted, _ := endpoint.EncryptText(
		"heritage=external-dns,external-dns/owner=example,external-dns/resource=ingress/default/example",
		key,
		nil,
	)
	decrypted, _, _ := endpoint.DecryptText(encrypted, key)
	fmt.Println(decrypted)
}
```

## Caching

The TXT registry can optionally cache DNS records read from the provider. This can mitigate
rate limits imposed by the provider.

Caching is enabled by specifying a cache duration with the `--txt-cache-interval` flag.
