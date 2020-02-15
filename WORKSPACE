load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/rules_go/releases/download/v0.20.3/rules_go-v0.20.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.20.3/rules_go-v0.20.3.tar.gz",
    ],
    sha256 = "e88471aea3a3a4f19ec1310a55ba94772d087e9ce46e41ae38ecebe17935de7b",
)

http_archive(
    name = "bazel_gazelle",
    urls = [
        "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.19.1/bazel-gazelle-v0.19.1.tar.gz",
    ],
    sha256 = "86c6d481b3f7aedc1d60c1c211c6f76da282ae197c3b3160f54bd3a8f847896f",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

########################
# TESTING DEPENDENCIES #
########################

go_repository(
    name="com_github_smartystreets_goconvey",
    commit="044398e4856c3a83fb98f9de0a8ed1c6495e2843",
    importpath="github.com/smartystreets/goconvey",
)

go_repository(
    name="com_github_smartystreets_assertions",
    importpath="github.com/smartystreets/assertions",
    version="v1.0.1",
    sum="h1:voD4ITNjPL5jjBfgR/r8fPIIBrliWrWHeiJApdr3r4w=",
)

go_repository(
    name="com_github_jtolds_gls",
    commit="b4936e06046bbecbb94cae9c18127ebe510a2cb9",
    importpath="github.com/jtolds/gls",
)
