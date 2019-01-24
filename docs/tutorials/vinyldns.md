# Setting up ExternalDNS for VinylDNS

## Create a sample deployment and service for external-dns to use

Run an application and expose it via a Kubernetes Service:

```console
$ kubectl run nginx --image=nginx --replicas=1 --port=80
$ kubectl expose deployment nginx --port=80 --target-port=80 --type=LoadBalancer
```

Annotate the Service with your desired external DNS name. Make sure to change `example.org` to your domain.

```console
$ kubectl annotate service nginx "external-dns.alpha.kubernetes.io/hostname=nginx.example.org."
```

After the service is up and running, it should get an EXTERNAL-IP. At first this may showing as `<pending>`

```console
$ kubectl get svc
NAME         CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
kubernetes   10.0.0.1     <none>        443/TCP        1h
nginx        10.0.0.115   <pending>     80:30543/TCP   10s
```

Once it's available

```console
% kubectl get svc
NAME         CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
kubernetes   10.0.0.1     <none>        443/TCP        1h
nginx        10.0.0.115   34.x.x.x      80:30543/TCP   2m
```

## Running a locally built version pointed to the above nginx service
Make sure your kubectl is configured correctly. Assuming you have the sources, build and run it like below.

The vinyl access details needs to exported to the environment before running.

```bash
make
# output skipped

export VINYLDNS_HOST=<fqdn of vinyl dns api>
export VINYLDNS_ACCESS_KEY=<access key>
export VINYLDNS_SECRET_KEY=<secret key>

./build/external-dns \
    --provider=vinyldns \
    --source=service \
    --domain-filter=elements.capsps.comcast.net. \
    --zone-id-filter=20e8bfd2-3a70-4e1b-8e11-c9c1948528d3 \
    --registry=txt \
    --txt-owner-id=grizz- \
    --txt-prefix=txt- \
    --namespace=default \
    --once \
    --dry-run \
    --log-level debug

INFO[0000] running in dry-run mode. No changes to DNS records will be made.
INFO[0000] Created Kubernetes client https://some-k8s-cluster.example.com
INFO[0001] Zone: [nginx.example.org.]
# output skipped
```

Having `--dry-run=true` and `--log-level=debug` is a great way to see _exactly_ what DynamicDNS is doing or is about to do.
