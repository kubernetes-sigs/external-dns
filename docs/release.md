# Release

## Release cycle

Currently we don't release regularly. Whenever we think it makes sense to release a new version we do it.
You might want to ask in our Slack channel [external-dns](https://kubernetes.slack.com/archives/C771MKDKQ) when the next release will come out.

## Staging Release cycle

A new staging image is released weekly and can be found at [gcr.io/k8s-staging-external-dns/external-dns](https://console.cloud.google.com/gcr/images/k8s-staging-external-dns/GLOBAL/external-dns?pli=1&inv=1&invt=AboL6Q).

> There is a time lag between merging changes into the master branch and the subsequent creation of the staging image.

Example command to fetch `10` most recent staging images:

```sh
export EXT_DNS_VERSION="v0.15.1"
curl -sLk https://gcr.io/v2/k8s-staging-external-dns/external-dns/tags/list | jq | grep "$EXT_DNS_VERSION" | tail -n 10
```

## Versioning convention

These are the conventions that we will be using for releases following `0.7.6`:

- **Patch** version should be updated if we need to merge bugfixes, e.g. provider a does need a fix in order make updates working again. I would see updating or improving documentation here.

- **Minor** version should be updated if new features are implemented in existing providers or new provider get introduced.

- **Major** version should be upgraded if we introduce breaking changes.

## How to release a new image

### Prerequisite

We use https://github.com/cli/cli to automate the release process. Please install it according to the [official documentation](https://github.com/cli/cli#installation).

You must be an official maintainer of the project to be able to do a release.

### Steps

- Run `scripts/releaser.sh` to create a new GitHub release. Alternatively you can create a release in the GitHub UI making sure to click on the autogenerate release node feature.
- The step above will trigger the Kubernetes based CI/CD system [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns). Verify that a new image was built and uploaded to `gcr.io/k8s-staging-external-dns/external-dns`.
- Create a PR in the [k8s.io repo](https://github.com/kubernetes/k8s.io) by taking the current staging image using the sha256 digest. Once the PR is merged, the image will be live with the corresponding tag specified in the PR.
  - See https://github.com/kubernetes/k8s.io/pull/540 for reference
- Verify that the image is pullable with the given tag (i.e. `v0.7.5`).
- Branch out from the default branch and run `scripts/kustomize-version-updater.sh` to update the image tag used in the kustomization.yaml.
- Create an issue to release the corresponding Helm chart via the chart release process (below) assigned to a chart maintainer
- Create a PR with the kustomize change.
- Create a PR to replace all versions for docker images in the tutorials. A possible script to use is `sd registry.k8s.io/external-dns/external-dns:v0.15.1`
- Once the PR is merged, all is done :-)

## How to release a new chart version

The chart needs to be released in response to an ExternalDNS image release or on an as-needed basis; this should be triggered by an issue to release the chart.

### Steps

- Create a PR to update _Chart.yaml_ with the ExternalDNS version in `appVersion`, agreed on chart release version in `version` and `annotations` showing the changes
- Validate that the chart linting is successful
- Merge the PR to trigger a GitHub action to release the chart
