# Developer Reference

The `external-dns` is the work of thousands of contributors, and is maintained by a small team within [kubernetes-sigs](https://github.com/kubernetes-sigs). This document covers basic needs to work with `external-dns` codebase. It contains instructions to build, run, and test `external-dns`.

## Tools

Building and/or testing `external-dns` requires additional tooling.

- [Git](https://git-scm.com/downloads)
- [Go 1.23+](https://golang.org/dl/)
- [Go modules](https://github.com/golang/go/wiki/Modules)
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [ko](https://ko.build/)
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl)
- [helm](https://helm.sh/docs/helm/helm_install/)
- [spectral](https://github.com/stoplightio/spectral)
- [python](https://www.python.org/downloads/)

## First Steps

***Configure Development Environment**

You must have a working [Go environment](https://go.dev/doc/install), compile the build, and set up testing.

```shell
git clone https://github.com/kubernetes-sigs/external-dns.git && cd external-dns
```

## Building & Testing

The project uses the make build system. It'll run code generators, tests and static code analysis.

Build, run tests and lint the code:

```shell
make go-lint
make test
make cover-html
```

If added any flags, re-generate flags documentation

```shell
make generate-flags-documentation
```

We require all changes to be covered by acceptance tests and/or unit tests, depending on the situation.
In the context of the `external-dns`, acceptance tests are tests of interactions with providers, such as creating, reading information about, and destroying DNS resources. In contrast, unit tests test functionality wholly within the codebase itself, such as function tests.

### Continuous Integration

When submitting a pull request, you'll notice that we run several automated processes on your proposed change. Some of these processes are tests to ensure your contribution aligns with our standards. While we strive for accuracy, some users may find these tests confusing.

## Execute code without building binary

The `external-dns` does not require `make build`. You could compile and run Go program with the command

```sh
go run main.go \
    --provider=aws \
    --registry=txt \
    --source=fake \
    --log-level=info
```

For this command to run successfully, it will require [AWS credentials](https://docs.aws.amazon.com/cli/v1/userguide/cli-configure-files.html) and access to local or remote access.

To run local cluster please refer to [running local cluster](#create-a-local-cluster)

## Deploying a local build

After building local images, it is often useful to deploy those images in a local cluster

We use [Minikube](https://minikube.sigs.k8s.io/docs/start/?arch=%2Fmacos%2Fx86-64%2Fstable%2Fbinary+download) but it could be [Kind](https://kind.sigs.k8s.io/) or any other solution.

- [Create local cluster](#create-a-local-cluster)
- [Build and load local images](#building-local-images)
- Deploy with Helm
- Deploy with kubernetes manifests

## Create a local cluster

For simplicity, [minikube](https://minikube.sigs.k8s.io) can be used to create a single
node cluster.

You can set a specific Kubernetes version by setting the node's container image.
See [basic controls](https://minikube.sigs.k8s.io/docs/handbook/controls/) within the documentation about configuration for more details on this.

Once you have a configuration in place, create the cluster with
that configuration:

```sh
minikube start \
  --profile=external-dns \
  --memory=2000 \
  --cpus=2 \
  --disk-size=5g \
  --kubernetes-version=v1.31 \
  --driver=docker

minikube profile external-dns
```

After the new Kubernetes cluster is ready, identify the cluster is running as the single node cluster:

```sh
❯❯ kubectl get nodes
NAME           STATUS   ROLES           AGE   VERSION
external-dns   Ready    control-plane   16s   v1.31.4
```

---

## Building local images

When building local images with ko you can't specify the registry used to create the image names. It will always be ko.local.

- [minikube handbooks](https://minikube.sigs.k8s.io/docs/handbook/pushing/)

> Note: You could skip this step if you build and push image to your private registry or using an official external-dns image

```sh
❯❯ export KO_DOCKER_REPO=ko.local
❯❯ export VERSION=v1
❯❯ docker context use rancher-desktop ## (optional) this command is only required when using rancher-desktop
❯❯ ls -al /var/run/docker.sock ## (optional) validate tha docker runtime is configured correctly and symlink exist

❯❯ ko build --tags ${VERSION}
❯❯ docker images
$$ ko.local/external-dns-9036f6870f30cbdefa42a10f30bada63   local-v1
```

***Push image to minikube***

Refer to [load image](https://minikube.sigs.k8s.io/docs/handbook/pushing/#7-loading-directly-to-in-cluster-container-runtime)

```sh
❯❯ minikube image load ko.local/external-dns-9036f6870f30cbdefa42a10f30bada63:local-v1
❯❯ minikube image ls
$$ registry.k8s.io/pause:3.10
$$ ...
$$ ko.local/external-dns-9036f6870f30cbdefa42a10f30bada63:local-v1
$$ ...
❯❯ kubectl run external-dns --image=ko.local/external-dns-9036f6870f30cbdefa42a10f30bada63:local-v1 --image-pull-policy=Never
```

***Build and push directly in minikube***

Any `docker` command you run in this current terminal will run against the docker inside minikube cluster.

Refer to [push directly](https://minikube.sigs.k8s.io/docs/handbook/pushing/#1-pushing-directly-to-the-in-cluster-docker-daemon-docker-env)

```sh
❯❯ eval $(minikube -p external-dns docker-env)
❯❯ echo $MINIKUBE_ACTIVE_DOCKERD
$$ external-dns
❯❯ export VERSION=v1
❯❯ ko build --local --tags ${VERSION}
❯❯ docker images
$$ REPOSITORY                                               TAG
$$ registry.k8s.io/kube-apiserver                           v1.31.4
$$ ....
$$ ko.local/external-dns-9036f6870f30cbdefa42a10f30bada63   minikube-v1
$$ ...
❯❯ eval $(minikube docker-env -u) ## unset minikube
```

***Pushing to an in-cluster using Registry addon***

Refer to [pushing images](https://minikube.sigs.k8s.io/docs/handbook/pushing/#4-pushing-to-an-in-cluster-using-registry-addon) for a full configuration

```sh
❯❯ export KO_DOCKER_REPO=$(minikube ip):5000
❯❯ export VERSION=registry-v1
❯❯ minikube addons enable registry
❯❯ ko build --tags ${VERSION}
```

## Building image and push to a registry

Build container image and push to a specific registry

```shell
make build.push IMAGE=your-registry/external-dns
```

---

## Deploy with Helm

Build local images if required, load them on a local cluster, and deploy helm charts, run:

Render chart templates locally and display the output

```sh
❯❯ helm lint --debug charts/external-dns
❯❯ helm template external-dns charts/external-dns --output-dir _scratch
```

Deploy manifests to a cluster with required values

```sh
❯❯ kubectl apply -f _scratch --recursive=true
```

Modify chart or values and validate the diff

```sh
❯❯ helm template external-dns charts/external-dns --output-dir _scratch
❯❯ kubectl diff -f _scratch/external-dns --recursive=true --show-managed-fields=false
```

### Helm Values

This helm chart comes with a JSON schema generated from values with [helm schema](https://github.com/losisin/helm-values-schema-json.git) plugin.

1. Install required plugin(s)

```sh
❯❯ scripts/helm-tools.sh --install
```

2. Ensure that the schema is always up-to-date

```sh
❯❯ scripts/helm-tools.sh --diff
```

3. When not up-to-date, update JSON schema

```sh
❯❯ scripts/helm-tools.sh --schema
```

4. Runs a series of tests to verify that the chart is well-formed, linted and JSON schema is valid

```sh
❯❯ scripts/helm-tools.sh --lint
```

5. Auto-generate documentation for helm charts into markdown files.

```sh
❯❯ scripts/helm-tools.sh --docs
```

6. Add an entry to the chart [CHANGELOG.md](../../charts/external-dns/CHANGELOG.md) under `## UNRELEASED` section and `open` pull request

## Deploy with kubernetes manifests

> Note; kubernetes manifest are not up to date. Consider to create an `examples` folder

```sh
kubectl apply -f kustomize --recursive=true --dry-run=client
```

## Contribute to documentation

All documentation is in `docs` folder. If new page is added or removed, make sure `mkdocs.yml` is also updated.

Install required dependencies. In order to not to break system packages, we are going to use virtual environments with [pipenv](https://pipenv.pypa.io/en/latest/installation.html).

```sh
❯❯ pipenv shell
❯❯ pip install -r docs/scripts/requirements.txt
❯❯ mkdocs serve
$$ ...
$$ Serving on http://127.0.0.1:8000/
```
