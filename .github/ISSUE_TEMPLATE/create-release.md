---
name: Create Release
about: Release template to track the next release
title: Release x.y
labels: area/release
assignees: ''

---

This Issue tracks the next `external-dns` release. Please follow the guideline below. If anything is missing or unclear, please add a comment to this issue so this can be improved after the release.

#### Preparation Tasks

- [ ] Run `scripts/releaser.sh` to create a new GitHub release. Alternatively you can create a release in the GitHub UI making sure to click on the autogenerate release node feature.
    - The step above will trigger the Kubernetes based CI/CD system [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns). Verify that a new image was built and uploaded to gcr.io/k8s-staging-external-dns/external-dns.
- [ ] Create a PR in the [k8s.io repo](https://github.com/kubernetes/k8s.io) by taking the current staging image using the sha256 digest. They can be obtained with `scripts/get-sha256.sh`. Once the PR is merged, the image will be live with the corresponding tag specified in the PR.
      - See https://github.com/kubernetes/k8s.io/pull/540 for reference
- [ ] Verify that the image is pullable with the given tag
   - `docker run registry.k8s.io/external-dns/external-dns:v0.16.0 --version`

#### Release Execution

- [ ] Branch out from the default branch and run scripts/version-updater.sh to update the image tag used in the kustomization.yaml and in documentation.
- [ ] Create the PR with this version change.
- [ ] Create an issue to release the corresponding Helm chart via the chart release process (below) assigned to a chart maintainer

#### After Release Tasks

- [ ] Announce release on `#external-dns` in Slack
