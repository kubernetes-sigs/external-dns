# Setting up external-dns for BlueCat

## Prerequisites
Install the BlueCat Gateway product and deploy the [community gateway workflows](https://github.com/bluecatlabs/gateway-workflows).

## Configuration Options

The options for configuring the Bluecat Provider are available through the json file provided to External-DNS via the flag `--bluecat-config-file`.

| Key               | Required           |
| ----------------- | ------------------ |
| gatewayHost       | Yes                |
| gatewayUsername   | Yes                |
| gatewayPassword   | Yes                |
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

Setup up deployment/service account:
```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: bluecat-example
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: bluecat-example
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
        image: k8s.gcr.io/external-dns/external-dns:$TAG # no released versions include the bluecat provider yet
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
```
