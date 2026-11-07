---
Do not merge unless explicitly authorized and required checks/reviews pass.
Otherwise wait for a maintainer to merge. Verify `mergedAt` remotely before
proceeding; an approved or closed PR alone does not mean it merged.

## 4. Wait for the production image to be pullable

After the promotion PR merges, poll
`crane digest "registry.k8s.io/external-dns/external-dns:$VERSION"`.
Require its index digest to equal the promoted staging digest.

Then run both commands against the release tag:

```sh
docker pull "registry.k8s.io/external-dns/external-dns:$VERSION"
docker run --rm --pull=always "registry.k8s.io/external-dns/external-dns:$VERSION" --version
```

Require a successful fresh pull and version output matching the requested release
(allow the binary's version formatting to omit the leading `v`). Record the
resolved digest and output. A cached local image, a successful build, a merged
promotion PR, or a resolvable manifest alone does not satisfy this gate.
If Docker cannot run, report the gate as blocked rather than skipping it.

## 5. Create and track the version-updater PR

**Only after step 4 passes**, use a clean branch based on the current ExternalDNS
default branch. Determine `PREV_TAG` from `kustomize/kustomization.yaml` and check
the existing documentation image references; do not guess it from a chart tag.

Run `scripts/version-updater.sh "$PREV_TAG" "$VERSION"`. This script stages files
and creates a signed-off commit itself. In environments requiring a pre-commit
secret scan, inspect and scan its prospective input files before running it;
do not bypass the environment's commit policy.

Review the resulting commit for only the intended kustomize/docs image changes,
including `docs/release.md`. Validate the changed manifests and documentation
using existing repository checks, and scan changed files before publishing.
Do not run this script on the release tag or try to move the tag to this commit.

Create a PR titled `chore(release): update manifests and docs for $VERSION`,
following the PR template and linking the release, promotion PR, and verified
production image digest. Watch checks and merge state as in step 3, merging only
with authorization. Verify the merged default-branch files reference the version.

## 6. Hand off the Helm chart release and report completion

Create or reuse an issue requesting the corresponding Helm chart release.
Link the image release, promotion PR, and merged version-updater PR, and assign
the confirmed chart maintainer. Do not invent an assignee or independently choose
a chart version: the chart has a separate release cycle.

The chart maintainer's follow-up PR updates `charts/external-dns/Chart.yaml`
(`appVersion` without the image tag's leading `v`, an agreed chart `version`,
and change annotations) and the chart changelog as required by the current
workflow. Chart linting must pass before merge.
`.github/workflows/release-chart.yaml` invokes `scripts/chart-release.sh` after
merge; do not invoke that publishing script locally as part of an image release.

Mark the image release complete only when the production pull/version gate
passed, the version-updater PR merged, and the chart handoff issue exists.
Report all links, the verified digest/version, any pending chart work, and remind
the maintainer to announce the release in `#external-dns` on Kubernetes Slack.
Do not claim an announcement was sent unless it actually was.

