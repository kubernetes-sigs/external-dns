# Release

## Release cycle

Currently we don't release regularly. Whenever we think it makes sense to release a new version we do it, but we aim to do a new release every month. You might want to ask in our Slack channel [external-dns](https://kubernetes.slack.com/archives/C771MKDKQ) when the next release will come out.

## How to release a new image

### Prerequisite

We use https://github.com/cli/cli to automate the release process. Please install it according to the [official documentation](https://github.com/cli/cli#installation).

You must be an official maintainer of the project to be able to do a release.

### Steps

- Run `scripts/releaser.sh` to create a new GitHub release.
- The step above will trigger the Kubernetes based CI/CD system [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns). Verify that a new image was built and uploaded to `gcr.io/k8s-staging-external-dns/external-dns`.
- Create a PR in the [k8s.io repo](https://github.com/kubernetes/k8s.io) (see https://github.com/kubernetes/k8s.io/pull/540 for reference) by taking the current staging image using the sha256 digest. Once the PR is merged, the image will be live with the corresponding tag specified in the PR.
- Verify that the image is pullable with the given tag (i.e. `v0.7.5`).
- Branch out from the default branch and run `scripts/kustomize-version-udapter.sh` to update the image tag used in the kustomization.yaml.
- Create a PR with the kustomize change.
- Once the PR is merged, all is done :-)
