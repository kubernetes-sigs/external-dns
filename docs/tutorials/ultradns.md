# Setting up ExternalDNS for Services on UltraDNS

This tutorial describes how to setup ExternalDNS for usage within a Kubernetes cluster using UltraDNS.

For this tutorial, please make sure that you are using a version **> 0.7.2** of ExternalDNS.

## Managing DNS with UltraDNS

If you would like to read-up on the UltraDNS service, you can find additional details here: [Introduction to UltraDNS](https://docs.ultradns.neustar)

Before proceeding, please create a new DNS Zone that you will create your records in for this tutorial process. For the examples in this tutorial, we will be using `example.com` as our Zone.

## Setting Up UltraDNS Credentials

The following environment variables will be needed to run ExternalDNS with UltraDNS.

`ULTRADNS_USERNAME`,`ULTRADNS_PASSWORD`, &`ULTRADNS_BASEURL`
`ULTRADNS_ACCOUNTNAME`(optional variable).

## Deploying ExternalDNS

Connect your `kubectl` client to the cluster you want to test ExternalDNS with.
Then, apply one of the following manifests file to deploy ExternalDNS.

- Note: We are assuming the zone is already present within UltraDNS.
- Note: While creating CNAMES as target endpoints, the `--txt-prefix` option is mandatory.
### Manifest (for clusters without RBAC enabled)

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
        - --source=service 
        - --source=ingress # ingress is also possible
        - --domain-filter=example.com # (Recommended) We recommend to use this filter as it minimize the time to propagate changes, as there are less number of zones to look into..
        - --provider=ultradns
        - --txt-prefix=txt-
        env:
        - name: ULTRADNS_USERNAME
          value: ""
        - name: ULTRADNS_PASSWORD  # The password is required to be BASE64 encrypted.
          value: ""
        - name: ULTRADNS_BASEURL
          value: "https://api.ultradns.com/"
        - name: ULTRADNS_ACCOUNTNAME
          value: ""
```

### Manifest (for clusters with RBAC enabled)

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups: [""]
  resources: ["services","endpoints","pods"]
  verbs: ["get","watch","list"]
- apiGroups: ["extensions"]
  resources: ["ingresses"]
  verbs: ["get","watch","list"]
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["list","watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns-viewer
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
- kind: ServiceAccount
  name: external-dns
  namespace: default
---
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
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.7.6
        args:
        - --source=service 
        - --source=ingress
        - --domain-filter=example.com #(Recommended) We recommend to use this filter as it minimize the time to propagate changes, as there are less number of zones to look into..
        - --provider=ultradns
        - --txt-prefix=txt-
        env:
        - name: ULTRADNS_USERNAME
          value: ""
        - name: ULTRADNS_PASSWORD # The password is required to be BASE64 encrypted.
          value: ""
        - name: ULTRADNS_BASEURL
          value: "https://api.ultradns.com/"
        - name: ULTRADNS_ACCOUNTNAME
          value: ""
```

## Deploying an Nginx Service

Create a service file called 'nginx.yaml' with the following contents:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        name: nginx
        ports:
        - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: nginx
  annotations:
    external-dns.alpha.kubernetes.io/hostname: my-app.example.com.
spec:
  selector:
    app: nginx
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
```

Please note the annotation on the service. Use the same hostname as the UltraDNS zone created above.

ExternalDNS uses this annotation to determine what services should be registered with DNS. Removing the annotation will cause ExternalDNS to remove the corresponding DNS records.

## Creating the Deployment and Service:

```console
$ kubectl create -f nginx.yaml
$ kubectl create -f external-dns.yaml
```

Depending on where you run your service from, it can take a few minutes for your cloud provider to create an external IP for the service.

Once the service has an external IP assigned, ExternalDNS will notice the new service IP address and will synchronize the UltraDNS records.

## Verifying UltraDNS Records

Please verify on the [UltraDNS UI](https://portal.ultradns.neustar) that the records are created under the zone "example.com".

For more information on UltraDNS UI, refer to (https://docs.ultradns.neustar/mspuserguide.html).

Select the zone that was created above (or select the appropriate zone if a different zone was used.)

The external IP address will be displayed as a CNAME record for your zone.

## Cleaning Up the Deployment and Service

Now that we have verified that ExternalDNS will automatically manage your UltraDNS records, you can delete example zones that you created in this tutorial:

```
$ kubectl delete service -f nginx.yaml
$ kubectl delete service -f externaldns.yaml
```
## Examples to Manage your Records
### Creating Multiple A Records Target
- First, you want to create a service file called 'apple-banana-echo.yaml' 
```yaml
---
kind: Pod
apiVersion: v1
metadata:
  name: example-app
  labels:
    app: apple
spec:
  containers:
    - name: example-app
      image: hashicorp/http-echo
      args:
        - "-text=apple"
---
kind: Service
apiVersion: v1
metadata:
  name: example-service
spec:
  selector:
    app: apple
  ports:
    - port: 5678 # Default port for image
```
- Then, create service file called 'expose-apple-banana-app.yaml' to expose the services. For more information to deploy ingress controller, refer to (https://kubernetes.github.io/ingress-nginx/deploy/)
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
  annotations:
    ingress.kubernetes.io/rewrite-target: /
    ingress.kubernetes.io/scheme: internet-facing
    external-dns.alpha.kubernetes.io/hostname: apple.example.com.
    external-dns.alpha.kubernetes.io/target: 10.10.10.1,10.10.10.23
spec:
  rules:
  - http:
      paths:
        - path: /apple
          pathType: Prefix
          backend:
            service:
              name: example-service
              port:
                number: 5678
```
- Then, create the deployment and service:
```console
$ kubectl create -f apple-banana-echo.yaml
$ kubectl create -f expose-apple-banana-app.yaml
$ kubectl create -f external-dns.yaml
```
- Depending on where you run your service from, it can take a few minutes for your cloud provider to create an external IP for the service.
- Please verify on the [UltraDNS UI](https://portal.ultradns.neustar) that the records have been created under the zone "example.com".
- Finally, you will need to clean up the deployment and service. Please verify on the UI afterwards that the records have been deleted from the zone "example.com":
```console
$ kubectl delete -f apple-banana-echo.yaml
$ kubectl delete -f expose-apple-banana-app.yaml
$ kubectl delete -f external-dns.yaml
```
### Creating CNAME Record
- Please note, that prior to deploying the external-dns service, you will need to add the option –txt-prefix=txt- into external-dns.yaml. If this not provided, your records will not be created.
-  First, create a service file called 'apple-banana-echo.yaml'
    - _Config File Example – kubernetes cluster is on-premise not on cloud_
    ```yaml
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app
      labels:
        app: apple
    spec:
      containers:
        - name: example-app
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service
    spec:
      selector:
        app: apple
      ports:
        - port: 5678 # Default port for image
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: apple.example.com.
        external-dns.alpha.kubernetes.io/target: apple.cname.com.
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service
                  port:
                    number: 5678
    ```
    - _Config File Example – Kubernetes cluster service from different cloud vendors_
    ```yaml
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app
      labels:
        app: apple
    spec:
      containers:
        - name: example-app
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service
      annotations:
        external-dns.alpha.kubernetes.io/hostname: my-app.example.com.
    spec:
      selector:
        app: apple
      type: LoadBalancer
      ports:
        - protocol: TCP
          port: 5678
          targetPort: 5678
    ```
- Then, create the deployment and service:
```console
$ kubectl create -f apple-banana-echo.yaml
$ kubectl create -f external-dns.yaml
```
- Depending on where you run your service from, it can take a few minutes for your cloud provider to create an external IP for the service.
- Please verify on the [UltraDNS UI](https://portal.ultradns.neustar), that the records have been created under the zone "example.com".
- Finally, you will need to clean up the deployment and service. Please verify on the UI afterwards that the records have been deleted from the zone "example.com":
```console
$ kubectl delete -f apple-banana-echo.yaml
$ kubectl delete -f external-dns.yaml
```
### Creating Multiple Types Of Records
- Please note, that prior to deploying the external-dns service, you will need to add the option –txt-prefix=txt- into external-dns.yaml. Since you will also be created a CNAME record, If this not provided, your records will not be created.
-  First, create a service file called 'apple-banana-echo.yaml'
    - _Config File Example – kubernetes cluster is on-premise not on cloud_
    ```yaml
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app
      labels:
        app: apple
    spec:
      containers:
        - name: example-app
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service
    spec:
      selector:
        app: apple
      ports:
        - port: 5678 # Default port for image
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app1
      labels:
        app: apple1
    spec:
      containers:
        - name: example-app1
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service1
    spec:
      selector:
        app: apple1
      ports:
        - port: 5679 # Default port for image
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app2
      labels:
        app: apple2
    spec:
      containers:
        - name: example-app2
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service2
    spec:
      selector:
        app: apple2
      ports:
        - port: 5680 # Default port for image
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: apple.example.com.
        external-dns.alpha.kubernetes.io/target: apple.cname.com.
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service
                  port:
                    number: 5678
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress1
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: apple-banana.example.com.
        external-dns.alpha.kubernetes.io/target: 10.10.10.3
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service1
                  port:
                    number: 5679
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress2
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: banana.example.com.
        external-dns.alpha.kubernetes.io/target: 10.10.10.3,10.10.10.20
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service2
                  port:
                    number: 5680
    ```
    - _Config File Example – Kubernetes cluster service from different cloud vendors_
    ```yaml
    ---
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: nginx
    spec:
      selector:
        matchLabels:
          app: nginx
      template:
        metadata:
          labels:
            app: nginx
        spec:
          containers:
          - image: nginx
            name: nginx
            ports:
            - containerPort: 80
    ---
    apiVersion: v1
    kind: Service
    metadata:
      name: nginx
      annotations:
        external-dns.alpha.kubernetes.io/hostname: my-app.example.com.
    spec:
      selector:
        app: nginx
      type: LoadBalancer
      ports:
        - protocol: TCP
          port: 80
          targetPort: 80
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app
      labels:
        app: apple
    spec:
      containers:
        - name: example-app
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service
    spec:
      selector:
        app: apple
      ports:
        - port: 5678 # Default port for image
    ---
    kind: Pod
    apiVersion: v1
    metadata:
      name: example-app1
      labels:
        app: apple1
    spec:
      containers:
        - name: example-app1
          image: hashicorp/http-echo
          args:
            - "-text=apple"
    ---
    apiVersion: extensions/v1beta1
    kind: Service
    apiVersion: v1
    metadata:
      name: example-service1
    spec:
      selector:
        app: apple1
      ports:
        - port: 5679 # Default port for image
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: apple.example.com.
        external-dns.alpha.kubernetes.io/target: 10.10.10.3,10.10.10.25
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service
                  port:
                    number: 5678
    ---
    apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: example-ingress1
      annotations:
        ingress.kubernetes.io/rewrite-target: /
        ingress.kubernetes.io/scheme: internet-facing
        external-dns.alpha.kubernetes.io/hostname: apple-banana.example.com.
        external-dns.alpha.kubernetes.io/target: 10.10.10.3
    spec:
      rules:
      - http:
          paths:
            - path: /apple
              backend:
                service:
                  name: example-service1
                  port:
                    number: 5679
    ```
- Then, create the deployment and service:
```console
$ kubectl create -f apple-banana-echo.yaml
$ kubectl create -f external-dns.yaml
```
- Depending on where you run your service from, it can take a few minutes for your cloud provider to create an external IP for the service.
- Please verify on the [UltraDNS UI](https://portal.ultradns.neustar), that the records have been created under the zone "example.com".
- Finally, you will need to clean up the deployment and service. Please verify on the UI afterwards that the records have been deleted from the zone "example.com":
```console 
$ kubectl delete -f apple-banana-echo.yaml
$ kubectl delete -f external-dns.yaml```
