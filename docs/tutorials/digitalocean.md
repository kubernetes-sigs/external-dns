# DigitalOcean DNS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using DigitalOcean DNS.

Make sure to use **>=0.4.2** version of ExternalDNS for this tutorial.

## Creating a DigitalOcean DNS zone

If you want to learn about how to use DigitalOcean's DNS service read the following tutorial series:

[An Introduction to Managing DNS](https://www.digitalocean.com/community/tutorial_series/an-introduction-to-managing-dns), and specifically [How To Set Up a Host Name with DigitalOcean DNS](https://www.digitalocean.com/community/tutorials/how-to-set-up-a-host-name-with-digitalocean)

Create a new DNS zone where you want to create your records in. Let's use `example.com` as an example here.

## Creating DigitalOcean Credentials

Generate a new personal token by going to [the API settings](https://cloud.digitalocean.com/settings/api/tokens) or follow [How To Use the DigitalOcean API v2](https://www.digitalocean.com/community/tutorials/how-to-use-the-digitalocean-api-v2) if you need more information.
Give the token a name and choose read and write access.
The token needs to be passed to ExternalDNS so make a note of it for later use.

The environment variable `DO_TOKEN` will be needed to run ExternalDNS with DigitalOcean.

## Deploy ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.

Begin by creating a Kubernetes secret to securely store your DigitalOcean API key. This key will enable ExternalDNS to authenticate with DigitalOcean:

```shell
kubectl create secret generic DO_TOKEN --from-literal=DO_TOKEN=YOUR_DIGITALOCEAN_API_KEY
```

Ensure to replace YOUR_DIGITALOCEAN_API_KEY with your actual DigitalOcean API key.

Then apply one of the following manifests file to deploy ExternalDNS.

## Using Helm

Create a values.yaml file to configure ExternalDNS to use DigitalOcean as the DNS provider. This file should include the necessary environment variables:

```shell
provider:
  name: digitalocean
env:
  - name: DO_TOKEN
    valueFrom:
      secretKeyRef:
        name: DO_TOKEN
        key: DO_TOKEN
```

### Manifest (for clusters without RBAC enabled)

```yaml
[[% include 'digitalocean/extdns-without-rbac.yaml' %]]
```

### Manifest (for clusters with RBAC enabled)

```yaml
[[% include 'digitalocean/extdns-with-rbac.yaml' %]]
```

## Deploying an Nginx Service

Create a service file called 'nginx.yaml' with the following contents:

```yaml
[[% include 'digitalocean/deploy-nginx.yaml' %]]
```

Note the annotation on the service; use the same hostname as the DigitalOcean DNS zone created above.

ExternalDNS uses this annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

Create the deployment and service:

```console
kubectl create -f nginx.yaml
```

Depending where you run your service it can take a little while for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and synchronize the DigitalOcean DNS records.

## Verifying DigitalOcean DNS records

Check your [DigitalOcean UI](https://cloud.digitalocean.com/networking/domains) to view the records for your DigitalOcean DNS zone.

Click on the zone for the one created above if a different domain was used.

This should show the external IP address of the service as the A record for your domain.

## Cleanup

Now that we have verified that ExternalDNS will automatically manage DigitalOcean DNS records, we can delete the tutorial's example:

```sh
kubectl delete service -f nginx.yaml
kubectl delete service -f externaldns.yaml
```

## Advanced Usage

### API Page Size

If you have a large number of domains and/or records within a domain, you may encounter API
rate limiting because of the number of API calls that external-dns must make to the DigitalOcean API to retrieve
the current DNS configuration during every reconciliation loop. If this is the case, use the
`--digitalocean-api-page-size` option to increase the size of the pages used when querying the DigitalOcean API.
(Note: external-dns uses a default of 50.)
