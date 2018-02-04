# Changelog

## Tue Jan 9 2018 - 0.6.0

- use VCR and fixtures for tests
- test ConvenientClient operations
- add support for zone create/delete operations

## Wed Aug 23 2017 - 0.5.3

- BUG-FIX: don't prepend dot for record with FQDN of Zone name

## Fri Aug 18 2017 - 0.5.2

- Handle errors reading response body in verbose mode (PR#20)

## Mon Jun 5 2017 - 0.5.1

- Update CHANGELOG

## Mon Jun 5 2017 - 0.5.0

- Add support for ALIAS, MX, NS, and SOA records, to the ConvenientClient
  (PR#17)

## Mon Jun 5 2017 - 0.4.1

- Handle rate limit errors

## Mon Jun 5 2017 - 0.4.0

- Fix nil-transport issue with the ConvenientClinent (PR#16)

## Fri Apr 21 2017 - 0.3.1

- Proxy support configurable with HTTP(S)_PROXY env variables
- BACKPORT: Handle rate limit errors

## Thu Sep 22 2016 - 0.3.0

- Verbose mode prints full url
- Handle Job redirections
- Support for unknown Content-Length
- Addition of ConvenientClient
- Support for Traffic Director (DSF) service

- BUGFIX: Don't override global log prefix

## Fri Nov 15 2013 - 0.2.0

- Fixed some struct field types
- Modified some of the tests
- Felt like it deserved a minor version bump

## Thu Nov 14 2013 - 0.1.9

- If verbosity is enabled, any unmarshaling errors will print the complete
  response body out, via logger

## Thu Nov 14 2013 - 0.1.8

## Wed Nov 13 2013 - 0.1.7

- Fixed a bug where empty request bodies would result in the API service
  responding with a 400 Bad Request
- Added some proper tests

## Wed Nov 13 2013 - 0.1.6

- Added a "verbose" mode to the client

## Tue Nov 12 2013 - 0.1.5

- Bug fixes
  - Logic bug in the *Client.Do() function, where it would not allow the
    POST /Session call if the client was logged out (POST /Session is used for
    logging in)

## Tue Nov 12 2013 - 0.1.4

- Includes 0.1.3
- Bug fixes
- Testing laid out, but there is not much there, as of right now

## Tue Nov 12 2013 - 0.1.2

- Bug fixes

## Tue Nov 12 2013 - 0.1.1

- Added structs for zone responses

## Tue Nov 12 2013 - 0.1.0

- Initial release
- The base client is complete; it will allow you to establish a session,
  terminate a session, and issue requests to the DynECT REST API endpoints
- TODO
  - Structs for marshaling and unmarshaling requests and responses still need
	to be done, as the current set of provided struct is all that is needed
	to be able to log in and create a session
  - More structs will be added on an "as I need them" basis
