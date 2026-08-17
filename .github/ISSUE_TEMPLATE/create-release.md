---
name: Create Release
about: Release template to track the next release
title: Release x.y
labels: area/release
assignees: ''

---

This Issue tracks the next `external-dns` release. Please follow the guideline below. If anything is missing or unclear, please add a comment to this issue so this can be improved after the release.

Full process: [docs/release.md](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/release.md).

## Preparation Tasks

- [ ] Confirm version bump type (patch/minor) per [versioning convention](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/release.md#versioning-convention)
- [ ] Release [steps](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/release.md#steps) reviewed

### Release Execution (order matters)

> **Chicken-and-egg:** the git tag must exist before the image is built/promoted; See [git tags vs kustomize](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/release.md#git-tags-vs-kustomize-manifests-known-lag).

- [ ] Create the GitHub release / git tag (`scripts/releaser.sh` or GitHub UI)
- [ ] Confirm staging image built for this tag (`gcr.io/k8s-staging-external-dns/external-dns`)
- [ ] Promote image via k8s.io PR (digest from `scripts/get-sha256.sh`) and wait for merge
- [ ] Verify pull: `docker run registry.k8s.io/external-dns/external-dns:vX.Y.Z --version`
- [ ] Run `scripts/version-updater.sh` on a branch from default; open PR for kustomize + docs image tags
- [ ] Merge version-updater PR (manifests on `master` now match the release image)
- [ ] Create an issue to release the Helm chart; assign a chart maintainer

### After Release Tasks

- [ ] Announce release on `#external-dns` in Slack
