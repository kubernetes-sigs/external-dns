# Setting up external-dns for BlueCat

The first external-dns release with with BlueCat provider support is v0.8.0.

## Prerequisites
Install the BlueCat Gateway product and deploy the [community gateway workflows](https://github.com/bluecatlabs/gateway-workflows).

## Configuration Options

There are two ways to pass configuration options to the Bluecat Provider JSON configuration file and command line flags. Currently if a valid configuration file is used all
BlueCat provider configurations will be taken from the configuration file. If a configuraiton file is not provided or cannot be read then all BlueCat provider configurations will
be taken from the command line flags. In the future an enhancement will be made to merge configuration options from the configuration file and command line flags if both are provided.

BlueCat provider supports getting the proxy URL from the environment variables. The format is the one specified by golang's [http.ProxyFromEnvironment](https://pkg.go.dev/net/http#ProxyFromEnvironment).

### Using CLI Flags
When using CLI flags to configure the Bluecat Provider the BlueCat Gateway credentials are passed in using environment variables `BLUECAT_USERNAME` and `BLUECAT_PASSWORD`.

#### Deploy
Setup up namespace, deployment, and service account:
```
kubectl create namespace bluecat-example
kubectl create secret generic bluecat-credentials --from-literal=username=bluecatuser --from-literal=password=bluecatpassword -n bluecat-example
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
      containers:
      - name: external-dns
        image: k8s.gcr.io/external-dns/external-dns:v0.8.0
        args:
        - --log-level=debug
        - --source=service
        - --provider=bluecat
        - --txt-owner-id=bluecat-example
        - --bluecat-dns-configuration=Example
        - --bluecat-dns-view=Internal
        - --bluecat-gateway-host=https://bluecatgw.example.com
        - --bluecat-root-zone=example.com
        env:
        - name: BLUECAT_USERNAME
          valueFrom:
            secretKeyRef:
              name: bluecat-credentials
              key: username
        - name: BLUECAT_PASSWORD
          valueFrom:
            secretKeyRef:
              name: bluecat-credentials
              key: password
EOF
kubectl apply -f ~/bluecat.yml -n bluecat-example
```


### Using JSON Configuration File
The options for configuring the Bluecat Provider are available through the JSON file provided to External-DNS via the flag `--bluecat-config-file`.

| Key               | Required           |
| ----------------- | ------------------ |
| gatewayHost       | Yes                |
| gatewayUsername   | No                 |
| gatewayPassword   | No                 |
| dnsConfiguration  | Yes                |
| dnsView           | Yes                |
| rootZone          | Yes                |
| dnsServerName     | No                 |
| dnsDeployType     | No                 |
| skipTLSVerify     | No (default false) |

#### Deploy
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
