# The TXT registry

The TXT registry is the default registry.
It stores DNS record metadata in TXT records, using the same provider.

## Record Format Options

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
external-dns --provider=aws --source=ingress --managed-record-types=A,TXT

# Only create new format records (alongside other required flags)
external-dns --provider=aws --source=ingress --managed-record-types=A,TXT --txt-new-format-only
```

The `--txt-new-format-only` flag should be used in addition to your existing external-dns configuration flags. It does not implicitly configure TXT record handling - you still need to specify `--managed-record-types=TXT` if you want external-dns to manage TXT records.

### Migration to New Format Only

When transitioning from dual-format to new-format-only records:

- Ensure all your `external-dns` instances support the new format
- Enable the `--txt-new-format-only` flag on your external-dns instances
Manually clean up any existing legacy format TXT records from your DNS provider

Note: `external-dns` will not automatically remove legacy format records when switching to new-format-only mode. You'll need to clean up the old records manually if desired.

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

## Wildcard Replacement

The `--txt-wildcard-replacement` flag specifies a string to use to replace the "*" in
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
