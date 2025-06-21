# Contributing Guidelines

Welcome to Kubernetes. We are excited about the prospect of you joining our [community](https://git.k8s.io/community)! The Kubernetes community abides by the CNCF [code of conduct](code-of-conduct.md). Here is an excerpt:

_In the interest of fostering an open and welcoming community, we pledge to respect all people who contribute through reporting issues, posting feature requests, updating documentation, submitting pull requests or other activities._

## Getting Started

We have full documentation on how to get started contributing here:

- [Contributor License Agreement](https://git.k8s.io/community/CLA.md) Kubernetes projects require that you sign a Contributor License Agreement (CLA) before we can accept your pull requests
- [Kubernetes Contributor Guide](https://git.k8s.io/community/contributors/guide) - Main contributor documentation, or you can just jump directly to the [contributing section](https://git.k8s.io/community/contributors/guide#contributing)
- [Contributor Cheat Sheet](https://git.k8s.io/community/contributors/guide/contributor-cheatsheet) - Common resources for existing developers

This project follows the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification on PR title. The explicit commit history is used, among other things, to provide a readable changelog in release notes.

## How to test a PR

On Linux (or WSL), a PR can be tested following this instruction with [gh](https://cli.github.com/) and [golang](https://go.dev/):

```bash
gh repo clone kubernetes-sigs/external-dns
cd external-dns
gh pr checkout XXX # <=== Set PR number here
go run main.go \
  --kubeconfig=<kubeconfig_path> \
  --log-format=text \
  --log-level=debug \
  --interval=1m
  --provider=xxx
  --source=yyy
```

## Mentorship

- [Mentoring Initiatives](https://git.k8s.io/community/mentoring) - We have a diverse set of mentorship programs available that are always looking for volunteers!

## Contact Information

- [Slack channel](https://kubernetes.slack.com/messages/external-dns)
- [Mailing list](https://groups.google.com/forum/#!forum/kubernetes-sig-network)
