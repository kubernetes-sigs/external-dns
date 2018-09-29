# CHANGELOG

#### master

- NEW: Added support for Collaborators methods (GH-48)
- NEW: Added region support to `ZoneRecord` (GH-47)
- NEW: Added support for Domains Pushes methods (GH-42)

- CHANGED: Renamed `DomainTransferRequest.AuthInfo` to `AuthCode` (GH-46)
- CHANGED: Updated registration, transfer, renewal response payload (dnsimple/dnsimple-developer#111, dnsimple/dnsimple-go#52).


#### Release 0.13.0

- NEW: Added support for Accounts methods (GH-29)
- NEW: Added support for Services methods (GH-30, GH-35)
- NEW: Added support for Certificates methods (GH-31)
- NEW: Added support for Vanity name servers methods (GH-34)
- NEW: Added support for delegation methods (GH-32)
- NEW: Added support for Templates methods (GH-36, GH-39)
- NEW: Added support for Template Records methods (GH-37)
- NEW: Added support for Zone files methods (GH-38)


#### Release 0.12.0

- CHANGED: Setting a custom user-agent no longer overrides the origina user-agent (GH-26)
- CHANGED: Renamed Contact#email_address to Contact#email (GH-27)


#### Release 0.11.0

- NEW: Added support for parsing ZoneRecord webhooks.
- NEW: Added support for listing options (GH-25).
- NEW: Added support for Template API (GH-21).


#### Release 0.10.0

Initial release.
