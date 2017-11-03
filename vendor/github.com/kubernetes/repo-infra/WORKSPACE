workspace(name = "io_kubernetes_build")

git_repository(
    name = "io_bazel_rules_go",
    commit = "7fefcfdf6ad108010bf0dca1b1aaa95e0d2c7ca4",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()
