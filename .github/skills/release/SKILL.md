---
name: release
description: Release ExternalDNS using the existing scripts, create and track image-promotion and version-update PRs, and wait for staging and production images to become pullable before updating manifests.
---

# Release ExternalDNS

Use this skill when a maintainer asks to perform or resume an ExternalDNS image
release. Follow `docs/release.md` and `.github/ISSUE_TEMPLATE/create-release.md`
in order. Execute the work, rather than just returning a checklist, subject to
the user's authorization and the environment's permissions.

Do not publish anything when asked only to explain, preview, or dry-run a release.
Do not start a release merely because this skill was added or edited.

## Inputs and preflight

Confirm these inputs before making remote changes:

- The image version, exactly `vMAJOR.MINOR.PATCH`, and patch/minor choice.
- The target repository and intended release commit on its default branch.
  Official releases use `kubernetes-sigs/external-dns`, not a personal fork.
- Authorization to publish the release and create the promotion PR, manifest PR,
  and chart-release issue. Obtain separate approval before merging PRs.
- The release-tracking issue, if any, and the chart maintainer to assign.

Read the current scripts before invoking them; their behavior is authoritative.
Require a clean working tree and authenticated GitHub access with the appropriate
maintainer permissions. Check for Bash, Git, `gh`, `jq`, `crane`, and a working
Docker daemon. `scripts/version-updater.sh` requires GNU-compatible `sed`.
Do not install tools, expose credentials, or change authentication without approval.

Use absolute checkout paths for local commands and run scripts from the
ExternalDNS repository root. Set `GH_REPO` explicitly for `scripts/releaser.sh`,
whose GitHub commands otherwise infer the repository from the checkout.
Use explicit repositories for all other GitHub operations as well.
Discover default branches instead of assuming `master` everywhere.

Check existing releases, tags, open/merged PRs, and chart issues for this version
before creating anything. Resume existing work; never recreate, retag, overwrite,
or delete a published release. If a tag exists, verify its commit matches the
approved commit. For a new tag, confirm the repository's current default-branch
HEAD is the approved commit: `releaser.sh` does not pass a target commit to GitHub.
Stop if these differ.

Use the environment's approved PR creation, commit/push, and merge tools.
Where unrestricted local tooling is available, GitHub CLI can manage the PR
lifecycle. Never bypass repository, branch, or cross-repository restrictions.
If the environment cannot edit `kubernetes/k8s.io`, request an authorized
checkout/session and report the blocked promotion step; do not claim it is done.

## Waiting and recovery

For each asynchronous gate below, actively recheck remote state at 60-second
intervals, with a 30-minute deadline per gate unless the maintainer chooses
another limit. Use supported background/wait mechanisms for long waits.
Do not substitute a fixed sleep or workflow success for an artifact check.

On timeout, denied access, failed CI, closed-unmerged PR, or digest/version
mismatch, stop before the next dependent step. Report the last observation,
relevant URLs, and the exact gate to resume. Never automatically merge to unblock
a wait, bypass checks, or recreate resources after an ambiguous write response.

Keep a checklist in the release-tracking issue or conversation with the release
URL, tag commit, staging digest, PR URLs, image verification results, and chart
issue URL. Revalidate these gates against live state when resuming.

## 1. Preview and publish the release

Run `scripts/releaser.sh` **without arguments** to preview the generated notes.
Review the version range, breaking changes, and exclusion of chart-only changes
with the maintainer before publication.

After approval, invoke `scripts/releaser.sh "$VERSION"` with the confirmed version
as its only argument. The script publishes a GitHub **prerelease** using `-p`;
it is not a draft or a dry run. Do not silently change that behavior. If a stable
release is required, confirm the intended status change with the maintainer.

Read back the release and tag, record their URL and commit, and verify the tag
matches the approved commit before continuing.

## 2. Wait for the staging image

Track the tag-triggered build in [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns).
Poll `crane digest "gcr.io/k8s-staging-external-dns/external-dns:$VERSION"`
until it resolves. Inspect the build/tag association to ensure this is the
approved release, not an unrelated weekly staging build.

Run `scripts/get-sha256.sh "gcr.io/k8s-staging-external-dns/external-dns:$VERSION"`.
Record the top-level `image:` sha256 digest and the architecture manifest list.
Use the **multi-architecture image/index digest**, not an individual platform
digest, for promotion. Stop if any command fails or the digest is invalid.

## 3. Create and track the image-promotion PR

In an authorized `kubernetes/k8s.io` checkout, read its contributor instructions
and locate the current external-dns image-promotion manifest under
`registry.k8s.io/images/k8s-staging-external-dns/`. Use the current manifest
schema; [promotion PR #8466](https://github.com/kubernetes/k8s.io/pull/8466)
is historical context, not a patch to copy blindly.

Branch from that repository's current default branch. Add the confirmed release
tag to the staging index digest, preserving previous promotions and unrelated
entries. If this tag is already mapped to another digest, stop for review.
Run that repository's required validation, inspect the diff, scan for secrets,
and commit/push using approved tools.

Create a PR titled `chore(external-dns): promote $VERSION`, following the target
repository's PR template. Include the release URL, source image, index digest,
and architecture evidence. Reuse an existing equivalent PR.

Watch checks and review state. If CI fails, inspect workflow runs and job logs
using GitHub tools; Prow checks must be investigated through their linked jobs.
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
