### TXT Registry migration to a new format ###

In order to support more record types and be able to track ownership without TXT record name clash, a new TXT record is introduced.
It contains record type it manages, e.g.:
* A record foo.example.com will be tracked with classic foo.example.com TXT record
* At the same time a new TXT record will be created a-foo.example.com

Prefix and suffix are extended with %{record_type} template where the user can control how prefixed/suffixed records should look like.

In order to maintain compatibility, both records will be maintained for some time, in order to have downgrade possibility.  
The controller will try to create the "new format" TXT records if they are not present to ease the migration from the versions < 0.12.0.

Later on, the old format will be dropped and only the new format will be kept (<record_type>-<endpoint_name>).

Cleanup will be done by controller itself.
