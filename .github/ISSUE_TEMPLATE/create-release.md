---
name: Create Release
about: Release template to track the next release
title: Release x.y
labels: area/release
assignees: ''

---

This Issue tracks the next `external-dns` release. Please follow the guideline below. If anything is missing or unclear, please add a comment to this issue so this can be improved after the release.

#### Preparation Tasks

- [ ] Release [steps](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/release.md#steps)

#### Release Execution

- [ ] Branch out from the default branch and run scripts/version-updater.sh to update the image tag used in the kustomization.yaml and in documentation.
- [ ] Create the PR with this version change.
- [ ] Create an issue to release the corresponding Helm chart via the chart release process (below) assigned to a chart maintainer

#### After Release Tasks

- [ ] Announce release on `#external-dns` in Slack
