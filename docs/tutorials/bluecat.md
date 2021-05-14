# Setting up external-dns for BlueCat

The first external-dns release with with BlueCat provider support is v0.8.0.

## Prerequisites
Install the BlueCat Gateway product and deploy the [community gateway workflows](https://github.com/bluecatlabs/gateway-workflows).

## Configuration Options

The options for configuring the Bluecat Provider are available through the json file provided to External-DNS via the flag `--bluecat-config-file`. The
BlueCat Gateway username and password can be supplied using the configuration file or environment variables `BLUECAT_USERNAME` and `BLUECAT_PASSWORD`.

| Key               | Required           |
| ----------------- | ------------------ |
| gatewayHost       | Yes                |
| gatewayUsername   | No                 |
| gatewayPassword   | No                 |
| dnsConfiguration  | Yes                |
| dnsView           | Yes                |
| rootZone          | Yes                |
| skipTLSVerify     | No (default false) |

## Deploy
Setup configuration file as k8s `Secret`.
```
cat << EOF > ~/bluecat.json
{
  "gatewayHost": "https://bluecatgw.example.com",
  "gatewayUsername": "user",
  "gatewayPassword": "pass",
  "dnsConfiguration": "Example",
  "dnsView": "Internal",
  "rootZone": "example.com",
  "skipTLSVerify": false
}
EOF
kubectl create secret generic bluecatconfig --from-file ~/bluecat.json -n bluecat-example
```

Setup up namespace, deployment, and service account:
```
kubectl create namespace bluecat-example
cat << EOF > ~/bluecat.yml
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
spec:
  selector:
    matchLabels:
      app: external-dns
  strategy:
    type: Recreate
  template:
    metadata:
      labels:
        app: external-dns
    spec:
      serviceAccountName: external-dns
      volumes:
        - name: bluecatconfig
          secret:
            secretName: bluecatconfig
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.8.0
        volumeMounts:
          - name: bluecatconfig
            mountPath: "/etc/external-dns/"
            readOnly: true
        args:
        - --log-level=debug
        - --source=service
        - --provider=bluecat
        - --txt-owner-id=bluecat-example
        - --bluecat-config-file=/etc/external-dns/bluecat.json
EOF
kubectl apply -f ~/bluecat.yml -n bluecat-example
```
