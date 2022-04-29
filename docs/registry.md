### TXT Registry migration to a new format ###

In order to support more record types and be able to track ownership without TXT record name clash, a new TXT record is introduced.
It contains record type it manages, e.g.:
* A record foo.example.com will be tracked with classic foo.example.com TXT record
* At the same time a new TXT record will be created a-foo.example.com

Prefix and suffix are extended with %{record_type} template where the user can control how prefixed/suffixed records should look like.

In order to maintain compatibility, both records will be maintained for some time, in order to have downgrade possibility.

If the implementation is to complete the downgrade, it should include the following modifications:
  * A new flag `-skip-ipv6-targets-in-a-record` has been added to allow filtering of IPv6 targets in type A endpoints so that a double-stacked rollback version of a source resource does not cause other effects. This feature should be merged in a previous release. This is because some providers (e.g. aliyundns) currently update A-record IPv6 targets as if they were A-records, and then are bound to encounter failures, and if they do they will delete the A-record at the same time.

Later on, the old format will be dropped and only the new format will be kept (<record_type>-<endpoint_name>).

Cleanup will be done by controller itself.
