# provider specific builds

## motivation

Usually clients will just use the external-dns for one specific provider. This means that big parts of the code are not needed during runtime. Having provider specific builds in order to produce binaries/images which just contain code which is needed for using one provider will reduce the size of the image and the build time.
However, imho the most important improvement comes for the clients regarding security. As user of provider "A" I don't want to be affected by a security issue in provider "B". Even if a project ( like this one) fixes security issues fast, it still needs the client to update his release version, which is not everywhere done automatically. So, the likelyhood of having a security issue in a provider specififc binary is much lower than in a binary which contains all providers.

### some numbers

| provider | binary size | build time | LOC     |
| -------- |------------ | ---------- | ------- |
| all      | 93 MB       | 1m15s      |  119K   |
| one      | 53 MB       | 47s        |   41K   |

So, a provider specific build
- is using **~60% less lines of code**,
- produces **~43% smaller binaries**,
- needs **~38% less build time**.

## implementation

The idea is to use [go build tags](https://golang.org/pkg/go/build/#hdr-Build_Constraints) to exclude provider specific code from the build. The build tag is the name of the provider and in order to support the build for all provider as it was before, there is the build tag `all` which is the default. The provider specific code is already well separated in the code base, but the `main.go` needs a refactoring and is splitted in provider specific go files ( ` main-<provider>.go`) in order to have here a clean separation as well. In order to prevent to flood the root directory with this main file variants, they are now located in `cmd/external-dns/`.

## backward compatibility

As mentioned above, the build tag `all` is the default in the `Makefile`, so all targets work as before. However, as a developer who is using an IDE like vscode or goland, you need to add the build tag `all` to the build configuration.

## provider specific images

The default build ( `make build.push` ), produces the same image name as before. The provider specific builds can be executed with setting the `BUILD_TAG` env variable, e.g. `BUILD_TAGS=aws make build.push`. The resulting image has a provider extension in the name, e.g. `us.gcr.io/k8s-artifacts-prod/external-dns/external-dns-aws`. There is a script [provider-builds.sh](./scripts/provider-builds.sh) which builds and pushes all provider specific images. Due to caching the build process of builds which are executed after the build for all providers is much faster.



