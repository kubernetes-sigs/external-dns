# Release

## Release cycle

Currently we don't release regularly. Whenever we think it makes sense to release a new version we do it. You might want to ask in our Slack channel [external-dns](https://kubernetes.slack.com/archives/C771MKDKQ) when the next release will come out.

## How to release a new image

When releasing a new version of external-dns, we tag the branch by using **vX.Y.Z** as tag name. This PR includes the updated **CHANGELOG.md** with the latest commits since last tag. As soon as we merge this PR into master, Kubernetes based CI/CD system [Prow](https://prow.k8s.io/?repo=kubernetes-sigs%2Fexternal-dns) will trigger a job to push the image. We're using the Google Container Registry for our Docker images.

The job itself looks at external-dns `cloudbuild.yaml` and executes the given steps. Inside it runs `make release.staging` which is basically only a `docker build` and `docker push`. The docker image is pushed `gcr.io/k8s-staging-external-dns/external-dns`, which is only a staging image and shouldn't be used. Promoting the official image we need to create another PR in [k8s.io](https://github.com/kubernetes/k8s.io), e.g. https://github.com/kubernetes/k8s.io/pull/540 by taking the current staging image using sha256.
