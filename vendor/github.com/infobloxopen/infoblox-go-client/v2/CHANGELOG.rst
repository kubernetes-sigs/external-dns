===================================
Infoblox Go Client Release Notes
===================================

.. contents:: Topics

v2.1.1
======

This is just a bugfix release.

v2.1.0
======

Release Summary
---------------

- Enhancements in Host Record functionality
- Code refinements

Minor Changes
-------------

- `SearchHostRecordByAltId` function to search Host Record By alternative ID from terraform.
- The code for every record has been seperated and added under a new file.


v2.0.0
======

Release Summary
---------------

Create, Update, Delete and Get operation on below records are being added or enhanced.

- Network View with comment and EAs
- IPv4 and IPv6 network containers with comment and EAs
- IPv4 and IPv6 network with comment and EAs
- Host Record with comment, EAs, EnableDns, EnableDhcp, UseTtl, Ttl, Alias attributes
- Fixed Address record with comment and EAs
- A record with comment, EAs, UseTtl, Ttl
- AAAA record with comment, EAs, UseTtl, Ttl
- PTR record with comment, EAs, UseTtl, Ttl
- Added IPv6 support for PTR record
- CNAME record with comment, EAs, UseTtl, Ttl
- Adds a compile-time check to the interface to make sure it stays in sync with the actual implementation.
- Added apt UTs and updated respective UTs

Minor Changes
-------------

- Added default value of network view in AllocateIP, CreateHostRecord and CreatePTRRecord Function

Bugfixes
-------------

- IPv6 Support `#86 <https://github.com/infobloxopen/infoblox-go-client/issues/86>`_
- Possibility to UPDATE a CNAME entry `#110 <https://github.com/infobloxopen/infoblox-go-client/issues/110>`_
- Feature Request: Ability to add comments `#116 <https://github.com/infobloxopen/infoblox-go-client/issues/116>`_
- Feature: Add ability to set extensible attributes when creating a network `#119 <https://github.com/infobloxopen/infoblox-go-client/issues/119>`_
- Feature request: add host aliases get/create/update `#126 <https://github.com/infobloxopen/infoblox-go-client/issues/126>`_
