# Webhook provider

The "Webhook" provider allows to integrate ExternalDNS with DNS providers via an HTTP interface.
The Webhook provider implements the Provider interface. Instead of implementing code specific to a provider, it implements an HTTP client that sends request to an HTTP API.
The idea behind it is that providers can be implemented in a separate program: these programs expose an HTTP API that the Webhook provider can interact with. The ideal setup for providers is to run as sidecars in the same pod of the ExternalDNS container and listen on localhost only. This is not strictly a requirement, but we would not recommend other setups.

## Architectural diagram

![Webhook provider](../img/webhook-provider.png)

## API guarantees

Providers implementing the HTTP API have to keep in sync with changes to the Go types `plan.Changes` and `endpoint.Endpoint`. We do not expect to make significant changes to those types given the maturity of the project, but we can't exclude that changes will need to happen. We commit to publishing changes to those in the release notes, to ensure that providers implementing the API can keep providers up to date quickly.

## Implementation requirements

The following table represents the methods to implement mapped to their HTTP method and route.

| Provider method | HTTP Method | Route |
| --- | --- | --- |
| Records | GET | /records |
| AdjustEndpoints | POST | /adjustendpoints |
| ApplyChanges | POST | /records |

The server needs to respond to those requests by reading the `Accept` header and responding with a corresponding `Content-Type` header specifying the supported media type format and version.

**NOTE**: only `5xx` responses will be retried and only `20x` will be considered as successful. All status codes different from those will be considered a failure on ExternalDNS's side.

## Provider registry

To simplify the discovery of providers, we will accept pull requests that will add links to providers in the [README](../../README.md) file. This list will serve the only purpose of simplifying finding providers and will not constitute an official endorsment of any of the externally implemented providers unless otherwise specified.


## Run the AWS provider with the webhook provider.

To test the Webhook provider and provide a reference implementation, we added the functionality to run the AWS provider as a webhook. To run the AWS provider as a webhook, you need the following flags:

```yaml
- --provider=webhook
- --run-aws-provider-as-webhook
```

What will happen behind the scenes is that the AWS provider will be be started as an HTTP server exposed only on localhost and the webhook provider will be configured to talk to it. This is the same setup that we recommend for other providers and a good way to test the Webhook provider.
