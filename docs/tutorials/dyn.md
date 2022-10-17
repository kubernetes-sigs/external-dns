# Setting up ExternalDNS for Dyn

## Creating a Dyn Configuration Secret

For ExternalDNS to access the Dyn API, create a Kubernetes secret.

To create the secret:

```
$ kubectl create secret generic external-dns \
      --from-literal=EXTERNAL_DNS_DYN_CUSTOMER_NAME=${DYN_CUSTOMER_NAME} \
      --from-literal=EXTERNAL_DNS_DYN_USERNAME=${DYN_USERNAME} \
      --from-literal=EXTERNAL_DNS_DYN_PASSWORD=${DYN_PASSWORD}
```

The credentials are the same ones created during account registration. As best practise, you are advised to
create an API-only user that is entitled to only the zones intended to be changed by ExternalDNS

## Deploy ExternalDNS
The rest of this tutorial assumes you own `example.com` domain and your DNS provider is Dyn. Change `example.com`
with a domain/zone that you really own.

In case of the dyn provider, the flag `--zone-id-filter` is mandatory as it specifies which zones to scan for records. Without it


Create a deployment file called `externaldns.yaml` with the following contents:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  strategy:
    type: Recreate
  selector:
    matchLabels:
      app: external-dns
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=ingress
        - --txt-prefix=_d
        - --namespace=example
        - --zone-id-filter=example.com
        - --domain-filter=example.com
        - --provider=dyn
        env:
        - name: EXTERNAL_DNS_DYN_CUSTOMER_NAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_DYN_CUSTOMER_NAME
        - name: EXTERNAL_DNS_DYN_USERNAME
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_DYN_USERNAME
        - name: EXTERNAL_DNS_DYN_PASSWORD
          valueFrom:
            secretKeyRef:
              name: external-dns
              key: EXTERNAL_DNS_DYN_PASSWORD
EOF
```

As we'll be creating an Ingress resource, you need `--txt-prefix=_d` as a CNAME cannot coexist with a TXT record. You can change the prefix to
any valid start of a FQDN.

Create the deployment for ExternalDNS:

```
$ kubectl create -f externaldns.yaml
```

## Running a locally build version
If you just want to test ExternalDNS in dry-run mode locally without doing the above deployment you can also do it.
Make sure your kubectl is configured correctly . Assuming you have the sources, build and run it like so:

```bash
make 
# output skipped

./build/external-dns \
    --provider=dyn \
    --dyn-customer-name=${DYN_CUSTOMER_NAME} \
    --dyn-username=${DYN_USERNAME} \
    --dyn-password=${DYN_PASSWORD} \
    --domain-filter=example.com \
    --zone-id-filter=example.com \
    --namespace=example \
    --log-level=debug \
    --txt-prefix=_ \
    --dry-run=true
INFO[0000] running in dry-run mode. No changes to DNS records will be made. 
INFO[0000] Connected to cluster at https://some-k8s-cluster.example.com 
INFO[0001] Zones: [example.com]
# output skipped
```

Having `--dry-run=true` and `--log-level=debug` is a great way to see _exactly_ what DynamicDNS is doing or is about to do.

## Deploying an Ingress Resource

Create a file called 'test-ingress.yaml' with the following contents:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:  
  name: test-ingress
  namespace: example
spec:
  rules:
  - host: test-ingress.example.com
    http:
      paths:
      - backend:
          service:
            name: my-awesome-service
            port:
              number: 8080
        pathType: Prefix
```

As the DNS name `test-ingress.example.com` matches the filter, external-dns will create two records:
a CNAME for test-ingress.example.com and TXT for _dtest-ingress.example.com. 

Create the Ingress:

```
$ kubectl create -f test-ingress.yaml
```

By default external-dns scans for changes every minute so give it some time to catch up with the 
## Verifying Dyn DNS records

Login to the console at https://portal.dynect.net/login/ and verify records are created

## Clean up

Login to the console at https://portal.dynect.net/login/ and delete the records created. Alternatively, just delete the sample
Ingress resources and external-dns will delete the records.
