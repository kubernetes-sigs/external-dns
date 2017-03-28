# Try out ExternalDNS 

ExternalDNS' current release is `v0.1` which allows to keep a managed zone in Google's [CloudDNS](https://cloud.google.com/dns/docs/) service synchronized with Services of `type=LoadBalancer` in your cluster.

In this release ExternalDNS is limited to a single managed zone and takes full ownership of it. That means if you have any existing records in that zone they will be removed. We encourage you to try out ExternalDNS in its own zone first to see if that model works for you. However, ExternalDNS, by default, runs in dryRun mode and won't make any changes to your infrastructure. So, as long as you don't change that flag, you're safe.

Make sure you meet the following prerequisites:
* You have a local Go 1.7+ development environment.
* You have access to a Google project with the DNS API enabled.
* You have access to a Kubernetes cluster that supports exposing Services, e.g. GKE.
* You have a properly setup, **unused** and **empty** hosted zone in Google CloudDNS.

First, get ExternalDNS.

```console
$ go get -u github.com/kubernetes-incubator/external-dns
```

Run an application and expose it via a Kubernetes Service.

```console
$ kubectl run nginx --image=nginx --replicas=1 --port=80
$ kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
```

Annotate the Service with your desired external DNS name. Make sure to change `example.org` to your domain and that it includes the trailing dot.

```console
$ kubectl annotate service nginx "external-dns.alpha.kubernetes.io/hostname=nginx.example.org."
```

Run a single sync loop of ExternalDNS locally. Make sure to change the Google project to one you control and the zone identifier to an **unused** and **empty** hosted zone in that project's Google CloudDNS.

```console
$ external-dns --zone example-org --google-project example-project --once
```

This should output the DNS records it's going to modify to match the managed zone with the DNS records you desire.

Once you're satisfied with the result you can run ExternalDNS like you would run it in your cluster: as a control loop and not in dryRun mode.

```console
$ external-dns --zone example-org --google-project example-project --dry-run=false
```

Check that ExternalDNS created the desired DNS record for your service and that it points to its load balancer's IP. Then try to resolve it.

```console
$ dig +short nginx.example.org.
104.155.60.49
```

Now you can experiment and watch how ExternalDNS makes sure that your DNS records are configured as desired. Here are a couple of things you can try out:
* Change the desired hostname by modifying the Service's annotation.
* Recreate the Service and see that the DNS record will be updated to point to the new load balancer IP.
* Add another Service to create more DNS records.
* Remove Services to clean up your managed zone.
