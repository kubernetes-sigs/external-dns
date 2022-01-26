# Setting up external-dns for 6connect ProVision

The first external-dns release with with ProVision provider support


## Configuration Options

The options for configuring the 6connect ProVision Provider are available through the json file provided to External-DNS via the flag `--pv6connect-config-file`. You can also use ENV variables or External DNS flags to set the required options

| Key               | Required           | Env                 | Flag                   |
| ----------------- | ------------------ | ------------------- | ---------------------- |
| provisionHost     | Yes                | PROVISION_HOST      | --pv6connect-host      |
| provisionUsername | Yes                | PROVISION_USERNAME  | --pv6connect-username  |
| provisionPassword | Yes                | PROVISION_PASSWORD  | --pv6connect-password  |
| zoneIDs           | Yes                | PROVISION_ZONEIDS   | --pv6connect-zoneids   |
| provisionPush     | No                 |                     | --pv6connect-push      |
| skipTLSVerify     | No (default false) |                     |                        |


### HTTP proxy

ProVision provider supports getting the proxy URL from the environment variables. The format is the one specified by golang's [http.ProxyFromEnvironment](https://pkg.go.dev/net/http#ProxyFromEnvironment).

## Deploy
Setup configuration file as k8s `Secret`.
```
cat << EOF > ~/provision.json
{
  "provisionHost": "https://example.com/6.2.0",
  "provisionUsername": "example@example.com",
  "provisionPassword": "password11",
  "zoneIDs": ["428964","10819","428972"],
  "skipTLSVerify": false
}
EOF
kubectl create secret generic provisionconfig --from-file ~/provision.json -n provision-example
```

Setup up namespace, deployment, and service account:
```
kubectl create namespace provision-example
cat << EOF > ~/provision.yml
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
        - name: provisionconfig
          secret:
            secretName: provisionconfig
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.8.0
        volumeMounts:
          - name: provisionconfig
            mountPath: "/etc/external-dns/"
            readOnly: true
        args:
        - --log-level=debug
        - --source=service
        - --provider=pv6connect
        - --txt-owner-id=provision-example
        - --pv6connect-config-file=/etc/external-dns/provision.json
EOF
kubectl apply -f ~/provision.yml -n provision-example
```
