# Release

## Release cycle

Currently we don't release regularly. Whenever we think it makes sense to release a new version we do it.
You might want to ask in our Slack channel [external-dns](https://kubernetes.slack.com/archives/C771MKDKQ) when the next release will come out.

## Staging Release cycle

A new staging image is released weekly and can be found at [gcr.io/k8s-staging-external-dns/external-dns](https://console.cloud.google.com/gcr/images/k8s-staging-external-dns/GLOBAL/external-dns?pli=1&inv=1&invt=AboL6Q).

> There is a time lag between merging changes into the master branch and the subsequent creation of the staging image.

Example command to fetch `10` most recent staging images:

```sh
export EXT_DNS_VERSION="v0.22.0"
curl -sLk https://gcr.io/v2/k8s-staging-external-dns/external-dns/tags/list | jq | grep "$EXT_DNS_VERSION" | tail -n 10
```

## Versioning convention

These are the conventions that we will be using for releases following `0.7.6`:

- **Patch** version should be updated if we need to merge bugfixes, e.g. provider a does need a fix in order make updates working again. I would see updating or improving documentation here.

- **Minor** version should be updated if new features are implemented in existing providers or new provider get introduced.

- **Major** version should be upgraded if we introduce breaking changes.

### Semantic Versioning Discipline

External-DNS follows semantic versioning principles:

- `0.x` → pre-stable, APIs subject to change.
- `1.x` → not yet considered.

> **Versioning & Releases**
> External-DNS opts to stay within `0.x` versioning scheme.
> We strive for stability, but reserve the right to introduce breaking changes in minor version bumps when necessary.

## How to release a new image

### Prerequisite

We use https://github.com/cli/cli to automate the release process. Please install it according to the [official documentation](https://github.com/cli/cli#installation).

You must be an official maintainer of the project to be able to do a release.

### Steps

1. Run `scripts/releaser.sh` to create a new GitHub release (and git tag). Alternatively create a release in the GitHub UI and use the autogenerate release notes feature.
2. The step above triggers the Kubernetes CI/CD system [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns). Verify that a new image was built and uploaded to `gcr.io/k8s-staging-external-dns/external-dns`.
3. Create a PR in the [k8s.io repo](https://github.com/kubernetes/k8s.io) promoting the staging image by **sha256 digest** (from `scripts/get-sha256.sh`). Once that PR merges, the image is available at `registry.k8s.io` under the release tag.
   - See https://github.com/kubernetes/k8s.io/pull/8466 for reference
4. Verify that the image is pullable with the given tag:
   - `docker run registry.k8s.io/external-dns/external-dns:v0.x.0 --version`
5. **Only after** the image is pullable: branch from the default branch and run `scripts/version-updater.sh` to update the image tag in `kustomize/` manifests and documentation.
6. Open and merge the version-updater PR.
7. Open an issue to release the corresponding Helm chart (chart process below), assigned to a chart maintainer.
8. Once the version-updater PR is merged, the release is complete.

### Git tags vs kustomize manifests (known lag)

Image promotion and in-repo manifest updates **cannot** happen in a single commit on the release tag. CI needs the git tag first to build and promote the image; kustomize/docs must not advertise a tag until that image is live on `registry.k8s.io`.

Consequences:

| Source | Image tag in manifests |
|--------|------------------------|
| Git **release tag** (e.g. `v0.18.0`) | Still points at the **previous** release image until a follow-up commit lands |
| **`master`** after the version-updater PR | Matches the new release image |
| **Helm chart** `appVersion` after chart release | Matches the new release image |

So:

- Do **not** treat the kustomize tree _on the git tag_ as the source of truth for that version’s image.
- Prefer the Helm chart, or the post-release version-updater commit on `master`, for install manifests that pin the new tag.

## How to release a new chart version

The chart needs to be released in response to an ExternalDNS image release or on an as-needed basis; this should be triggered by an issue to release the chart.

### Steps

- Create a PR to update _Chart.yaml_ with the ExternalDNS version in `appVersion`, agreed on chart release version in `version` and `annotations` showing the changes
- Validate that the chart linting is successful
- Merge the PR to trigger a GitHub action to release the chart
