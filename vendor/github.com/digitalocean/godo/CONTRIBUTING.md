# Contributing

We love contributions! You are welcome to open a pull request, but it's a good idea to
open an issue and discuss your idea with us first.

Once you are ready to open a PR, please keep the following guidelines in mind:

1. Code should be `go fmt` compliant.
1. Types, structs and funcs should be documented.
1. Tests pass.

## Getting set up

`godo` uses go modules. Just fork this repo, clone your fork and off you go!

## Running tests

When working on code in this repository, tests can be run via:

```sh
go test -mod=vendor .
```

## Versioning

Godo follows [semver](https://www.semver.org) versioning semantics.
New functionality should be accompanied by increment to the minor
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
version number. Any code merged to main is subject to release.

## Releasing

Releasing a new version of godo is currently a manual process.

Submit a separate pull request for the version change from the pull
request with your changes.

1. Update the `CHANGELOG.md` with your changes. If a version header
   for the next (unreleased) version does not exist, create one.
   Include one bullet point for each piece of new functionality in the
   release, including the pull request ID, description, and author(s).
   For example:

```
## [v1.8.0] - 2019-03-13

- #210 - @jcodybaker - Expose tags on storage volume create/list/get.
- #123 - @digitalocean - Update test dependencies
```

   To generate a list of changes since the previous release in the correct
   format, you can use [github-changelog-generator](https://github.com/digitalocean/github-changelog-generator).
   It can be installed from source by running:

```
go get -u github.com/digitalocean/github-changelog-generator
```

   Next, list the changes by running:

```
github-changelog-generator -org digitalocean -repo godo
```

2. Update the `libraryVersion` number in `godo.go`.
3. Make a pull request with these changes.  This PR should be separate from the PR containing the godo changes.
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
version number. Any code merged to master is subject to release.
||||||| parent of 5ce8c7613 (update vendored files)
version number. Any code merged to master is subject to release.
=======
version number. Any code merged to main is subject to release.
>>>>>>> 5ce8c7613 (update vendored files)

## Releasing

Releasing a new version of godo is currently a manual process.

Submit a separate pull request for the version change from the pull
request with your changes.

1. Update the `CHANGELOG.md` with your changes. If a version header
   for the next (unreleased) version does not exist, create one.
   Include one bullet point for each piece of new functionality in the
   release, including the pull request ID, description, and author(s).
   For example:

```
## [v1.8.0] - 2019-03-13

- #210 - @jcodybaker - Expose tags on storage volume create/list/get.
- #123 - @digitalocean - Update test dependencies
```

   To generate a list of changes since the previous release in the correct
   format, you can use [github-changelog-generator](https://github.com/digitalocean/github-changelog-generator).
   It can be installed from source by running:

```
go get -u github.com/digitalocean/github-changelog-generator
```

   Next, list the changes by running:

```
github-changelog-generator -org digitalocean -repo godo
```

2. Update the `libraryVersion` number in `godo.go`.
3. Make a pull request with these changes.  This PR should be separate from the PR containing the godo changes.
<<<<<<< HEAD
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
=======
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
version number. Any code merged to master is subject to release.
||||||| parent of 6b7ce455e (update vendored files)
version number. Any code merged to master is subject to release.
=======
version number. Any code merged to main is subject to release.
>>>>>>> 6b7ce455e (update vendored files)

## Releasing

Releasing a new version of godo is currently a manual process.

Submit a separate pull request for the version change from the pull
request with your changes.

1. Update the `CHANGELOG.md` with your changes. If a version header
   for the next (unreleased) version does not exist, create one.
   Include one bullet point for each piece of new functionality in the
   release, including the pull request ID, description, and author(s).
   For example:

```
## [v1.8.0] - 2019-03-13

- #210 - @jcodybaker - Expose tags on storage volume create/list/get.
- #123 - @digitalocean - Update test dependencies
```

   To generate a list of changes since the previous release in the correct
   format, you can use [github-changelog-generator](https://github.com/digitalocean/github-changelog-generator).
   It can be installed from source by running:

```
go get -u github.com/digitalocean/github-changelog-generator
```

   Next, list the changes by running:

```
github-changelog-generator -org digitalocean -repo godo
```

2. Update the `libraryVersion` number in `godo.go`.
3. Make a pull request with these changes.  This PR should be separate from the PR containing the godo changes.
<<<<<<< HEAD
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
=======
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
version number. Any code merged to master is subject to release.
||||||| parent of 4d7e5ad26 (update vendored files)
version number. Any code merged to master is subject to release.
=======
version number. Any code merged to main is subject to release.
>>>>>>> 4d7e5ad26 (update vendored files)

## Releasing

Releasing a new version of godo is currently a manual process.

Submit a separate pull request for the version change from the pull
request with your changes.

1. Update the `CHANGELOG.md` with your changes. If a version header
   for the next (unreleased) version does not exist, create one.
   Include one bullet point for each piece of new functionality in the
   release, including the pull request ID, description, and author(s).
   For example:

```
## [v1.8.0] - 2019-03-13

- #210 - @jcodybaker - Expose tags on storage volume create/list/get.
- #123 - @digitalocean - Update test dependencies
```

   To generate a list of changes since the previous release in the correct
   format, you can use [github-changelog-generator](https://github.com/digitalocean/github-changelog-generator).
   It can be installed from source by running:

```
go get -u github.com/digitalocean/github-changelog-generator
```

   Next, list the changes by running:

```
github-changelog-generator -org digitalocean -repo godo
```

2. Update the `libraryVersion` number in `godo.go`.
3. Make a pull request with these changes.  This PR should be separate from the PR containing the godo changes.
<<<<<<< HEAD
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
=======
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
version number. Any code merged to master is subject to release.
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
version number. Any code merged to master is subject to release.
=======
version number. Any code merged to main is subject to release.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

## Releasing

Releasing a new version of godo is currently a manual process.

Submit a separate pull request for the version change from the pull
request with your changes.

1. Update the `CHANGELOG.md` with your changes. If a version header
   for the next (unreleased) version does not exist, create one.
   Include one bullet point for each piece of new functionality in the
   release, including the pull request ID, description, and author(s).
   For example:

```
## [v1.8.0] - 2019-03-13

- #210 - @jcodybaker - Expose tags on storage volume create/list/get.
- #123 - @digitalocean - Update test dependencies
```

   To generate a list of changes since the previous release in the correct
   format, you can use [github-changelog-generator](https://github.com/digitalocean/github-changelog-generator).
   It can be installed from source by running:

```
go get -u github.com/digitalocean/github-changelog-generator
```

   Next, list the changes by running:

```
github-changelog-generator -org digitalocean -repo godo
```

2. Update the `libraryVersion` number in `godo.go`.
3. Make a pull request with these changes.  This PR should be separate from the PR containing the godo changes.
<<<<<<< HEAD
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).  
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.  
=======
4. Once the pull request has been merged, [draft a new release](https://github.com/digitalocean/godo/releases/new).
5. Update the `Tag version` and `Release title` field with the new godo version.  Be sure the version has a `v` prefixed in both places. Ex `v1.8.0`.
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
6. Copy the changelog bullet points to the description field.
7. Publish the release.

## Go Version Support

This project follows the support [policy of Go](https://go.dev/doc/devel/release#policy)
as its support policy. The two latest major releases of Go are supported by the project.
[CI workflows](.github/workflows/ci.yml) should test against both supported versions.
[go.mod](./go.mod) should specify the oldest of the supported versions to give
downstream users of godo flexibility.
